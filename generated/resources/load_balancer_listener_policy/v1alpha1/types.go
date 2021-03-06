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

// LoadBalancerListenerPolicy is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type LoadBalancerListenerPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LoadBalancerListenerPolicySpec   `json:"spec"`
	Status LoadBalancerListenerPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LoadBalancerListenerPolicy contains a list of LoadBalancerListenerPolicyList
type LoadBalancerListenerPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LoadBalancerListenerPolicy `json:"items"`
}

// A LoadBalancerListenerPolicySpec defines the desired state of a LoadBalancerListenerPolicy
type LoadBalancerListenerPolicySpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  LoadBalancerListenerPolicyParameters `json:"forProvider"`
}

// A LoadBalancerListenerPolicyParameters defines the desired state of a LoadBalancerListenerPolicy
type LoadBalancerListenerPolicyParameters struct {
	LoadBalancerName string   `json:"load_balancer_name"`
	LoadBalancerPort int64    `json:"load_balancer_port"`
	PolicyNames      []string `json:"policy_names"`
}

// A LoadBalancerListenerPolicyStatus defines the observed state of a LoadBalancerListenerPolicy
type LoadBalancerListenerPolicyStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     LoadBalancerListenerPolicyObservation `json:"atProvider"`
}

// A LoadBalancerListenerPolicyObservation records the observed state of a LoadBalancerListenerPolicy
type LoadBalancerListenerPolicyObservation struct{}