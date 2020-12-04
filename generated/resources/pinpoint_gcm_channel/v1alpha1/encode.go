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

func EncodePinpointGcmChannel(r PinpointGcmChannel) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodePinpointGcmChannel_ApiKey(r.Spec.ForProvider, ctyVal)
	EncodePinpointGcmChannel_ApplicationId(r.Spec.ForProvider, ctyVal)
	EncodePinpointGcmChannel_Enabled(r.Spec.ForProvider, ctyVal)
	EncodePinpointGcmChannel_Id(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodePinpointGcmChannel_ApiKey(p PinpointGcmChannelParameters, vals map[string]cty.Value) {
	vals["api_key"] = cty.StringVal(p.ApiKey)
}

func EncodePinpointGcmChannel_ApplicationId(p PinpointGcmChannelParameters, vals map[string]cty.Value) {
	vals["application_id"] = cty.StringVal(p.ApplicationId)
}

func EncodePinpointGcmChannel_Enabled(p PinpointGcmChannelParameters, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodePinpointGcmChannel_Id(p PinpointGcmChannelParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}