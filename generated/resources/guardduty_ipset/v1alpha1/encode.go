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

func EncodeGuarddutyIpset(r GuarddutyIpset) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeGuarddutyIpset_Location(r.Spec.ForProvider, ctyVal)
	EncodeGuarddutyIpset_Name(r.Spec.ForProvider, ctyVal)
	EncodeGuarddutyIpset_Tags(r.Spec.ForProvider, ctyVal)
	EncodeGuarddutyIpset_Activate(r.Spec.ForProvider, ctyVal)
	EncodeGuarddutyIpset_DetectorId(r.Spec.ForProvider, ctyVal)
	EncodeGuarddutyIpset_Format(r.Spec.ForProvider, ctyVal)
	EncodeGuarddutyIpset_Id(r.Spec.ForProvider, ctyVal)
	EncodeGuarddutyIpset_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeGuarddutyIpset_Location(p GuarddutyIpsetParameters, vals map[string]cty.Value) {
	vals["location"] = cty.StringVal(p.Location)
}

func EncodeGuarddutyIpset_Name(p GuarddutyIpsetParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeGuarddutyIpset_Tags(p GuarddutyIpsetParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeGuarddutyIpset_Activate(p GuarddutyIpsetParameters, vals map[string]cty.Value) {
	vals["activate"] = cty.BoolVal(p.Activate)
}

func EncodeGuarddutyIpset_DetectorId(p GuarddutyIpsetParameters, vals map[string]cty.Value) {
	vals["detector_id"] = cty.StringVal(p.DetectorId)
}

func EncodeGuarddutyIpset_Format(p GuarddutyIpsetParameters, vals map[string]cty.Value) {
	vals["format"] = cty.StringVal(p.Format)
}

func EncodeGuarddutyIpset_Id(p GuarddutyIpsetParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeGuarddutyIpset_Arn(p GuarddutyIpsetObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}