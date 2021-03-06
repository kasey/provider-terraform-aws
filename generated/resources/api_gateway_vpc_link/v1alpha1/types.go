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

// ApiGatewayVpcLink is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ApiGatewayVpcLink struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApiGatewayVpcLinkSpec   `json:"spec"`
	Status ApiGatewayVpcLinkStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApiGatewayVpcLink contains a list of ApiGatewayVpcLinkList
type ApiGatewayVpcLinkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiGatewayVpcLink `json:"items"`
}

// A ApiGatewayVpcLinkSpec defines the desired state of a ApiGatewayVpcLink
type ApiGatewayVpcLinkSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  ApiGatewayVpcLinkParameters `json:"forProvider"`
}

// A ApiGatewayVpcLinkParameters defines the desired state of a ApiGatewayVpcLink
type ApiGatewayVpcLinkParameters struct {
	Description string            `json:"description"`
	Name        string            `json:"name"`
	Tags        map[string]string `json:"tags"`
	TargetArns  []string          `json:"target_arns"`
}

// A ApiGatewayVpcLinkStatus defines the observed state of a ApiGatewayVpcLink
type ApiGatewayVpcLinkStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     ApiGatewayVpcLinkObservation `json:"atProvider"`
}

// A ApiGatewayVpcLinkObservation records the observed state of a ApiGatewayVpcLink
type ApiGatewayVpcLinkObservation struct {
	Arn string `json:"arn"`
}