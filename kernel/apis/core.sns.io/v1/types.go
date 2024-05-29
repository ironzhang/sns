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

package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// State is a type which represents the endpoint's state.
type State string

// The State const values defined here.
const (
	Enabled  State = "enabled"
	Disabled State = "disabled"
)

// Endpoint is a type which represents endpoint.
type Endpoint struct {
	// +optional
	Addr string `json:"addr,omitempty" protobuf:"bytes,1,opt,name=addr"`

	// +optional
	State State `json:"state,omitempty" protobuf:"bytes,2,opt,name=state"`

	// +optional
	Weight int `json:"weight,omitempty" protobuf:"varint,3,opt,name=weight"`

	// +optional
	// +patchStrategy=merge,retainKeys
	Tags map[string]string `json:"tags,omitempty" patchStrategy:"merge,retainKeys" protobuf:"bytes,4,rep,name=tags"`
}

// ClusterSpec is a specification of a cluster.
type ClusterSpec struct {
	// A map is used to store cluster's tags.
	// +optional
	// +patchStrategy=merge,retainKeys
	Tags map[string]string `json:"tags,omitempty" patchStrategy:"merge,retainKeys" protobuf:"bytes,1,rep,name=tags"`

	// An endpoint list of the cluster.
	// +optional
	// +patchMergeKey=addr
	// +patchStrategy=merge,retainKeys
	// +listType=map
	// +listMapKey=addr
	Endpoints []Endpoint `json:"endpoints,omitempty" patchStrategy:"merge,retainKeys" patchMergeKey:"addr" protobuf:"bytes,2,rep,name=endpoints"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SNSCluster is a top-level type which represents cluster.
type SNSCluster struct {
	metav1.TypeMeta `json:",inline"`

	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// A specification of a cluster.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	// +optional
	Spec ClusterSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SNSClusterList is a top-level list of clusters.
type SNSClusterList struct {
	metav1.TypeMeta `json:",inline"`

	// Standard list metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	// +optional
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// List of clusters.
	Items []SNSCluster `json:"items" protobuf:"bytes,2,rep,name=items"`
}
