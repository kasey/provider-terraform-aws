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

// LicensemanagerLicenseConfiguration is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type LicensemanagerLicenseConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LicensemanagerLicenseConfigurationSpec   `json:"spec"`
	Status LicensemanagerLicenseConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LicensemanagerLicenseConfiguration contains a list of LicensemanagerLicenseConfigurationList
type LicensemanagerLicenseConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LicensemanagerLicenseConfiguration `json:"items"`
}

// A LicensemanagerLicenseConfigurationSpec defines the desired state of a LicensemanagerLicenseConfiguration
type LicensemanagerLicenseConfigurationSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  LicensemanagerLicenseConfigurationParameters `json:",inline"`
}

// A LicensemanagerLicenseConfigurationParameters defines the desired state of a LicensemanagerLicenseConfiguration
type LicensemanagerLicenseConfigurationParameters struct {
	LicenseRules          []string `json:"license_rules"`
	Name                  string   `json:"name"`
	Description           string   `json:"description"`
	LicenseCount          int      `json:"license_count"`
	LicenseCountHardLimit bool     `json:"license_count_hard_limit"`
	LicenseCountingType   string   `json:"license_counting_type"`
}

// A LicensemanagerLicenseConfigurationStatus defines the observed state of a LicensemanagerLicenseConfiguration
type LicensemanagerLicenseConfigurationStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     LicensemanagerLicenseConfigurationObservation `json:",inline"`
}

// A LicensemanagerLicenseConfigurationObservation records the observed state of a LicensemanagerLicenseConfiguration
type LicensemanagerLicenseConfigurationObservation struct {
	Id string `json:"id"`
}