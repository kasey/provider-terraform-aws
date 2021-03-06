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

// Eip is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Eip struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EipSpec   `json:"spec"`
	Status EipStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Eip contains a list of EipList
type EipList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Eip `json:"items"`
}

// A EipSpec defines the desired state of a Eip
type EipSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  EipParameters `json:"forProvider"`
}

// A EipParameters defines the desired state of a Eip
type EipParameters struct {
	Instance               string            `json:"instance"`
	CustomerOwnedIpv4Pool  string            `json:"customer_owned_ipv4_pool"`
	Vpc                    bool              `json:"vpc"`
	AssociateWithPrivateIp string            `json:"associate_with_private_ip"`
	NetworkInterface       string            `json:"network_interface"`
	PublicIpv4Pool         string            `json:"public_ipv4_pool"`
	Tags                   map[string]string `json:"tags"`
	Timeouts               Timeouts          `json:"timeouts"`
}

type Timeouts struct {
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}

// A EipStatus defines the observed state of a Eip
type EipStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     EipObservation `json:"atProvider"`
}

// A EipObservation records the observed state of a Eip
type EipObservation struct {
	PrivateIp       string `json:"private_ip"`
	AllocationId    string `json:"allocation_id"`
	CustomerOwnedIp string `json:"customer_owned_ip"`
	PublicDns       string `json:"public_dns"`
	PrivateDns      string `json:"private_dns"`
	AssociationId   string `json:"association_id"`
	Domain          string `json:"domain"`
	PublicIp        string `json:"public_ip"`
}