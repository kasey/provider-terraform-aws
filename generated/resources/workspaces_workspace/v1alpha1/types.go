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

// WorkspacesWorkspace is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type WorkspacesWorkspace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WorkspacesWorkspaceSpec   `json:"spec"`
	Status WorkspacesWorkspaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WorkspacesWorkspace contains a list of WorkspacesWorkspaceList
type WorkspacesWorkspaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WorkspacesWorkspace `json:"items"`
}

// A WorkspacesWorkspaceSpec defines the desired state of a WorkspacesWorkspace
type WorkspacesWorkspaceSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  WorkspacesWorkspaceParameters `json:",inline"`
}

// A WorkspacesWorkspaceParameters defines the desired state of a WorkspacesWorkspace
type WorkspacesWorkspaceParameters struct {
	BundleId                    string `json:"bundle_id"`
	UserName                    string `json:"user_name"`
	UserVolumeEncryptionEnabled bool   `json:"user_volume_encryption_enabled"`
	VolumeEncryptionKey         string `json:"volume_encryption_key"`
	DirectoryId                 string `json:"directory_id"`
	RootVolumeEncryptionEnabled bool   `json:"root_volume_encryption_enabled"`
}

// A WorkspacesWorkspaceStatus defines the observed state of a WorkspacesWorkspace
type WorkspacesWorkspaceStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     WorkspacesWorkspaceObservation `json:",inline"`
}

// A WorkspacesWorkspaceObservation records the observed state of a WorkspacesWorkspace
type WorkspacesWorkspaceObservation struct {
	Id           string `json:"id"`
	State        string `json:"state"`
	ComputerName string `json:"computer_name"`
	IpAddress    string `json:"ip_address"`
}