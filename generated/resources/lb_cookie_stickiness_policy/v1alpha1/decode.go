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
	r, ok := mr.(*LbCookieStickinessPolicy)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeLbCookieStickinessPolicy(r, ctyValue)
}

func DecodeLbCookieStickinessPolicy(prev *LbCookieStickinessPolicy, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeLbCookieStickinessPolicy_LbPort(&new.Spec.ForProvider, valMap)
	DecodeLbCookieStickinessPolicy_LoadBalancer(&new.Spec.ForProvider, valMap)
	DecodeLbCookieStickinessPolicy_Name(&new.Spec.ForProvider, valMap)
	DecodeLbCookieStickinessPolicy_CookieExpirationPeriod(&new.Spec.ForProvider, valMap)
	DecodeLbCookieStickinessPolicy_Id(&new.Spec.ForProvider, valMap)

	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodeLbCookieStickinessPolicy_LbPort(p *LbCookieStickinessPolicyParameters, vals map[string]cty.Value) {
	p.LbPort = ctwhy.ValueAsInt64(vals["lb_port"])
}

func DecodeLbCookieStickinessPolicy_LoadBalancer(p *LbCookieStickinessPolicyParameters, vals map[string]cty.Value) {
	p.LoadBalancer = ctwhy.ValueAsString(vals["load_balancer"])
}

func DecodeLbCookieStickinessPolicy_Name(p *LbCookieStickinessPolicyParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

func DecodeLbCookieStickinessPolicy_CookieExpirationPeriod(p *LbCookieStickinessPolicyParameters, vals map[string]cty.Value) {
	p.CookieExpirationPeriod = ctwhy.ValueAsInt64(vals["cookie_expiration_period"])
}

func DecodeLbCookieStickinessPolicy_Id(p *LbCookieStickinessPolicyParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}