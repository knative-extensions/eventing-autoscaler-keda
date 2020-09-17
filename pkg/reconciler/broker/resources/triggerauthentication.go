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
	"fmt"

	kedav1alpha1 "github.com/kedacore/keda/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"knative.dev/eventing/pkg/apis/eventing"
	eventingv1 "knative.dev/eventing/pkg/apis/eventing/v1"
	"knative.dev/pkg/kmeta"
)

func MakeTriggerAuthentication(b *eventingv1.Broker, secretName, secretKey string) *kedav1alpha1.TriggerAuthentication {
	return &kedav1alpha1.TriggerAuthentication{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: b.Namespace,
			Name:      fmt.Sprintf("%s-trigger-auth", b.Name),
			OwnerReferences: []metav1.OwnerReference{
				*kmeta.NewControllerRef(b),
			},
			Labels: map[string]string{
				eventing.BrokerLabelKey:           b.Name,
				"eventing.knative.dev/brokerRole": "keda-scaler",
			},
		},
		Spec: kedav1alpha1.TriggerAuthenticationSpec{
			SecretTargetRef: []kedav1alpha1.AuthSecretTargetRef{
				{
					Parameter: "host",
					Name:      secretName,
					Key:       secretKey,
				},
			},
			Env: []kedav1alpha1.AuthEnvironment{},
			// HACK. Without this doesn't work.
			HashiCorpVault: kedav1alpha1.HashiCorpVault{Secrets: []kedav1alpha1.VaultSecret{}},
		},
	}
}
