/*
Copyright (c) 2024 Seldon Technologies Ltd.

Use of this software is governed BY
(1) the license included in the LICENSE file or
(2) if the license included in the LICENSE file is the Business Source License 1.1,
the Change License after the Change Date as each is defined in accordance with the LICENSE file.
*/
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha2 "github.com/seldonio/seldon-core/operator/apis/machinelearning.seldon.io/v1alpha2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSeldonDeployments implements SeldonDeploymentInterface
type FakeSeldonDeployments struct {
	Fake *FakeMachinelearningV1alpha2
	ns   string
}

var seldondeploymentsResource = schema.GroupVersionResource{Group: "machinelearning.seldon.io", Version: "v1alpha2", Resource: "seldondeployments"}

var seldondeploymentsKind = schema.GroupVersionKind{Group: "machinelearning.seldon.io", Version: "v1alpha2", Kind: "SeldonDeployment"}

// Get takes name of the seldonDeployment, and returns the corresponding seldonDeployment object, and an error if there is any.
func (c *FakeSeldonDeployments) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.SeldonDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(seldondeploymentsResource, c.ns, name), &v1alpha2.SeldonDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.SeldonDeployment), err
}

// List takes label and field selectors, and returns the list of SeldonDeployments that match those selectors.
func (c *FakeSeldonDeployments) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.SeldonDeploymentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(seldondeploymentsResource, seldondeploymentsKind, c.ns, opts), &v1alpha2.SeldonDeploymentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha2.SeldonDeploymentList{ListMeta: obj.(*v1alpha2.SeldonDeploymentList).ListMeta}
	for _, item := range obj.(*v1alpha2.SeldonDeploymentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested seldonDeployments.
func (c *FakeSeldonDeployments) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(seldondeploymentsResource, c.ns, opts))

}

// Create takes the representation of a seldonDeployment and creates it.  Returns the server's representation of the seldonDeployment, and an error, if there is any.
func (c *FakeSeldonDeployments) Create(ctx context.Context, seldonDeployment *v1alpha2.SeldonDeployment, opts v1.CreateOptions) (result *v1alpha2.SeldonDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(seldondeploymentsResource, c.ns, seldonDeployment), &v1alpha2.SeldonDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.SeldonDeployment), err
}

// Update takes the representation of a seldonDeployment and updates it. Returns the server's representation of the seldonDeployment, and an error, if there is any.
func (c *FakeSeldonDeployments) Update(ctx context.Context, seldonDeployment *v1alpha2.SeldonDeployment, opts v1.UpdateOptions) (result *v1alpha2.SeldonDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(seldondeploymentsResource, c.ns, seldonDeployment), &v1alpha2.SeldonDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.SeldonDeployment), err
}

// Delete takes name of the seldonDeployment and deletes it. Returns an error if one occurs.
func (c *FakeSeldonDeployments) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(seldondeploymentsResource, c.ns, name), &v1alpha2.SeldonDeployment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSeldonDeployments) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(seldondeploymentsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha2.SeldonDeploymentList{})
	return err
}

// Patch applies the patch and returns the patched seldonDeployment.
func (c *FakeSeldonDeployments) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.SeldonDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(seldondeploymentsResource, c.ns, name, pt, data, subresources...), &v1alpha2.SeldonDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.SeldonDeployment), err
}
