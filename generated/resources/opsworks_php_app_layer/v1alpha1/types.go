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

// OpsworksPhpAppLayer is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type OpsworksPhpAppLayer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OpsworksPhpAppLayerSpec   `json:"spec"`
	Status OpsworksPhpAppLayerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OpsworksPhpAppLayer contains a list of OpsworksPhpAppLayerList
type OpsworksPhpAppLayerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OpsworksPhpAppLayer `json:"items"`
}

// A OpsworksPhpAppLayerSpec defines the desired state of a OpsworksPhpAppLayer
type OpsworksPhpAppLayerSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  OpsworksPhpAppLayerParameters `json:",inline"`
}

// A OpsworksPhpAppLayerParameters defines the desired state of a OpsworksPhpAppLayer
type OpsworksPhpAppLayerParameters struct {
	CustomJson               string   `json:"custom_json"`
	ElasticLoadBalancer      string   `json:"elastic_load_balancer"`
	Name                     string   `json:"name"`
	StackId                  string   `json:"stack_id"`
	SystemPackages           []string `json:"system_packages"`
	CustomSetupRecipes       []string `json:"custom_setup_recipes"`
	AutoAssignElasticIps     bool     `json:"auto_assign_elastic_ips"`
	AutoHealing              bool     `json:"auto_healing"`
	CustomConfigureRecipes   []string `json:"custom_configure_recipes"`
	InstanceShutdownTimeout  int      `json:"instance_shutdown_timeout"`
	AutoAssignPublicIps      bool     `json:"auto_assign_public_ips"`
	CustomDeployRecipes      []string `json:"custom_deploy_recipes"`
	CustomInstanceProfileArn string   `json:"custom_instance_profile_arn"`
	CustomSecurityGroupIds   []string `json:"custom_security_group_ids"`
	CustomShutdownRecipes    []string `json:"custom_shutdown_recipes"`
	CustomUndeployRecipes    []string `json:"custom_undeploy_recipes"`
	DrainElbOnShutdown       bool     `json:"drain_elb_on_shutdown"`
	InstallUpdatesOnBoot     bool     `json:"install_updates_on_boot"`
	UseEbsOptimizedInstances bool     `json:"use_ebs_optimized_instances"`
}

// A OpsworksPhpAppLayerStatus defines the observed state of a OpsworksPhpAppLayer
type OpsworksPhpAppLayerStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     OpsworksPhpAppLayerObservation `json:",inline"`
}

// A OpsworksPhpAppLayerObservation records the observed state of a OpsworksPhpAppLayer
type OpsworksPhpAppLayerObservation struct {
	Arn string `json:"arn"`
	Id  string `json:"id"`
}