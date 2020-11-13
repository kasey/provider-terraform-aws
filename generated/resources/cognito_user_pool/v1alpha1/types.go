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

// CognitoUserPool is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type CognitoUserPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CognitoUserPoolSpec   `json:"spec"`
	Status CognitoUserPoolStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CognitoUserPool contains a list of CognitoUserPoolList
type CognitoUserPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CognitoUserPool `json:"items"`
}

// A CognitoUserPoolSpec defines the desired state of a CognitoUserPool
type CognitoUserPoolSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  CognitoUserPoolParameters `json:",inline"`
}

// A CognitoUserPoolParameters defines the desired state of a CognitoUserPool
type CognitoUserPoolParameters struct {
	MfaConfiguration              string                        `json:"mfa_configuration"`
	Name                          string                        `json:"name"`
	AutoVerifiedAttributes        []string                      `json:"auto_verified_attributes"`
	EmailVerificationMessage      string                        `json:"email_verification_message"`
	EmailVerificationSubject      string                        `json:"email_verification_subject"`
	Id                            string                        `json:"id"`
	SmsAuthenticationMessage      string                        `json:"sms_authentication_message"`
	Tags                          map[string]string             `json:"tags"`
	UsernameAttributes            []string                      `json:"username_attributes"`
	AliasAttributes               []string                      `json:"alias_attributes"`
	SmsVerificationMessage        string                        `json:"sms_verification_message"`
	UserPoolAddOns                UserPoolAddOns                `json:"user_pool_add_ons"`
	UsernameConfiguration         UsernameConfiguration         `json:"username_configuration"`
	VerificationMessageTemplate   VerificationMessageTemplate   `json:"verification_message_template"`
	EmailConfiguration            EmailConfiguration            `json:"email_configuration"`
	PasswordPolicy                PasswordPolicy                `json:"password_policy"`
	LambdaConfig                  LambdaConfig                  `json:"lambda_config"`
	Schema                        []Schema                      `json:"schema"`
	SmsConfiguration              SmsConfiguration              `json:"sms_configuration"`
	SoftwareTokenMfaConfiguration SoftwareTokenMfaConfiguration `json:"software_token_mfa_configuration"`
	AdminCreateUserConfig         AdminCreateUserConfig         `json:"admin_create_user_config"`
	DeviceConfiguration           DeviceConfiguration           `json:"device_configuration"`
}

type UserPoolAddOns struct {
	AdvancedSecurityMode string `json:"advanced_security_mode"`
}

type UsernameConfiguration struct {
	CaseSensitive bool `json:"case_sensitive"`
}

type VerificationMessageTemplate struct {
	EmailMessage       string `json:"email_message"`
	EmailMessageByLink string `json:"email_message_by_link"`
	EmailSubject       string `json:"email_subject"`
	EmailSubjectByLink string `json:"email_subject_by_link"`
	SmsMessage         string `json:"sms_message"`
	DefaultEmailOption string `json:"default_email_option"`
}

type EmailConfiguration struct {
	SourceArn           string `json:"source_arn"`
	EmailSendingAccount string `json:"email_sending_account"`
	FromEmailAddress    string `json:"from_email_address"`
	ReplyToEmailAddress string `json:"reply_to_email_address"`
}

type PasswordPolicy struct {
	MinimumLength                 int  `json:"minimum_length"`
	RequireLowercase              bool `json:"require_lowercase"`
	RequireNumbers                bool `json:"require_numbers"`
	RequireSymbols                bool `json:"require_symbols"`
	RequireUppercase              bool `json:"require_uppercase"`
	TemporaryPasswordValidityDays int  `json:"temporary_password_validity_days"`
}

type LambdaConfig struct {
	CustomMessage               string `json:"custom_message"`
	PostAuthentication          string `json:"post_authentication"`
	PostConfirmation            string `json:"post_confirmation"`
	PreAuthentication           string `json:"pre_authentication"`
	PreSignUp                   string `json:"pre_sign_up"`
	CreateAuthChallenge         string `json:"create_auth_challenge"`
	DefineAuthChallenge         string `json:"define_auth_challenge"`
	PreTokenGeneration          string `json:"pre_token_generation"`
	UserMigration               string `json:"user_migration"`
	VerifyAuthChallengeResponse string `json:"verify_auth_challenge_response"`
}

type Schema struct {
	AttributeDataType          string                     `json:"attribute_data_type"`
	DeveloperOnlyAttribute     bool                       `json:"developer_only_attribute"`
	Mutable                    bool                       `json:"mutable"`
	Name                       string                     `json:"name"`
	Required                   bool                       `json:"required"`
	NumberAttributeConstraints NumberAttributeConstraints `json:"number_attribute_constraints"`
	StringAttributeConstraints StringAttributeConstraints `json:"string_attribute_constraints"`
}

type NumberAttributeConstraints struct {
	MaxValue string `json:"max_value"`
	MinValue string `json:"min_value"`
}

type StringAttributeConstraints struct {
	MinLength string `json:"min_length"`
	MaxLength string `json:"max_length"`
}

type SmsConfiguration struct {
	SnsCallerArn string `json:"sns_caller_arn"`
	ExternalId   string `json:"external_id"`
}

type SoftwareTokenMfaConfiguration struct {
	Enabled bool `json:"enabled"`
}

type AdminCreateUserConfig struct {
	AllowAdminCreateUserOnly bool                  `json:"allow_admin_create_user_only"`
	InviteMessageTemplate    InviteMessageTemplate `json:"invite_message_template"`
}

type InviteMessageTemplate struct {
	EmailSubject string `json:"email_subject"`
	SmsMessage   string `json:"sms_message"`
	EmailMessage string `json:"email_message"`
}

type DeviceConfiguration struct {
	ChallengeRequiredOnNewDevice     bool `json:"challenge_required_on_new_device"`
	DeviceOnlyRememberedOnUserPrompt bool `json:"device_only_remembered_on_user_prompt"`
}

// A CognitoUserPoolStatus defines the observed state of a CognitoUserPool
type CognitoUserPoolStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     CognitoUserPoolObservation `json:",inline"`
}

// A CognitoUserPoolObservation records the observed state of a CognitoUserPool
type CognitoUserPoolObservation struct {
	Endpoint         string `json:"endpoint"`
	LastModifiedDate string `json:"last_modified_date"`
	Arn              string `json:"arn"`
	CreationDate     string `json:"creation_date"`
}