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

// Route53Zone is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Route53Zone struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Route53ZoneSpec   `json:"spec"`
	Status Route53ZoneStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Route53Zone contains a list of Route53ZoneList
type Route53ZoneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Route53Zone `json:"items"`
}

// A Route53ZoneSpec defines the desired state of a Route53Zone
type Route53ZoneSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  Route53ZoneParameters `json:",inline"`
}

// A Route53ZoneParameters defines the desired state of a Route53Zone
type Route53ZoneParameters struct {
	Name            string `json:"name"`
	Comment         string `json:"comment"`
	DelegationSetId string `json:"delegation_set_id"`
	ForceDestroy    bool   `json:"force_destroy"`
}

// A Route53ZoneStatus defines the observed state of a Route53Zone
type Route53ZoneStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     Route53ZoneObservation `json:",inline"`
}

// A Route53ZoneObservation records the observed state of a Route53Zone
type Route53ZoneObservation struct {
	NameServers []string `json:"name_servers"`
	ZoneId      string   `json:"zone_id"`
	Id          string   `json:"id"`
}