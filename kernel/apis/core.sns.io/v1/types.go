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

// Destination is a type which represents route destination.
type Destination struct {
	// +optional
	Cluster string `json:"cluster,omitempty" protobuf:"bytes,1,opt,name=cluster"`

	// +optional
	Percent float64 `json:"percent,omitempty" protobuf:"bytes,2,opt,name=percent"`
}

// RouteScript is a type which represents route script.
type RouteScript struct {
	// +optional
	Enable bool `json:"enable,omitempty" protobuf:"bytes,1,opt,name=enable"`

	// +optional
	Content string `json:"content,omitempty" protobuf:"bytes,2,opt,name=content"`
}

// ClusterSpec is a specification of a cluster.
type ClusterSpec struct {
	// cluster kind
	// +optional
	Kind string `json:"kind,omitempty" protobuf:"bytes,1,opt,name=kind"`

	// A map is used to store cluster's labels.
	// +optional
	// +patchStrategy=merge,retainKeys
	Labels map[string]string `json:"labels,omitempty" patchStrategy:"merge,retainKeys" protobuf:"bytes,2,rep,name=labels"`

	// An endpoint list of the cluster.
	// +optional
	// +patchMergeKey=addr
	// +patchStrategy=merge,retainKeys
	// +listType=map
	// +listMapKey=addr
	Endpoints []Endpoint `json:"endpoints,omitempty" patchStrategy:"merge,retainKeys" patchMergeKey:"addr" protobuf:"bytes,3,rep,name=endpoints"`
}

// RoutePolicySpec is a specification of a route policy.
type RoutePolicySpec struct {
	// route script
	// +optional
	Script RouteScript `json:"script,omitempty" protobuf:"bytes,1,opt,name=script"`

	// a default destination list
	// +optional
	DefaultDestinations []Destination `json:"defaultDestinations,omitempty" protobuf:"bytes,2,rep,name=defaultDestinations"`
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

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type SNSRoutePolicy struct {
	metav1.TypeMeta `json:",inline"`

	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// A specification of a route policy.
	// +optional
	Spec RoutePolicySpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type SNSRoutePolicyList struct {
	metav1.TypeMeta `json:",inline"`

	// Standard list metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	// +optional
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// List of route policies.
	Items []SNSRoutePolicy `json:"items" protobuf:"bytes,2,rep,name=items"`
}
