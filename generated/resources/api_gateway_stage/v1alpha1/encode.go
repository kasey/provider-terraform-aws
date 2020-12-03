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

package v1alpha1func EncodeApiGatewayStage(r ApiGatewayStage) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeApiGatewayStage_Id(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayStage_DeploymentId(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayStage_Description(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayStage_DocumentationVersion(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayStage_ClientCertificateId(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayStage_StageName(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayStage_CacheClusterEnabled(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayStage_CacheClusterSize(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayStage_Variables(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayStage_XrayTracingEnabled(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayStage_RestApiId(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayStage_Tags(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayStage_AccessLogSettings(r.Spec.ForProvider.AccessLogSettings, ctyVal)
	EncodeApiGatewayStage_InvokeUrl(r.Status.AtProvider, ctyVal)
	EncodeApiGatewayStage_ExecutionArn(r.Status.AtProvider, ctyVal)
	EncodeApiGatewayStage_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeApiGatewayStage_Id(p *ApiGatewayStageParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeApiGatewayStage_DeploymentId(p *ApiGatewayStageParameters, vals map[string]cty.Value) {
	vals["deployment_id"] = cty.StringVal(p.DeploymentId)
}

func EncodeApiGatewayStage_Description(p *ApiGatewayStageParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeApiGatewayStage_DocumentationVersion(p *ApiGatewayStageParameters, vals map[string]cty.Value) {
	vals["documentation_version"] = cty.StringVal(p.DocumentationVersion)
}

func EncodeApiGatewayStage_ClientCertificateId(p *ApiGatewayStageParameters, vals map[string]cty.Value) {
	vals["client_certificate_id"] = cty.StringVal(p.ClientCertificateId)
}

func EncodeApiGatewayStage_StageName(p *ApiGatewayStageParameters, vals map[string]cty.Value) {
	vals["stage_name"] = cty.StringVal(p.StageName)
}

func EncodeApiGatewayStage_CacheClusterEnabled(p *ApiGatewayStageParameters, vals map[string]cty.Value) {
	vals["cache_cluster_enabled"] = cty.BoolVal(p.CacheClusterEnabled)
}

func EncodeApiGatewayStage_CacheClusterSize(p *ApiGatewayStageParameters, vals map[string]cty.Value) {
	vals["cache_cluster_size"] = cty.StringVal(p.CacheClusterSize)
}

func EncodeApiGatewayStage_Variables(p *ApiGatewayStageParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Variables {
		mVals[key] = cty.StringVal(value)
	}
	vals["variables"] = cty.MapVal(mVals)
}

func EncodeApiGatewayStage_XrayTracingEnabled(p *ApiGatewayStageParameters, vals map[string]cty.Value) {
	vals["xray_tracing_enabled"] = cty.BoolVal(p.XrayTracingEnabled)
}

func EncodeApiGatewayStage_RestApiId(p *ApiGatewayStageParameters, vals map[string]cty.Value) {
	vals["rest_api_id"] = cty.StringVal(p.RestApiId)
}

func EncodeApiGatewayStage_Tags(p *ApiGatewayStageParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeApiGatewayStage_AccessLogSettings(p *AccessLogSettings, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AccessLogSettings {
		ctyVal = make(map[string]cty.Value)
		EncodeApiGatewayStage_AccessLogSettings_DestinationArn(v, ctyVal)
		EncodeApiGatewayStage_AccessLogSettings_Format(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["access_log_settings"] = cty.ListVal(valsForCollection)
}

func EncodeApiGatewayStage_AccessLogSettings_DestinationArn(p *AccessLogSettings, vals map[string]cty.Value) {
	vals["destination_arn"] = cty.StringVal(p.DestinationArn)
}

func EncodeApiGatewayStage_AccessLogSettings_Format(p *AccessLogSettings, vals map[string]cty.Value) {
	vals["format"] = cty.StringVal(p.Format)
}

func EncodeApiGatewayStage_InvokeUrl(p *ApiGatewayStageObservation, vals map[string]cty.Value) {
	vals["invoke_url"] = cty.StringVal(p.InvokeUrl)
}

func EncodeApiGatewayStage_ExecutionArn(p *ApiGatewayStageObservation, vals map[string]cty.Value) {
	vals["execution_arn"] = cty.StringVal(p.ExecutionArn)
}

func EncodeApiGatewayStage_Arn(p *ApiGatewayStageObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}