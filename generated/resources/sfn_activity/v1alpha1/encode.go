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

func EncodeSfnActivity(r SfnActivity) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeSfnActivity_Id(r.Spec.ForProvider, ctyVal)
	EncodeSfnActivity_Name(r.Spec.ForProvider, ctyVal)
	EncodeSfnActivity_Tags(r.Spec.ForProvider, ctyVal)
	EncodeSfnActivity_CreationDate(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeSfnActivity_Id(p SfnActivityParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeSfnActivity_Name(p SfnActivityParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeSfnActivity_Tags(p SfnActivityParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeSfnActivity_CreationDate(p SfnActivityObservation, vals map[string]cty.Value) {
	vals["creation_date"] = cty.StringVal(p.CreationDate)
}