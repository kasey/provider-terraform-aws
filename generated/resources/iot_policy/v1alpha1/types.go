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

// IotPolicy is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type IotPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IotPolicySpec   `json:"spec"`
	Status IotPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IotPolicy contains a list of IotPolicyList
type IotPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IotPolicy `json:"items"`
}

// A IotPolicySpec defines the desired state of a IotPolicy
type IotPolicySpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  IotPolicyParameters `json:"forProvider"`
}

// A IotPolicyParameters defines the desired state of a IotPolicy
type IotPolicyParameters struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Policy string `json:"policy"`
}

// A IotPolicyStatus defines the observed state of a IotPolicy
type IotPolicyStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     IotPolicyObservation `json:"atProvider"`
}

// A IotPolicyObservation records the observed state of a IotPolicy
type IotPolicyObservation struct {
	DefaultVersionId string `json:"default_version_id"`
	Arn              string `json:"arn"`
}