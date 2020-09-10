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
	awssqsv1alpha1 "knative.dev/eventing-contrib/awssqs/pkg/apis/sources/v1alpha1"
	kafkav1beta1 "knative.dev/eventing-contrib/kafka/source/pkg/apis/sources/v1beta1"
	"knative.dev/eventing/pkg/logging"
	duckv1 "knative.dev/pkg/apis/duck/v1"
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
		logging.FromContext(ctx).Error("invalid resource key")
		return nil
	}

	// Get the Source resource with this namespace/name
	runtimeObj, err := r.sourceLister.ByNamespace(namespace).Get(name)
	if err != nil {
		logging.FromContext(ctx).Error("not able to get runtime object")
		return nil
	}
	if apierrors.IsNotFound(err) {
		// The resource may no longer exist, in which case we stop processing.
		logging.FromContext(ctx).Error("Source in work queue no longer exists")
		return nil
	} else if err != nil {
		return err
	}

	logging.FromContext(ctx).Info("RECONCILE kind: " + runtimeObj.GetObjectKind().GroupVersionKind().String())

	var ok bool
	if _, ok = runtimeObj.(*duckv1.Source); !ok {
		logging.FromContext(ctx).Error("runtime object is not convertible to Source duck type")
		// Avoid re-enqueuing.
		return nil
	}

	if runtimeObj.GetObjectKind().GroupVersionKind() != r.gvk {
		logging.FromContext(ctx).Error("runtime object is GVK doesn't match GVK specified for Reconciler")
		// Avoid re-enqueuing.
		return nil
	}

	unstructuredSource, err := r.sourceInterface.Namespace(namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		logging.FromContext(ctx).Error("Failed to get Unstructured Source:", zap.Error(err))
		return nil
	}

	switch r.gvk.Kind {
	case "KafkaSource":
		var kafkaSource = &kafkav1beta1.KafkaSource{}
		// TODO move scheme register up
		kafkav1beta1.AddToScheme(scheme.Scheme)
		if err := scheme.Scheme.Convert(unstructuredSource, kafkaSource, nil); err != nil {
			logging.FromContext(ctx).Error("Failed to convert Unstructured Source to KafkaSource", zap.Error(err))
			return err
		}
		return r.reconcileKafkaSource(ctx, kafkaSource)
	case "AwsSqsSource":
		var awsSqsSource = &awssqsv1alpha1.AwsSqsSource{}
		// TODO move scheme register up
		awssqsv1alpha1.AddToScheme(scheme.Scheme)
		if err := scheme.Scheme.Convert(unstructuredSource, awsSqsSource, nil); err != nil {
			logging.FromContext(ctx).Error("Failed to convert Unstructured Source to AwsSqsSource", zap.Error(err))
			return err
		}
		return r.reconcileAwsSqsSource(ctx, awsSqsSource)
	}

	return nil
}

func (r *Reconciler) reconcileKafkaSource(ctx context.Context, src *kafkav1beta1.KafkaSource) error {
	triggers, err := kafka.GenerateScaleTriggers(src)
	if err != nil {
		return err
	}
	scaledObject, err := keda.GenerateScaledObject(src, r.gvk, kafka.GenerateScaleTargetName(src), triggers)
	if err != nil {
		return err
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
			return scaleObjectDeploymentFailed(scaledObject.Namespace, scaledObject.Name, err)
		}
		return scaleObjectCreated(scaledObject.Namespace, scaledObject.Name)
	} else if err != nil {
		logging.FromContext(ctx).Error("Unable to get an existing ScaledObject", zap.Error(err))
		return err
	} else if !metav1.IsControlledBy(scaledObject, obj) {
		return fmt.Errorf("ScaledObject %q is not owned by %q", scaledObject.Name, obj)
	} else if !equality.Semantic.DeepDerivative(scaledObject.Spec, expectedScaledObject.Spec) {
		scaledObject.Spec = expectedScaledObject.Spec
		if _, err = r.kedaClient.KedaV1alpha1().ScaledObjects(expectedScaledObject.Namespace).Update(ctx, scaledObject, metav1.UpdateOptions{}); err != nil {
			return err
		}
		return scaleObjectUpdated(scaledObject.Namespace, scaledObject.Name)
	} else {
		logging.FromContext(ctx).Debug("Reusing existing ScaledObject", zap.Any("ScaledObject", scaledObject))
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
