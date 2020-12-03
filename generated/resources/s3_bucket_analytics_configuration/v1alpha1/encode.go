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

package v1alpha1func EncodeS3BucketAnalyticsConfiguration(r S3BucketAnalyticsConfiguration) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeS3BucketAnalyticsConfiguration_Bucket(r.Spec.ForProvider, ctyVal)
	EncodeS3BucketAnalyticsConfiguration_Id(r.Spec.ForProvider, ctyVal)
	EncodeS3BucketAnalyticsConfiguration_Name(r.Spec.ForProvider, ctyVal)
	EncodeS3BucketAnalyticsConfiguration_Filter(r.Spec.ForProvider.Filter, ctyVal)
	EncodeS3BucketAnalyticsConfiguration_StorageClassAnalysis(r.Spec.ForProvider.StorageClassAnalysis, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeS3BucketAnalyticsConfiguration_Bucket(p *S3BucketAnalyticsConfigurationParameters, vals map[string]cty.Value) {
	vals["bucket"] = cty.StringVal(p.Bucket)
}

func EncodeS3BucketAnalyticsConfiguration_Id(p *S3BucketAnalyticsConfigurationParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeS3BucketAnalyticsConfiguration_Name(p *S3BucketAnalyticsConfigurationParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeS3BucketAnalyticsConfiguration_Filter(p *Filter, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Filter {
		ctyVal = make(map[string]cty.Value)
		EncodeS3BucketAnalyticsConfiguration_Filter_Prefix(v, ctyVal)
		EncodeS3BucketAnalyticsConfiguration_Filter_Tags(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["filter"] = cty.ListVal(valsForCollection)
}

func EncodeS3BucketAnalyticsConfiguration_Filter_Prefix(p *Filter, vals map[string]cty.Value) {
	vals["prefix"] = cty.StringVal(p.Prefix)
}

func EncodeS3BucketAnalyticsConfiguration_Filter_Tags(p *Filter, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeS3BucketAnalyticsConfiguration_StorageClassAnalysis(p *StorageClassAnalysis, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.StorageClassAnalysis {
		ctyVal = make(map[string]cty.Value)
		EncodeS3BucketAnalyticsConfiguration_StorageClassAnalysis_DataExport(v.DataExport, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["storage_class_analysis"] = cty.ListVal(valsForCollection)
}

func EncodeS3BucketAnalyticsConfiguration_StorageClassAnalysis_DataExport(p *DataExport, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.DataExport {
		ctyVal = make(map[string]cty.Value)
		EncodeS3BucketAnalyticsConfiguration_StorageClassAnalysis_DataExport_OutputSchemaVersion(v, ctyVal)
		EncodeS3BucketAnalyticsConfiguration_StorageClassAnalysis_DataExport_Destination(v.Destination, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["data_export"] = cty.ListVal(valsForCollection)
}

func EncodeS3BucketAnalyticsConfiguration_StorageClassAnalysis_DataExport_OutputSchemaVersion(p *DataExport, vals map[string]cty.Value) {
	vals["output_schema_version"] = cty.StringVal(p.OutputSchemaVersion)
}

func EncodeS3BucketAnalyticsConfiguration_StorageClassAnalysis_DataExport_Destination(p *Destination, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Destination {
		ctyVal = make(map[string]cty.Value)
		EncodeS3BucketAnalyticsConfiguration_StorageClassAnalysis_DataExport_Destination_S3BucketDestination(v.S3BucketDestination, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["destination"] = cty.ListVal(valsForCollection)
}

func EncodeS3BucketAnalyticsConfiguration_StorageClassAnalysis_DataExport_Destination_S3BucketDestination(p *S3BucketDestination, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.S3BucketDestination {
		ctyVal = make(map[string]cty.Value)
		EncodeS3BucketAnalyticsConfiguration_StorageClassAnalysis_DataExport_Destination_S3BucketDestination_BucketAccountId(v, ctyVal)
		EncodeS3BucketAnalyticsConfiguration_StorageClassAnalysis_DataExport_Destination_S3BucketDestination_BucketArn(v, ctyVal)
		EncodeS3BucketAnalyticsConfiguration_StorageClassAnalysis_DataExport_Destination_S3BucketDestination_Format(v, ctyVal)
		EncodeS3BucketAnalyticsConfiguration_StorageClassAnalysis_DataExport_Destination_S3BucketDestination_Prefix(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["s3_bucket_destination"] = cty.ListVal(valsForCollection)
}

func EncodeS3BucketAnalyticsConfiguration_StorageClassAnalysis_DataExport_Destination_S3BucketDestination_BucketAccountId(p *S3BucketDestination, vals map[string]cty.Value) {
	vals["bucket_account_id"] = cty.StringVal(p.BucketAccountId)
}

func EncodeS3BucketAnalyticsConfiguration_StorageClassAnalysis_DataExport_Destination_S3BucketDestination_BucketArn(p *S3BucketDestination, vals map[string]cty.Value) {
	vals["bucket_arn"] = cty.StringVal(p.BucketArn)
}

func EncodeS3BucketAnalyticsConfiguration_StorageClassAnalysis_DataExport_Destination_S3BucketDestination_Format(p *S3BucketDestination, vals map[string]cty.Value) {
	vals["format"] = cty.StringVal(p.Format)
}

func EncodeS3BucketAnalyticsConfiguration_StorageClassAnalysis_DataExport_Destination_S3BucketDestination_Prefix(p *S3BucketDestination, vals map[string]cty.Value) {
	vals["prefix"] = cty.StringVal(p.Prefix)
}