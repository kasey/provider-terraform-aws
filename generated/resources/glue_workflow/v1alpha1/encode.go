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
	r, ok := mr.(*GlueWorkflow)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a GlueWorkflow.")
	}
	return EncodeGlueWorkflow(*r), nil
}

func EncodeGlueWorkflow(r GlueWorkflow) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeGlueWorkflow_MaxConcurrentRuns(r.Spec.ForProvider, ctyVal)
	EncodeGlueWorkflow_Name(r.Spec.ForProvider, ctyVal)
	EncodeGlueWorkflow_Tags(r.Spec.ForProvider, ctyVal)
	EncodeGlueWorkflow_DefaultRunProperties(r.Spec.ForProvider, ctyVal)
	EncodeGlueWorkflow_Description(r.Spec.ForProvider, ctyVal)
	EncodeGlueWorkflow_Arn(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeGlueWorkflow_MaxConcurrentRuns(p GlueWorkflowParameters, vals map[string]cty.Value) {
	vals["max_concurrent_runs"] = cty.NumberIntVal(p.MaxConcurrentRuns)
}

func EncodeGlueWorkflow_Name(p GlueWorkflowParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeGlueWorkflow_Tags(p GlueWorkflowParameters, vals map[string]cty.Value) {
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

func EncodeGlueWorkflow_DefaultRunProperties(p GlueWorkflowParameters, vals map[string]cty.Value) {
	if len(p.DefaultRunProperties) == 0 {
		vals["default_run_properties"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.DefaultRunProperties {
		mVals[key] = cty.StringVal(value)
	}
	vals["default_run_properties"] = cty.MapVal(mVals)
}

func EncodeGlueWorkflow_Description(p GlueWorkflowParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeGlueWorkflow_Arn(p GlueWorkflowObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}