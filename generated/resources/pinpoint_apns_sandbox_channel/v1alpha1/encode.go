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
	"fmt"

	"github.com/zclconf/go-cty/cty"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
)

type ctyEncoder struct{}

func (e *ctyEncoder) EncodeCty(mr resource.Managed, schema *providers.Schema) (cty.Value, error) {
	r, ok := mr.(*PinpointApnsSandboxChannel)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a PinpointApnsSandboxChannel.")
	}
	return EncodePinpointApnsSandboxChannel(*r), nil
}

func EncodePinpointApnsSandboxChannel(r PinpointApnsSandboxChannel) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodePinpointApnsSandboxChannel_PrivateKey(r.Spec.ForProvider, ctyVal)
	EncodePinpointApnsSandboxChannel_TokenKey(r.Spec.ForProvider, ctyVal)
	EncodePinpointApnsSandboxChannel_TokenKeyId(r.Spec.ForProvider, ctyVal)
	EncodePinpointApnsSandboxChannel_ApplicationId(r.Spec.ForProvider, ctyVal)
	EncodePinpointApnsSandboxChannel_BundleId(r.Spec.ForProvider, ctyVal)
	EncodePinpointApnsSandboxChannel_Enabled(r.Spec.ForProvider, ctyVal)
	EncodePinpointApnsSandboxChannel_TeamId(r.Spec.ForProvider, ctyVal)
	EncodePinpointApnsSandboxChannel_Certificate(r.Spec.ForProvider, ctyVal)
	EncodePinpointApnsSandboxChannel_DefaultAuthenticationMethod(r.Spec.ForProvider, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodePinpointApnsSandboxChannel_PrivateKey(p PinpointApnsSandboxChannelParameters, vals map[string]cty.Value) {
	vals["private_key"] = cty.StringVal(p.PrivateKey)
}

func EncodePinpointApnsSandboxChannel_TokenKey(p PinpointApnsSandboxChannelParameters, vals map[string]cty.Value) {
	vals["token_key"] = cty.StringVal(p.TokenKey)
}

func EncodePinpointApnsSandboxChannel_TokenKeyId(p PinpointApnsSandboxChannelParameters, vals map[string]cty.Value) {
	vals["token_key_id"] = cty.StringVal(p.TokenKeyId)
}

func EncodePinpointApnsSandboxChannel_ApplicationId(p PinpointApnsSandboxChannelParameters, vals map[string]cty.Value) {
	vals["application_id"] = cty.StringVal(p.ApplicationId)
}

func EncodePinpointApnsSandboxChannel_BundleId(p PinpointApnsSandboxChannelParameters, vals map[string]cty.Value) {
	vals["bundle_id"] = cty.StringVal(p.BundleId)
}

func EncodePinpointApnsSandboxChannel_Enabled(p PinpointApnsSandboxChannelParameters, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodePinpointApnsSandboxChannel_TeamId(p PinpointApnsSandboxChannelParameters, vals map[string]cty.Value) {
	vals["team_id"] = cty.StringVal(p.TeamId)
}

func EncodePinpointApnsSandboxChannel_Certificate(p PinpointApnsSandboxChannelParameters, vals map[string]cty.Value) {
	vals["certificate"] = cty.StringVal(p.Certificate)
}

func EncodePinpointApnsSandboxChannel_DefaultAuthenticationMethod(p PinpointApnsSandboxChannelParameters, vals map[string]cty.Value) {
	vals["default_authentication_method"] = cty.StringVal(p.DefaultAuthenticationMethod)
}