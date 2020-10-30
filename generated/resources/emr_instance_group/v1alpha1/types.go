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

// EmrInstanceGroup is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type EmrInstanceGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EmrInstanceGroupSpec   `json:"spec"`
	Status EmrInstanceGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EmrInstanceGroup contains a list of EmrInstanceGroupList
type EmrInstanceGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EmrInstanceGroup `json:"items"`
}

// A EmrInstanceGroupSpec defines the desired state of a EmrInstanceGroup
type EmrInstanceGroupSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  EmrInstanceGroupParameters `json:",inline"`
}

// A EmrInstanceGroupParameters defines the desired state of a EmrInstanceGroup
type EmrInstanceGroupParameters struct {
	BidPrice           string `json:"bid_price"`
	EbsOptimized       bool   `json:"ebs_optimized"`
	InstanceCount      int    `json:"instance_count"`
	Name               string `json:"name"`
	AutoscalingPolicy  string `json:"autoscaling_policy"`
	ClusterId          string `json:"cluster_id"`
	ConfigurationsJson string `json:"configurations_json"`
	InstanceType       string `json:"instance_type"`
}

// A EmrInstanceGroupStatus defines the observed state of a EmrInstanceGroup
type EmrInstanceGroupStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     EmrInstanceGroupObservation `json:",inline"`
}

// A EmrInstanceGroupObservation records the observed state of a EmrInstanceGroup
type EmrInstanceGroupObservation struct {
	RunningInstanceCount int    `json:"running_instance_count"`
	Status               string `json:"status"`
	Id                   string `json:"id"`
}