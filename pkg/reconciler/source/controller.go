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
	"time"

	"go.uber.org/zap"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/tools/cache"

	sourceinformer "knative.dev/pkg/client/injection/ducks/duck/v1/source"
	kubeclient "knative.dev/pkg/client/injection/kube/client"
	"knative.dev/pkg/configmap"
	"knative.dev/pkg/controller"
	"knative.dev/pkg/injection"
	"knative.dev/pkg/injection/clients/dynamicclient"
	"knative.dev/pkg/logging"
	pkgreconciler "knative.dev/pkg/reconciler"

	kedaclient "knative.dev/eventing-autoscaler-keda/pkg/client/injection/keda/client"
	//	scaledobjectinformer "knative.dev/eventing-autoscaler-keda/pkg/client/injection/keda/informers/keda/v1alpha1/scaledobject"
	kedaresources "knative.dev/eventing-autoscaler-keda/pkg/reconciler/keda"
)

const (
	// ReconcilerName is the name of the reconciler.
	ReconcilerName = "KEDASourceDucks"
)

// NewController returns a function that initializes the controller and
// Registers event handlers to enqueue events
func NewController(crd string, gvr schema.GroupVersionResource, gvk schema.GroupVersionKind) injection.ControllerConstructor {
	return func(ctx context.Context,
		cmw configmap.Watcher,
	) *controller.Impl {
		logger := logging.FromContext(ctx)
		sourceduckInformer := sourceinformer.Get(ctx)

		var sourceInformer cache.SharedIndexInformer
		var sourceLister cache.GenericLister

		var err error
		for i := 0; i < 10; i++ {
			sourceInformer, sourceLister, err = sourceduckInformer.Get(ctx, gvr)
			if err == nil {
				break
			} else if apierrors.IsNotFound(err) {
				logger.Debug("SourceDuckInformer not found -> waiting", zap.String("GVR", gvr.String()), zap.Error(err))
				time.Sleep(1 * time.Second)
			} else {
				logger.Errorw("Error getting source informer", zap.String("GVR", gvr.String()), zap.Error(err))
				return nil
			}
		}

		//	scaledobjectInformer := scaledobjectinformer.Get(ctx)

		r := &Reconciler{
			kubeClient:      kubeclient.Get(ctx),
			kedaClient:      kedaclient.Get(ctx),
			sourceInterface: dynamicclient.Get(ctx).Resource(gvr),
			sourceLister:    sourceLister,
			gvk:             gvk,
			gvr:             gvr,
		}
		impl := controller.NewImpl(r, logger, ReconcilerName)

		logger.Info("Setting up event handlers")
		sourceInformer.AddEventHandler(cache.FilteringResourceEventHandler{
			FilterFunc: pkgreconciler.AnnotationFilterFunc(kedaresources.AutoscalingClassAnnotation, kedaresources.KEDA, false),
			Handler:    controller.HandleAll(impl.Enqueue),
		})

		// FIXME don't handle updates on ScaledObject.Status field
		// scaledobjectInformer.Informer().AddEventHandler(cache.FilteringResourceEventHandler{
		// 	FilterFunc: controller.FilterControllerGVK(gvk),
		// 	Handler:    controller.HandleAll(impl.EnqueueControllerOf),
		// })

		return impl
	}
}
