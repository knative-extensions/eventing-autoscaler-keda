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

	"k8s.io/client-go/tools/cache"

	eventingv1 "knative.dev/eventing/pkg/apis/eventing/v1"

	brokerinformer "knative.dev/eventing/pkg/client/injection/informers/eventing/v1/broker"
	brokerreconciler "knative.dev/eventing/pkg/client/injection/reconciler/eventing/v1/broker"
	"knative.dev/pkg/configmap"
	"knative.dev/pkg/controller"
	"knative.dev/pkg/logging"
	pkgreconciler "knative.dev/pkg/reconciler"

	kedaclient "knative.dev/eventing-autoscaler-keda/pkg/client/injection/keda/client"
	triggerauthenticationinformer "knative.dev/eventing-autoscaler-keda/pkg/client/injection/keda/informers/keda/v1alpha1/triggerauthentication"
)

// Only can scale RabbitMQBrokers
const brokerClass = "RabbitMQBroker"

// NewController initializes the controller and is called by the generated code
// Registers event handlers to enqueue events
func NewController(
	ctx context.Context,
	cmw configmap.Watcher,
) *controller.Impl {
	brokerInformer := brokerinformer.Get(ctx)
	triggerAuthenticationInformer := triggerauthenticationinformer.Get(ctx)

	r := &Reconciler{
		kedaClientset:               kedaclient.Get(ctx),
		triggerAuthenticationLister: triggerAuthenticationInformer.Lister(),
	}

	impl := brokerreconciler.NewImpl(ctx, r, brokerClass)

	logging.FromContext(ctx).Info("Setting up event handlers")

	brokerInformer.Informer().AddEventHandler(cache.FilteringResourceEventHandler{
		FilterFunc: pkgreconciler.AnnotationFilterFunc(brokerreconciler.ClassAnnotationKey, brokerClass, false /*allowUnset*/),
		Handler:    controller.HandleAll(impl.Enqueue),
	})

	triggerAuthenticationInformer.Informer().AddEventHandler(cache.FilteringResourceEventHandler{
		FilterFunc: controller.FilterControllerGK(eventingv1.Kind("Broker")),
		Handler:    controller.HandleAll(impl.EnqueueControllerOf),
	})

	return impl
}
