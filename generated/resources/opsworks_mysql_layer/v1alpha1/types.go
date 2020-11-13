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

// OpsworksMysqlLayer is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type OpsworksMysqlLayer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OpsworksMysqlLayerSpec   `json:"spec"`
	Status OpsworksMysqlLayerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OpsworksMysqlLayer contains a list of OpsworksMysqlLayerList
type OpsworksMysqlLayerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OpsworksMysqlLayer `json:"items"`
}

// A OpsworksMysqlLayerSpec defines the desired state of a OpsworksMysqlLayer
type OpsworksMysqlLayerSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  OpsworksMysqlLayerParameters `json:",inline"`
}

// A OpsworksMysqlLayerParameters defines the desired state of a OpsworksMysqlLayer
type OpsworksMysqlLayerParameters struct {
	CustomUndeployRecipes      []string          `json:"custom_undeploy_recipes"`
	DrainElbOnShutdown         bool              `json:"drain_elb_on_shutdown"`
	ElasticLoadBalancer        string            `json:"elastic_load_balancer"`
	SystemPackages             []string          `json:"system_packages"`
	AutoAssignElasticIps       bool              `json:"auto_assign_elastic_ips"`
	CustomInstanceProfileArn   string            `json:"custom_instance_profile_arn"`
	CustomSetupRecipes         []string          `json:"custom_setup_recipes"`
	CustomShutdownRecipes      []string          `json:"custom_shutdown_recipes"`
	InstallUpdatesOnBoot       bool              `json:"install_updates_on_boot"`
	Name                       string            `json:"name"`
	RootPassword               string            `json:"root_password"`
	RootPasswordOnAllInstances bool              `json:"root_password_on_all_instances"`
	StackId                    string            `json:"stack_id"`
	UseEbsOptimizedInstances   bool              `json:"use_ebs_optimized_instances"`
	CustomConfigureRecipes     []string          `json:"custom_configure_recipes"`
	CustomDeployRecipes        []string          `json:"custom_deploy_recipes"`
	CustomJson                 string            `json:"custom_json"`
	CustomSecurityGroupIds     []string          `json:"custom_security_group_ids"`
	Id                         string            `json:"id"`
	Tags                       map[string]string `json:"tags"`
	AutoAssignPublicIps        bool              `json:"auto_assign_public_ips"`
	AutoHealing                bool              `json:"auto_healing"`
	InstanceShutdownTimeout    int               `json:"instance_shutdown_timeout"`
	EbsVolume                  []EbsVolume       `json:"ebs_volume"`
}

type EbsVolume struct {
	Type          string `json:"type"`
	Encrypted     bool   `json:"encrypted"`
	Iops          int    `json:"iops"`
	MountPoint    string `json:"mount_point"`
	NumberOfDisks int    `json:"number_of_disks"`
	RaidLevel     string `json:"raid_level"`
	Size          int    `json:"size"`
}

// A OpsworksMysqlLayerStatus defines the observed state of a OpsworksMysqlLayer
type OpsworksMysqlLayerStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     OpsworksMysqlLayerObservation `json:",inline"`
}

// A OpsworksMysqlLayerObservation records the observed state of a OpsworksMysqlLayer
type OpsworksMysqlLayerObservation struct {
	Arn string `json:"arn"`
}