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
	r, ok := mr.(*OrganizationsAccount)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a OrganizationsAccount.")
	}
	return EncodeOrganizationsAccount(*r), nil
}

func EncodeOrganizationsAccount(r OrganizationsAccount) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeOrganizationsAccount_Name(r.Spec.ForProvider, ctyVal)
	EncodeOrganizationsAccount_ParentId(r.Spec.ForProvider, ctyVal)
	EncodeOrganizationsAccount_RoleName(r.Spec.ForProvider, ctyVal)
	EncodeOrganizationsAccount_Tags(r.Spec.ForProvider, ctyVal)
	EncodeOrganizationsAccount_Email(r.Spec.ForProvider, ctyVal)
	EncodeOrganizationsAccount_IamUserAccessToBilling(r.Spec.ForProvider, ctyVal)
	EncodeOrganizationsAccount_Status(r.Status.AtProvider, ctyVal)
	EncodeOrganizationsAccount_Arn(r.Status.AtProvider, ctyVal)
	EncodeOrganizationsAccount_JoinedTimestamp(r.Status.AtProvider, ctyVal)
	EncodeOrganizationsAccount_JoinedMethod(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeOrganizationsAccount_Name(p OrganizationsAccountParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeOrganizationsAccount_ParentId(p OrganizationsAccountParameters, vals map[string]cty.Value) {
	vals["parent_id"] = cty.StringVal(p.ParentId)
}

func EncodeOrganizationsAccount_RoleName(p OrganizationsAccountParameters, vals map[string]cty.Value) {
	vals["role_name"] = cty.StringVal(p.RoleName)
}

func EncodeOrganizationsAccount_Tags(p OrganizationsAccountParameters, vals map[string]cty.Value) {
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

func EncodeOrganizationsAccount_Email(p OrganizationsAccountParameters, vals map[string]cty.Value) {
	vals["email"] = cty.StringVal(p.Email)
}

func EncodeOrganizationsAccount_IamUserAccessToBilling(p OrganizationsAccountParameters, vals map[string]cty.Value) {
	vals["iam_user_access_to_billing"] = cty.StringVal(p.IamUserAccessToBilling)
}

func EncodeOrganizationsAccount_Status(p OrganizationsAccountObservation, vals map[string]cty.Value) {
	vals["status"] = cty.StringVal(p.Status)
}

func EncodeOrganizationsAccount_Arn(p OrganizationsAccountObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeOrganizationsAccount_JoinedTimestamp(p OrganizationsAccountObservation, vals map[string]cty.Value) {
	vals["joined_timestamp"] = cty.StringVal(p.JoinedTimestamp)
}

func EncodeOrganizationsAccount_JoinedMethod(p OrganizationsAccountObservation, vals map[string]cty.Value) {
	vals["joined_method"] = cty.StringVal(p.JoinedMethod)
}