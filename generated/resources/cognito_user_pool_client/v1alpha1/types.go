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

// CognitoUserPoolClient is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type CognitoUserPoolClient struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CognitoUserPoolClientSpec   `json:"spec"`
	Status CognitoUserPoolClientStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CognitoUserPoolClient contains a list of CognitoUserPoolClientList
type CognitoUserPoolClientList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CognitoUserPoolClient `json:"items"`
}

// A CognitoUserPoolClientSpec defines the desired state of a CognitoUserPoolClient
type CognitoUserPoolClientSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  CognitoUserPoolClientParameters `json:",inline"`
}

// A CognitoUserPoolClientParameters defines the desired state of a CognitoUserPoolClient
type CognitoUserPoolClientParameters struct {
	ExplicitAuthFlows               []string               `json:"explicit_auth_flows"`
	GenerateSecret                  bool                   `json:"generate_secret"`
	ReadAttributes                  []string               `json:"read_attributes"`
	SupportedIdentityProviders      []string               `json:"supported_identity_providers"`
	AllowedOauthFlows               []string               `json:"allowed_oauth_flows"`
	AllowedOauthFlowsUserPoolClient bool                   `json:"allowed_oauth_flows_user_pool_client"`
	LogoutUrls                      []string               `json:"logout_urls"`
	UserPoolId                      string                 `json:"user_pool_id"`
	DefaultRedirectUri              string                 `json:"default_redirect_uri"`
	Id                              string                 `json:"id"`
	Name                            string                 `json:"name"`
	PreventUserExistenceErrors      string                 `json:"prevent_user_existence_errors"`
	RefreshTokenValidity            int                    `json:"refresh_token_validity"`
	AllowedOauthScopes              []string               `json:"allowed_oauth_scopes"`
	CallbackUrls                    []string               `json:"callback_urls"`
	WriteAttributes                 []string               `json:"write_attributes"`
	AnalyticsConfiguration          AnalyticsConfiguration `json:"analytics_configuration"`
}

type AnalyticsConfiguration struct {
	ExternalId     string `json:"external_id"`
	RoleArn        string `json:"role_arn"`
	UserDataShared bool   `json:"user_data_shared"`
	ApplicationId  string `json:"application_id"`
}

// A CognitoUserPoolClientStatus defines the observed state of a CognitoUserPoolClient
type CognitoUserPoolClientStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     CognitoUserPoolClientObservation `json:",inline"`
}

// A CognitoUserPoolClientObservation records the observed state of a CognitoUserPoolClient
type CognitoUserPoolClientObservation struct {
	ClientSecret string `json:"client_secret"`
}