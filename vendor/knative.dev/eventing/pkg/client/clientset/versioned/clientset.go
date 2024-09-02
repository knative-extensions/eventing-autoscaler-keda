/*
Copyright 2021 The Knative Authors

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

// Code generated by client-gen. DO NOT EDIT.

package versioned

import (
	"fmt"
	"net/http"

	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
	eventingv1 "knative.dev/eventing/pkg/client/clientset/versioned/typed/eventing/v1"
	eventingv1alpha1 "knative.dev/eventing/pkg/client/clientset/versioned/typed/eventing/v1alpha1"
	eventingv1beta1 "knative.dev/eventing/pkg/client/clientset/versioned/typed/eventing/v1beta1"
	eventingv1beta2 "knative.dev/eventing/pkg/client/clientset/versioned/typed/eventing/v1beta2"
	eventingv1beta3 "knative.dev/eventing/pkg/client/clientset/versioned/typed/eventing/v1beta3"
	flowsv1 "knative.dev/eventing/pkg/client/clientset/versioned/typed/flows/v1"
	messagingv1 "knative.dev/eventing/pkg/client/clientset/versioned/typed/messaging/v1"
	sinksv1alpha1 "knative.dev/eventing/pkg/client/clientset/versioned/typed/sinks/v1alpha1"
	sourcesv1 "knative.dev/eventing/pkg/client/clientset/versioned/typed/sources/v1"
	sourcesv1beta2 "knative.dev/eventing/pkg/client/clientset/versioned/typed/sources/v1beta2"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	EventingV1() eventingv1.EventingV1Interface
	EventingV1alpha1() eventingv1alpha1.EventingV1alpha1Interface
	EventingV1beta1() eventingv1beta1.EventingV1beta1Interface
	EventingV1beta2() eventingv1beta2.EventingV1beta2Interface
	EventingV1beta3() eventingv1beta3.EventingV1beta3Interface
	FlowsV1() flowsv1.FlowsV1Interface
	MessagingV1() messagingv1.MessagingV1Interface
	SinksV1alpha1() sinksv1alpha1.SinksV1alpha1Interface
	SourcesV1() sourcesv1.SourcesV1Interface
	SourcesV1beta2() sourcesv1beta2.SourcesV1beta2Interface
}

// Clientset contains the clients for groups.
type Clientset struct {
	*discovery.DiscoveryClient
	eventingV1       *eventingv1.EventingV1Client
	eventingV1alpha1 *eventingv1alpha1.EventingV1alpha1Client
	eventingV1beta1  *eventingv1beta1.EventingV1beta1Client
	eventingV1beta2  *eventingv1beta2.EventingV1beta2Client
	eventingV1beta3  *eventingv1beta3.EventingV1beta3Client
	flowsV1          *flowsv1.FlowsV1Client
	messagingV1      *messagingv1.MessagingV1Client
	sinksV1alpha1    *sinksv1alpha1.SinksV1alpha1Client
	sourcesV1        *sourcesv1.SourcesV1Client
	sourcesV1beta2   *sourcesv1beta2.SourcesV1beta2Client
}

// EventingV1 retrieves the EventingV1Client
func (c *Clientset) EventingV1() eventingv1.EventingV1Interface {
	return c.eventingV1
}

// EventingV1alpha1 retrieves the EventingV1alpha1Client
func (c *Clientset) EventingV1alpha1() eventingv1alpha1.EventingV1alpha1Interface {
	return c.eventingV1alpha1
}

// EventingV1beta1 retrieves the EventingV1beta1Client
func (c *Clientset) EventingV1beta1() eventingv1beta1.EventingV1beta1Interface {
	return c.eventingV1beta1
}

// EventingV1beta2 retrieves the EventingV1beta2Client
func (c *Clientset) EventingV1beta2() eventingv1beta2.EventingV1beta2Interface {
	return c.eventingV1beta2
}

// EventingV1beta3 retrieves the EventingV1beta3Client
func (c *Clientset) EventingV1beta3() eventingv1beta3.EventingV1beta3Interface {
	return c.eventingV1beta3
}

// FlowsV1 retrieves the FlowsV1Client
func (c *Clientset) FlowsV1() flowsv1.FlowsV1Interface {
	return c.flowsV1
}

// MessagingV1 retrieves the MessagingV1Client
func (c *Clientset) MessagingV1() messagingv1.MessagingV1Interface {
	return c.messagingV1
}

// SinksV1alpha1 retrieves the SinksV1alpha1Client
func (c *Clientset) SinksV1alpha1() sinksv1alpha1.SinksV1alpha1Interface {
	return c.sinksV1alpha1
}

// SourcesV1 retrieves the SourcesV1Client
func (c *Clientset) SourcesV1() sourcesv1.SourcesV1Interface {
	return c.sourcesV1
}

// SourcesV1beta2 retrieves the SourcesV1beta2Client
func (c *Clientset) SourcesV1beta2() sourcesv1beta2.SourcesV1beta2Interface {
	return c.sourcesV1beta2
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfig will generate a rate-limiter in configShallowCopy.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c

	if configShallowCopy.UserAgent == "" {
		configShallowCopy.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	// share the transport between all clients
	httpClient, err := rest.HTTPClientFor(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	return NewForConfigAndClient(&configShallowCopy, httpClient)
}

// NewForConfigAndClient creates a new Clientset for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfigAndClient will generate a rate-limiter in configShallowCopy.
func NewForConfigAndClient(c *rest.Config, httpClient *http.Client) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}

	var cs Clientset
	var err error
	cs.eventingV1, err = eventingv1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.eventingV1alpha1, err = eventingv1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.eventingV1beta1, err = eventingv1beta1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.eventingV1beta2, err = eventingv1beta2.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.eventingV1beta3, err = eventingv1beta3.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.flowsV1, err = flowsv1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.messagingV1, err = messagingv1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.sinksV1alpha1, err = sinksv1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.sourcesV1, err = sourcesv1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.sourcesV1beta2, err = sourcesv1beta2.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	cs, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.eventingV1 = eventingv1.New(c)
	cs.eventingV1alpha1 = eventingv1alpha1.New(c)
	cs.eventingV1beta1 = eventingv1beta1.New(c)
	cs.eventingV1beta2 = eventingv1beta2.New(c)
	cs.eventingV1beta3 = eventingv1beta3.New(c)
	cs.flowsV1 = flowsv1.New(c)
	cs.messagingV1 = messagingv1.New(c)
	cs.sinksV1alpha1 = sinksv1alpha1.New(c)
	cs.sourcesV1 = sourcesv1.New(c)
	cs.sourcesV1beta2 = sourcesv1beta2.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
