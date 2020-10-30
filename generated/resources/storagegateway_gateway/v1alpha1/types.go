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

// StoragegatewayGateway is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type StoragegatewayGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StoragegatewayGatewaySpec   `json:"spec"`
	Status StoragegatewayGatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StoragegatewayGateway contains a list of StoragegatewayGatewayList
type StoragegatewayGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StoragegatewayGateway `json:"items"`
}

// A StoragegatewayGatewaySpec defines the desired state of a StoragegatewayGateway
type StoragegatewayGatewaySpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  StoragegatewayGatewayParameters `json:",inline"`
}

// A StoragegatewayGatewayParameters defines the desired state of a StoragegatewayGateway
type StoragegatewayGatewayParameters struct {
	GatewayName                          string `json:"gateway_name"`
	GatewayType                          string `json:"gateway_type"`
	MediumChangerType                    string `json:"medium_changer_type"`
	SmbGuestPassword                     string `json:"smb_guest_password"`
	AverageUploadRateLimitInBitsPerSec   int    `json:"average_upload_rate_limit_in_bits_per_sec"`
	GatewayTimezone                      string `json:"gateway_timezone"`
	TapeDriveType                        string `json:"tape_drive_type"`
	AverageDownloadRateLimitInBitsPerSec int    `json:"average_download_rate_limit_in_bits_per_sec"`
	CloudwatchLogGroupArn                string `json:"cloudwatch_log_group_arn"`
	GatewayVpcEndpoint                   string `json:"gateway_vpc_endpoint"`
}

// A StoragegatewayGatewayStatus defines the observed state of a StoragegatewayGateway
type StoragegatewayGatewayStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     StoragegatewayGatewayObservation `json:",inline"`
}

// A StoragegatewayGatewayObservation records the observed state of a StoragegatewayGateway
type StoragegatewayGatewayObservation struct {
	Id                  string `json:"id"`
	Arn                 string `json:"arn"`
	GatewayId           string `json:"gateway_id"`
	SmbSecurityStrategy string `json:"smb_security_strategy"`
	ActivationKey       string `json:"activation_key"`
	GatewayIpAddress    string `json:"gateway_ip_address"`
}