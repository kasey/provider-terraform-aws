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

// SesReceiptRuleSet is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type SesReceiptRuleSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SesReceiptRuleSetSpec   `json:"spec"`
	Status SesReceiptRuleSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SesReceiptRuleSet contains a list of SesReceiptRuleSetList
type SesReceiptRuleSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SesReceiptRuleSet `json:"items"`
}

// A SesReceiptRuleSetSpec defines the desired state of a SesReceiptRuleSet
type SesReceiptRuleSetSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  SesReceiptRuleSetParameters `json:"forProvider"`
}

// A SesReceiptRuleSetParameters defines the desired state of a SesReceiptRuleSet
type SesReceiptRuleSetParameters struct {
	RuleSetName string `json:"rule_set_name"`
}

// A SesReceiptRuleSetStatus defines the observed state of a SesReceiptRuleSet
type SesReceiptRuleSetStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     SesReceiptRuleSetObservation `json:"atProvider"`
}

// A SesReceiptRuleSetObservation records the observed state of a SesReceiptRuleSet
type SesReceiptRuleSetObservation struct{}