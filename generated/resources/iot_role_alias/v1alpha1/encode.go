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
	r, ok := mr.(*IotRoleAlias)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a IotRoleAlias.")
	}
	return EncodeIotRoleAlias(*r), nil
}

func EncodeIotRoleAlias(r IotRoleAlias) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeIotRoleAlias_CredentialDuration(r.Spec.ForProvider, ctyVal)
	EncodeIotRoleAlias_RoleArn(r.Spec.ForProvider, ctyVal)
	EncodeIotRoleAlias_Alias(r.Spec.ForProvider, ctyVal)
	EncodeIotRoleAlias_Arn(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeIotRoleAlias_CredentialDuration(p IotRoleAliasParameters, vals map[string]cty.Value) {
	vals["credential_duration"] = cty.NumberIntVal(p.CredentialDuration)
}

func EncodeIotRoleAlias_RoleArn(p IotRoleAliasParameters, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeIotRoleAlias_Alias(p IotRoleAliasParameters, vals map[string]cty.Value) {
	vals["alias"] = cty.StringVal(p.Alias)
}

func EncodeIotRoleAlias_Arn(p IotRoleAliasObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}