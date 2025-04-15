/*
 * Copyright 2021 The Knative Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	context "context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
	messagingv1beta1 "knative.dev/eventing-kafka-broker/control-plane/pkg/apis/messaging/v1beta1"
	scheme "knative.dev/eventing-kafka-broker/control-plane/pkg/client/clientset/versioned/scheme"
)

// KafkaChannelsGetter has a method to return a KafkaChannelInterface.
// A group's client should implement this interface.
type KafkaChannelsGetter interface {
	KafkaChannels(namespace string) KafkaChannelInterface
}

// KafkaChannelInterface has methods to work with KafkaChannel resources.
type KafkaChannelInterface interface {
	Create(ctx context.Context, kafkaChannel *messagingv1beta1.KafkaChannel, opts v1.CreateOptions) (*messagingv1beta1.KafkaChannel, error)
	Update(ctx context.Context, kafkaChannel *messagingv1beta1.KafkaChannel, opts v1.UpdateOptions) (*messagingv1beta1.KafkaChannel, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, kafkaChannel *messagingv1beta1.KafkaChannel, opts v1.UpdateOptions) (*messagingv1beta1.KafkaChannel, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*messagingv1beta1.KafkaChannel, error)
	List(ctx context.Context, opts v1.ListOptions) (*messagingv1beta1.KafkaChannelList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *messagingv1beta1.KafkaChannel, err error)
	KafkaChannelExpansion
}

// kafkaChannels implements KafkaChannelInterface
type kafkaChannels struct {
	*gentype.ClientWithList[*messagingv1beta1.KafkaChannel, *messagingv1beta1.KafkaChannelList]
}

// newKafkaChannels returns a KafkaChannels
func newKafkaChannels(c *MessagingV1beta1Client, namespace string) *kafkaChannels {
	return &kafkaChannels{
		gentype.NewClientWithList[*messagingv1beta1.KafkaChannel, *messagingv1beta1.KafkaChannelList](
			"kafkachannels",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *messagingv1beta1.KafkaChannel { return &messagingv1beta1.KafkaChannel{} },
			func() *messagingv1beta1.KafkaChannelList { return &messagingv1beta1.KafkaChannelList{} },
		),
	}
}
