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
	r, ok := mr.(*GuarddutyInviteAccepter)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a GuarddutyInviteAccepter.")
	}
	return EncodeGuarddutyInviteAccepter(*r), nil
}

func EncodeGuarddutyInviteAccepter(r GuarddutyInviteAccepter) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeGuarddutyInviteAccepter_MasterAccountId(r.Spec.ForProvider, ctyVal)
	EncodeGuarddutyInviteAccepter_DetectorId(r.Spec.ForProvider, ctyVal)
	EncodeGuarddutyInviteAccepter_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeGuarddutyInviteAccepter_MasterAccountId(p GuarddutyInviteAccepterParameters, vals map[string]cty.Value) {
	vals["master_account_id"] = cty.StringVal(p.MasterAccountId)
}

func EncodeGuarddutyInviteAccepter_DetectorId(p GuarddutyInviteAccepterParameters, vals map[string]cty.Value) {
	vals["detector_id"] = cty.StringVal(p.DetectorId)
}

func EncodeGuarddutyInviteAccepter_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeGuarddutyInviteAccepter_Timeouts_Create(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeGuarddutyInviteAccepter_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}