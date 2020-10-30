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

// DocdbCluster is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DocdbCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DocdbClusterSpec   `json:"spec"`
	Status DocdbClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DocdbCluster contains a list of DocdbClusterList
type DocdbClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DocdbCluster `json:"items"`
}

// A DocdbClusterSpec defines the desired state of a DocdbCluster
type DocdbClusterSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  DocdbClusterParameters `json:",inline"`
}

// A DocdbClusterParameters defines the desired state of a DocdbCluster
type DocdbClusterParameters struct {
	Port                         int      `json:"port"`
	FinalSnapshotIdentifier      string   `json:"final_snapshot_identifier"`
	DeletionProtection           bool     `json:"deletion_protection"`
	StorageEncrypted             bool     `json:"storage_encrypted"`
	EnabledCloudwatchLogsExports []string `json:"enabled_cloudwatch_logs_exports"`
	BackupRetentionPeriod        int      `json:"backup_retention_period"`
	Engine                       string   `json:"engine"`
	MasterPassword               string   `json:"master_password"`
	SkipFinalSnapshot            bool     `json:"skip_final_snapshot"`
	SnapshotIdentifier           string   `json:"snapshot_identifier"`
}

// A DocdbClusterStatus defines the observed state of a DocdbCluster
type DocdbClusterStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DocdbClusterObservation `json:",inline"`
}

// A DocdbClusterObservation records the observed state of a DocdbCluster
type DocdbClusterObservation struct {
	PreferredBackupWindow       string   `json:"preferred_backup_window"`
	ClusterMembers              []string `json:"cluster_members"`
	DbClusterParameterGroupName string   `json:"db_cluster_parameter_group_name"`
	Id                          string   `json:"id"`
	ApplyImmediately            bool     `json:"apply_immediately"`
	ClusterResourceId           string   `json:"cluster_resource_id"`
	KmsKeyId                    string   `json:"kms_key_id"`
	MasterUsername              string   `json:"master_username"`
	ClusterIdentifier           string   `json:"cluster_identifier"`
	ClusterIdentifierPrefix     string   `json:"cluster_identifier_prefix"`
	Arn                         string   `json:"arn"`
	Endpoint                    string   `json:"endpoint"`
	EngineVersion               string   `json:"engine_version"`
	AvailabilityZones           []string `json:"availability_zones"`
	DbSubnetGroupName           string   `json:"db_subnet_group_name"`
	ReaderEndpoint              string   `json:"reader_endpoint"`
	VpcSecurityGroupIds         []string `json:"vpc_security_group_ids"`
	HostedZoneId                string   `json:"hosted_zone_id"`
	PreferredMaintenanceWindow  string   `json:"preferred_maintenance_window"`
}