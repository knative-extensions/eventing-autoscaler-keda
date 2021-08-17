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

package redisstream

import (
	"fmt"
	"strconv"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	redisstreamv1alpha1 "knative.dev/eventing-redis/source/pkg/apis/sources/v1alpha1"
	"knative.dev/pkg/kmeta"

	kedav1alpha1 "knative.dev/eventing-autoscaler-keda/third_party/pkg/apis/keda/v1alpha1"

	"knative.dev/eventing-autoscaler-keda/pkg/reconciler/keda"
)

const (
	defaultRedisStreamPendingEntriesCount = 5
)

func GenerateScaleTarget(src *redisstreamv1alpha1.RedisStreamSource) *kedav1alpha1.ScaleTarget {
	return &kedav1alpha1.ScaleTarget{
		Name:       src.Name,
		APIVersion: redisstreamv1alpha1.SchemeGroupVersion.String(),
		Kind:       "RedisStreamSource",
	}
}

func GenerateScaleTriggers(src *redisstreamv1alpha1.RedisStreamSource, triggerAuthentication *kedav1alpha1.TriggerAuthentication) ([]kedav1alpha1.ScaleTriggers, error) {
	pendingEntriesCount, err := keda.GetInt32ValueFromMap(src.Annotations, keda.KedaAutoscalingRedisStreamPendingEntriesCount, defaultRedisStreamPendingEntriesCount)
	if err != nil {
		return nil, err
	}

	triggerMetadata := map[string]string{
		"address":             src.Spec.Address,
		"stream":              src.Spec.Stream,
		"consumerGroup":       src.Spec.Group, //TODO: this is an optional field. ConsumerGroup name is set to the pod name, if one not provided.
		"pendingEntriesCount": strconv.Itoa(int(*pendingEntriesCount)),
	}

	trigger := kedav1alpha1.ScaleTriggers{
		Type:              "redis-streams",
		Metadata:          triggerMetadata,
		AuthenticationRef: &kedav1alpha1.ScaledObjectAuthRef{},
	}

	if triggerAuthentication != nil {
		trigger.AuthenticationRef.Name = triggerAuthentication.Name
	}

	return []kedav1alpha1.ScaleTriggers{
		trigger,
	}, nil
}

func GenerateTriggerAuthentication(src *redisstreamv1alpha1.RedisStreamSource) (*kedav1alpha1.TriggerAuthentication, *corev1.Secret) {

	secretTargetRefs := []kedav1alpha1.AuthSecretTargetRef{}
	secretKey := "redis_password"

	secret := corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprintf("%s-secret", src.Name),
			Namespace: src.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*kmeta.NewControllerRef(src),
			},
		},
	}

	secret.StringData[secretKey] = src.Spec.Options.Password.String()

	password := kedav1alpha1.AuthSecretTargetRef{
		Parameter: "password",
		Name:      secret.Name,
		Key:       secretKey,
	}

	secretTargetRefs = append(secretTargetRefs, password)

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
	}, &secret
}
