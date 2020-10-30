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

// SesEmailIdentity is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type SesEmailIdentity struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SesEmailIdentitySpec   `json:"spec"`
	Status SesEmailIdentityStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SesEmailIdentity contains a list of SesEmailIdentityList
type SesEmailIdentityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SesEmailIdentity `json:"items"`
}

// A SesEmailIdentitySpec defines the desired state of a SesEmailIdentity
type SesEmailIdentitySpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  SesEmailIdentityParameters `json:",inline"`
}

// A SesEmailIdentityParameters defines the desired state of a SesEmailIdentity
type SesEmailIdentityParameters struct {
	Email string `json:"email"`
}

// A SesEmailIdentityStatus defines the observed state of a SesEmailIdentity
type SesEmailIdentityStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     SesEmailIdentityObservation `json:",inline"`
}

// A SesEmailIdentityObservation records the observed state of a SesEmailIdentity
type SesEmailIdentityObservation struct {
	Arn string `json:"arn"`
	Id  string `json:"id"`
}