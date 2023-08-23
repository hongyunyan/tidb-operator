// Copyright PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/pingcap/tidb-operator/pkg/apis/pingcap/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTidbInitializers implements TidbInitializerInterface
type FakeTidbInitializers struct {
	Fake *FakePingcapV1alpha1
	ns   string
}

var tidbinitializersResource = schema.GroupVersionResource{Group: "pingcap.com", Version: "v1alpha1", Resource: "tidbinitializers"}

var tidbinitializersKind = schema.GroupVersionKind{Group: "pingcap.com", Version: "v1alpha1", Kind: "TidbInitializer"}

// Get takes name of the tidbInitializer, and returns the corresponding tidbInitializer object, and an error if there is any.
func (c *FakeTidbInitializers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.TidbInitializer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(tidbinitializersResource, c.ns, name), &v1alpha1.TidbInitializer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TidbInitializer), err
}

// List takes label and field selectors, and returns the list of TidbInitializers that match those selectors.
func (c *FakeTidbInitializers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.TidbInitializerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(tidbinitializersResource, tidbinitializersKind, c.ns, opts), &v1alpha1.TidbInitializerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.TidbInitializerList{ListMeta: obj.(*v1alpha1.TidbInitializerList).ListMeta}
	for _, item := range obj.(*v1alpha1.TidbInitializerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested tidbInitializers.
func (c *FakeTidbInitializers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(tidbinitializersResource, c.ns, opts))

}

// Create takes the representation of a tidbInitializer and creates it.  Returns the server's representation of the tidbInitializer, and an error, if there is any.
func (c *FakeTidbInitializers) Create(ctx context.Context, tidbInitializer *v1alpha1.TidbInitializer, opts v1.CreateOptions) (result *v1alpha1.TidbInitializer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(tidbinitializersResource, c.ns, tidbInitializer), &v1alpha1.TidbInitializer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TidbInitializer), err
}

// Update takes the representation of a tidbInitializer and updates it. Returns the server's representation of the tidbInitializer, and an error, if there is any.
func (c *FakeTidbInitializers) Update(ctx context.Context, tidbInitializer *v1alpha1.TidbInitializer, opts v1.UpdateOptions) (result *v1alpha1.TidbInitializer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(tidbinitializersResource, c.ns, tidbInitializer), &v1alpha1.TidbInitializer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TidbInitializer), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTidbInitializers) UpdateStatus(ctx context.Context, tidbInitializer *v1alpha1.TidbInitializer, opts v1.UpdateOptions) (*v1alpha1.TidbInitializer, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(tidbinitializersResource, "status", c.ns, tidbInitializer), &v1alpha1.TidbInitializer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TidbInitializer), err
}

// Delete takes name of the tidbInitializer and deletes it. Returns an error if one occurs.
func (c *FakeTidbInitializers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(tidbinitializersResource, c.ns, name, opts), &v1alpha1.TidbInitializer{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTidbInitializers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(tidbinitializersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.TidbInitializerList{})
	return err
}

// Patch applies the patch and returns the patched tidbInitializer.
func (c *FakeTidbInitializers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.TidbInitializer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(tidbinitializersResource, c.ns, name, pt, data, subresources...), &v1alpha1.TidbInitializer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TidbInitializer), err
}
