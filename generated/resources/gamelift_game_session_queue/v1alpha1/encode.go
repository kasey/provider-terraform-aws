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

package v1alpha1func EncodeGameliftGameSessionQueue(r GameliftGameSessionQueue) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeGameliftGameSessionQueue_Name(r.Spec.ForProvider, ctyVal)
	EncodeGameliftGameSessionQueue_Tags(r.Spec.ForProvider, ctyVal)
	EncodeGameliftGameSessionQueue_TimeoutInSeconds(r.Spec.ForProvider, ctyVal)
	EncodeGameliftGameSessionQueue_Destinations(r.Spec.ForProvider, ctyVal)
	EncodeGameliftGameSessionQueue_Id(r.Spec.ForProvider, ctyVal)
	EncodeGameliftGameSessionQueue_PlayerLatencyPolicy(r.Spec.ForProvider.PlayerLatencyPolicy, ctyVal)
	EncodeGameliftGameSessionQueue_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeGameliftGameSessionQueue_Name(p *GameliftGameSessionQueueParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeGameliftGameSessionQueue_Tags(p *GameliftGameSessionQueueParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeGameliftGameSessionQueue_TimeoutInSeconds(p *GameliftGameSessionQueueParameters, vals map[string]cty.Value) {
	vals["timeout_in_seconds"] = cty.IntVal(p.TimeoutInSeconds)
}

func EncodeGameliftGameSessionQueue_Destinations(p *GameliftGameSessionQueueParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Destinations {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["destinations"] = cty.ListVal(colVals)
}

func EncodeGameliftGameSessionQueue_Id(p *GameliftGameSessionQueueParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeGameliftGameSessionQueue_PlayerLatencyPolicy(p *PlayerLatencyPolicy, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.PlayerLatencyPolicy {
		ctyVal = make(map[string]cty.Value)
		EncodeGameliftGameSessionQueue_PlayerLatencyPolicy_PolicyDurationSeconds(v, ctyVal)
		EncodeGameliftGameSessionQueue_PlayerLatencyPolicy_MaximumIndividualPlayerLatencyMilliseconds(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["player_latency_policy"] = cty.ListVal(valsForCollection)
}

func EncodeGameliftGameSessionQueue_PlayerLatencyPolicy_PolicyDurationSeconds(p *PlayerLatencyPolicy, vals map[string]cty.Value) {
	vals["policy_duration_seconds"] = cty.IntVal(p.PolicyDurationSeconds)
}

func EncodeGameliftGameSessionQueue_PlayerLatencyPolicy_MaximumIndividualPlayerLatencyMilliseconds(p *PlayerLatencyPolicy, vals map[string]cty.Value) {
	vals["maximum_individual_player_latency_milliseconds"] = cty.IntVal(p.MaximumIndividualPlayerLatencyMilliseconds)
}

func EncodeGameliftGameSessionQueue_Arn(p *GameliftGameSessionQueueObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}