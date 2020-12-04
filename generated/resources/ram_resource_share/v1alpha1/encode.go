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

func EncodeRamResourceShare(r RamResourceShare) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeRamResourceShare_Id(r.Spec.ForProvider, ctyVal)
	EncodeRamResourceShare_Name(r.Spec.ForProvider, ctyVal)
	EncodeRamResourceShare_Tags(r.Spec.ForProvider, ctyVal)
	EncodeRamResourceShare_AllowExternalPrincipals(r.Spec.ForProvider, ctyVal)
	EncodeRamResourceShare_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeRamResourceShare_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeRamResourceShare_Id(p RamResourceShareParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeRamResourceShare_Name(p RamResourceShareParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeRamResourceShare_Tags(p RamResourceShareParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeRamResourceShare_AllowExternalPrincipals(p RamResourceShareParameters, vals map[string]cty.Value) {
	vals["allow_external_principals"] = cty.BoolVal(p.AllowExternalPrincipals)
}

func EncodeRamResourceShare_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeRamResourceShare_Timeouts_Create(p, ctyVal)
	EncodeRamResourceShare_Timeouts_Delete(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeRamResourceShare_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeRamResourceShare_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeRamResourceShare_Arn(p RamResourceShareObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}