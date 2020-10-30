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

// S3BucketObject is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type S3BucketObject struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   S3BucketObjectSpec   `json:"spec"`
	Status S3BucketObjectStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// S3BucketObject contains a list of S3BucketObjectList
type S3BucketObjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []S3BucketObject `json:"items"`
}

// A S3BucketObjectSpec defines the desired state of a S3BucketObject
type S3BucketObjectSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  S3BucketObjectParameters `json:",inline"`
}

// A S3BucketObjectParameters defines the desired state of a S3BucketObject
type S3BucketObjectParameters struct {
	ContentLanguage           string `json:"content_language"`
	KmsKeyId                  string `json:"kms_key_id"`
	ObjectLockMode            string `json:"object_lock_mode"`
	WebsiteRedirect           string `json:"website_redirect"`
	Acl                       string `json:"acl"`
	Bucket                    string `json:"bucket"`
	ContentDisposition        string `json:"content_disposition"`
	ForceDestroy              bool   `json:"force_destroy"`
	Key                       string `json:"key"`
	CacheControl              string `json:"cache_control"`
	ContentEncoding           string `json:"content_encoding"`
	ObjectLockLegalHoldStatus string `json:"object_lock_legal_hold_status"`
	Content                   string `json:"content"`
	ContentBase64             string `json:"content_base64"`
	ObjectLockRetainUntilDate string `json:"object_lock_retain_until_date"`
	Source                    string `json:"source"`
}

// A S3BucketObjectStatus defines the observed state of a S3BucketObject
type S3BucketObjectStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     S3BucketObjectObservation `json:",inline"`
}

// A S3BucketObjectObservation records the observed state of a S3BucketObject
type S3BucketObjectObservation struct {
	ContentType          string `json:"content_type"`
	Etag                 string `json:"etag"`
	Id                   string `json:"id"`
	ServerSideEncryption string `json:"server_side_encryption"`
	StorageClass         string `json:"storage_class"`
	VersionId            string `json:"version_id"`
}