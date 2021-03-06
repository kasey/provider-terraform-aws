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

// DbClusterSnapshot is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DbClusterSnapshot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DbClusterSnapshotSpec   `json:"spec"`
	Status DbClusterSnapshotStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DbClusterSnapshot contains a list of DbClusterSnapshotList
type DbClusterSnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DbClusterSnapshot `json:"items"`
}

// A DbClusterSnapshotSpec defines the desired state of a DbClusterSnapshot
type DbClusterSnapshotSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  DbClusterSnapshotParameters `json:"forProvider"`
}

// A DbClusterSnapshotParameters defines the desired state of a DbClusterSnapshot
type DbClusterSnapshotParameters struct {
	Tags                        map[string]string `json:"tags"`
	DbClusterIdentifier         string            `json:"db_cluster_identifier"`
	DbClusterSnapshotIdentifier string            `json:"db_cluster_snapshot_identifier"`
	Timeouts                    Timeouts          `json:"timeouts"`
}

type Timeouts struct {
	Create string `json:"create"`
}

// A DbClusterSnapshotStatus defines the observed state of a DbClusterSnapshot
type DbClusterSnapshotStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DbClusterSnapshotObservation `json:"atProvider"`
}

// A DbClusterSnapshotObservation records the observed state of a DbClusterSnapshot
type DbClusterSnapshotObservation struct {
	EngineVersion              string   `json:"engine_version"`
	Port                       int64    `json:"port"`
	Status                     string   `json:"status"`
	DbClusterSnapshotArn       string   `json:"db_cluster_snapshot_arn"`
	StorageEncrypted           bool     `json:"storage_encrypted"`
	AllocatedStorage           int64    `json:"allocated_storage"`
	LicenseModel               string   `json:"license_model"`
	SnapshotType               string   `json:"snapshot_type"`
	SourceDbClusterSnapshotArn string   `json:"source_db_cluster_snapshot_arn"`
	VpcId                      string   `json:"vpc_id"`
	KmsKeyId                   string   `json:"kms_key_id"`
	Engine                     string   `json:"engine"`
	AvailabilityZones          []string `json:"availability_zones"`
}