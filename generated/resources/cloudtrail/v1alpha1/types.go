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

// Cloudtrail is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Cloudtrail struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudtrailSpec   `json:"spec"`
	Status CloudtrailStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Cloudtrail contains a list of CloudtrailList
type CloudtrailList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cloudtrail `json:"items"`
}

// A CloudtrailSpec defines the desired state of a Cloudtrail
type CloudtrailSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  CloudtrailParameters `json:",inline"`
}

// A CloudtrailParameters defines the desired state of a Cloudtrail
type CloudtrailParameters struct {
	EnableLogFileValidation    bool   `json:"enable_log_file_validation"`
	IsOrganizationTrail        bool   `json:"is_organization_trail"`
	CloudWatchLogsRoleArn      string `json:"cloud_watch_logs_role_arn"`
	IsMultiRegionTrail         bool   `json:"is_multi_region_trail"`
	SnsTopicName               string `json:"sns_topic_name"`
	Name                       string `json:"name"`
	S3KeyPrefix                string `json:"s3_key_prefix"`
	S3BucketName               string `json:"s3_bucket_name"`
	CloudWatchLogsGroupArn     string `json:"cloud_watch_logs_group_arn"`
	EnableLogging              bool   `json:"enable_logging"`
	IncludeGlobalServiceEvents bool   `json:"include_global_service_events"`
	KmsKeyId                   string `json:"kms_key_id"`
}

// A CloudtrailStatus defines the observed state of a Cloudtrail
type CloudtrailStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     CloudtrailObservation `json:",inline"`
}

// A CloudtrailObservation records the observed state of a Cloudtrail
type CloudtrailObservation struct {
	Arn        string `json:"arn"`
	HomeRegion string `json:"home_region"`
	Id         string `json:"id"`
}