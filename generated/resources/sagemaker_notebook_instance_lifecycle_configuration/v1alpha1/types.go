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

// SagemakerNotebookInstanceLifecycleConfiguration is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type SagemakerNotebookInstanceLifecycleConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SagemakerNotebookInstanceLifecycleConfigurationSpec   `json:"spec"`
	Status SagemakerNotebookInstanceLifecycleConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SagemakerNotebookInstanceLifecycleConfiguration contains a list of SagemakerNotebookInstanceLifecycleConfigurationList
type SagemakerNotebookInstanceLifecycleConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SagemakerNotebookInstanceLifecycleConfiguration `json:"items"`
}

// A SagemakerNotebookInstanceLifecycleConfigurationSpec defines the desired state of a SagemakerNotebookInstanceLifecycleConfiguration
type SagemakerNotebookInstanceLifecycleConfigurationSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  SagemakerNotebookInstanceLifecycleConfigurationParameters `json:",inline"`
}

// A SagemakerNotebookInstanceLifecycleConfigurationParameters defines the desired state of a SagemakerNotebookInstanceLifecycleConfiguration
type SagemakerNotebookInstanceLifecycleConfigurationParameters struct {
	OnCreate string `json:"on_create"`
	OnStart  string `json:"on_start"`
	Name     string `json:"name"`
}

// A SagemakerNotebookInstanceLifecycleConfigurationStatus defines the observed state of a SagemakerNotebookInstanceLifecycleConfiguration
type SagemakerNotebookInstanceLifecycleConfigurationStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     SagemakerNotebookInstanceLifecycleConfigurationObservation `json:",inline"`
}

// A SagemakerNotebookInstanceLifecycleConfigurationObservation records the observed state of a SagemakerNotebookInstanceLifecycleConfiguration
type SagemakerNotebookInstanceLifecycleConfigurationObservation struct {
	Arn string `json:"arn"`
	Id  string `json:"id"`
}