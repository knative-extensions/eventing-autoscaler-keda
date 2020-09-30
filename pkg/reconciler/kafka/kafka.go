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

package kafka

import (
	"fmt"
	"strconv"
	"strings"

	"knative.dev/pkg/kmeta"

	kedav1alpha1 "github.com/kedacore/keda/api/v1alpha1"
	"knative.dev/eventing-autoscaler-keda/pkg/reconciler/keda"
	kafkav1beta1 "knative.dev/eventing-contrib/kafka/source/pkg/apis/sources/v1beta1"
)

const (
	defaultKafkaLagThreshold = 10
)

func GenerateScaleTargetName(src *kafkav1beta1.KafkaSource) string {
	return kmeta.ChildName(fmt.Sprintf("kafkasource-%s-", src.Name), string(src.GetUID()))
}

func GenerateScaleTriggers(src *kafkav1beta1.KafkaSource) ([]kedav1alpha1.ScaleTriggers, error) {
	triggers := []kedav1alpha1.ScaleTriggers{}
	bootstrapServers := strings.Join(src.Spec.BootstrapServers[:], ",")
	consumerGroup := src.Spec.ConsumerGroup

	lagThreshold, err := keda.GetInt32ValueFromMap(src.Annotations, keda.KedaAutoscalingKafkaLagThreshold, defaultKafkaLagThreshold)
	if err != nil {
		return nil, err
	}

	for _, topic := range src.Spec.Topics {
		triggerMetadata := map[string]string{
			"bootstrapServers": bootstrapServers,
			"consumerGroup":    consumerGroup,
			"topic":            topic,
			"lagThreshold":     strconv.Itoa(int(*lagThreshold)),
		}

		trigger := kedav1alpha1.ScaleTriggers{
			Type:     "kafka",
			Metadata: triggerMetadata,
		}
		triggers = append(triggers, trigger)
	}

	return triggers, nil
}
