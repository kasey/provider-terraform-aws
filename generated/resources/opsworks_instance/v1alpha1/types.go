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

// OpsworksInstance is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type OpsworksInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OpsworksInstanceSpec   `json:"spec"`
	Status OpsworksInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OpsworksInstance contains a list of OpsworksInstanceList
type OpsworksInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OpsworksInstance `json:"items"`
}

// A OpsworksInstanceSpec defines the desired state of a OpsworksInstance
type OpsworksInstanceSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  OpsworksInstanceParameters `json:",inline"`
}

// A OpsworksInstanceParameters defines the desired state of a OpsworksInstance
type OpsworksInstanceParameters struct {
	ReportedOsName           string                 `json:"reported_os_name"`
	RootDeviceVolumeId       string                 `json:"root_device_volume_id"`
	SecurityGroupIds         []string               `json:"security_group_ids"`
	Tenancy                  string                 `json:"tenancy"`
	AvailabilityZone         string                 `json:"availability_zone"`
	InstallUpdatesOnBoot     bool                   `json:"install_updates_on_boot"`
	CreatedAt                string                 `json:"created_at"`
	Os                       string                 `json:"os"`
	Platform                 string                 `json:"platform"`
	RootDeviceType           string                 `json:"root_device_type"`
	AgentVersion             string                 `json:"agent_version"`
	AmiId                    string                 `json:"ami_id"`
	Hostname                 string                 `json:"hostname"`
	InstanceProfileArn       string                 `json:"instance_profile_arn"`
	LastServiceErrorId       string                 `json:"last_service_error_id"`
	ReportedOsFamily         string                 `json:"reported_os_family"`
	AutoScalingType          string                 `json:"auto_scaling_type"`
	ElasticIp                string                 `json:"elastic_ip"`
	Id                       string                 `json:"id"`
	LayerIds                 []string               `json:"layer_ids"`
	PublicDns                string                 `json:"public_dns"`
	SshHostRsaKeyFingerprint string                 `json:"ssh_host_rsa_key_fingerprint"`
	Status                   string                 `json:"status"`
	SubnetId                 string                 `json:"subnet_id"`
	Architecture             string                 `json:"architecture"`
	DeleteEbs                bool                   `json:"delete_ebs"`
	SshKeyName               string                 `json:"ssh_key_name"`
	PrivateDns               string                 `json:"private_dns"`
	ReportedOsVersion        string                 `json:"reported_os_version"`
	SshHostDsaKeyFingerprint string                 `json:"ssh_host_dsa_key_fingerprint"`
	State                    string                 `json:"state"`
	InfrastructureClass      string                 `json:"infrastructure_class"`
	PrivateIp                string                 `json:"private_ip"`
	PublicIp                 string                 `json:"public_ip"`
	ReportedAgentVersion     string                 `json:"reported_agent_version"`
	StackId                  string                 `json:"stack_id"`
	VirtualizationType       string                 `json:"virtualization_type"`
	DeleteEip                bool                   `json:"delete_eip"`
	EbsOptimized             bool                   `json:"ebs_optimized"`
	InstanceType             string                 `json:"instance_type"`
	RegisteredBy             string                 `json:"registered_by"`
	EcsClusterArn            string                 `json:"ecs_cluster_arn"`
	RootBlockDevice          []RootBlockDevice      `json:"root_block_device"`
	Timeouts                 []Timeouts             `json:"timeouts"`
	EbsBlockDevice           []EbsBlockDevice       `json:"ebs_block_device"`
	EphemeralBlockDevice     []EphemeralBlockDevice `json:"ephemeral_block_device"`
}

type RootBlockDevice struct {
	DeleteOnTermination bool   `json:"delete_on_termination"`
	Iops                int    `json:"iops"`
	VolumeSize          int    `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
}

type Timeouts struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

type EbsBlockDevice struct {
	Iops                int    `json:"iops"`
	SnapshotId          string `json:"snapshot_id"`
	VolumeSize          int    `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
	DeleteOnTermination bool   `json:"delete_on_termination"`
	DeviceName          string `json:"device_name"`
}

type EphemeralBlockDevice struct {
	DeviceName  string `json:"device_name"`
	VirtualName string `json:"virtual_name"`
}

// A OpsworksInstanceStatus defines the observed state of a OpsworksInstance
type OpsworksInstanceStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     OpsworksInstanceObservation `json:",inline"`
}

// A OpsworksInstanceObservation records the observed state of a OpsworksInstance
type OpsworksInstanceObservation struct {
	Ec2InstanceId string `json:"ec2_instance_id"`
}