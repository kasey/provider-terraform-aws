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
	"github.com/zclconf/go-cty/cty"
)

func EncodeGlueWorkflow(r GlueWorkflow) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeGlueWorkflow_MaxConcurrentRuns(r.Spec.ForProvider, ctyVal)
	EncodeGlueWorkflow_Name(r.Spec.ForProvider, ctyVal)
	EncodeGlueWorkflow_Tags(r.Spec.ForProvider, ctyVal)
	EncodeGlueWorkflow_DefaultRunProperties(r.Spec.ForProvider, ctyVal)
	EncodeGlueWorkflow_Description(r.Spec.ForProvider, ctyVal)
	EncodeGlueWorkflow_Id(r.Spec.ForProvider, ctyVal)
	EncodeGlueWorkflow_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeGlueWorkflow_MaxConcurrentRuns(p GlueWorkflowParameters, vals map[string]cty.Value) {
	vals["max_concurrent_runs"] = cty.NumberIntVal(p.MaxConcurrentRuns)
}

func EncodeGlueWorkflow_Name(p GlueWorkflowParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeGlueWorkflow_Tags(p GlueWorkflowParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeGlueWorkflow_DefaultRunProperties(p GlueWorkflowParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.DefaultRunProperties {
		mVals[key] = cty.StringVal(value)
	}
	vals["default_run_properties"] = cty.MapVal(mVals)
}

func EncodeGlueWorkflow_Description(p GlueWorkflowParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeGlueWorkflow_Id(p GlueWorkflowParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeGlueWorkflow_Arn(p GlueWorkflowObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}