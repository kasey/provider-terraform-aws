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

package v1alpha1func EncodeMskConfiguration(r MskConfiguration) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeMskConfiguration_ServerProperties(r.Spec.ForProvider, ctyVal)
	EncodeMskConfiguration_Description(r.Spec.ForProvider, ctyVal)
	EncodeMskConfiguration_Id(r.Spec.ForProvider, ctyVal)
	EncodeMskConfiguration_KafkaVersions(r.Spec.ForProvider, ctyVal)
	EncodeMskConfiguration_Name(r.Spec.ForProvider, ctyVal)
	EncodeMskConfiguration_Arn(r.Status.AtProvider, ctyVal)
	EncodeMskConfiguration_LatestRevision(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeMskConfiguration_ServerProperties(p *MskConfigurationParameters, vals map[string]cty.Value) {
	vals["server_properties"] = cty.StringVal(p.ServerProperties)
}

func EncodeMskConfiguration_Description(p *MskConfigurationParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeMskConfiguration_Id(p *MskConfigurationParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeMskConfiguration_KafkaVersions(p *MskConfigurationParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.KafkaVersions {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["kafka_versions"] = cty.SetVal(colVals)
}

func EncodeMskConfiguration_Name(p *MskConfigurationParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeMskConfiguration_Arn(p *MskConfigurationObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeMskConfiguration_LatestRevision(p *MskConfigurationObservation, vals map[string]cty.Value) {
	vals["latest_revision"] = cty.IntVal(p.LatestRevision)
}