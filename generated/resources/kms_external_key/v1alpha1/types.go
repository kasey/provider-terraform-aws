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

// KmsExternalKey is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type KmsExternalKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KmsExternalKeySpec   `json:"spec"`
	Status KmsExternalKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KmsExternalKey contains a list of KmsExternalKeyList
type KmsExternalKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KmsExternalKey `json:"items"`
}

// A KmsExternalKeySpec defines the desired state of a KmsExternalKey
type KmsExternalKeySpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  KmsExternalKeyParameters `json:",inline"`
}

// A KmsExternalKeyParameters defines the desired state of a KmsExternalKey
type KmsExternalKeyParameters struct {
	DeletionWindowInDays int    `json:"deletion_window_in_days"`
	Description          string `json:"description"`
	KeyMaterialBase64    string `json:"key_material_base64"`
	ValidTo              string `json:"valid_to"`
}

// A KmsExternalKeyStatus defines the observed state of a KmsExternalKey
type KmsExternalKeyStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     KmsExternalKeyObservation `json:",inline"`
}

// A KmsExternalKeyObservation records the observed state of a KmsExternalKey
type KmsExternalKeyObservation struct {
	Enabled         bool   `json:"enabled"`
	ExpirationModel string `json:"expiration_model"`
	Id              string `json:"id"`
	KeyState        string `json:"key_state"`
	Arn             string `json:"arn"`
	Policy          string `json:"policy"`
	KeyUsage        string `json:"key_usage"`
}