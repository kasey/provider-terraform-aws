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

// OrganizationsAccount is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type OrganizationsAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OrganizationsAccountSpec   `json:"spec"`
	Status OrganizationsAccountStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OrganizationsAccount contains a list of OrganizationsAccountList
type OrganizationsAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OrganizationsAccount `json:"items"`
}

// A OrganizationsAccountSpec defines the desired state of a OrganizationsAccount
type OrganizationsAccountSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  OrganizationsAccountParameters `json:",inline"`
}

// A OrganizationsAccountParameters defines the desired state of a OrganizationsAccount
type OrganizationsAccountParameters struct {
	Email                  string            `json:"email"`
	Name                   string            `json:"name"`
	ParentId               string            `json:"parent_id"`
	IamUserAccessToBilling string            `json:"iam_user_access_to_billing"`
	Id                     string            `json:"id"`
	RoleName               string            `json:"role_name"`
	Tags                   map[string]string `json:"tags"`
}

// A OrganizationsAccountStatus defines the observed state of a OrganizationsAccount
type OrganizationsAccountStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     OrganizationsAccountObservation `json:",inline"`
}

// A OrganizationsAccountObservation records the observed state of a OrganizationsAccount
type OrganizationsAccountObservation struct {
	Status          string `json:"status"`
	JoinedMethod    string `json:"joined_method"`
	JoinedTimestamp string `json:"joined_timestamp"`
	Arn             string `json:"arn"`
}