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

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kedav1alpha1 "knative.dev/eventing-autoscaler-keda/third_party/pkg/apis/keda/v1alpha1"
	kafkav1beta1 "knative.dev/eventing-kafka/pkg/apis/sources/v1beta1"
	"knative.dev/pkg/kmeta"

	"knative.dev/eventing-autoscaler-keda/pkg/reconciler/keda"
)

const (
	defaultKafkaLagThreshold = 10
)

func GenerateScaleTarget(src *kafkav1beta1.KafkaSource) *kedav1alpha1.ScaleTarget {
	return &kedav1alpha1.ScaleTarget{
		Name:       src.Name,
		APIVersion: kafkav1beta1.SchemeGroupVersion.String(),
		Kind:       "KafkaSource",
	}
}

func GenerateScaleTriggers(src *kafkav1beta1.KafkaSource, triggerAuthentication *kedav1alpha1.TriggerAuthentication) ([]kedav1alpha1.ScaleTriggers, error) {
	triggers := []kedav1alpha1.ScaleTriggers{}
	bootstrapServers := strings.Join(src.Spec.BootstrapServers, ",")
	consumerGroup := src.Spec.ConsumerGroup

	lagThreshold, err := keda.GetInt32ValueFromMap(src.Annotations, keda.KedaAutoscalingKafkaLagThreshold, defaultKafkaLagThreshold)
	if err != nil {
		return nil, err
	}

	allowIdleConsumers := "false"
	if src.Status.Placements != nil {
		// KafkaSource is being managed by the multi-tenant controller.
		allowIdleConsumers = "true"
	}

	for _, topic := range src.Spec.Topics {
		triggerMetadata := map[string]string{
			"bootstrapServers":   bootstrapServers,
			"consumerGroup":      consumerGroup,
			"topic":              topic,
			"lagThreshold":       strconv.Itoa(int(*lagThreshold)),
			"allowIdleConsumers": allowIdleConsumers,
		}

		trigger := kedav1alpha1.ScaleTriggers{
			Type:              "kafka",
			Metadata:          triggerMetadata,
			AuthenticationRef: &kedav1alpha1.ScaledObjectAuthRef{},
		}

		if triggerAuthentication != nil {
			trigger.AuthenticationRef.Name = triggerAuthentication.Name
		}

		triggers = append(triggers, trigger)
	}

	return triggers, nil
}

func GenerateTriggerAuthentication(src *kafkav1beta1.KafkaSource, saslType *string) (*kedav1alpha1.TriggerAuthentication, *corev1.Secret, error) {

	secretTargetRefs := []kedav1alpha1.AuthSecretTargetRef{}

	secret := corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprintf("%s-secret", src.Name),
			Namespace: src.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*kmeta.NewControllerRef(src),
			},
		},
		Data:       make(map[string][]byte),
		StringData: make(map[string]string),
	}

	if src.Spec.KafkaAuthSpec.Net.SASL.Enable {

		if saslType != nil {
			switch *saslType {
			case "SCRAM-SHA-256":
				secret.StringData["sasl"] = "scram_sha256"
			case "SCRAM-SHA-512":
				secret.StringData["sasl"] = "scram_sha512"
			case "PLAIN":
				secret.StringData["sasl"] = "plaintext"
			default:
				return nil, nil, fmt.Errorf("SASL type value %q is not supported", *saslType)
			}
		} else {
			secret.StringData["sasl"] = "plaintext" //default
		}

		sasl := kedav1alpha1.AuthSecretTargetRef{Parameter: "sasl", Name: secret.Name, Key: "sasl"}

		username := kedav1alpha1.AuthSecretTargetRef{
			Parameter: "username",
			Name:      src.Spec.KafkaAuthSpec.Net.SASL.User.SecretKeyRef.Name,
			Key:       src.Spec.KafkaAuthSpec.Net.SASL.User.SecretKeyRef.Key,
		}
		password := kedav1alpha1.AuthSecretTargetRef{
			Parameter: "password",
			Name:      src.Spec.KafkaAuthSpec.Net.SASL.Password.SecretKeyRef.Name,
			Key:       src.Spec.KafkaAuthSpec.Net.SASL.Password.SecretKeyRef.Key,
		}

		secretTargetRefs = append(secretTargetRefs, sasl, username, password)
	}

	if src.Spec.KafkaAuthSpec.Net.TLS.Enable {
		secret.StringData["tls"] = "enable"
		tls := kedav1alpha1.AuthSecretTargetRef{Parameter: "tls", Name: secret.Name, Key: "tls"}
		secretTargetRefs = append(secretTargetRefs, tls)

		if src.Spec.KafkaAuthSpec.Net.TLS.CACert.SecretKeyRef != nil {
			ca := kedav1alpha1.AuthSecretTargetRef{
				Parameter: "ca",
				Name:      src.Spec.KafkaAuthSpec.Net.TLS.CACert.SecretKeyRef.Name,
				Key:       src.Spec.KafkaAuthSpec.Net.TLS.CACert.SecretKeyRef.Key,
			}

			secretTargetRefs = append(secretTargetRefs, ca)
		}

		if src.Spec.KafkaAuthSpec.Net.TLS.Cert.SecretKeyRef != nil {
			cert := kedav1alpha1.AuthSecretTargetRef{
				Parameter: "cert",
				Name:      src.Spec.KafkaAuthSpec.Net.TLS.Cert.SecretKeyRef.Name,
				Key:       src.Spec.KafkaAuthSpec.Net.TLS.Cert.SecretKeyRef.Key,
			}
			secretTargetRefs = append(secretTargetRefs, cert)
		}

		if src.Spec.KafkaAuthSpec.Net.TLS.Key.SecretKeyRef != nil {
			key := kedav1alpha1.AuthSecretTargetRef{
				Parameter: "key",
				Name:      src.Spec.KafkaAuthSpec.Net.TLS.Key.SecretKeyRef.Name,
				Key:       src.Spec.KafkaAuthSpec.Net.TLS.Key.SecretKeyRef.Key,
			}
			secretTargetRefs = append(secretTargetRefs, key)
		}
	}

	return &kedav1alpha1.TriggerAuthentication{
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprintf("%s-trigger-auth", src.Name),
			Namespace: src.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*kmeta.NewControllerRef(src),
			},
		},
		Spec: kedav1alpha1.TriggerAuthenticationSpec{
			SecretTargetRef: secretTargetRefs,
		},
	}, &secret, nil
}
