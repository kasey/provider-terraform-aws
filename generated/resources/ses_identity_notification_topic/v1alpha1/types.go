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

// SesIdentityNotificationTopic is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type SesIdentityNotificationTopic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SesIdentityNotificationTopicSpec   `json:"spec"`
	Status SesIdentityNotificationTopicStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SesIdentityNotificationTopic contains a list of SesIdentityNotificationTopicList
type SesIdentityNotificationTopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SesIdentityNotificationTopic `json:"items"`
}

// A SesIdentityNotificationTopicSpec defines the desired state of a SesIdentityNotificationTopic
type SesIdentityNotificationTopicSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  SesIdentityNotificationTopicParameters `json:"forProvider"`
}

// A SesIdentityNotificationTopicParameters defines the desired state of a SesIdentityNotificationTopic
type SesIdentityNotificationTopicParameters struct {
	Identity               string `json:"identity"`
	IncludeOriginalHeaders bool   `json:"include_original_headers"`
	NotificationType       string `json:"notification_type"`
	TopicArn               string `json:"topic_arn"`
}

// A SesIdentityNotificationTopicStatus defines the observed state of a SesIdentityNotificationTopic
type SesIdentityNotificationTopicStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     SesIdentityNotificationTopicObservation `json:"atProvider"`
}

// A SesIdentityNotificationTopicObservation records the observed state of a SesIdentityNotificationTopic
type SesIdentityNotificationTopicObservation struct{}