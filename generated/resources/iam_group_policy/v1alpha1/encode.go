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

func EncodeIamGroupPolicy(r IamGroupPolicy) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeIamGroupPolicy_Group(r.Spec.ForProvider, ctyVal)
	EncodeIamGroupPolicy_Id(r.Spec.ForProvider, ctyVal)
	EncodeIamGroupPolicy_Name(r.Spec.ForProvider, ctyVal)
	EncodeIamGroupPolicy_NamePrefix(r.Spec.ForProvider, ctyVal)
	EncodeIamGroupPolicy_Policy(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeIamGroupPolicy_Group(p IamGroupPolicyParameters, vals map[string]cty.Value) {
	vals["group"] = cty.StringVal(p.Group)
}

func EncodeIamGroupPolicy_Id(p IamGroupPolicyParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeIamGroupPolicy_Name(p IamGroupPolicyParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeIamGroupPolicy_NamePrefix(p IamGroupPolicyParameters, vals map[string]cty.Value) {
	vals["name_prefix"] = cty.StringVal(p.NamePrefix)
}

func EncodeIamGroupPolicy_Policy(p IamGroupPolicyParameters, vals map[string]cty.Value) {
	vals["policy"] = cty.StringVal(p.Policy)
}