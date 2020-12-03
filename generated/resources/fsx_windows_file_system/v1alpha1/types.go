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

// FsxWindowsFileSystem is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type FsxWindowsFileSystem struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FsxWindowsFileSystemSpec   `json:"spec"`
	Status FsxWindowsFileSystemStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FsxWindowsFileSystem contains a list of FsxWindowsFileSystemList
type FsxWindowsFileSystemList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FsxWindowsFileSystem `json:"items"`
}

// A FsxWindowsFileSystemSpec defines the desired state of a FsxWindowsFileSystem
type FsxWindowsFileSystemSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  FsxWindowsFileSystemParameters `json:",inline"`
}

// A FsxWindowsFileSystemParameters defines the desired state of a FsxWindowsFileSystem
type FsxWindowsFileSystemParameters struct {
	DailyAutomaticBackupStartTime string                     `json:"daily_automatic_backup_start_time"`
	AutomaticBackupRetentionDays  int                        `json:"automatic_backup_retention_days"`
	Id                            string                     `json:"id"`
	SkipFinalBackup               bool                       `json:"skip_final_backup"`
	Tags                          map[string]string          `json:"tags"`
	PreferredSubnetId             string                     `json:"preferred_subnet_id"`
	SecurityGroupIds              []string                   `json:"security_group_ids"`
	StorageCapacity               int                        `json:"storage_capacity"`
	StorageType                   string                     `json:"storage_type"`
	SubnetIds                     []string                   `json:"subnet_ids"`
	ActiveDirectoryId             string                     `json:"active_directory_id"`
	CopyTagsToBackups             bool                       `json:"copy_tags_to_backups"`
	DeploymentType                string                     `json:"deployment_type"`
	KmsKeyId                      string                     `json:"kms_key_id"`
	ThroughputCapacity            int                        `json:"throughput_capacity"`
	WeeklyMaintenanceStartTime    string                     `json:"weekly_maintenance_start_time"`
	SelfManagedActiveDirectory    SelfManagedActiveDirectory `json:"self_managed_active_directory"`
	Timeouts                      []Timeouts                 `json:"timeouts"`
}

type SelfManagedActiveDirectory struct {
	DnsIps                              []string `json:"dns_ips"`
	DomainName                          string   `json:"domain_name"`
	FileSystemAdministratorsGroup       string   `json:"file_system_administrators_group"`
	OrganizationalUnitDistinguishedName string   `json:"organizational_unit_distinguished_name"`
	Password                            string   `json:"password"`
	Username                            string   `json:"username"`
}

type Timeouts struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
}

// A FsxWindowsFileSystemStatus defines the observed state of a FsxWindowsFileSystem
type FsxWindowsFileSystemStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     FsxWindowsFileSystemObservation `json:",inline"`
}

// A FsxWindowsFileSystemObservation records the observed state of a FsxWindowsFileSystem
type FsxWindowsFileSystemObservation struct {
	Arn                          string   `json:"arn"`
	DnsName                      string   `json:"dns_name"`
	PreferredFileServerIp        string   `json:"preferred_file_server_ip"`
	VpcId                        string   `json:"vpc_id"`
	NetworkInterfaceIds          []string `json:"network_interface_ids"`
	RemoteAdministrationEndpoint string   `json:"remote_administration_endpoint"`
	OwnerId                      string   `json:"owner_id"`
}