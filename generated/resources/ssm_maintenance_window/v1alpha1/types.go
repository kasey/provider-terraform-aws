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

// SsmMaintenanceWindow is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type SsmMaintenanceWindow struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SsmMaintenanceWindowSpec   `json:"spec"`
	Status SsmMaintenanceWindowStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SsmMaintenanceWindow contains a list of SsmMaintenanceWindowList
type SsmMaintenanceWindowList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SsmMaintenanceWindow `json:"items"`
}

// A SsmMaintenanceWindowSpec defines the desired state of a SsmMaintenanceWindow
type SsmMaintenanceWindowSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  SsmMaintenanceWindowParameters `json:",inline"`
}

// A SsmMaintenanceWindowParameters defines the desired state of a SsmMaintenanceWindow
type SsmMaintenanceWindowParameters struct {
	EndDate                  string `json:"end_date"`
	Name                     string `json:"name"`
	ScheduleTimezone         string `json:"schedule_timezone"`
	AllowUnassociatedTargets bool   `json:"allow_unassociated_targets"`
	Description              string `json:"description"`
	Enabled                  bool   `json:"enabled"`
	StartDate                string `json:"start_date"`
	Cutoff                   int    `json:"cutoff"`
	Duration                 int    `json:"duration"`
	Schedule                 string `json:"schedule"`
}

// A SsmMaintenanceWindowStatus defines the observed state of a SsmMaintenanceWindow
type SsmMaintenanceWindowStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     SsmMaintenanceWindowObservation `json:",inline"`
}

// A SsmMaintenanceWindowObservation records the observed state of a SsmMaintenanceWindow
type SsmMaintenanceWindowObservation struct {
	Id string `json:"id"`
}