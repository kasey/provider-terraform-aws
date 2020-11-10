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

// ApiGatewayDocumentationVersion is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ApiGatewayDocumentationVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApiGatewayDocumentationVersionSpec   `json:"spec"`
	Status ApiGatewayDocumentationVersionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApiGatewayDocumentationVersion contains a list of ApiGatewayDocumentationVersionList
type ApiGatewayDocumentationVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiGatewayDocumentationVersion `json:"items"`
}

// A ApiGatewayDocumentationVersionSpec defines the desired state of a ApiGatewayDocumentationVersion
type ApiGatewayDocumentationVersionSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  ApiGatewayDocumentationVersionParameters `json:",inline"`
}

// A ApiGatewayDocumentationVersionParameters defines the desired state of a ApiGatewayDocumentationVersion
type ApiGatewayDocumentationVersionParameters struct {
	RestApiId   string `json:"rest_api_id"`
	Version     string `json:"version"`
	Description string `json:"description"`
}

// A ApiGatewayDocumentationVersionStatus defines the observed state of a ApiGatewayDocumentationVersion
type ApiGatewayDocumentationVersionStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     ApiGatewayDocumentationVersionObservation `json:",inline"`
}

// A ApiGatewayDocumentationVersionObservation records the observed state of a ApiGatewayDocumentationVersion
type ApiGatewayDocumentationVersionObservation struct {
	Id string `json:"id"`
}