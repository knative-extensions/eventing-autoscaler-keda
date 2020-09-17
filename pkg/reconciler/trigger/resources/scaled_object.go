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
	"knative.dev/eventing-autoscaler-keda/pkg/reconciler/keda"
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

// Consider refactoring this so that we can reuse more of the pkg/reconciler/keda/resources
const (
	defaultPollingInterval = 30
	defaultCooldownPeriod  = 300
	defaultMinReplicaCount = 0
	defaultMaxReplicaCount = 50
	defaultQueueLength     = "1"
)

func MakeDispatcherScaledObject(ctx context.Context, b *v1.Broker, t *v1.Trigger) (*kedav1alpha1.ScaledObject, error) {
	cooldownPeriod, err := keda.GetInt32ValueFromMap(b.GetAnnotations(), keda.KedaAutoscalingCooldownPeriodAnnotation, defaultCooldownPeriod)
	if err != nil {
		return nil, err
	}
	pollingInterval, err := keda.GetInt32ValueFromMap(b.GetAnnotations(), keda.KedaAutoscalingPollingIntervalAnnotation, defaultPollingInterval)
	if err != nil {
		return nil, err
	}
	minReplicaCount, err := keda.GetInt32ValueFromMap(b.GetAnnotations(), keda.AutoscalingMinScaleAnnotation, defaultMinReplicaCount)
	if err != nil {
		return nil, err
	}
	maxReplicaCount, err := keda.GetInt32ValueFromMap(b.GetAnnotations(), keda.AutoscalingMaxScaleAnnotation, defaultMaxReplicaCount)
	if err != nil {
		return nil, err
	}

	queueLength := b.GetAnnotations()[keda.KedaAutoscalingRabbitMQQueueLength]
	if queueLength == "" {
		queueLength = defaultQueueLength
	}

	// TODO(vaikas): https://github.com/knative-sandbox/eventing-rabbitmq/issues/61
	// queueName := fmt.Sprintf("%s/%s", t.Namespace, t.Name)
	queueName := fmt.Sprintf("%s-%s", t.Namespace, t.Name)
	deploymentName := fmt.Sprintf("%s-dispatcher", t.Name)
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
				Name:       deploymentName,
				APIVersion: "apps/v1",
				Kind:       "Deployment",
			},
			PollingInterval: pollingInterval,
			CooldownPeriod:  cooldownPeriod,
			MinReplicaCount: minReplicaCount,
			MaxReplicaCount: maxReplicaCount,
			Triggers: []kedav1alpha1.ScaleTriggers{
				{
					Type: "rabbitmq",
					Metadata: map[string]string{
						"queueName":   queueName,
						"host":        brokerURLSecretKey,
						"queueLength": queueLength,
					},
					AuthenticationRef: &kedav1alpha1.ScaledObjectAuthRef{
						Name: triggerAuthenticationName,
					},
				},
			},
		},
	}, nil
}
