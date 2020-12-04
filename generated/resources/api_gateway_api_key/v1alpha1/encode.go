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

func EncodeApiGatewayApiKey(r ApiGatewayApiKey) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeApiGatewayApiKey_Description(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayApiKey_Enabled(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayApiKey_Id(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayApiKey_Tags(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayApiKey_Value(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayApiKey_Name(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayApiKey_Arn(r.Status.AtProvider, ctyVal)
	EncodeApiGatewayApiKey_CreatedDate(r.Status.AtProvider, ctyVal)
	EncodeApiGatewayApiKey_LastUpdatedDate(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeApiGatewayApiKey_Description(p ApiGatewayApiKeyParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeApiGatewayApiKey_Enabled(p ApiGatewayApiKeyParameters, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeApiGatewayApiKey_Id(p ApiGatewayApiKeyParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeApiGatewayApiKey_Tags(p ApiGatewayApiKeyParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeApiGatewayApiKey_Value(p ApiGatewayApiKeyParameters, vals map[string]cty.Value) {
	vals["value"] = cty.StringVal(p.Value)
}

func EncodeApiGatewayApiKey_Name(p ApiGatewayApiKeyParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeApiGatewayApiKey_Arn(p ApiGatewayApiKeyObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeApiGatewayApiKey_CreatedDate(p ApiGatewayApiKeyObservation, vals map[string]cty.Value) {
	vals["created_date"] = cty.StringVal(p.CreatedDate)
}

func EncodeApiGatewayApiKey_LastUpdatedDate(p ApiGatewayApiKeyObservation, vals map[string]cty.Value) {
	vals["last_updated_date"] = cty.StringVal(p.LastUpdatedDate)
}