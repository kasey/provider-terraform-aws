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

// OpsworksRdsDbInstance is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type OpsworksRdsDbInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OpsworksRdsDbInstanceSpec   `json:"spec"`
	Status OpsworksRdsDbInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OpsworksRdsDbInstance contains a list of OpsworksRdsDbInstanceList
type OpsworksRdsDbInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OpsworksRdsDbInstance `json:"items"`
}

// A OpsworksRdsDbInstanceSpec defines the desired state of a OpsworksRdsDbInstance
type OpsworksRdsDbInstanceSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  OpsworksRdsDbInstanceParameters `json:",inline"`
}

// A OpsworksRdsDbInstanceParameters defines the desired state of a OpsworksRdsDbInstance
type OpsworksRdsDbInstanceParameters struct {
	StackId          string `json:"stack_id"`
	DbPassword       string `json:"db_password"`
	DbUser           string `json:"db_user"`
	Id               string `json:"id"`
	RdsDbInstanceArn string `json:"rds_db_instance_arn"`
}

// A OpsworksRdsDbInstanceStatus defines the observed state of a OpsworksRdsDbInstance
type OpsworksRdsDbInstanceStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     OpsworksRdsDbInstanceObservation `json:",inline"`
}

// A OpsworksRdsDbInstanceObservation records the observed state of a OpsworksRdsDbInstance
type OpsworksRdsDbInstanceObservation struct{}