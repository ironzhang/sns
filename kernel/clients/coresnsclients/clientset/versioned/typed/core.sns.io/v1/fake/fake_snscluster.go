/*
Copyright The Kubernetes Authors.

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

	coresnsiov1 "github.com/ironzhang/sns/kernel/apis/core.sns.io/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSNSClusters implements SNSClusterInterface
type FakeSNSClusters struct {
	Fake *FakeCoreV1
	ns   string
}

var snsclustersResource = schema.GroupVersionResource{Group: "core.sns.io", Version: "v1", Resource: "snsclusters"}

var snsclustersKind = schema.GroupVersionKind{Group: "core.sns.io", Version: "v1", Kind: "SNSCluster"}

// Get takes name of the sNSCluster, and returns the corresponding sNSCluster object, and an error if there is any.
func (c *FakeSNSClusters) Get(ctx context.Context, name string, options v1.GetOptions) (result *coresnsiov1.SNSCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(snsclustersResource, c.ns, name), &coresnsiov1.SNSCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*coresnsiov1.SNSCluster), err
}

// List takes label and field selectors, and returns the list of SNSClusters that match those selectors.
func (c *FakeSNSClusters) List(ctx context.Context, opts v1.ListOptions) (result *coresnsiov1.SNSClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(snsclustersResource, snsclustersKind, c.ns, opts), &coresnsiov1.SNSClusterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &coresnsiov1.SNSClusterList{ListMeta: obj.(*coresnsiov1.SNSClusterList).ListMeta}
	for _, item := range obj.(*coresnsiov1.SNSClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested sNSClusters.
func (c *FakeSNSClusters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(snsclustersResource, c.ns, opts))

}

// Create takes the representation of a sNSCluster and creates it.  Returns the server's representation of the sNSCluster, and an error, if there is any.
func (c *FakeSNSClusters) Create(ctx context.Context, sNSCluster *coresnsiov1.SNSCluster, opts v1.CreateOptions) (result *coresnsiov1.SNSCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(snsclustersResource, c.ns, sNSCluster), &coresnsiov1.SNSCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*coresnsiov1.SNSCluster), err
}

// Update takes the representation of a sNSCluster and updates it. Returns the server's representation of the sNSCluster, and an error, if there is any.
func (c *FakeSNSClusters) Update(ctx context.Context, sNSCluster *coresnsiov1.SNSCluster, opts v1.UpdateOptions) (result *coresnsiov1.SNSCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(snsclustersResource, c.ns, sNSCluster), &coresnsiov1.SNSCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*coresnsiov1.SNSCluster), err
}

// Delete takes name of the sNSCluster and deletes it. Returns an error if one occurs.
func (c *FakeSNSClusters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(snsclustersResource, c.ns, name), &coresnsiov1.SNSCluster{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSNSClusters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(snsclustersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &coresnsiov1.SNSClusterList{})
	return err
}

// Patch applies the patch and returns the patched sNSCluster.
func (c *FakeSNSClusters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *coresnsiov1.SNSCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(snsclustersResource, c.ns, name, pt, data, subresources...), &coresnsiov1.SNSCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*coresnsiov1.SNSCluster), err
}
