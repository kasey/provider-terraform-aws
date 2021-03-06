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

// DatapipelinePipeline is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DatapipelinePipeline struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DatapipelinePipelineSpec   `json:"spec"`
	Status DatapipelinePipelineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatapipelinePipeline contains a list of DatapipelinePipelineList
type DatapipelinePipelineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatapipelinePipeline `json:"items"`
}

// A DatapipelinePipelineSpec defines the desired state of a DatapipelinePipeline
type DatapipelinePipelineSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  DatapipelinePipelineParameters `json:"forProvider"`
}

// A DatapipelinePipelineParameters defines the desired state of a DatapipelinePipeline
type DatapipelinePipelineParameters struct {
	Name        string            `json:"name"`
	Tags        map[string]string `json:"tags"`
	Description string            `json:"description"`
}

// A DatapipelinePipelineStatus defines the observed state of a DatapipelinePipeline
type DatapipelinePipelineStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DatapipelinePipelineObservation `json:"atProvider"`
}

// A DatapipelinePipelineObservation records the observed state of a DatapipelinePipeline
type DatapipelinePipelineObservation struct{}