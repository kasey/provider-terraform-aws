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

// CurReportDefinition is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type CurReportDefinition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CurReportDefinitionSpec   `json:"spec"`
	Status CurReportDefinitionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CurReportDefinition contains a list of CurReportDefinitionList
type CurReportDefinitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CurReportDefinition `json:"items"`
}

// A CurReportDefinitionSpec defines the desired state of a CurReportDefinition
type CurReportDefinitionSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  CurReportDefinitionParameters `json:"forProvider"`
}

// A CurReportDefinitionParameters defines the desired state of a CurReportDefinition
type CurReportDefinitionParameters struct {
	AdditionalSchemaElements []string `json:"additional_schema_elements"`
	Compression              string   `json:"compression"`
	RefreshClosedReports     bool     `json:"refresh_closed_reports"`
	S3Prefix                 string   `json:"s3_prefix"`
	TimeUnit                 string   `json:"time_unit"`
	AdditionalArtifacts      []string `json:"additional_artifacts"`
	Format                   string   `json:"format"`
	ReportName               string   `json:"report_name"`
	ReportVersioning         string   `json:"report_versioning"`
	S3Bucket                 string   `json:"s3_bucket"`
	S3Region                 string   `json:"s3_region"`
}

// A CurReportDefinitionStatus defines the observed state of a CurReportDefinition
type CurReportDefinitionStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     CurReportDefinitionObservation `json:"atProvider"`
}

// A CurReportDefinitionObservation records the observed state of a CurReportDefinition
type CurReportDefinitionObservation struct{}