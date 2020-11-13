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

// ElasticacheReplicationGroup is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ElasticacheReplicationGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ElasticacheReplicationGroupSpec   `json:"spec"`
	Status ElasticacheReplicationGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ElasticacheReplicationGroup contains a list of ElasticacheReplicationGroupList
type ElasticacheReplicationGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ElasticacheReplicationGroup `json:"items"`
}

// A ElasticacheReplicationGroupSpec defines the desired state of a ElasticacheReplicationGroup
type ElasticacheReplicationGroupSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  ElasticacheReplicationGroupParameters `json:",inline"`
}

// A ElasticacheReplicationGroupParameters defines the desired state of a ElasticacheReplicationGroup
type ElasticacheReplicationGroupParameters struct {
	MaintenanceWindow           string            `json:"maintenance_window"`
	NumberCacheClusters         int               `json:"number_cache_clusters"`
	Port                        int               `json:"port"`
	SecurityGroupNames          []string          `json:"security_group_names"`
	AtRestEncryptionEnabled     bool              `json:"at_rest_encryption_enabled"`
	AutoMinorVersionUpgrade     bool              `json:"auto_minor_version_upgrade"`
	NodeType                    string            `json:"node_type"`
	ReplicationGroupId          string            `json:"replication_group_id"`
	ParameterGroupName          string            `json:"parameter_group_name"`
	SnapshotName                string            `json:"snapshot_name"`
	SnapshotWindow              string            `json:"snapshot_window"`
	ApplyImmediately            bool              `json:"apply_immediately"`
	Id                          string            `json:"id"`
	Tags                        map[string]string `json:"tags"`
	KmsKeyId                    string            `json:"kms_key_id"`
	SecurityGroupIds            []string          `json:"security_group_ids"`
	NotificationTopicArn        string            `json:"notification_topic_arn"`
	AuthToken                   string            `json:"auth_token"`
	EngineVersion               string            `json:"engine_version"`
	SubnetGroupName             string            `json:"subnet_group_name"`
	AutomaticFailoverEnabled    bool              `json:"automatic_failover_enabled"`
	AvailabilityZones           []string          `json:"availability_zones"`
	Engine                      string            `json:"engine"`
	ReplicationGroupDescription string            `json:"replication_group_description"`
	SnapshotArns                []string          `json:"snapshot_arns"`
	SnapshotRetentionLimit      int               `json:"snapshot_retention_limit"`
	TransitEncryptionEnabled    bool              `json:"transit_encryption_enabled"`
	Timeouts                    []Timeouts        `json:"timeouts"`
	ClusterMode                 ClusterMode       `json:"cluster_mode"`
}

type Timeouts struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

type ClusterMode struct {
	ReplicasPerNodeGroup int `json:"replicas_per_node_group"`
	NumNodeGroups        int `json:"num_node_groups"`
}

// A ElasticacheReplicationGroupStatus defines the observed state of a ElasticacheReplicationGroup
type ElasticacheReplicationGroupStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     ElasticacheReplicationGroupObservation `json:",inline"`
}

// A ElasticacheReplicationGroupObservation records the observed state of a ElasticacheReplicationGroup
type ElasticacheReplicationGroupObservation struct {
	MemberClusters               []string `json:"member_clusters"`
	ConfigurationEndpointAddress string   `json:"configuration_endpoint_address"`
	PrimaryEndpointAddress       string   `json:"primary_endpoint_address"`
}