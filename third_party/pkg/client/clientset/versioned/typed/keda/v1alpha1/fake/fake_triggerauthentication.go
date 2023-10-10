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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "knative.dev/eventing-autoscaler-keda/third_party/pkg/apis/keda/v1alpha1"
)

// FakeTriggerAuthentications implements TriggerAuthenticationInterface
type FakeTriggerAuthentications struct {
	Fake *FakeKedaV1alpha1
	ns   string
}

var triggerauthenticationsResource = v1alpha1.SchemeGroupVersion.WithResource("triggerauthentications")

var triggerauthenticationsKind = v1alpha1.SchemeGroupVersion.WithKind("TriggerAuthentication")

// Get takes name of the triggerAuthentication, and returns the corresponding triggerAuthentication object, and an error if there is any.
func (c *FakeTriggerAuthentications) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.TriggerAuthentication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(triggerauthenticationsResource, c.ns, name), &v1alpha1.TriggerAuthentication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TriggerAuthentication), err
}

// List takes label and field selectors, and returns the list of TriggerAuthentications that match those selectors.
func (c *FakeTriggerAuthentications) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.TriggerAuthenticationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(triggerauthenticationsResource, triggerauthenticationsKind, c.ns, opts), &v1alpha1.TriggerAuthenticationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.TriggerAuthenticationList{ListMeta: obj.(*v1alpha1.TriggerAuthenticationList).ListMeta}
	for _, item := range obj.(*v1alpha1.TriggerAuthenticationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested triggerAuthentications.
func (c *FakeTriggerAuthentications) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(triggerauthenticationsResource, c.ns, opts))

}

// Create takes the representation of a triggerAuthentication and creates it.  Returns the server's representation of the triggerAuthentication, and an error, if there is any.
func (c *FakeTriggerAuthentications) Create(ctx context.Context, triggerAuthentication *v1alpha1.TriggerAuthentication, opts v1.CreateOptions) (result *v1alpha1.TriggerAuthentication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(triggerauthenticationsResource, c.ns, triggerAuthentication), &v1alpha1.TriggerAuthentication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TriggerAuthentication), err
}

// Update takes the representation of a triggerAuthentication and updates it. Returns the server's representation of the triggerAuthentication, and an error, if there is any.
func (c *FakeTriggerAuthentications) Update(ctx context.Context, triggerAuthentication *v1alpha1.TriggerAuthentication, opts v1.UpdateOptions) (result *v1alpha1.TriggerAuthentication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(triggerauthenticationsResource, c.ns, triggerAuthentication), &v1alpha1.TriggerAuthentication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TriggerAuthentication), err
}

// Delete takes name of the triggerAuthentication and deletes it. Returns an error if one occurs.
func (c *FakeTriggerAuthentications) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(triggerauthenticationsResource, c.ns, name, opts), &v1alpha1.TriggerAuthentication{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTriggerAuthentications) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(triggerauthenticationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.TriggerAuthenticationList{})
	return err
}

// Patch applies the patch and returns the patched triggerAuthentication.
func (c *FakeTriggerAuthentications) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.TriggerAuthentication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(triggerauthenticationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.TriggerAuthentication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TriggerAuthentication), err
}
