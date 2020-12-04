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

func EncodeDbSubnetGroup(r DbSubnetGroup) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDbSubnetGroup_NamePrefix(r.Spec.ForProvider, ctyVal)
	EncodeDbSubnetGroup_SubnetIds(r.Spec.ForProvider, ctyVal)
	EncodeDbSubnetGroup_Tags(r.Spec.ForProvider, ctyVal)
	EncodeDbSubnetGroup_Description(r.Spec.ForProvider, ctyVal)
	EncodeDbSubnetGroup_Id(r.Spec.ForProvider, ctyVal)
	EncodeDbSubnetGroup_Name(r.Spec.ForProvider, ctyVal)
	EncodeDbSubnetGroup_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeDbSubnetGroup_NamePrefix(p DbSubnetGroupParameters, vals map[string]cty.Value) {
	vals["name_prefix"] = cty.StringVal(p.NamePrefix)
}

func EncodeDbSubnetGroup_SubnetIds(p DbSubnetGroupParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SubnetIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["subnet_ids"] = cty.SetVal(colVals)
}

func EncodeDbSubnetGroup_Tags(p DbSubnetGroupParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeDbSubnetGroup_Description(p DbSubnetGroupParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeDbSubnetGroup_Id(p DbSubnetGroupParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeDbSubnetGroup_Name(p DbSubnetGroupParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeDbSubnetGroup_Arn(p DbSubnetGroupObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}