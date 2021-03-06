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
	r, ok := mr.(*KeyPair)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeKeyPair(r, ctyValue)
}

func DecodeKeyPair(prev *KeyPair, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeKeyPair_PublicKey(&new.Spec.ForProvider, valMap)
	DecodeKeyPair_Tags(&new.Spec.ForProvider, valMap)
	DecodeKeyPair_KeyName(&new.Spec.ForProvider, valMap)
	DecodeKeyPair_KeyNamePrefix(&new.Spec.ForProvider, valMap)
	DecodeKeyPair_KeyPairId(&new.Status.AtProvider, valMap)
	DecodeKeyPair_Arn(&new.Status.AtProvider, valMap)
	DecodeKeyPair_Fingerprint(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeKeyPair_PublicKey(p *KeyPairParameters, vals map[string]cty.Value) {
	p.PublicKey = ctwhy.ValueAsString(vals["public_key"])
}

//primitiveMapTypeDecodeTemplate
func DecodeKeyPair_Tags(p *KeyPairParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

//primitiveTypeDecodeTemplate
func DecodeKeyPair_KeyName(p *KeyPairParameters, vals map[string]cty.Value) {
	p.KeyName = ctwhy.ValueAsString(vals["key_name"])
}

//primitiveTypeDecodeTemplate
func DecodeKeyPair_KeyNamePrefix(p *KeyPairParameters, vals map[string]cty.Value) {
	p.KeyNamePrefix = ctwhy.ValueAsString(vals["key_name_prefix"])
}

//primitiveTypeDecodeTemplate
func DecodeKeyPair_KeyPairId(p *KeyPairObservation, vals map[string]cty.Value) {
	p.KeyPairId = ctwhy.ValueAsString(vals["key_pair_id"])
}

//primitiveTypeDecodeTemplate
func DecodeKeyPair_Arn(p *KeyPairObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}

//primitiveTypeDecodeTemplate
func DecodeKeyPair_Fingerprint(p *KeyPairObservation, vals map[string]cty.Value) {
	p.Fingerprint = ctwhy.ValueAsString(vals["fingerprint"])
}