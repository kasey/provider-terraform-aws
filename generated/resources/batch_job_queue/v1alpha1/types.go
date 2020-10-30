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

// BatchJobQueue is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type BatchJobQueue struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BatchJobQueueSpec   `json:"spec"`
	Status BatchJobQueueStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BatchJobQueue contains a list of BatchJobQueueList
type BatchJobQueueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BatchJobQueue `json:"items"`
}

// A BatchJobQueueSpec defines the desired state of a BatchJobQueue
type BatchJobQueueSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  BatchJobQueueParameters `json:",inline"`
}

// A BatchJobQueueParameters defines the desired state of a BatchJobQueue
type BatchJobQueueParameters struct {
	ComputeEnvironments []string `json:"compute_environments"`
	Name                string   `json:"name"`
	Priority            int      `json:"priority"`
	State               string   `json:"state"`
}

// A BatchJobQueueStatus defines the observed state of a BatchJobQueue
type BatchJobQueueStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     BatchJobQueueObservation `json:",inline"`
}

// A BatchJobQueueObservation records the observed state of a BatchJobQueue
type BatchJobQueueObservation struct {
	Id  string `json:"id"`
	Arn string `json:"arn"`
}