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
	r, ok := mr.(*CustomerGateway)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a CustomerGateway.")
	}
	return EncodeCustomerGateway(*r), nil
}

func EncodeCustomerGateway(r CustomerGateway) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeCustomerGateway_BgpAsn(r.Spec.ForProvider, ctyVal)
	EncodeCustomerGateway_IpAddress(r.Spec.ForProvider, ctyVal)
	EncodeCustomerGateway_Tags(r.Spec.ForProvider, ctyVal)
	EncodeCustomerGateway_Type(r.Spec.ForProvider, ctyVal)
	EncodeCustomerGateway_Arn(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeCustomerGateway_BgpAsn(p CustomerGatewayParameters, vals map[string]cty.Value) {
	vals["bgp_asn"] = cty.StringVal(p.BgpAsn)
}

func EncodeCustomerGateway_IpAddress(p CustomerGatewayParameters, vals map[string]cty.Value) {
	vals["ip_address"] = cty.StringVal(p.IpAddress)
}

func EncodeCustomerGateway_Tags(p CustomerGatewayParameters, vals map[string]cty.Value) {
	if len(p.Tags) == 0 {
		vals["tags"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeCustomerGateway_Type(p CustomerGatewayParameters, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeCustomerGateway_Arn(p CustomerGatewayObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}