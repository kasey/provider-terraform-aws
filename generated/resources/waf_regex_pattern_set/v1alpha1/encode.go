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

func EncodeWafRegexPatternSet(r WafRegexPatternSet) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeWafRegexPatternSet_Name(r.Spec.ForProvider, ctyVal)
	EncodeWafRegexPatternSet_RegexPatternStrings(r.Spec.ForProvider, ctyVal)
	EncodeWafRegexPatternSet_Id(r.Spec.ForProvider, ctyVal)
	EncodeWafRegexPatternSet_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeWafRegexPatternSet_Name(p WafRegexPatternSetParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafRegexPatternSet_RegexPatternStrings(p WafRegexPatternSetParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.RegexPatternStrings {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["regex_pattern_strings"] = cty.SetVal(colVals)
}

func EncodeWafRegexPatternSet_Id(p WafRegexPatternSetParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeWafRegexPatternSet_Arn(p WafRegexPatternSetObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}