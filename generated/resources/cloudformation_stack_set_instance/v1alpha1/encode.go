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

func EncodeCloudformationStackSetInstance(r CloudformationStackSetInstance) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeCloudformationStackSetInstance_Id(r.Spec.ForProvider, ctyVal)
	EncodeCloudformationStackSetInstance_ParameterOverrides(r.Spec.ForProvider, ctyVal)
	EncodeCloudformationStackSetInstance_Region(r.Spec.ForProvider, ctyVal)
	EncodeCloudformationStackSetInstance_RetainStack(r.Spec.ForProvider, ctyVal)
	EncodeCloudformationStackSetInstance_StackSetName(r.Spec.ForProvider, ctyVal)
	EncodeCloudformationStackSetInstance_AccountId(r.Spec.ForProvider, ctyVal)
	EncodeCloudformationStackSetInstance_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeCloudformationStackSetInstance_StackId(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeCloudformationStackSetInstance_Id(p CloudformationStackSetInstanceParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeCloudformationStackSetInstance_ParameterOverrides(p CloudformationStackSetInstanceParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.ParameterOverrides {
		mVals[key] = cty.StringVal(value)
	}
	vals["parameter_overrides"] = cty.MapVal(mVals)
}

func EncodeCloudformationStackSetInstance_Region(p CloudformationStackSetInstanceParameters, vals map[string]cty.Value) {
	vals["region"] = cty.StringVal(p.Region)
}

func EncodeCloudformationStackSetInstance_RetainStack(p CloudformationStackSetInstanceParameters, vals map[string]cty.Value) {
	vals["retain_stack"] = cty.BoolVal(p.RetainStack)
}

func EncodeCloudformationStackSetInstance_StackSetName(p CloudformationStackSetInstanceParameters, vals map[string]cty.Value) {
	vals["stack_set_name"] = cty.StringVal(p.StackSetName)
}

func EncodeCloudformationStackSetInstance_AccountId(p CloudformationStackSetInstanceParameters, vals map[string]cty.Value) {
	vals["account_id"] = cty.StringVal(p.AccountId)
}

func EncodeCloudformationStackSetInstance_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeCloudformationStackSetInstance_Timeouts_Create(p, ctyVal)
	EncodeCloudformationStackSetInstance_Timeouts_Delete(p, ctyVal)
	EncodeCloudformationStackSetInstance_Timeouts_Update(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeCloudformationStackSetInstance_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeCloudformationStackSetInstance_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeCloudformationStackSetInstance_Timeouts_Update(p Timeouts, vals map[string]cty.Value) {
	vals["update"] = cty.StringVal(p.Update)
}

func EncodeCloudformationStackSetInstance_StackId(p CloudformationStackSetInstanceObservation, vals map[string]cty.Value) {
	vals["stack_id"] = cty.StringVal(p.StackId)
}