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

// IamGroupPolicy is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type IamGroupPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IamGroupPolicySpec   `json:"spec"`
	Status IamGroupPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IamGroupPolicy contains a list of IamGroupPolicyList
type IamGroupPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IamGroupPolicy `json:"items"`
}

// A IamGroupPolicySpec defines the desired state of a IamGroupPolicy
type IamGroupPolicySpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  IamGroupPolicyParameters `json:",inline"`
}

// A IamGroupPolicyParameters defines the desired state of a IamGroupPolicy
type IamGroupPolicyParameters struct {
	Group      string `json:"group"`
	NamePrefix string `json:"name_prefix"`
	Policy     string `json:"policy"`
}

// A IamGroupPolicyStatus defines the observed state of a IamGroupPolicy
type IamGroupPolicyStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     IamGroupPolicyObservation `json:",inline"`
}

// A IamGroupPolicyObservation records the observed state of a IamGroupPolicy
type IamGroupPolicyObservation struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}