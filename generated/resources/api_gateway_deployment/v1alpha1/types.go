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

// ApiGatewayDeployment is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ApiGatewayDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApiGatewayDeploymentSpec   `json:"spec"`
	Status ApiGatewayDeploymentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApiGatewayDeployment contains a list of ApiGatewayDeploymentList
type ApiGatewayDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiGatewayDeployment `json:"items"`
}

// A ApiGatewayDeploymentSpec defines the desired state of a ApiGatewayDeployment
type ApiGatewayDeploymentSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  ApiGatewayDeploymentParameters `json:"forProvider"`
}

// A ApiGatewayDeploymentParameters defines the desired state of a ApiGatewayDeployment
type ApiGatewayDeploymentParameters struct {
	RestApiId        string            `json:"rest_api_id"`
	StageDescription string            `json:"stage_description"`
	StageName        string            `json:"stage_name"`
	Triggers         map[string]string `json:"triggers"`
	Description      string            `json:"description"`
	Variables        map[string]string `json:"variables"`
}

// A ApiGatewayDeploymentStatus defines the observed state of a ApiGatewayDeployment
type ApiGatewayDeploymentStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     ApiGatewayDeploymentObservation `json:"atProvider"`
}

// A ApiGatewayDeploymentObservation records the observed state of a ApiGatewayDeployment
type ApiGatewayDeploymentObservation struct {
	CreatedDate  string `json:"created_date"`
	ExecutionArn string `json:"execution_arn"`
	InvokeUrl    string `json:"invoke_url"`
}