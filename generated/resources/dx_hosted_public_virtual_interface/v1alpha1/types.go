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

// DxHostedPublicVirtualInterface is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DxHostedPublicVirtualInterface struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DxHostedPublicVirtualInterfaceSpec   `json:"spec"`
	Status DxHostedPublicVirtualInterfaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DxHostedPublicVirtualInterface contains a list of DxHostedPublicVirtualInterfaceList
type DxHostedPublicVirtualInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DxHostedPublicVirtualInterface `json:"items"`
}

// A DxHostedPublicVirtualInterfaceSpec defines the desired state of a DxHostedPublicVirtualInterface
type DxHostedPublicVirtualInterfaceSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  DxHostedPublicVirtualInterfaceParameters `json:"forProvider"`
}

// A DxHostedPublicVirtualInterfaceParameters defines the desired state of a DxHostedPublicVirtualInterface
type DxHostedPublicVirtualInterfaceParameters struct {
	BgpAuthKey          string   `json:"bgp_auth_key"`
	Name                string   `json:"name"`
	OwnerAccountId      string   `json:"owner_account_id"`
	RouteFilterPrefixes []string `json:"route_filter_prefixes"`
	AmazonAddress       string   `json:"amazon_address"`
	Vlan                int64    `json:"vlan"`
	CustomerAddress     string   `json:"customer_address"`
	BgpAsn              int64    `json:"bgp_asn"`
	ConnectionId        string   `json:"connection_id"`
	AddressFamily       string   `json:"address_family"`
	Timeouts            Timeouts `json:"timeouts"`
}

type Timeouts struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
}

// A DxHostedPublicVirtualInterfaceStatus defines the observed state of a DxHostedPublicVirtualInterface
type DxHostedPublicVirtualInterfaceStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DxHostedPublicVirtualInterfaceObservation `json:"atProvider"`
}

// A DxHostedPublicVirtualInterfaceObservation records the observed state of a DxHostedPublicVirtualInterface
type DxHostedPublicVirtualInterfaceObservation struct {
	Arn           string `json:"arn"`
	AwsDevice     string `json:"aws_device"`
	AmazonSideAsn string `json:"amazon_side_asn"`
}