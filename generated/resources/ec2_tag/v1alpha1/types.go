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

// Ec2Tag is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Ec2Tag struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Ec2TagSpec   `json:"spec"`
	Status Ec2TagStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2Tag contains a list of Ec2TagList
type Ec2TagList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Ec2Tag `json:"items"`
}

// A Ec2TagSpec defines the desired state of a Ec2Tag
type Ec2TagSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  Ec2TagParameters `json:",inline"`
}

// A Ec2TagParameters defines the desired state of a Ec2Tag
type Ec2TagParameters struct {
	Value      string `json:"value"`
	Key        string `json:"key"`
	ResourceId string `json:"resource_id"`
}

// A Ec2TagStatus defines the observed state of a Ec2Tag
type Ec2TagStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     Ec2TagObservation `json:",inline"`
}

// A Ec2TagObservation records the observed state of a Ec2Tag
type Ec2TagObservation struct {
	Id string `json:"id"`
}