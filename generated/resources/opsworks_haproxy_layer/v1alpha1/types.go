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

// OpsworksHaproxyLayer is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type OpsworksHaproxyLayer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OpsworksHaproxyLayerSpec   `json:"spec"`
	Status OpsworksHaproxyLayerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OpsworksHaproxyLayer contains a list of OpsworksHaproxyLayerList
type OpsworksHaproxyLayerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OpsworksHaproxyLayer `json:"items"`
}

// A OpsworksHaproxyLayerSpec defines the desired state of a OpsworksHaproxyLayer
type OpsworksHaproxyLayerSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  OpsworksHaproxyLayerParameters `json:",inline"`
}

// A OpsworksHaproxyLayerParameters defines the desired state of a OpsworksHaproxyLayer
type OpsworksHaproxyLayerParameters struct {
	AutoAssignElasticIps     bool     `json:"auto_assign_elastic_ips"`
	AutoAssignPublicIps      bool     `json:"auto_assign_public_ips"`
	CustomUndeployRecipes    []string `json:"custom_undeploy_recipes"`
	Name                     string   `json:"name"`
	CustomConfigureRecipes   []string `json:"custom_configure_recipes"`
	CustomShutdownRecipes    []string `json:"custom_shutdown_recipes"`
	ElasticLoadBalancer      string   `json:"elastic_load_balancer"`
	HealthcheckUrl           string   `json:"healthcheck_url"`
	StatsEnabled             bool     `json:"stats_enabled"`
	SystemPackages           []string `json:"system_packages"`
	CustomInstanceProfileArn string   `json:"custom_instance_profile_arn"`
	CustomJson               string   `json:"custom_json"`
	InstallUpdatesOnBoot     bool     `json:"install_updates_on_boot"`
	StackId                  string   `json:"stack_id"`
	StatsUser                string   `json:"stats_user"`
	StatsUrl                 string   `json:"stats_url"`
	AutoHealing              bool     `json:"auto_healing"`
	CustomDeployRecipes      []string `json:"custom_deploy_recipes"`
	CustomSecurityGroupIds   []string `json:"custom_security_group_ids"`
	DrainElbOnShutdown       bool     `json:"drain_elb_on_shutdown"`
	InstanceShutdownTimeout  int      `json:"instance_shutdown_timeout"`
	UseEbsOptimizedInstances bool     `json:"use_ebs_optimized_instances"`
	CustomSetupRecipes       []string `json:"custom_setup_recipes"`
	HealthcheckMethod        string   `json:"healthcheck_method"`
	StatsPassword            string   `json:"stats_password"`
}

// A OpsworksHaproxyLayerStatus defines the observed state of a OpsworksHaproxyLayer
type OpsworksHaproxyLayerStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     OpsworksHaproxyLayerObservation `json:",inline"`
}

// A OpsworksHaproxyLayerObservation records the observed state of a OpsworksHaproxyLayer
type OpsworksHaproxyLayerObservation struct {
	Id  string `json:"id"`
	Arn string `json:"arn"`
}