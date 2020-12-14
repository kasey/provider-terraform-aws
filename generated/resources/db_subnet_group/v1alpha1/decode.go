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
	r, ok := mr.(*DbSubnetGroup)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeDbSubnetGroup(r, ctyValue)
}

func DecodeDbSubnetGroup(prev *DbSubnetGroup, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeDbSubnetGroup_NamePrefix(&new.Spec.ForProvider, valMap)
	DecodeDbSubnetGroup_SubnetIds(&new.Spec.ForProvider, valMap)
	DecodeDbSubnetGroup_Tags(&new.Spec.ForProvider, valMap)
	DecodeDbSubnetGroup_Description(&new.Spec.ForProvider, valMap)
	DecodeDbSubnetGroup_Id(&new.Spec.ForProvider, valMap)
	DecodeDbSubnetGroup_Name(&new.Spec.ForProvider, valMap)
	DecodeDbSubnetGroup_Arn(&new.Status.AtProvider, valMap)
	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodeDbSubnetGroup_NamePrefix(p *DbSubnetGroupParameters, vals map[string]cty.Value) {
	p.NamePrefix = ctwhy.ValueAsString(vals["name_prefix"])
}

func DecodeDbSubnetGroup_SubnetIds(p *DbSubnetGroupParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["subnet_ids"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.SubnetIds = goVals
}

func DecodeDbSubnetGroup_Tags(p *DbSubnetGroupParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

func DecodeDbSubnetGroup_Description(p *DbSubnetGroupParameters, vals map[string]cty.Value) {
	p.Description = ctwhy.ValueAsString(vals["description"])
}

func DecodeDbSubnetGroup_Id(p *DbSubnetGroupParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

func DecodeDbSubnetGroup_Name(p *DbSubnetGroupParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

func DecodeDbSubnetGroup_Arn(p *DbSubnetGroupObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}