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

// Route53QueryLog is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Route53QueryLog struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Route53QueryLogSpec   `json:"spec"`
	Status Route53QueryLogStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Route53QueryLog contains a list of Route53QueryLogList
type Route53QueryLogList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Route53QueryLog `json:"items"`
}

// A Route53QueryLogSpec defines the desired state of a Route53QueryLog
type Route53QueryLogSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  Route53QueryLogParameters `json:"forProvider"`
}

// A Route53QueryLogParameters defines the desired state of a Route53QueryLog
type Route53QueryLogParameters struct {
	CloudwatchLogGroupArn string `json:"cloudwatch_log_group_arn"`
	ZoneId                string `json:"zone_id"`
}

// A Route53QueryLogStatus defines the observed state of a Route53QueryLog
type Route53QueryLogStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     Route53QueryLogObservation `json:"atProvider"`
}

// A Route53QueryLogObservation records the observed state of a Route53QueryLog
type Route53QueryLogObservation struct{}