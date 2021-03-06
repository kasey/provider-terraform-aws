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

// RedshiftEventSubscription is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type RedshiftEventSubscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RedshiftEventSubscriptionSpec   `json:"spec"`
	Status RedshiftEventSubscriptionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RedshiftEventSubscription contains a list of RedshiftEventSubscriptionList
type RedshiftEventSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedshiftEventSubscription `json:"items"`
}

// A RedshiftEventSubscriptionSpec defines the desired state of a RedshiftEventSubscription
type RedshiftEventSubscriptionSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  RedshiftEventSubscriptionParameters `json:"forProvider"`
}

// A RedshiftEventSubscriptionParameters defines the desired state of a RedshiftEventSubscription
type RedshiftEventSubscriptionParameters struct {
	Enabled         bool              `json:"enabled"`
	SnsTopicArn     string            `json:"sns_topic_arn"`
	SourceIds       []string          `json:"source_ids"`
	Tags            map[string]string `json:"tags"`
	Severity        string            `json:"severity"`
	SourceType      string            `json:"source_type"`
	EventCategories []string          `json:"event_categories"`
	Name            string            `json:"name"`
	Timeouts        Timeouts          `json:"timeouts"`
}

type Timeouts struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

// A RedshiftEventSubscriptionStatus defines the observed state of a RedshiftEventSubscription
type RedshiftEventSubscriptionStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     RedshiftEventSubscriptionObservation `json:"atProvider"`
}

// A RedshiftEventSubscriptionObservation records the observed state of a RedshiftEventSubscription
type RedshiftEventSubscriptionObservation struct {
	Status        string `json:"status"`
	Arn           string `json:"arn"`
	CustomerAwsId string `json:"customer_aws_id"`
}