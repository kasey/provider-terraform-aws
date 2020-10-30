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

// DmsReplicationInstance is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DmsReplicationInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DmsReplicationInstanceSpec   `json:"spec"`
	Status DmsReplicationInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DmsReplicationInstance contains a list of DmsReplicationInstanceList
type DmsReplicationInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DmsReplicationInstance `json:"items"`
}

// A DmsReplicationInstanceSpec defines the desired state of a DmsReplicationInstance
type DmsReplicationInstanceSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  DmsReplicationInstanceParameters `json:",inline"`
}

// A DmsReplicationInstanceParameters defines the desired state of a DmsReplicationInstance
type DmsReplicationInstanceParameters struct {
	ReplicationInstanceClass string `json:"replication_instance_class"`
	ReplicationInstanceId    string `json:"replication_instance_id"`
	AllowMajorVersionUpgrade bool   `json:"allow_major_version_upgrade"`
	ApplyImmediately         bool   `json:"apply_immediately"`
}

// A DmsReplicationInstanceStatus defines the observed state of a DmsReplicationInstance
type DmsReplicationInstanceStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DmsReplicationInstanceObservation `json:",inline"`
}

// A DmsReplicationInstanceObservation records the observed state of a DmsReplicationInstance
type DmsReplicationInstanceObservation struct {
	VpcSecurityGroupIds           []string `json:"vpc_security_group_ids"`
	AutoMinorVersionUpgrade       bool     `json:"auto_minor_version_upgrade"`
	EngineVersion                 string   `json:"engine_version"`
	KmsKeyArn                     string   `json:"kms_key_arn"`
	PreferredMaintenanceWindow    string   `json:"preferred_maintenance_window"`
	ReplicationInstancePrivateIps []string `json:"replication_instance_private_ips"`
	ReplicationSubnetGroupId      string   `json:"replication_subnet_group_id"`
	MultiAz                       bool     `json:"multi_az"`
	AvailabilityZone              string   `json:"availability_zone"`
	Id                            string   `json:"id"`
	PubliclyAccessible            bool     `json:"publicly_accessible"`
	ReplicationInstanceArn        string   `json:"replication_instance_arn"`
	ReplicationInstancePublicIps  []string `json:"replication_instance_public_ips"`
	AllocatedStorage              int      `json:"allocated_storage"`
}