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

	kedav1aplha1 "github.com/kedacore/keda/api/v1alpha1"
	kafkav1beta1 "knative.dev/eventing-contrib/kafka/source/pkg/apis/sources/v1beta1"
	eventingutils "knative.dev/eventing/pkg/utils"

	"knative.dev/eventing-autoscaler-keda/pkg/reconciler/keda"
)

const (
	defaultKafkaLagThreshold = 10
)

func GenerateScaleTargetName(src *kafkav1beta1.KafkaSource) string {
	return eventingutils.GenerateFixedName(src, fmt.Sprintf("kafkasource-%s", src.Name))
}

func GenerateScaleTriggers(src *kafkav1beta1.KafkaSource) ([]kedav1aplha1.ScaleTriggers, error) {
	triggers := []kedav1aplha1.ScaleTriggers{}
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

		trigger := kedav1aplha1.ScaleTriggers{
			Type:     "kafka",
			Metadata: triggerMetadata,
		}
		triggers = append(triggers, trigger)
	}

	return triggers, nil
}
