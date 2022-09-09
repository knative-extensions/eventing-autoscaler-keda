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

package broker

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/api/equality"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "knative.dev/eventing/pkg/apis/eventing/v1"
	brokerreconciler "knative.dev/eventing/pkg/client/injection/reconciler/eventing/v1/broker"
	"knative.dev/pkg/logging"
	pkgreconciler "knative.dev/pkg/reconciler"

	kedaclientset "knative.dev/eventing-autoscaler-keda/third_party/pkg/client/clientset/versioned"
	kedalisters "knative.dev/eventing-autoscaler-keda/third_party/pkg/client/listers/keda/v1alpha1"

	"knative.dev/eventing-autoscaler-keda/pkg/reconciler/broker/resources"
	"knative.dev/eventing-autoscaler-keda/pkg/reconciler/keda"
)

type Reconciler struct {
	kedaClientset kedaclientset.Interface

	// listers index properties about resources
	triggerAuthenticationLister kedalisters.TriggerAuthenticationLister
}

// Check that our Reconciler implements Interface
var _ brokerreconciler.Interface = (*Reconciler)(nil)

// This has to stay in sync with:
// / https://github.com/knative-sandbox/eventing-rabbitmq/blob/master/pkg/reconciler/broker/resources/secret.go#L29
const brokerURLSecretKey = "brokerURL"

// This has to stay in sync with:
// https://github.com/knative-sandbox/eventing-rabbitmq/blob/master/pkg/reconciler/broker/resources/secret.go#L49
func secretName(brokerName string) string {
	return fmt.Sprintf("%s-broker-rabbit", brokerName)
}

func (r *Reconciler) ReconcileKind(ctx context.Context, b *v1.Broker) pkgreconciler.Event {
	logging.FromContext(ctx).Debugw("Reconciling", zap.Any("Broker", b))

	// TODO: Check if there are any KEDA annotations before proceeding...
	// If they get updated / deleted, need to clean up.

	// Just reconciles a Broker and KEDA TriggerAuthentication to be used by ScaledObjects
	// that are then managed by ../trigger/
	err := r.reconcileScaleTriggerAuthentication(ctx, b)
	if err != nil {
		logging.FromContext(ctx).Errorw("Problem creating TriggerAuthentication", zap.Error(err))
		return err
	}
	return nil
}

func (r *Reconciler) reconcileScaleTriggerAuthentication(ctx context.Context, b *v1.Broker) error {
	namespace := b.Namespace
	// Check the annotation to see if the Brokers Triggers should even be scaled.
	doAutoscale := b.GetAnnotations()[keda.AutoscalingClassAnnotation] == keda.KEDA

	triggerAuthentication := resources.MakeTriggerAuthentication(b, secretName(b.Name), brokerURLSecretKey)

	current, err := r.triggerAuthenticationLister.TriggerAuthentications(namespace).Get(triggerAuthentication.Name)
	if apierrs.IsNotFound(err) {
		if !doAutoscale {
			// Ok, not there, not wanted, we're good...
			return nil
		}
		logging.FromContext(ctx).Errorw("Creating TriggerAuthentication", zap.Any("triggerauthentication", triggerAuthentication))
		_, err = r.kedaClientset.KedaV1alpha1().TriggerAuthentications(namespace).Create(ctx, triggerAuthentication, metav1.CreateOptions{})
		return err
	}
	if err != nil {
		return err
	}
	// It's there, should it be?
	if !doAutoscale {
		logging.FromContext(ctx).Infof("Deleting TriggerAuthentication %s/%s", namespace, triggerAuthentication.Name)
		err = r.kedaClientset.KedaV1alpha1().TriggerAuthentications(namespace).Delete(ctx, triggerAuthentication.Name, metav1.DeleteOptions{})
		if err != nil {
			logging.FromContext(ctx).Errorw("Failed to delete TriggerAuthentication", zap.Error(err))
		}
		return err
	}
	if !equality.Semantic.DeepDerivative(triggerAuthentication.Spec, triggerAuthentication.Spec) {
		// Don't modify the informers copy.
		desired := current.DeepCopy()
		desired.Spec = triggerAuthentication.Spec
		_, err = r.kedaClientset.KedaV1alpha1().TriggerAuthentications(namespace).Update(ctx, desired, metav1.UpdateOptions{})
		return err
	}
	return nil
}
