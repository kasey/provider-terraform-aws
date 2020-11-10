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

// RdsClusterInstance is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type RdsClusterInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RdsClusterInstanceSpec   `json:"spec"`
	Status RdsClusterInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RdsClusterInstance contains a list of RdsClusterInstanceList
type RdsClusterInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RdsClusterInstance `json:"items"`
}

// A RdsClusterInstanceSpec defines the desired state of a RdsClusterInstance
type RdsClusterInstanceSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  RdsClusterInstanceParameters `json:",inline"`
}

// A RdsClusterInstanceParameters defines the desired state of a RdsClusterInstance
type RdsClusterInstanceParameters struct {
	Engine                  string            `json:"engine"`
	Tags                    map[string]string `json:"tags"`
	ClusterIdentifier       string            `json:"cluster_identifier"`
	CopyTagsToSnapshot      bool              `json:"copy_tags_to_snapshot"`
	MonitoringInterval      int               `json:"monitoring_interval"`
	AutoMinorVersionUpgrade bool              `json:"auto_minor_version_upgrade"`
	InstanceClass           string            `json:"instance_class"`
	PromotionTier           int               `json:"promotion_tier"`
	PubliclyAccessible      bool              `json:"publicly_accessible"`
	Timeouts                []Timeouts        `json:"timeouts"`
}

type Timeouts struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

// A RdsClusterInstanceStatus defines the observed state of a RdsClusterInstance
type RdsClusterInstanceStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     RdsClusterInstanceObservation `json:",inline"`
}

// A RdsClusterInstanceObservation records the observed state of a RdsClusterInstance
type RdsClusterInstanceObservation struct {
	AvailabilityZone            string `json:"availability_zone"`
	CaCertIdentifier            string `json:"ca_cert_identifier"`
	Arn                         string `json:"arn"`
	Endpoint                    string `json:"endpoint"`
	PerformanceInsightsEnabled  bool   `json:"performance_insights_enabled"`
	PerformanceInsightsKmsKeyId string `json:"performance_insights_kms_key_id"`
	ApplyImmediately            bool   `json:"apply_immediately"`
	DbSubnetGroupName           string `json:"db_subnet_group_name"`
	IdentifierPrefix            string `json:"identifier_prefix"`
	KmsKeyId                    string `json:"kms_key_id"`
	Port                        int    `json:"port"`
	PreferredMaintenanceWindow  string `json:"preferred_maintenance_window"`
	StorageEncrypted            bool   `json:"storage_encrypted"`
	Id                          string `json:"id"`
	MonitoringRoleArn           string `json:"monitoring_role_arn"`
	DbParameterGroupName        string `json:"db_parameter_group_name"`
	DbiResourceId               string `json:"dbi_resource_id"`
	PreferredBackupWindow       string `json:"preferred_backup_window"`
	EngineVersion               string `json:"engine_version"`
	Identifier                  string `json:"identifier"`
	Writer                      bool   `json:"writer"`
}