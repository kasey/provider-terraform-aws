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
	r, ok := mr.(*DefaultVpcDhcpOptions)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeDefaultVpcDhcpOptions(r, ctyValue)
}

func DecodeDefaultVpcDhcpOptions(prev *DefaultVpcDhcpOptions, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeDefaultVpcDhcpOptions_NetbiosNameServers(&new.Spec.ForProvider, valMap)
	DecodeDefaultVpcDhcpOptions_NetbiosNodeType(&new.Spec.ForProvider, valMap)
	DecodeDefaultVpcDhcpOptions_Tags(&new.Spec.ForProvider, valMap)
	DecodeDefaultVpcDhcpOptions_OwnerId(&new.Status.AtProvider, valMap)
	DecodeDefaultVpcDhcpOptions_DomainName(&new.Status.AtProvider, valMap)
	DecodeDefaultVpcDhcpOptions_DomainNameServers(&new.Status.AtProvider, valMap)
	DecodeDefaultVpcDhcpOptions_Arn(&new.Status.AtProvider, valMap)
	DecodeDefaultVpcDhcpOptions_NtpServers(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveCollectionTypeDecodeTemplate
func DecodeDefaultVpcDhcpOptions_NetbiosNameServers(p *DefaultVpcDhcpOptionsParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsList(vals["netbios_name_servers"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.NetbiosNameServers = goVals
}

//primitiveTypeDecodeTemplate
func DecodeDefaultVpcDhcpOptions_NetbiosNodeType(p *DefaultVpcDhcpOptionsParameters, vals map[string]cty.Value) {
	p.NetbiosNodeType = ctwhy.ValueAsString(vals["netbios_node_type"])
}

//primitiveMapTypeDecodeTemplate
func DecodeDefaultVpcDhcpOptions_Tags(p *DefaultVpcDhcpOptionsParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

//primitiveTypeDecodeTemplate
func DecodeDefaultVpcDhcpOptions_OwnerId(p *DefaultVpcDhcpOptionsObservation, vals map[string]cty.Value) {
	p.OwnerId = ctwhy.ValueAsString(vals["owner_id"])
}

//primitiveTypeDecodeTemplate
func DecodeDefaultVpcDhcpOptions_DomainName(p *DefaultVpcDhcpOptionsObservation, vals map[string]cty.Value) {
	p.DomainName = ctwhy.ValueAsString(vals["domain_name"])
}

//primitiveTypeDecodeTemplate
func DecodeDefaultVpcDhcpOptions_DomainNameServers(p *DefaultVpcDhcpOptionsObservation, vals map[string]cty.Value) {
	p.DomainNameServers = ctwhy.ValueAsString(vals["domain_name_servers"])
}

//primitiveTypeDecodeTemplate
func DecodeDefaultVpcDhcpOptions_Arn(p *DefaultVpcDhcpOptionsObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}

//primitiveTypeDecodeTemplate
func DecodeDefaultVpcDhcpOptions_NtpServers(p *DefaultVpcDhcpOptionsObservation, vals map[string]cty.Value) {
	p.NtpServers = ctwhy.ValueAsString(vals["ntp_servers"])
}