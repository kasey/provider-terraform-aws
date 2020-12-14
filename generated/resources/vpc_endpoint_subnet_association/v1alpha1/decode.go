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
	r, ok := mr.(*VpcEndpointSubnetAssociation)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeVpcEndpointSubnetAssociation(r, ctyValue)
}

func DecodeVpcEndpointSubnetAssociation(prev *VpcEndpointSubnetAssociation, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeVpcEndpointSubnetAssociation_VpcEndpointId(&new.Spec.ForProvider, valMap)
	DecodeVpcEndpointSubnetAssociation_Id(&new.Spec.ForProvider, valMap)
	DecodeVpcEndpointSubnetAssociation_SubnetId(&new.Spec.ForProvider, valMap)
	DecodeVpcEndpointSubnetAssociation_Timeouts(&new.Spec.ForProvider.Timeouts, valMap)

	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodeVpcEndpointSubnetAssociation_VpcEndpointId(p *VpcEndpointSubnetAssociationParameters, vals map[string]cty.Value) {
	p.VpcEndpointId = ctwhy.ValueAsString(vals["vpc_endpoint_id"])
}

func DecodeVpcEndpointSubnetAssociation_Id(p *VpcEndpointSubnetAssociationParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

func DecodeVpcEndpointSubnetAssociation_SubnetId(p *VpcEndpointSubnetAssociationParameters, vals map[string]cty.Value) {
	p.SubnetId = ctwhy.ValueAsString(vals["subnet_id"])
}

func DecodeVpcEndpointSubnetAssociation_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	valMap := vals["timeouts"].AsValueMap()
	DecodeVpcEndpointSubnetAssociation_Timeouts_Create(p, valMap)
	DecodeVpcEndpointSubnetAssociation_Timeouts_Delete(p, valMap)
}

func DecodeVpcEndpointSubnetAssociation_Timeouts_Create(p *Timeouts, vals map[string]cty.Value) {
	p.Create = ctwhy.ValueAsString(vals["create"])
}

func DecodeVpcEndpointSubnetAssociation_Timeouts_Delete(p *Timeouts, vals map[string]cty.Value) {
	p.Delete = ctwhy.ValueAsString(vals["delete"])
}