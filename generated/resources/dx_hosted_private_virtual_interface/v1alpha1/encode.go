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
	"fmt"

	"github.com/zclconf/go-cty/cty"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
)

type ctyEncoder struct{}

func (e *ctyEncoder) EncodeCty(mr resource.Managed, schema *providers.Schema) (cty.Value, error) {
	r, ok := mr.(*DxHostedPrivateVirtualInterface)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a DxHostedPrivateVirtualInterface.")
	}
	return EncodeDxHostedPrivateVirtualInterface(*r), nil
}

func EncodeDxHostedPrivateVirtualInterface(r DxHostedPrivateVirtualInterface) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDxHostedPrivateVirtualInterface_ConnectionId(r.Spec.ForProvider, ctyVal)
	EncodeDxHostedPrivateVirtualInterface_CustomerAddress(r.Spec.ForProvider, ctyVal)
	EncodeDxHostedPrivateVirtualInterface_OwnerAccountId(r.Spec.ForProvider, ctyVal)
	EncodeDxHostedPrivateVirtualInterface_AddressFamily(r.Spec.ForProvider, ctyVal)
	EncodeDxHostedPrivateVirtualInterface_BgpAuthKey(r.Spec.ForProvider, ctyVal)
	EncodeDxHostedPrivateVirtualInterface_Mtu(r.Spec.ForProvider, ctyVal)
	EncodeDxHostedPrivateVirtualInterface_Name(r.Spec.ForProvider, ctyVal)
	EncodeDxHostedPrivateVirtualInterface_BgpAsn(r.Spec.ForProvider, ctyVal)
	EncodeDxHostedPrivateVirtualInterface_Vlan(r.Spec.ForProvider, ctyVal)
	EncodeDxHostedPrivateVirtualInterface_AmazonAddress(r.Spec.ForProvider, ctyVal)
	EncodeDxHostedPrivateVirtualInterface_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeDxHostedPrivateVirtualInterface_Arn(r.Status.AtProvider, ctyVal)
	EncodeDxHostedPrivateVirtualInterface_JumboFrameCapable(r.Status.AtProvider, ctyVal)
	EncodeDxHostedPrivateVirtualInterface_AmazonSideAsn(r.Status.AtProvider, ctyVal)
	EncodeDxHostedPrivateVirtualInterface_AwsDevice(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeDxHostedPrivateVirtualInterface_ConnectionId(p DxHostedPrivateVirtualInterfaceParameters, vals map[string]cty.Value) {
	vals["connection_id"] = cty.StringVal(p.ConnectionId)
}

func EncodeDxHostedPrivateVirtualInterface_CustomerAddress(p DxHostedPrivateVirtualInterfaceParameters, vals map[string]cty.Value) {
	vals["customer_address"] = cty.StringVal(p.CustomerAddress)
}

func EncodeDxHostedPrivateVirtualInterface_OwnerAccountId(p DxHostedPrivateVirtualInterfaceParameters, vals map[string]cty.Value) {
	vals["owner_account_id"] = cty.StringVal(p.OwnerAccountId)
}

func EncodeDxHostedPrivateVirtualInterface_AddressFamily(p DxHostedPrivateVirtualInterfaceParameters, vals map[string]cty.Value) {
	vals["address_family"] = cty.StringVal(p.AddressFamily)
}

func EncodeDxHostedPrivateVirtualInterface_BgpAuthKey(p DxHostedPrivateVirtualInterfaceParameters, vals map[string]cty.Value) {
	vals["bgp_auth_key"] = cty.StringVal(p.BgpAuthKey)
}

func EncodeDxHostedPrivateVirtualInterface_Mtu(p DxHostedPrivateVirtualInterfaceParameters, vals map[string]cty.Value) {
	vals["mtu"] = cty.NumberIntVal(p.Mtu)
}

func EncodeDxHostedPrivateVirtualInterface_Name(p DxHostedPrivateVirtualInterfaceParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeDxHostedPrivateVirtualInterface_BgpAsn(p DxHostedPrivateVirtualInterfaceParameters, vals map[string]cty.Value) {
	vals["bgp_asn"] = cty.NumberIntVal(p.BgpAsn)
}

func EncodeDxHostedPrivateVirtualInterface_Vlan(p DxHostedPrivateVirtualInterfaceParameters, vals map[string]cty.Value) {
	vals["vlan"] = cty.NumberIntVal(p.Vlan)
}

func EncodeDxHostedPrivateVirtualInterface_AmazonAddress(p DxHostedPrivateVirtualInterfaceParameters, vals map[string]cty.Value) {
	vals["amazon_address"] = cty.StringVal(p.AmazonAddress)
}

func EncodeDxHostedPrivateVirtualInterface_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeDxHostedPrivateVirtualInterface_Timeouts_Update(p, ctyVal)
	EncodeDxHostedPrivateVirtualInterface_Timeouts_Create(p, ctyVal)
	EncodeDxHostedPrivateVirtualInterface_Timeouts_Delete(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeDxHostedPrivateVirtualInterface_Timeouts_Update(p Timeouts, vals map[string]cty.Value) {
	vals["update"] = cty.StringVal(p.Update)
}

func EncodeDxHostedPrivateVirtualInterface_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeDxHostedPrivateVirtualInterface_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeDxHostedPrivateVirtualInterface_Arn(p DxHostedPrivateVirtualInterfaceObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeDxHostedPrivateVirtualInterface_JumboFrameCapable(p DxHostedPrivateVirtualInterfaceObservation, vals map[string]cty.Value) {
	vals["jumbo_frame_capable"] = cty.BoolVal(p.JumboFrameCapable)
}

func EncodeDxHostedPrivateVirtualInterface_AmazonSideAsn(p DxHostedPrivateVirtualInterfaceObservation, vals map[string]cty.Value) {
	vals["amazon_side_asn"] = cty.StringVal(p.AmazonSideAsn)
}

func EncodeDxHostedPrivateVirtualInterface_AwsDevice(p DxHostedPrivateVirtualInterfaceObservation, vals map[string]cty.Value) {
	vals["aws_device"] = cty.StringVal(p.AwsDevice)
}