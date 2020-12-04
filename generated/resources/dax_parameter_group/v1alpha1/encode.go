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

func EncodeDaxParameterGroup(r DaxParameterGroup) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDaxParameterGroup_Description(r.Spec.ForProvider, ctyVal)
	EncodeDaxParameterGroup_Id(r.Spec.ForProvider, ctyVal)
	EncodeDaxParameterGroup_Name(r.Spec.ForProvider, ctyVal)
	EncodeDaxParameterGroup_Parameters(r.Spec.ForProvider.Parameters, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeDaxParameterGroup_Description(p DaxParameterGroupParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeDaxParameterGroup_Id(p DaxParameterGroupParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeDaxParameterGroup_Name(p DaxParameterGroupParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeDaxParameterGroup_Parameters(p Parameters, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeDaxParameterGroup_Parameters_Name(p, ctyVal)
	EncodeDaxParameterGroup_Parameters_Value(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["parameters"] = cty.SetVal(valsForCollection)
}

func EncodeDaxParameterGroup_Parameters_Name(p Parameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeDaxParameterGroup_Parameters_Value(p Parameters, vals map[string]cty.Value) {
	vals["value"] = cty.StringVal(p.Value)
}