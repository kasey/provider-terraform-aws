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

	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
	"github.com/zclconf/go-cty/cty"
	ctwhy "github.com/crossplane-contrib/terraform-runtime/pkg/plugin/cty"
)

type ctyDecoder struct{}

func (e *ctyDecoder) DecodeCty(mr resource.Managed, ctyValue cty.Value, schema *providers.Schema) (resource.Managed, error) {
	r, ok := mr.(*PinpointApnsSandboxChannel)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodePinpointApnsSandboxChannel(r, ctyValue)
}

func DecodePinpointApnsSandboxChannel(prev *PinpointApnsSandboxChannel, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodePinpointApnsSandboxChannel_BundleId(&new.Spec.ForProvider, valMap)
	DecodePinpointApnsSandboxChannel_DefaultAuthenticationMethod(&new.Spec.ForProvider, valMap)
	DecodePinpointApnsSandboxChannel_Id(&new.Spec.ForProvider, valMap)
	DecodePinpointApnsSandboxChannel_ApplicationId(&new.Spec.ForProvider, valMap)
	DecodePinpointApnsSandboxChannel_Enabled(&new.Spec.ForProvider, valMap)
	DecodePinpointApnsSandboxChannel_PrivateKey(&new.Spec.ForProvider, valMap)
	DecodePinpointApnsSandboxChannel_TeamId(&new.Spec.ForProvider, valMap)
	DecodePinpointApnsSandboxChannel_TokenKey(&new.Spec.ForProvider, valMap)
	DecodePinpointApnsSandboxChannel_TokenKeyId(&new.Spec.ForProvider, valMap)
	DecodePinpointApnsSandboxChannel_Certificate(&new.Spec.ForProvider, valMap)

	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodePinpointApnsSandboxChannel_BundleId(p *PinpointApnsSandboxChannelParameters, vals map[string]cty.Value) {
	p.BundleId = ctwhy.ValueAsString(vals["bundle_id"])
}

func DecodePinpointApnsSandboxChannel_DefaultAuthenticationMethod(p *PinpointApnsSandboxChannelParameters, vals map[string]cty.Value) {
	p.DefaultAuthenticationMethod = ctwhy.ValueAsString(vals["default_authentication_method"])
}

func DecodePinpointApnsSandboxChannel_Id(p *PinpointApnsSandboxChannelParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

func DecodePinpointApnsSandboxChannel_ApplicationId(p *PinpointApnsSandboxChannelParameters, vals map[string]cty.Value) {
	p.ApplicationId = ctwhy.ValueAsString(vals["application_id"])
}

func DecodePinpointApnsSandboxChannel_Enabled(p *PinpointApnsSandboxChannelParameters, vals map[string]cty.Value) {
	p.Enabled = ctwhy.ValueAsBool(vals["enabled"])
}

func DecodePinpointApnsSandboxChannel_PrivateKey(p *PinpointApnsSandboxChannelParameters, vals map[string]cty.Value) {
	p.PrivateKey = ctwhy.ValueAsString(vals["private_key"])
}

func DecodePinpointApnsSandboxChannel_TeamId(p *PinpointApnsSandboxChannelParameters, vals map[string]cty.Value) {
	p.TeamId = ctwhy.ValueAsString(vals["team_id"])
}

func DecodePinpointApnsSandboxChannel_TokenKey(p *PinpointApnsSandboxChannelParameters, vals map[string]cty.Value) {
	p.TokenKey = ctwhy.ValueAsString(vals["token_key"])
}

func DecodePinpointApnsSandboxChannel_TokenKeyId(p *PinpointApnsSandboxChannelParameters, vals map[string]cty.Value) {
	p.TokenKeyId = ctwhy.ValueAsString(vals["token_key_id"])
}

func DecodePinpointApnsSandboxChannel_Certificate(p *PinpointApnsSandboxChannelParameters, vals map[string]cty.Value) {
	p.Certificate = ctwhy.ValueAsString(vals["certificate"])
}