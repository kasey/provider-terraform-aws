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

func EncodeBatchJobQueue(r BatchJobQueue) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeBatchJobQueue_Id(r.Spec.ForProvider, ctyVal)
	EncodeBatchJobQueue_Name(r.Spec.ForProvider, ctyVal)
	EncodeBatchJobQueue_Priority(r.Spec.ForProvider, ctyVal)
	EncodeBatchJobQueue_State(r.Spec.ForProvider, ctyVal)
	EncodeBatchJobQueue_ComputeEnvironments(r.Spec.ForProvider, ctyVal)
	EncodeBatchJobQueue_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeBatchJobQueue_Id(p BatchJobQueueParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeBatchJobQueue_Name(p BatchJobQueueParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeBatchJobQueue_Priority(p BatchJobQueueParameters, vals map[string]cty.Value) {
	vals["priority"] = cty.NumberIntVal(p.Priority)
}

func EncodeBatchJobQueue_State(p BatchJobQueueParameters, vals map[string]cty.Value) {
	vals["state"] = cty.StringVal(p.State)
}

func EncodeBatchJobQueue_ComputeEnvironments(p BatchJobQueueParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.ComputeEnvironments {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["compute_environments"] = cty.ListVal(colVals)
}

func EncodeBatchJobQueue_Arn(p BatchJobQueueObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}