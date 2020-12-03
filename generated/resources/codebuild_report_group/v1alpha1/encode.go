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

package v1alpha1func EncodeCodebuildReportGroup(r CodebuildReportGroup) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeCodebuildReportGroup_Tags(r.Spec.ForProvider, ctyVal)
	EncodeCodebuildReportGroup_Type(r.Spec.ForProvider, ctyVal)
	EncodeCodebuildReportGroup_Id(r.Spec.ForProvider, ctyVal)
	EncodeCodebuildReportGroup_Name(r.Spec.ForProvider, ctyVal)
	EncodeCodebuildReportGroup_ExportConfig(r.Spec.ForProvider.ExportConfig, ctyVal)
	EncodeCodebuildReportGroup_Arn(r.Status.AtProvider, ctyVal)
	EncodeCodebuildReportGroup_Created(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeCodebuildReportGroup_Tags(p *CodebuildReportGroupParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeCodebuildReportGroup_Type(p *CodebuildReportGroupParameters, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeCodebuildReportGroup_Id(p *CodebuildReportGroupParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeCodebuildReportGroup_Name(p *CodebuildReportGroupParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeCodebuildReportGroup_ExportConfig(p *ExportConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.ExportConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeCodebuildReportGroup_ExportConfig_Type(v, ctyVal)
		EncodeCodebuildReportGroup_ExportConfig_S3Destination(v.S3Destination, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["export_config"] = cty.ListVal(valsForCollection)
}

func EncodeCodebuildReportGroup_ExportConfig_Type(p *ExportConfig, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeCodebuildReportGroup_ExportConfig_S3Destination(p *S3Destination, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.S3Destination {
		ctyVal = make(map[string]cty.Value)
		EncodeCodebuildReportGroup_ExportConfig_S3Destination_Bucket(v, ctyVal)
		EncodeCodebuildReportGroup_ExportConfig_S3Destination_EncryptionDisabled(v, ctyVal)
		EncodeCodebuildReportGroup_ExportConfig_S3Destination_EncryptionKey(v, ctyVal)
		EncodeCodebuildReportGroup_ExportConfig_S3Destination_Packaging(v, ctyVal)
		EncodeCodebuildReportGroup_ExportConfig_S3Destination_Path(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["s3_destination"] = cty.ListVal(valsForCollection)
}

func EncodeCodebuildReportGroup_ExportConfig_S3Destination_Bucket(p *S3Destination, vals map[string]cty.Value) {
	vals["bucket"] = cty.StringVal(p.Bucket)
}

func EncodeCodebuildReportGroup_ExportConfig_S3Destination_EncryptionDisabled(p *S3Destination, vals map[string]cty.Value) {
	vals["encryption_disabled"] = cty.BoolVal(p.EncryptionDisabled)
}

func EncodeCodebuildReportGroup_ExportConfig_S3Destination_EncryptionKey(p *S3Destination, vals map[string]cty.Value) {
	vals["encryption_key"] = cty.StringVal(p.EncryptionKey)
}

func EncodeCodebuildReportGroup_ExportConfig_S3Destination_Packaging(p *S3Destination, vals map[string]cty.Value) {
	vals["packaging"] = cty.StringVal(p.Packaging)
}

func EncodeCodebuildReportGroup_ExportConfig_S3Destination_Path(p *S3Destination, vals map[string]cty.Value) {
	vals["path"] = cty.StringVal(p.Path)
}

func EncodeCodebuildReportGroup_Arn(p *CodebuildReportGroupObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeCodebuildReportGroup_Created(p *CodebuildReportGroupObservation, vals map[string]cty.Value) {
	vals["created"] = cty.StringVal(p.Created)
}