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
	r, ok := mr.(*XraySamplingRule)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a XraySamplingRule.")
	}
	return EncodeXraySamplingRule(*r), nil
}

func EncodeXraySamplingRule(r XraySamplingRule) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeXraySamplingRule_Priority(r.Spec.ForProvider, ctyVal)
	EncodeXraySamplingRule_ReservoirSize(r.Spec.ForProvider, ctyVal)
	EncodeXraySamplingRule_ServiceType(r.Spec.ForProvider, ctyVal)
	EncodeXraySamplingRule_FixedRate(r.Spec.ForProvider, ctyVal)
	EncodeXraySamplingRule_Host(r.Spec.ForProvider, ctyVal)
	EncodeXraySamplingRule_HttpMethod(r.Spec.ForProvider, ctyVal)
	EncodeXraySamplingRule_Version(r.Spec.ForProvider, ctyVal)
	EncodeXraySamplingRule_RuleName(r.Spec.ForProvider, ctyVal)
	EncodeXraySamplingRule_ServiceName(r.Spec.ForProvider, ctyVal)
	EncodeXraySamplingRule_Tags(r.Spec.ForProvider, ctyVal)
	EncodeXraySamplingRule_Attributes(r.Spec.ForProvider, ctyVal)
	EncodeXraySamplingRule_ResourceArn(r.Spec.ForProvider, ctyVal)
	EncodeXraySamplingRule_UrlPath(r.Spec.ForProvider, ctyVal)
	EncodeXraySamplingRule_Arn(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeXraySamplingRule_Priority(p XraySamplingRuleParameters, vals map[string]cty.Value) {
	vals["priority"] = cty.NumberIntVal(p.Priority)
}

func EncodeXraySamplingRule_ReservoirSize(p XraySamplingRuleParameters, vals map[string]cty.Value) {
	vals["reservoir_size"] = cty.NumberIntVal(p.ReservoirSize)
}

func EncodeXraySamplingRule_ServiceType(p XraySamplingRuleParameters, vals map[string]cty.Value) {
	vals["service_type"] = cty.StringVal(p.ServiceType)
}

func EncodeXraySamplingRule_FixedRate(p XraySamplingRuleParameters, vals map[string]cty.Value) {
	vals["fixed_rate"] = cty.NumberIntVal(p.FixedRate)
}

func EncodeXraySamplingRule_Host(p XraySamplingRuleParameters, vals map[string]cty.Value) {
	vals["host"] = cty.StringVal(p.Host)
}

func EncodeXraySamplingRule_HttpMethod(p XraySamplingRuleParameters, vals map[string]cty.Value) {
	vals["http_method"] = cty.StringVal(p.HttpMethod)
}

func EncodeXraySamplingRule_Version(p XraySamplingRuleParameters, vals map[string]cty.Value) {
	vals["version"] = cty.NumberIntVal(p.Version)
}

func EncodeXraySamplingRule_RuleName(p XraySamplingRuleParameters, vals map[string]cty.Value) {
	vals["rule_name"] = cty.StringVal(p.RuleName)
}

func EncodeXraySamplingRule_ServiceName(p XraySamplingRuleParameters, vals map[string]cty.Value) {
	vals["service_name"] = cty.StringVal(p.ServiceName)
}

func EncodeXraySamplingRule_Tags(p XraySamplingRuleParameters, vals map[string]cty.Value) {
	if len(p.Tags) == 0 {
		vals["tags"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeXraySamplingRule_Attributes(p XraySamplingRuleParameters, vals map[string]cty.Value) {
	if len(p.Attributes) == 0 {
		vals["attributes"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.Attributes {
		mVals[key] = cty.StringVal(value)
	}
	vals["attributes"] = cty.MapVal(mVals)
}

func EncodeXraySamplingRule_ResourceArn(p XraySamplingRuleParameters, vals map[string]cty.Value) {
	vals["resource_arn"] = cty.StringVal(p.ResourceArn)
}

func EncodeXraySamplingRule_UrlPath(p XraySamplingRuleParameters, vals map[string]cty.Value) {
	vals["url_path"] = cty.StringVal(p.UrlPath)
}

func EncodeXraySamplingRule_Arn(p XraySamplingRuleObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}