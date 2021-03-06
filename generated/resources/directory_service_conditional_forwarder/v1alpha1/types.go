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

// DirectoryServiceConditionalForwarder is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DirectoryServiceConditionalForwarder struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DirectoryServiceConditionalForwarderSpec   `json:"spec"`
	Status DirectoryServiceConditionalForwarderStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DirectoryServiceConditionalForwarder contains a list of DirectoryServiceConditionalForwarderList
type DirectoryServiceConditionalForwarderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DirectoryServiceConditionalForwarder `json:"items"`
}

// A DirectoryServiceConditionalForwarderSpec defines the desired state of a DirectoryServiceConditionalForwarder
type DirectoryServiceConditionalForwarderSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  DirectoryServiceConditionalForwarderParameters `json:"forProvider"`
}

// A DirectoryServiceConditionalForwarderParameters defines the desired state of a DirectoryServiceConditionalForwarder
type DirectoryServiceConditionalForwarderParameters struct {
	RemoteDomainName string   `json:"remote_domain_name"`
	DirectoryId      string   `json:"directory_id"`
	DnsIps           []string `json:"dns_ips"`
}

// A DirectoryServiceConditionalForwarderStatus defines the observed state of a DirectoryServiceConditionalForwarder
type DirectoryServiceConditionalForwarderStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DirectoryServiceConditionalForwarderObservation `json:"atProvider"`
}

// A DirectoryServiceConditionalForwarderObservation records the observed state of a DirectoryServiceConditionalForwarder
type DirectoryServiceConditionalForwarderObservation struct{}