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

package v1alpha1func EncodeIotThingType(r IotThingType) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeIotThingType_Deprecated(r.Spec.ForProvider, ctyVal)
	EncodeIotThingType_Id(r.Spec.ForProvider, ctyVal)
	EncodeIotThingType_Name(r.Spec.ForProvider, ctyVal)
	EncodeIotThingType_Properties(r.Spec.ForProvider.Properties, ctyVal)
	EncodeIotThingType_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeIotThingType_Deprecated(p *IotThingTypeParameters, vals map[string]cty.Value) {
	vals["deprecated"] = cty.BoolVal(p.Deprecated)
}

func EncodeIotThingType_Id(p *IotThingTypeParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeIotThingType_Name(p *IotThingTypeParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeIotThingType_Properties(p *Properties, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Properties {
		ctyVal = make(map[string]cty.Value)
		EncodeIotThingType_Properties_Description(v, ctyVal)
		EncodeIotThingType_Properties_SearchableAttributes(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["properties"] = cty.ListVal(valsForCollection)
}

func EncodeIotThingType_Properties_Description(p *Properties, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeIotThingType_Properties_SearchableAttributes(p *Properties, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SearchableAttributes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["searchable_attributes"] = cty.SetVal(colVals)
}

func EncodeIotThingType_Arn(p *IotThingTypeObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}