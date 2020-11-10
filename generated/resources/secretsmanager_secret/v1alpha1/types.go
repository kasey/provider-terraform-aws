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

// SecretsmanagerSecret is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type SecretsmanagerSecret struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SecretsmanagerSecretSpec   `json:"spec"`
	Status SecretsmanagerSecretStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecretsmanagerSecret contains a list of SecretsmanagerSecretList
type SecretsmanagerSecretList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecretsmanagerSecret `json:"items"`
}

// A SecretsmanagerSecretSpec defines the desired state of a SecretsmanagerSecret
type SecretsmanagerSecretSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  SecretsmanagerSecretParameters `json:",inline"`
}

// A SecretsmanagerSecretParameters defines the desired state of a SecretsmanagerSecret
type SecretsmanagerSecretParameters struct {
	RecoveryWindowInDays int               `json:"recovery_window_in_days"`
	Tags                 map[string]string `json:"tags"`
	KmsKeyId             string            `json:"kms_key_id"`
	Policy               string            `json:"policy"`
	Description          string            `json:"description"`
	RotationRules        RotationRules     `json:"rotation_rules"`
}

type RotationRules struct {
	AutomaticallyAfterDays int `json:"automatically_after_days"`
}

// A SecretsmanagerSecretStatus defines the observed state of a SecretsmanagerSecret
type SecretsmanagerSecretStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     SecretsmanagerSecretObservation `json:",inline"`
}

// A SecretsmanagerSecretObservation records the observed state of a SecretsmanagerSecret
type SecretsmanagerSecretObservation struct {
	RotationLambdaArn string `json:"rotation_lambda_arn"`
	Id                string `json:"id"`
	NamePrefix        string `json:"name_prefix"`
	Name              string `json:"name"`
	RotationEnabled   bool   `json:"rotation_enabled"`
	Arn               string `json:"arn"`
}