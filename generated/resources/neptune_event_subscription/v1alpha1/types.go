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

// NeptuneEventSubscription is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type NeptuneEventSubscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NeptuneEventSubscriptionSpec   `json:"spec"`
	Status NeptuneEventSubscriptionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NeptuneEventSubscription contains a list of NeptuneEventSubscriptionList
type NeptuneEventSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NeptuneEventSubscription `json:"items"`
}

// A NeptuneEventSubscriptionSpec defines the desired state of a NeptuneEventSubscription
type NeptuneEventSubscriptionSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  NeptuneEventSubscriptionParameters `json:"forProvider"`
}

// A NeptuneEventSubscriptionParameters defines the desired state of a NeptuneEventSubscription
type NeptuneEventSubscriptionParameters struct {
	SnsTopicArn     string            `json:"sns_topic_arn"`
	SourceType      string            `json:"source_type"`
	Name            string            `json:"name"`
	NamePrefix      string            `json:"name_prefix"`
	SourceIds       []string          `json:"source_ids"`
	Tags            map[string]string `json:"tags"`
	Enabled         bool              `json:"enabled"`
	EventCategories []string          `json:"event_categories"`
	Timeouts        Timeouts          `json:"timeouts"`
}

type Timeouts struct {
	Delete string `json:"delete"`
	Update string `json:"update"`
	Create string `json:"create"`
}

// A NeptuneEventSubscriptionStatus defines the observed state of a NeptuneEventSubscription
type NeptuneEventSubscriptionStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     NeptuneEventSubscriptionObservation `json:"atProvider"`
}

// A NeptuneEventSubscriptionObservation records the observed state of a NeptuneEventSubscription
type NeptuneEventSubscriptionObservation struct {
	CustomerAwsId string `json:"customer_aws_id"`
	Arn           string `json:"arn"`
}