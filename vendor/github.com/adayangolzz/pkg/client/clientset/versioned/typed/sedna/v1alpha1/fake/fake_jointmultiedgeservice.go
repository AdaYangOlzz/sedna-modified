/*
Copyright The KubeEdge Authors.

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

	// v1alpha1 "github.com/adayangolzz/sedna-modified/pkg/apis/sedna/v1alpha1"
	// TODO: v1alpha1路径确认 (DONE.)
	v1alpha1 "github.com/adayangolzz/sedna-modified/pkg/apis/sedna/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeJointInferenceServices implements JointInferenceServiceInterface
type FakeJointMultiEdgeServices struct {
	Fake *FakeSednaV1alpha1
	ns   string
}

var jointmultiedgeservicesResource = schema.GroupVersionResource{Group: "sedna.io", Version: "v1alpha1", Resource: "jointmultiedgeservices"}

var jointmultiedgeservicesKind = schema.GroupVersionKind{Group: "sedna.io", Version: "v1alpha1", Kind: "JointMultiEdgeService"}

// Get takes name of the jointInferenceService, and returns the corresponding jointInferenceService object, and an error if there is any.
func (c *FakeJointMultiEdgeServices) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.JointMultiEdgeService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(jointmultiedgeservicesResource, c.ns, name), &v1alpha1.JointMultiEdgeService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.JointMultiEdgeService), err
}

// List takes label and field selectors, and returns the list of JointInferenceServices that match those selectors.
func (c *FakeJointMultiEdgeServices) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.JointMultiEdgeServiceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(jointmultiedgeservicesResource, jointmultiedgeservicesKind, c.ns, opts), &v1alpha1.JointMultiEdgeServiceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.JointMultiEdgeServiceList{ListMeta: obj.(*v1alpha1.JointMultiEdgeServiceList).ListMeta}
	for _, item := range obj.(*v1alpha1.JointMultiEdgeServiceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested jointInferenceServices.
func (c *FakeJointMultiEdgeServices) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(jointmultiedgeservicesResource, c.ns, opts))

}

// Create takes the representation of a jointInferenceService and creates it.  Returns the server's representation of the jointInferenceService, and an error, if there is any.
func (c *FakeJointMultiEdgeServices) Create(ctx context.Context, JointMultiEdgeService *v1alpha1.JointMultiEdgeService, opts v1.CreateOptions) (result *v1alpha1.JointMultiEdgeService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(jointmultiedgeservicesResource, c.ns, JointMultiEdgeService), &v1alpha1.JointMultiEdgeService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.JointMultiEdgeService), err
}

// Update takes the representation of a jointInferenceService and updates it. Returns the server's representation of the jointInferenceService, and an error, if there is any.
func (c *FakeJointMultiEdgeServices) Update(ctx context.Context, JointMultiEdgeService *v1alpha1.JointMultiEdgeService, opts v1.UpdateOptions) (result *v1alpha1.JointMultiEdgeService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(jointmultiedgeservicesResource, c.ns, JointMultiEdgeService), &v1alpha1.JointMultiEdgeService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.JointMultiEdgeService), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeJointMultiEdgeServices) UpdateStatus(ctx context.Context, JointMultiEdgeService *v1alpha1.JointMultiEdgeService, opts v1.UpdateOptions) (*v1alpha1.JointMultiEdgeService, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(jointmultiedgeservicesResource, "status", c.ns, JointMultiEdgeService), &v1alpha1.JointMultiEdgeService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.JointMultiEdgeService), err
}

// Delete takes name of the jointInferenceService and deletes it. Returns an error if one occurs.
func (c *FakeJointMultiEdgeServices) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(jointmultiedgeservicesResource, c.ns, name), &v1alpha1.JointMultiEdgeService{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeJointMultiEdgeServices) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(jointmultiedgeservicesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.JointMultiEdgeServiceList{})
	return err
}

// Patch applies the patch and returns the patched jointInferenceService.
func (c *FakeJointMultiEdgeServices) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.JointMultiEdgeService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(jointmultiedgeservicesResource, c.ns, name, pt, data, subresources...), &v1alpha1.JointMultiEdgeService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.JointMultiEdgeService), err
}