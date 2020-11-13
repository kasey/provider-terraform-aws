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

// MqConfiguration is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type MqConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MqConfigurationSpec   `json:"spec"`
	Status MqConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MqConfiguration contains a list of MqConfigurationList
type MqConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MqConfiguration `json:"items"`
}

// A MqConfigurationSpec defines the desired state of a MqConfiguration
type MqConfigurationSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  MqConfigurationParameters `json:",inline"`
}

// A MqConfigurationParameters defines the desired state of a MqConfiguration
type MqConfigurationParameters struct {
	EngineVersion string            `json:"engine_version"`
	Id            string            `json:"id"`
	Data          string            `json:"data"`
	Description   string            `json:"description"`
	EngineType    string            `json:"engine_type"`
	Name          string            `json:"name"`
	Tags          map[string]string `json:"tags"`
}

// A MqConfigurationStatus defines the observed state of a MqConfiguration
type MqConfigurationStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     MqConfigurationObservation `json:",inline"`
}

// A MqConfigurationObservation records the observed state of a MqConfiguration
type MqConfigurationObservation struct {
	LatestRevision int    `json:"latest_revision"`
	Arn            string `json:"arn"`
}