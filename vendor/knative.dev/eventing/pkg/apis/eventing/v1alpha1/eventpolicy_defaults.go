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

package v1alpha1

import (
	"context"

	"knative.dev/pkg/apis"
)

func (ep *EventPolicy) SetDefaults(ctx context.Context) {
	ctx = apis.WithinParent(ctx, ep.ObjectMeta)
	ep.Spec.SetDefaults(ctx)
}

func (ets *EventPolicySpec) SetDefaults(ctx context.Context) {
	for i := range ets.From {
		ets.From[i].SetDefaults(ctx)
	}
}

func (from *EventPolicySpecFrom) SetDefaults(ctx context.Context) {
	if from.Ref != nil && from.Ref.Namespace == "" {
		// default to event policies namespace
		from.Ref.Namespace = apis.ParentMeta(ctx).Namespace
	}
}
