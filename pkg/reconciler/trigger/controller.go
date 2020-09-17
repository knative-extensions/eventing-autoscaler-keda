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
	"log"

	kedaclient "knative.dev/eventing-autoscaler-keda/pkg/client/injection/keda/client"

	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"

	scaledobjectinformer "knative.dev/eventing-autoscaler-keda/pkg/client/injection/keda/informers/keda/v1alpha1/scaledobject"
	v1 "knative.dev/eventing/pkg/apis/eventing/v1"
	brokerinformer "knative.dev/eventing/pkg/client/injection/informers/eventing/v1/broker"
	triggerinformer "knative.dev/eventing/pkg/client/injection/informers/eventing/v1/trigger"
	brokerreconciler "knative.dev/eventing/pkg/client/injection/reconciler/eventing/v1/broker"
	triggerreconciler "knative.dev/eventing/pkg/client/injection/reconciler/eventing/v1/trigger"

	"knative.dev/pkg/configmap"
	"knative.dev/pkg/controller"
	"knative.dev/pkg/logging"
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
	triggerInformer := triggerinformer.Get(ctx)
	scaledObjectInformer := scaledobjectinformer.Get(ctx)

	r := &Reconciler{
		kedaClientset:      kedaclient.Get(ctx),
		brokerLister:       brokerInformer.Lister(),
		scaledObjectLister: scaledObjectInformer.Lister(),
		brokerClass:        brokerClass,
	}

	impl := triggerreconciler.NewImpl(ctx, r)

	logging.FromContext(ctx).Info("Setting up event handlers")

	brokerInformer.Informer().AddEventHandler(controller.HandleAll(
		func(obj interface{}) {
			if broker, ok := obj.(*v1.Broker); ok {
				triggers, err := triggerInformer.Lister().Triggers(broker.Namespace).List(labels.Everything())
				if err != nil {
					log.Print("Failed to lookup Triggers for Broker", zap.Error(err))
				} else {
					for _, t := range triggers {
						if t.Spec.Broker == broker.Name {
							impl.Enqueue(t)
						}
					}
				}
			}
		},
	))

	triggerInformer.Informer().AddEventHandler(controller.HandleAll(
		func(obj interface{}) {
			if trigger, ok := obj.(*v1.Trigger); ok {
				broker, err := brokerInformer.Lister().Brokers(trigger.Namespace).Get(trigger.Spec.Broker)
				if err != nil {
					log.Print("Failed to lookup Broker for Trigger", zap.Error(err))
				} else {
					label := broker.ObjectMeta.Annotations[brokerreconciler.ClassAnnotationKey]
					if label == brokerClass {
						impl.Enqueue(obj)
					}
				}
			}
		},
	))

	scaledObjectInformer.Informer().AddEventHandler(cache.FilteringResourceEventHandler{
		FilterFunc: controller.FilterControllerGK(v1.Kind("Trigger")),
		Handler:    controller.HandleAll(impl.EnqueueControllerOf),
	})

	return impl
}
