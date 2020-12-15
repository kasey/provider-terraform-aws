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

// Wafv2IpSet is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Wafv2IpSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Wafv2IpSetSpec   `json:"spec"`
	Status Wafv2IpSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Wafv2IpSet contains a list of Wafv2IpSetList
type Wafv2IpSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Wafv2IpSet `json:"items"`
}

// A Wafv2IpSetSpec defines the desired state of a Wafv2IpSet
type Wafv2IpSetSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  Wafv2IpSetParameters `json:"forProvider"`
}

// A Wafv2IpSetParameters defines the desired state of a Wafv2IpSet
type Wafv2IpSetParameters struct {
	Addresses        []string          `json:"addresses"`
	Description      string            `json:"description"`
	IpAddressVersion string            `json:"ip_address_version"`
	Scope            string            `json:"scope"`
	Name             string            `json:"name"`
	Tags             map[string]string `json:"tags"`
}

// A Wafv2IpSetStatus defines the observed state of a Wafv2IpSet
type Wafv2IpSetStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     Wafv2IpSetObservation `json:"atProvider"`
}

// A Wafv2IpSetObservation records the observed state of a Wafv2IpSet
type Wafv2IpSetObservation struct {
	Arn       string `json:"arn"`
	LockToken string `json:"lock_token"`
}