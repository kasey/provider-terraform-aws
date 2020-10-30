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

// AppsyncGraphqlApi is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type AppsyncGraphqlApi struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AppsyncGraphqlApiSpec   `json:"spec"`
	Status AppsyncGraphqlApiStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppsyncGraphqlApi contains a list of AppsyncGraphqlApiList
type AppsyncGraphqlApiList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppsyncGraphqlApi `json:"items"`
}

// A AppsyncGraphqlApiSpec defines the desired state of a AppsyncGraphqlApi
type AppsyncGraphqlApiSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  AppsyncGraphqlApiParameters `json:",inline"`
}

// A AppsyncGraphqlApiParameters defines the desired state of a AppsyncGraphqlApi
type AppsyncGraphqlApiParameters struct {
	AuthenticationType string `json:"authentication_type"`
	Name               string `json:"name"`
	Schema             string `json:"schema"`
	XrayEnabled        bool   `json:"xray_enabled"`
}

// A AppsyncGraphqlApiStatus defines the observed state of a AppsyncGraphqlApi
type AppsyncGraphqlApiStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     AppsyncGraphqlApiObservation `json:",inline"`
}

// A AppsyncGraphqlApiObservation records the observed state of a AppsyncGraphqlApi
type AppsyncGraphqlApiObservation struct {
	Id  string `json:"id"`
	Arn string `json:"arn"`
}