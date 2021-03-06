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

// LambdaPermission is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type LambdaPermission struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LambdaPermissionSpec   `json:"spec"`
	Status LambdaPermissionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LambdaPermission contains a list of LambdaPermissionList
type LambdaPermissionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LambdaPermission `json:"items"`
}

// A LambdaPermissionSpec defines the desired state of a LambdaPermission
type LambdaPermissionSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  LambdaPermissionParameters `json:"forProvider"`
}

// A LambdaPermissionParameters defines the desired state of a LambdaPermission
type LambdaPermissionParameters struct {
	EventSourceToken  string `json:"event_source_token"`
	Principal         string `json:"principal"`
	StatementId       string `json:"statement_id"`
	Action            string `json:"action"`
	FunctionName      string `json:"function_name"`
	Qualifier         string `json:"qualifier"`
	SourceAccount     string `json:"source_account"`
	SourceArn         string `json:"source_arn"`
	StatementIdPrefix string `json:"statement_id_prefix"`
}

// A LambdaPermissionStatus defines the observed state of a LambdaPermission
type LambdaPermissionStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     LambdaPermissionObservation `json:"atProvider"`
}

// A LambdaPermissionObservation records the observed state of a LambdaPermission
type LambdaPermissionObservation struct{}