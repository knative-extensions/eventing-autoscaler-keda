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

package trigger

import (
	"context"

	"go.uber.org/zap"

	"k8s.io/apimachinery/pkg/api/equality"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"knative.dev/pkg/logging"

	"knative.dev/eventing-autoscaler-keda/pkg/reconciler/keda"
	"knative.dev/eventing-autoscaler-keda/pkg/reconciler/trigger/resources"
	"knative.dev/eventing/pkg/apis/eventing"
	v1 "knative.dev/eventing/pkg/apis/eventing/v1"
	triggerreconciler "knative.dev/eventing/pkg/client/injection/reconciler/eventing/v1/trigger"
	eventinglisters "knative.dev/eventing/pkg/client/listers/eventing/v1"

	pkgreconciler "knative.dev/pkg/reconciler"

	kedaclientset "github.com/kedacore/keda/pkg/generated/clientset/versioned"
	kedalisters "github.com/kedacore/keda/pkg/generated/listers/keda/v1alpha1"
)

const (
	// Name of the corev1.Events emitted from the Trigger reconciliation process.
	triggerReconciled = "TriggerReconciled"
)

type Reconciler struct {
	kedaClientset kedaclientset.Interface

	// listers index properties about resources
	brokerLister       eventinglisters.BrokerLister
	scaledObjectLister kedalisters.ScaledObjectLister
	brokerClass        string
}

// Check that our Reconciler implements Interface
var _ triggerreconciler.Interface = (*Reconciler)(nil)

// ReconcilerArgs are the arguments needed to create a broker.Reconciler.
type ReconcilerArgs struct {
	DispatcherImage              string
	DispatcherServiceAccountName string
}

func (r *Reconciler) ReconcileKind(ctx context.Context, t *v1.Trigger) pkgreconciler.Event {
	logging.FromContext(ctx).Debug("Reconciling", zap.Any("Trigger", t))

	broker, err := r.brokerLister.Brokers(t.Namespace).Get(t.Spec.Broker)
	if err != nil {
		if apierrs.IsNotFound(err) {
			// Ok to return nil here. Once the Broker comes available, or Trigger changes, we get requeued.
			return nil
		}
		logging.FromContext(ctx).Errorf("Failed to get Broker: \"%s/%s\" : %s", t.Spec.Broker, t.Namespace, err)
		return nil
	}

	// If it's not my brokerclass, ignore
	if broker.Annotations[eventing.BrokerClassKey] != r.brokerClass {
		logging.FromContext(ctx).Infof("Ignoring trigger %s/%s", t.Namespace, t.Name)
		return nil
	}

	return r.reconcileScaledObject(ctx, broker, t)
}

func (r *Reconciler) reconcileScaledObject(ctx context.Context, broker *v1.Broker, trigger *v1.Trigger) error {
	// Check the annotation to see if the Brokers Triggers should even be scaled.
	doAutoscale := broker.GetAnnotations()[keda.AutoscalingClassAnnotation] == keda.KEDA

	so, err := resources.MakeDispatcherScaledObject(ctx, broker, trigger)
	if err != nil {
		logging.FromContext(ctx).Errorw("Failed to create scaled object resource", zap.Error(err))
		return err
	}

	current, err := r.scaledObjectLister.ScaledObjects(so.Namespace).Get(so.Name)
	if apierrs.IsNotFound(err) {
		if !doAutoscale {
			// Ok, not there, not wanted, we're good...
			return nil
		}
		logging.FromContext(ctx).Infof("Creating ScaledObject %s/%s", so.Namespace, so.Name)
		_, err = r.kedaClientset.KedaV1alpha1().ScaledObjects(so.Namespace).Create(ctx, so, metav1.CreateOptions{})
		return err
	}
	if err != nil {
		return err
	}

	// It's there, should it be?
	if !doAutoscale {
		logging.FromContext(ctx).Infof("Deleting ScaledObject %s/%s", so.Namespace, so.Name)
		err = r.kedaClientset.KedaV1alpha1().ScaledObjects(so.Namespace).Delete(ctx, so.Name, metav1.DeleteOptions{})
		if err != nil {
			logging.FromContext(ctx).Errorw("Failed to delete ScaledObject", zap.Error(err))
		}
		return err
	}
	if !equality.Semantic.DeepDerivative(so.Spec, current.Spec) {
		// Don't modify the informers copy.
		desired := current.DeepCopy()
		desired.Spec = so.Spec
		_, err = r.kedaClientset.KedaV1alpha1().ScaledObjects(desired.Namespace).Update(ctx, desired, metav1.UpdateOptions{})
		return err
	}
	return nil
}
