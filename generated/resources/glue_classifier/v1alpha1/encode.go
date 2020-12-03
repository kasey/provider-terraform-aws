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

package v1alpha1func EncodeGlueClassifier(r GlueClassifier) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeGlueClassifier_Id(r.Spec.ForProvider, ctyVal)
	EncodeGlueClassifier_Name(r.Spec.ForProvider, ctyVal)
	EncodeGlueClassifier_CsvClassifier(r.Spec.ForProvider.CsvClassifier, ctyVal)
	EncodeGlueClassifier_GrokClassifier(r.Spec.ForProvider.GrokClassifier, ctyVal)
	EncodeGlueClassifier_JsonClassifier(r.Spec.ForProvider.JsonClassifier, ctyVal)
	EncodeGlueClassifier_XmlClassifier(r.Spec.ForProvider.XmlClassifier, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeGlueClassifier_Id(p *GlueClassifierParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeGlueClassifier_Name(p *GlueClassifierParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeGlueClassifier_CsvClassifier(p *CsvClassifier, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.CsvClassifier {
		ctyVal = make(map[string]cty.Value)
		EncodeGlueClassifier_CsvClassifier_AllowSingleColumn(v, ctyVal)
		EncodeGlueClassifier_CsvClassifier_ContainsHeader(v, ctyVal)
		EncodeGlueClassifier_CsvClassifier_Delimiter(v, ctyVal)
		EncodeGlueClassifier_CsvClassifier_DisableValueTrimming(v, ctyVal)
		EncodeGlueClassifier_CsvClassifier_Header(v, ctyVal)
		EncodeGlueClassifier_CsvClassifier_QuoteSymbol(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["csv_classifier"] = cty.ListVal(valsForCollection)
}

func EncodeGlueClassifier_CsvClassifier_AllowSingleColumn(p *CsvClassifier, vals map[string]cty.Value) {
	vals["allow_single_column"] = cty.BoolVal(p.AllowSingleColumn)
}

func EncodeGlueClassifier_CsvClassifier_ContainsHeader(p *CsvClassifier, vals map[string]cty.Value) {
	vals["contains_header"] = cty.StringVal(p.ContainsHeader)
}

func EncodeGlueClassifier_CsvClassifier_Delimiter(p *CsvClassifier, vals map[string]cty.Value) {
	vals["delimiter"] = cty.StringVal(p.Delimiter)
}

func EncodeGlueClassifier_CsvClassifier_DisableValueTrimming(p *CsvClassifier, vals map[string]cty.Value) {
	vals["disable_value_trimming"] = cty.BoolVal(p.DisableValueTrimming)
}

func EncodeGlueClassifier_CsvClassifier_Header(p *CsvClassifier, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Header {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["header"] = cty.ListVal(colVals)
}

func EncodeGlueClassifier_CsvClassifier_QuoteSymbol(p *CsvClassifier, vals map[string]cty.Value) {
	vals["quote_symbol"] = cty.StringVal(p.QuoteSymbol)
}

func EncodeGlueClassifier_GrokClassifier(p *GrokClassifier, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.GrokClassifier {
		ctyVal = make(map[string]cty.Value)
		EncodeGlueClassifier_GrokClassifier_Classification(v, ctyVal)
		EncodeGlueClassifier_GrokClassifier_CustomPatterns(v, ctyVal)
		EncodeGlueClassifier_GrokClassifier_GrokPattern(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["grok_classifier"] = cty.ListVal(valsForCollection)
}

func EncodeGlueClassifier_GrokClassifier_Classification(p *GrokClassifier, vals map[string]cty.Value) {
	vals["classification"] = cty.StringVal(p.Classification)
}

func EncodeGlueClassifier_GrokClassifier_CustomPatterns(p *GrokClassifier, vals map[string]cty.Value) {
	vals["custom_patterns"] = cty.StringVal(p.CustomPatterns)
}

func EncodeGlueClassifier_GrokClassifier_GrokPattern(p *GrokClassifier, vals map[string]cty.Value) {
	vals["grok_pattern"] = cty.StringVal(p.GrokPattern)
}

func EncodeGlueClassifier_JsonClassifier(p *JsonClassifier, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.JsonClassifier {
		ctyVal = make(map[string]cty.Value)
		EncodeGlueClassifier_JsonClassifier_JsonPath(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["json_classifier"] = cty.ListVal(valsForCollection)
}

func EncodeGlueClassifier_JsonClassifier_JsonPath(p *JsonClassifier, vals map[string]cty.Value) {
	vals["json_path"] = cty.StringVal(p.JsonPath)
}

func EncodeGlueClassifier_XmlClassifier(p *XmlClassifier, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.XmlClassifier {
		ctyVal = make(map[string]cty.Value)
		EncodeGlueClassifier_XmlClassifier_Classification(v, ctyVal)
		EncodeGlueClassifier_XmlClassifier_RowTag(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["xml_classifier"] = cty.ListVal(valsForCollection)
}

func EncodeGlueClassifier_XmlClassifier_Classification(p *XmlClassifier, vals map[string]cty.Value) {
	vals["classification"] = cty.StringVal(p.Classification)
}

func EncodeGlueClassifier_XmlClassifier_RowTag(p *XmlClassifier, vals map[string]cty.Value) {
	vals["row_tag"] = cty.StringVal(p.RowTag)
}