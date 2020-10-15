/*
Copyright 2020 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package source

import (
	"context"
	"fmt"

	kedav1alpha1 "github.com/kedacore/keda/api/v1alpha1"
	kedaclientset "github.com/kedacore/keda/pkg/generated/clientset/versioned"
	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/equality"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/cache"
	awssqsv1alpha1 "knative.dev/eventing-awssqs/pkg/apis/sources/v1alpha1"
	kafkav1beta1 "knative.dev/eventing-kafka/pkg/apis/sources/v1beta1"
	duckv1 "knative.dev/pkg/apis/duck/v1"
	"knative.dev/pkg/logging"
	pkgreconciler "knative.dev/pkg/reconciler"

	"knative.dev/eventing-autoscaler-keda/pkg/reconciler/awssqs"
	"knative.dev/eventing-autoscaler-keda/pkg/reconciler/kafka"
	"knative.dev/eventing-autoscaler-keda/pkg/reconciler/keda"
)

type Reconciler struct {
	kubeClient      kubernetes.Interface
	kedaClient      kedaclientset.Interface
	sourceInterface dynamic.NamespaceableResourceInterface
	sourceLister    cache.GenericLister
	gvk             schema.GroupVersionKind
	gvr             schema.GroupVersionResource
}

func (r *Reconciler) Reconcile(ctx context.Context, key string) error {

	// Convert the namespace/name string into a distinct namespace and name
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		logging.FromContext(ctx).Errorw("invalid resource key")
		return nil
	}

	// Get the Source resource with this namespace/name
	runtimeObj, err := r.sourceLister.ByNamespace(namespace).Get(name)
	if err != nil {
		logging.FromContext(ctx).Errorw("not able to get runtime object")
		return nil
	}
	if apierrors.IsNotFound(err) {
		// The resource may no longer exist, in which case we stop processing.
		logging.FromContext(ctx).Errorw("Source in work queue no longer exists")
		return nil
	} else if err != nil {
		return err
	}

	logging.FromContext(ctx).Info("RECONCILE kind: " + runtimeObj.GetObjectKind().GroupVersionKind().String())

	var ok bool
	if _, ok = runtimeObj.(*duckv1.Source); !ok {
		logging.FromContext(ctx).Errorw("runtime object is not convertible to Source duck type")
		// Avoid re-enqueuing.
		return nil
	}

	if runtimeObj.GetObjectKind().GroupVersionKind() != r.gvk {
		logging.FromContext(ctx).Errorw("runtime object is GVK doesn't match GVK specified for Reconciler")
		// Avoid re-enqueuing.
		return nil
	}

	unstructuredSource, err := r.sourceInterface.Namespace(namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		logging.FromContext(ctx).Errorw("Failed to get Unstructured Source:", zap.Error(err))
		return nil
	}

	switch r.gvk.Kind {
	case "KafkaSource":
		var kafkaSource = &kafkav1beta1.KafkaSource{}
		// TODO move scheme register up
		kafkav1beta1.AddToScheme(scheme.Scheme)
		if err := scheme.Scheme.Convert(unstructuredSource, kafkaSource, nil); err != nil {
			logging.FromContext(ctx).Errorw("Failed to convert Unstructured Source to KafkaSource", zap.Error(err))
			return err
		}
		return r.reconcileKafkaSource(ctx, kafkaSource)
	case "AwsSqsSource":
		var awsSqsSource = &awssqsv1alpha1.AwsSqsSource{}
		// TODO move scheme register up
		awssqsv1alpha1.AddToScheme(scheme.Scheme)
		if err := scheme.Scheme.Convert(unstructuredSource, awsSqsSource, nil); err != nil {
			logging.FromContext(ctx).Errorw("Failed to convert Unstructured Source to AwsSqsSource", zap.Error(err))
			return err
		}
		return r.reconcileAwsSqsSource(ctx, awsSqsSource)
	}

	return nil
}

func (r *Reconciler) reconcileKafkaSource(ctx context.Context, src *kafkav1beta1.KafkaSource) error {

	var triggerAuthentication *kedav1alpha1.TriggerAuthentication
	var secret *corev1.Secret
	if src.Spec.KafkaAuthSpec.Net.SASL.Enable || src.Spec.KafkaAuthSpec.Net.TLS.Enable {
		triggerAuthentication, secret = kafka.GenerateTriggerAuthentication(src)
	}

	triggers, err := kafka.GenerateScaleTriggers(src, triggerAuthentication)
	if err != nil {
		return err
	}
	scaledObject, err := keda.GenerateScaledObject(src, r.gvk, kafka.GenerateScaleTargetName(src), triggers)
	if err != nil {
		return err
	}

	if triggerAuthentication != nil && secret != nil {
		err = r.reconcileSecret(ctx, secret, src)
		if err != nil {
			return err
		}
		err = r.reconcileTriggerAuthentication(ctx, triggerAuthentication, src)
		if err != nil {
			return err
		}
	}

	return r.reconcileScaledObject(ctx, scaledObject, src)
}

func (r *Reconciler) reconcileAwsSqsSource(ctx context.Context, src *awssqsv1alpha1.AwsSqsSource) error {
	scaletarget, err := awssqs.GenerateScaleTargetName(ctx, r.kubeClient, src)
	if err != nil {
		return err
	}
	triggers, err := awssqs.GenerateScaleTriggers(src)
	if err != nil {
		return err
	}
	scaledObject, err := keda.GenerateScaledObject(src, r.gvk, scaletarget, triggers)
	if err != nil {
		return err
	}

	return r.reconcileScaledObject(ctx, scaledObject, src)
}

func (r *Reconciler) reconcileScaledObject(ctx context.Context, expectedScaledObject *kedav1alpha1.ScaledObject, obj metav1.Object) error {
	scaledObject, err := r.kedaClient.KedaV1alpha1().ScaledObjects(expectedScaledObject.Namespace).Get(ctx, expectedScaledObject.Name, metav1.GetOptions{})
	if apierrors.IsNotFound(err) {
		scaledObject, err = r.kedaClient.KedaV1alpha1().ScaledObjects(expectedScaledObject.Namespace).Create(ctx, expectedScaledObject, metav1.CreateOptions{})
		if err != nil {
			return scaleObjectDeploymentFailed(expectedScaledObject.Namespace, expectedScaledObject.Name, err)
		}
		return scaleObjectCreated(scaledObject.Namespace, scaledObject.Name)
	} else if err != nil {
		logging.FromContext(ctx).Errorw("Unable to get an existing ScaledObject", zap.Error(err))
		return err
	} else if !metav1.IsControlledBy(scaledObject, obj) {
		return fmt.Errorf("ScaledObject %q is not owned by %q", scaledObject.Name, obj)
	} else if !equality.Semantic.DeepDerivative(scaledObject.Spec, expectedScaledObject.Spec) {
		logging.FromContext(ctx).Debug(fmt.Sprintf("ScaledObject changed, found: %#v expected: %#v", scaledObject.Spec, expectedScaledObject.Spec))
		scaledObject.Spec = expectedScaledObject.Spec
		if _, err = r.kedaClient.KedaV1alpha1().ScaledObjects(expectedScaledObject.Namespace).Update(ctx, scaledObject, metav1.UpdateOptions{}); err != nil {
			return err
		}
		return scaleObjectUpdated(scaledObject.Namespace, scaledObject.Name)
	} else {
		logging.FromContext(ctx).Debugw("Reusing existing ScaledObject", zap.Any("ScaledObject", scaledObject))
	}

	return nil
}

func (r *Reconciler) reconcileTriggerAuthentication(ctx context.Context, expectedTriggerAuth *kedav1alpha1.TriggerAuthentication, obj metav1.Object) error {
	triggerAuth, err := r.kedaClient.KedaV1alpha1().TriggerAuthentications(expectedTriggerAuth.Namespace).Get(ctx, expectedTriggerAuth.Name, metav1.GetOptions{})
	if apierrors.IsNotFound(err) {
		triggerAuth, err = r.kedaClient.KedaV1alpha1().TriggerAuthentications(expectedTriggerAuth.Namespace).Create(ctx, expectedTriggerAuth, metav1.CreateOptions{})
		if err != nil {
			return pkgreconciler.NewEvent(corev1.EventTypeWarning, "TriggerAuthenticationFailed", "TriggerAuthentication deployment failed to: \"%s/%s\", %w",
				expectedTriggerAuth.Namespace, expectedTriggerAuth.Name, err)
		}
		return pkgreconciler.NewEvent(corev1.EventTypeNormal, "TriggerAuthenticationCreated", "TriggerAuthentication created: \"%s/%s\"",
			triggerAuth.Namespace, triggerAuth.Name)
	} else if err != nil {
		logging.FromContext(ctx).Errorw("Unable to get an existing ScaledObject", zap.Error(err))
		return err
	} else if !metav1.IsControlledBy(triggerAuth, obj) {
		return fmt.Errorf("ScaledObject %q is not owned by %q", triggerAuth.Name, obj)
	} else if !equality.Semantic.DeepDerivative(triggerAuth.Spec, expectedTriggerAuth.Spec) {
		logging.FromContext(ctx).Debug(fmt.Sprintf("TriggerAuthentication changed, found: %#v expected: %#v", triggerAuth.Spec, expectedTriggerAuth.Spec))
		triggerAuth.Spec = expectedTriggerAuth.Spec
		if _, err = r.kedaClient.KedaV1alpha1().TriggerAuthentications(expectedTriggerAuth.Namespace).Update(ctx, triggerAuth, metav1.UpdateOptions{}); err != nil {
			return err
		}
		return pkgreconciler.NewEvent(corev1.EventTypeNormal, "TriggerAuthenticationUpdated", "TriggerAuthentication updated: \"%s/%s\"",
			triggerAuth.Namespace, triggerAuth.Name)
	} else {
		logging.FromContext(ctx).Debugw("Reusing existing TriggerAuthentication", zap.Any("TriggerAuthentication", triggerAuth))
	}

	return nil
}

func (r *Reconciler) reconcileSecret(ctx context.Context, expectedSecret *corev1.Secret, obj metav1.Object) error {
	secret, err := r.kubeClient.CoreV1().Secrets(expectedSecret.Namespace).Get(ctx, expectedSecret.Name, metav1.GetOptions{})
	if apierrors.IsNotFound(err) {
		secret, err = r.kubeClient.CoreV1().Secrets(expectedSecret.Namespace).Create(ctx, expectedSecret, metav1.CreateOptions{})
		if err != nil {
			return pkgreconciler.NewEvent(corev1.EventTypeWarning, "SecretDeploymentFailed", "Secret deployment failed to: \"%s/%s\", %w",
				expectedSecret.Namespace, expectedSecret.Name, err)
		}
		return pkgreconciler.NewEvent(corev1.EventTypeNormal, "SecretCreated", "Secret created: \"%s/%s\"", secret.Namespace, secret.Name)
	} else if err != nil {
		logging.FromContext(ctx).Errorw("Unable to get an existing ScaledObject", zap.Error(err))
		return err
	} else if !metav1.IsControlledBy(secret, obj) {
		return fmt.Errorf("Secret %q is not owned by %q", secret.Name, obj)
	} else if !equality.Semantic.DeepDerivative(secret.Data, expectedSecret.Data) {
		logging.FromContext(ctx).Debug(fmt.Sprintf("Secret changed, found: %#v expected: %#v", secret.Data, expectedSecret.Data))
		secret.Data = expectedSecret.Data
		if _, err = r.kubeClient.CoreV1().Secrets(expectedSecret.Namespace).Update(ctx, secret, metav1.UpdateOptions{}); err != nil {
			return err
		}
		return pkgreconciler.NewEvent(corev1.EventTypeNormal, "SecretUpdated", "Secret updated: \"%s/%s\"", secret.Namespace, secret.Name)
	} else {
		logging.FromContext(ctx).Debugw("Reusing existing Secret", zap.Any("Secret", secret))
	}

	return nil
}

// scaleObjectCreated makes a new reconciler event with event type Normal, and
// reason ScaledObjectCreated.
func scaleObjectCreated(namespace, name string) pkgreconciler.Event {
	return pkgreconciler.NewEvent(corev1.EventTypeNormal, "ScaledObjectCreated", "ScaledObject created: \"%s/%s\"", namespace, name)
}

// scaleObjectUpdated makes a new reconciler event with event type Normal, and
// reason ScaledObjectUpdated.
func scaleObjectUpdated(namespace, name string) pkgreconciler.Event {
	return pkgreconciler.NewEvent(corev1.EventTypeNormal, "ScaledObjectUpdated", "ScaledObject updated: \"%s/%s\"", namespace, name)
}

// scaleObjectDeploymentFailed makes a new reconciler event with event type Warning, and
// reason ScaleObjectDeploymentFailed.
func scaleObjectDeploymentFailed(namespace, name string, err error) pkgreconciler.Event {
	return pkgreconciler.NewEvent(corev1.EventTypeWarning, "ScaleObjectDeploymentFailed", "ScaledObject deployment failed to: \"%s/%s\", %w", namespace, name, err)
}
