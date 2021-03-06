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
	r, ok := mr.(*LightsailStaticIp)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeLightsailStaticIp(r, ctyValue)
}

func DecodeLightsailStaticIp(prev *LightsailStaticIp, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeLightsailStaticIp_Name(&new.Spec.ForProvider, valMap)
	DecodeLightsailStaticIp_Arn(&new.Status.AtProvider, valMap)
	DecodeLightsailStaticIp_IpAddress(&new.Status.AtProvider, valMap)
	DecodeLightsailStaticIp_SupportCode(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeLightsailStaticIp_Name(p *LightsailStaticIpParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//primitiveTypeDecodeTemplate
func DecodeLightsailStaticIp_Arn(p *LightsailStaticIpObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}

//primitiveTypeDecodeTemplate
func DecodeLightsailStaticIp_IpAddress(p *LightsailStaticIpObservation, vals map[string]cty.Value) {
	p.IpAddress = ctwhy.ValueAsString(vals["ip_address"])
}

//primitiveTypeDecodeTemplate
func DecodeLightsailStaticIp_SupportCode(p *LightsailStaticIpObservation, vals map[string]cty.Value) {
	p.SupportCode = ctwhy.ValueAsString(vals["support_code"])
}