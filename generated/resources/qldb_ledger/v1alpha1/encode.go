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

func EncodeQldbLedger(r QldbLedger) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeQldbLedger_DeletionProtection(r.Spec.ForProvider, ctyVal)
	EncodeQldbLedger_Id(r.Spec.ForProvider, ctyVal)
	EncodeQldbLedger_Name(r.Spec.ForProvider, ctyVal)
	EncodeQldbLedger_Tags(r.Spec.ForProvider, ctyVal)
	EncodeQldbLedger_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeQldbLedger_DeletionProtection(p QldbLedgerParameters, vals map[string]cty.Value) {
	vals["deletion_protection"] = cty.BoolVal(p.DeletionProtection)
}

func EncodeQldbLedger_Id(p QldbLedgerParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeQldbLedger_Name(p QldbLedgerParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeQldbLedger_Tags(p QldbLedgerParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeQldbLedger_Arn(p QldbLedgerObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}