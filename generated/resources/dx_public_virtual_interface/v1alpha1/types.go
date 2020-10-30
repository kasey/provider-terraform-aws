/*
	Copyright 2019 The Crossplane Authors.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	runtimev1alpha1 "github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
)

// +kubebuilder:object:root=true

// DxPublicVirtualInterface is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DxPublicVirtualInterface struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DxPublicVirtualInterfaceSpec   `json:"spec"`
	Status DxPublicVirtualInterfaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DxPublicVirtualInterface contains a list of DxPublicVirtualInterfaceList
type DxPublicVirtualInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DxPublicVirtualInterface `json:"items"`
}

// A DxPublicVirtualInterfaceSpec defines the desired state of a DxPublicVirtualInterface
type DxPublicVirtualInterfaceSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  DxPublicVirtualInterfaceParameters `json:",inline"`
}

// A DxPublicVirtualInterfaceParameters defines the desired state of a DxPublicVirtualInterface
type DxPublicVirtualInterfaceParameters struct {
	Name                string   `json:"name"`
	RouteFilterPrefixes []string `json:"route_filter_prefixes"`
	AddressFamily       string   `json:"address_family"`
	BgpAsn              int      `json:"bgp_asn"`
	ConnectionId        string   `json:"connection_id"`
	Vlan                int      `json:"vlan"`
}

// A DxPublicVirtualInterfaceStatus defines the observed state of a DxPublicVirtualInterface
type DxPublicVirtualInterfaceStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DxPublicVirtualInterfaceObservation `json:",inline"`
}

// A DxPublicVirtualInterfaceObservation records the observed state of a DxPublicVirtualInterface
type DxPublicVirtualInterfaceObservation struct {
	AmazonAddress   string `json:"amazon_address"`
	AmazonSideAsn   string `json:"amazon_side_asn"`
	Arn             string `json:"arn"`
	AwsDevice       string `json:"aws_device"`
	BgpAuthKey      string `json:"bgp_auth_key"`
	CustomerAddress string `json:"customer_address"`
	Id              string `json:"id"`
}