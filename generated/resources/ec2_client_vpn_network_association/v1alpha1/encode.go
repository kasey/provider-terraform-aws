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
	r, ok := mr.(*Ec2ClientVpnNetworkAssociation)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a Ec2ClientVpnNetworkAssociation.")
	}
	return EncodeEc2ClientVpnNetworkAssociation(*r), nil
}

func EncodeEc2ClientVpnNetworkAssociation(r Ec2ClientVpnNetworkAssociation) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeEc2ClientVpnNetworkAssociation_SubnetId(r.Spec.ForProvider, ctyVal)
	EncodeEc2ClientVpnNetworkAssociation_ClientVpnEndpointId(r.Spec.ForProvider, ctyVal)
	EncodeEc2ClientVpnNetworkAssociation_SecurityGroups(r.Spec.ForProvider, ctyVal)
	EncodeEc2ClientVpnNetworkAssociation_Status(r.Status.AtProvider, ctyVal)
	EncodeEc2ClientVpnNetworkAssociation_VpcId(r.Status.AtProvider, ctyVal)
	EncodeEc2ClientVpnNetworkAssociation_AssociationId(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeEc2ClientVpnNetworkAssociation_SubnetId(p Ec2ClientVpnNetworkAssociationParameters, vals map[string]cty.Value) {
	vals["subnet_id"] = cty.StringVal(p.SubnetId)
}

func EncodeEc2ClientVpnNetworkAssociation_ClientVpnEndpointId(p Ec2ClientVpnNetworkAssociationParameters, vals map[string]cty.Value) {
	vals["client_vpn_endpoint_id"] = cty.StringVal(p.ClientVpnEndpointId)
}

func EncodeEc2ClientVpnNetworkAssociation_SecurityGroups(p Ec2ClientVpnNetworkAssociationParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SecurityGroups {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["security_groups"] = cty.SetVal(colVals)
}

func EncodeEc2ClientVpnNetworkAssociation_Status(p Ec2ClientVpnNetworkAssociationObservation, vals map[string]cty.Value) {
	vals["status"] = cty.StringVal(p.Status)
}

func EncodeEc2ClientVpnNetworkAssociation_VpcId(p Ec2ClientVpnNetworkAssociationObservation, vals map[string]cty.Value) {
	vals["vpc_id"] = cty.StringVal(p.VpcId)
}

func EncodeEc2ClientVpnNetworkAssociation_AssociationId(p Ec2ClientVpnNetworkAssociationObservation, vals map[string]cty.Value) {
	vals["association_id"] = cty.StringVal(p.AssociationId)
}