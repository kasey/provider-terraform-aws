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
	r, ok := mr.(*SwfDomain)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeSwfDomain(r, ctyValue)
}

func DecodeSwfDomain(prev *SwfDomain, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeSwfDomain_Id(&new.Spec.ForProvider, valMap)
	DecodeSwfDomain_Name(&new.Spec.ForProvider, valMap)
	DecodeSwfDomain_NamePrefix(&new.Spec.ForProvider, valMap)
	DecodeSwfDomain_Tags(&new.Spec.ForProvider, valMap)
	DecodeSwfDomain_WorkflowExecutionRetentionPeriodInDays(&new.Spec.ForProvider, valMap)
	DecodeSwfDomain_Description(&new.Spec.ForProvider, valMap)
	DecodeSwfDomain_Arn(&new.Status.AtProvider, valMap)
	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodeSwfDomain_Id(p *SwfDomainParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

func DecodeSwfDomain_Name(p *SwfDomainParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

func DecodeSwfDomain_NamePrefix(p *SwfDomainParameters, vals map[string]cty.Value) {
	p.NamePrefix = ctwhy.ValueAsString(vals["name_prefix"])
}

func DecodeSwfDomain_Tags(p *SwfDomainParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

func DecodeSwfDomain_WorkflowExecutionRetentionPeriodInDays(p *SwfDomainParameters, vals map[string]cty.Value) {
	p.WorkflowExecutionRetentionPeriodInDays = ctwhy.ValueAsString(vals["workflow_execution_retention_period_in_days"])
}

func DecodeSwfDomain_Description(p *SwfDomainParameters, vals map[string]cty.Value) {
	p.Description = ctwhy.ValueAsString(vals["description"])
}

func DecodeSwfDomain_Arn(p *SwfDomainObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}