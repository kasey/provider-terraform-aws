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
	r, ok := mr.(*Wafv2IpSet)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeWafv2IpSet(r, ctyValue)
}

func DecodeWafv2IpSet(prev *Wafv2IpSet, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeWafv2IpSet_Addresses(&new.Spec.ForProvider, valMap)
	DecodeWafv2IpSet_Description(&new.Spec.ForProvider, valMap)
	DecodeWafv2IpSet_IpAddressVersion(&new.Spec.ForProvider, valMap)
	DecodeWafv2IpSet_Name(&new.Spec.ForProvider, valMap)
	DecodeWafv2IpSet_Id(&new.Spec.ForProvider, valMap)
	DecodeWafv2IpSet_Scope(&new.Spec.ForProvider, valMap)
	DecodeWafv2IpSet_Tags(&new.Spec.ForProvider, valMap)
	DecodeWafv2IpSet_Arn(&new.Status.AtProvider, valMap)
	DecodeWafv2IpSet_LockToken(&new.Status.AtProvider, valMap)
	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodeWafv2IpSet_Addresses(p *Wafv2IpSetParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["addresses"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.Addresses = goVals
}

func DecodeWafv2IpSet_Description(p *Wafv2IpSetParameters, vals map[string]cty.Value) {
	p.Description = ctwhy.ValueAsString(vals["description"])
}

func DecodeWafv2IpSet_IpAddressVersion(p *Wafv2IpSetParameters, vals map[string]cty.Value) {
	p.IpAddressVersion = ctwhy.ValueAsString(vals["ip_address_version"])
}

func DecodeWafv2IpSet_Name(p *Wafv2IpSetParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

func DecodeWafv2IpSet_Id(p *Wafv2IpSetParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

func DecodeWafv2IpSet_Scope(p *Wafv2IpSetParameters, vals map[string]cty.Value) {
	p.Scope = ctwhy.ValueAsString(vals["scope"])
}

func DecodeWafv2IpSet_Tags(p *Wafv2IpSetParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

func DecodeWafv2IpSet_Arn(p *Wafv2IpSetObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}

func DecodeWafv2IpSet_LockToken(p *Wafv2IpSetObservation, vals map[string]cty.Value) {
	p.LockToken = ctwhy.ValueAsString(vals["lock_token"])
}