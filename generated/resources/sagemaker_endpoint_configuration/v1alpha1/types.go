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

// SagemakerEndpointConfiguration is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type SagemakerEndpointConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SagemakerEndpointConfigurationSpec   `json:"spec"`
	Status SagemakerEndpointConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SagemakerEndpointConfiguration contains a list of SagemakerEndpointConfigurationList
type SagemakerEndpointConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SagemakerEndpointConfiguration `json:"items"`
}

// A SagemakerEndpointConfigurationSpec defines the desired state of a SagemakerEndpointConfiguration
type SagemakerEndpointConfigurationSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  SagemakerEndpointConfigurationParameters `json:",inline"`
}

// A SagemakerEndpointConfigurationParameters defines the desired state of a SagemakerEndpointConfiguration
type SagemakerEndpointConfigurationParameters struct {
	KmsKeyArn string `json:"kms_key_arn"`
}

// A SagemakerEndpointConfigurationStatus defines the observed state of a SagemakerEndpointConfiguration
type SagemakerEndpointConfigurationStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     SagemakerEndpointConfigurationObservation `json:",inline"`
}

// A SagemakerEndpointConfigurationObservation records the observed state of a SagemakerEndpointConfiguration
type SagemakerEndpointConfigurationObservation struct {
	Name string `json:"name"`
	Arn  string `json:"arn"`
	Id   string `json:"id"`
}