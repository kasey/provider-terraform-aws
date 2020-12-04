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

func EncodePinpointApp(r PinpointApp) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodePinpointApp_Id(r.Spec.ForProvider, ctyVal)
	EncodePinpointApp_Name(r.Spec.ForProvider, ctyVal)
	EncodePinpointApp_NamePrefix(r.Spec.ForProvider, ctyVal)
	EncodePinpointApp_Tags(r.Spec.ForProvider, ctyVal)
	EncodePinpointApp_CampaignHook(r.Spec.ForProvider.CampaignHook, ctyVal)
	EncodePinpointApp_Limits(r.Spec.ForProvider.Limits, ctyVal)
	EncodePinpointApp_QuietTime(r.Spec.ForProvider.QuietTime, ctyVal)
	EncodePinpointApp_ApplicationId(r.Status.AtProvider, ctyVal)
	EncodePinpointApp_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodePinpointApp_Id(p PinpointAppParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodePinpointApp_Name(p PinpointAppParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodePinpointApp_NamePrefix(p PinpointAppParameters, vals map[string]cty.Value) {
	vals["name_prefix"] = cty.StringVal(p.NamePrefix)
}

func EncodePinpointApp_Tags(p PinpointAppParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodePinpointApp_CampaignHook(p CampaignHook, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodePinpointApp_CampaignHook_LambdaFunctionName(p, ctyVal)
	EncodePinpointApp_CampaignHook_Mode(p, ctyVal)
	EncodePinpointApp_CampaignHook_WebUrl(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["campaign_hook"] = cty.ListVal(valsForCollection)
}

func EncodePinpointApp_CampaignHook_LambdaFunctionName(p CampaignHook, vals map[string]cty.Value) {
	vals["lambda_function_name"] = cty.StringVal(p.LambdaFunctionName)
}

func EncodePinpointApp_CampaignHook_Mode(p CampaignHook, vals map[string]cty.Value) {
	vals["mode"] = cty.StringVal(p.Mode)
}

func EncodePinpointApp_CampaignHook_WebUrl(p CampaignHook, vals map[string]cty.Value) {
	vals["web_url"] = cty.StringVal(p.WebUrl)
}

func EncodePinpointApp_Limits(p Limits, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodePinpointApp_Limits_MessagesPerSecond(p, ctyVal)
	EncodePinpointApp_Limits_Total(p, ctyVal)
	EncodePinpointApp_Limits_Daily(p, ctyVal)
	EncodePinpointApp_Limits_MaximumDuration(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["limits"] = cty.ListVal(valsForCollection)
}

func EncodePinpointApp_Limits_MessagesPerSecond(p Limits, vals map[string]cty.Value) {
	vals["messages_per_second"] = cty.NumberIntVal(p.MessagesPerSecond)
}

func EncodePinpointApp_Limits_Total(p Limits, vals map[string]cty.Value) {
	vals["total"] = cty.NumberIntVal(p.Total)
}

func EncodePinpointApp_Limits_Daily(p Limits, vals map[string]cty.Value) {
	vals["daily"] = cty.NumberIntVal(p.Daily)
}

func EncodePinpointApp_Limits_MaximumDuration(p Limits, vals map[string]cty.Value) {
	vals["maximum_duration"] = cty.NumberIntVal(p.MaximumDuration)
}

func EncodePinpointApp_QuietTime(p QuietTime, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodePinpointApp_QuietTime_End(p, ctyVal)
	EncodePinpointApp_QuietTime_Start(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["quiet_time"] = cty.ListVal(valsForCollection)
}

func EncodePinpointApp_QuietTime_End(p QuietTime, vals map[string]cty.Value) {
	vals["end"] = cty.StringVal(p.End)
}

func EncodePinpointApp_QuietTime_Start(p QuietTime, vals map[string]cty.Value) {
	vals["start"] = cty.StringVal(p.Start)
}

func EncodePinpointApp_ApplicationId(p PinpointAppObservation, vals map[string]cty.Value) {
	vals["application_id"] = cty.StringVal(p.ApplicationId)
}

func EncodePinpointApp_Arn(p PinpointAppObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}