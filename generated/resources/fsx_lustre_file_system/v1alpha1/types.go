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

// FsxLustreFileSystem is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type FsxLustreFileSystem struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FsxLustreFileSystemSpec   `json:"spec"`
	Status FsxLustreFileSystemStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FsxLustreFileSystem contains a list of FsxLustreFileSystemList
type FsxLustreFileSystemList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FsxLustreFileSystem `json:"items"`
}

// A FsxLustreFileSystemSpec defines the desired state of a FsxLustreFileSystem
type FsxLustreFileSystemSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  FsxLustreFileSystemParameters `json:",inline"`
}

// A FsxLustreFileSystemParameters defines the desired state of a FsxLustreFileSystem
type FsxLustreFileSystemParameters struct {
	DeploymentType                string            `json:"deployment_type"`
	DriveCacheType                string            `json:"drive_cache_type"`
	Tags                          map[string]string `json:"tags"`
	StorageType                   string            `json:"storage_type"`
	AutomaticBackupRetentionDays  int               `json:"automatic_backup_retention_days"`
	ExportPath                    string            `json:"export_path"`
	ImportPath                    string            `json:"import_path"`
	KmsKeyId                      string            `json:"kms_key_id"`
	SubnetIds                     []string          `json:"subnet_ids"`
	AutoImportPolicy              string            `json:"auto_import_policy"`
	Id                            string            `json:"id"`
	ImportedFileChunkSize         int               `json:"imported_file_chunk_size"`
	WeeklyMaintenanceStartTime    string            `json:"weekly_maintenance_start_time"`
	DailyAutomaticBackupStartTime string            `json:"daily_automatic_backup_start_time"`
	PerUnitStorageThroughput      int               `json:"per_unit_storage_throughput"`
	SecurityGroupIds              []string          `json:"security_group_ids"`
	StorageCapacity               int               `json:"storage_capacity"`
	Timeouts                      []Timeouts        `json:"timeouts"`
}

type Timeouts struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
}

// A FsxLustreFileSystemStatus defines the observed state of a FsxLustreFileSystem
type FsxLustreFileSystemStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     FsxLustreFileSystemObservation `json:",inline"`
}

// A FsxLustreFileSystemObservation records the observed state of a FsxLustreFileSystem
type FsxLustreFileSystemObservation struct {
	VpcId               string   `json:"vpc_id"`
	NetworkInterfaceIds []string `json:"network_interface_ids"`
	MountName           string   `json:"mount_name"`
	OwnerId             string   `json:"owner_id"`
	Arn                 string   `json:"arn"`
	DnsName             string   `json:"dns_name"`
}