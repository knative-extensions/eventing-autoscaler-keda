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

package resources

import (
	"context"
	"fmt"

	kedav1alpha1 "github.com/kedacore/keda/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"knative.dev/eventing/pkg/apis/eventing"
	v1 "knative.dev/eventing/pkg/apis/eventing/v1"
	"knative.dev/pkg/kmeta"
)

// This has to stay in sync with:
// https://github.com/knative-sandbox/eventing-rabbitmq/blob/master/pkg/reconciler/broker/resources/secret.go#L29
const brokerURLSecretKey = "brokerURL"

// DispatcherLabels generates the labels present on all resources representing the dispatcher of the given
// Broker.
func dispatcherLabels(brokerName string) map[string]string {
	return map[string]string{
		eventing.BrokerLabelKey:           brokerName,
		"eventing.knative.dev/brokerRole": "dispatcher",
	}
}

func MakeDispatcherScaledObject(ctx context.Context, t *v1.Trigger) *kedav1alpha1.ScaledObject {
	// TODO: plumb the Broker in here and get the values from there that we need to use.
	zero := int32(0)
	one := int32(1)
	five := int32(5)
	thirty := int32(30)

	queueName := fmt.Sprintf("%s/%s", t.Namespace, t.Name)
	triggerAuthenticationName := fmt.Sprintf("%s-trigger-auth", t.Spec.Broker)

	return &kedav1alpha1.ScaledObject{
		ObjectMeta: metav1.ObjectMeta{
			Name:      t.Name,
			Namespace: t.Namespace,
			Labels:    dispatcherLabels(t.Spec.Broker),
			OwnerReferences: []metav1.OwnerReference{
				*kmeta.NewControllerRef(t),
			},
		},
		Spec: kedav1alpha1.ScaledObjectSpec{
			ScaleTargetRef: &kedav1alpha1.ScaleTarget{
				Name:       t.Name,
				APIVersion: "apps/v1",
				Kind:       "Deployment",
			},
			PollingInterval: &five,   // seconds
			CooldownPeriod:  &thirty, // seconds
			MinReplicaCount: &zero,
			MaxReplicaCount: &one, // for now
			Triggers: []kedav1alpha1.ScaleTriggers{
				{
					Type: "rabbitmq",
					Metadata: map[string]string{
						"queueName":   queueName,
						"host":        brokerURLSecretKey,
						"queueLength": "1",
					},
					AuthenticationRef: &kedav1alpha1.ScaledObjectAuthRef{
						Name: triggerAuthenticationName,
					},
				},
			},
		},
	}
}
