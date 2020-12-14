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

	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
	"github.com/zclconf/go-cty/cty"
	ctwhy "github.com/crossplane-contrib/terraform-runtime/pkg/plugin/cty"
)

type ctyDecoder struct{}

func (e *ctyDecoder) DecodeCty(mr resource.Managed, ctyValue cty.Value, schema *providers.Schema) (resource.Managed, error) {
	r, ok := mr.(*Eip)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeEip(r, ctyValue)
}

func DecodeEip(prev *Eip, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeEip_Id(&new.Spec.ForProvider, valMap)
	DecodeEip_Instance(&new.Spec.ForProvider, valMap)
	DecodeEip_NetworkInterface(&new.Spec.ForProvider, valMap)
	DecodeEip_PublicIpv4Pool(&new.Spec.ForProvider, valMap)
	DecodeEip_CustomerOwnedIpv4Pool(&new.Spec.ForProvider, valMap)
	DecodeEip_Vpc(&new.Spec.ForProvider, valMap)
	DecodeEip_Tags(&new.Spec.ForProvider, valMap)
	DecodeEip_AssociateWithPrivateIp(&new.Spec.ForProvider, valMap)
	DecodeEip_Timeouts(&new.Spec.ForProvider.Timeouts, valMap)
	DecodeEip_PublicDns(&new.Status.AtProvider, valMap)
	DecodeEip_AssociationId(&new.Status.AtProvider, valMap)
	DecodeEip_Domain(&new.Status.AtProvider, valMap)
	DecodeEip_PrivateDns(&new.Status.AtProvider, valMap)
	DecodeEip_CustomerOwnedIp(&new.Status.AtProvider, valMap)
	DecodeEip_PublicIp(&new.Status.AtProvider, valMap)
	DecodeEip_PrivateIp(&new.Status.AtProvider, valMap)
	DecodeEip_AllocationId(&new.Status.AtProvider, valMap)
	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodeEip_Id(p *EipParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

func DecodeEip_Instance(p *EipParameters, vals map[string]cty.Value) {
	p.Instance = ctwhy.ValueAsString(vals["instance"])
}

func DecodeEip_NetworkInterface(p *EipParameters, vals map[string]cty.Value) {
	p.NetworkInterface = ctwhy.ValueAsString(vals["network_interface"])
}

func DecodeEip_PublicIpv4Pool(p *EipParameters, vals map[string]cty.Value) {
	p.PublicIpv4Pool = ctwhy.ValueAsString(vals["public_ipv4_pool"])
}

func DecodeEip_CustomerOwnedIpv4Pool(p *EipParameters, vals map[string]cty.Value) {
	p.CustomerOwnedIpv4Pool = ctwhy.ValueAsString(vals["customer_owned_ipv4_pool"])
}

func DecodeEip_Vpc(p *EipParameters, vals map[string]cty.Value) {
	p.Vpc = ctwhy.ValueAsBool(vals["vpc"])
}

func DecodeEip_Tags(p *EipParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

func DecodeEip_AssociateWithPrivateIp(p *EipParameters, vals map[string]cty.Value) {
	p.AssociateWithPrivateIp = ctwhy.ValueAsString(vals["associate_with_private_ip"])
}

func DecodeEip_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	valMap := vals["timeouts"].AsValueMap()
	DecodeEip_Timeouts_Delete(p, valMap)
	DecodeEip_Timeouts_Read(p, valMap)
	DecodeEip_Timeouts_Update(p, valMap)
}

func DecodeEip_Timeouts_Delete(p *Timeouts, vals map[string]cty.Value) {
	p.Delete = ctwhy.ValueAsString(vals["delete"])
}

func DecodeEip_Timeouts_Read(p *Timeouts, vals map[string]cty.Value) {
	p.Read = ctwhy.ValueAsString(vals["read"])
}

func DecodeEip_Timeouts_Update(p *Timeouts, vals map[string]cty.Value) {
	p.Update = ctwhy.ValueAsString(vals["update"])
}

func DecodeEip_PublicDns(p *EipObservation, vals map[string]cty.Value) {
	p.PublicDns = ctwhy.ValueAsString(vals["public_dns"])
}

func DecodeEip_AssociationId(p *EipObservation, vals map[string]cty.Value) {
	p.AssociationId = ctwhy.ValueAsString(vals["association_id"])
}

func DecodeEip_Domain(p *EipObservation, vals map[string]cty.Value) {
	p.Domain = ctwhy.ValueAsString(vals["domain"])
}

func DecodeEip_PrivateDns(p *EipObservation, vals map[string]cty.Value) {
	p.PrivateDns = ctwhy.ValueAsString(vals["private_dns"])
}

func DecodeEip_CustomerOwnedIp(p *EipObservation, vals map[string]cty.Value) {
	p.CustomerOwnedIp = ctwhy.ValueAsString(vals["customer_owned_ip"])
}

func DecodeEip_PublicIp(p *EipObservation, vals map[string]cty.Value) {
	p.PublicIp = ctwhy.ValueAsString(vals["public_ip"])
}

func DecodeEip_PrivateIp(p *EipObservation, vals map[string]cty.Value) {
	p.PrivateIp = ctwhy.ValueAsString(vals["private_ip"])
}

func DecodeEip_AllocationId(p *EipObservation, vals map[string]cty.Value) {
	p.AllocationId = ctwhy.ValueAsString(vals["allocation_id"])
}