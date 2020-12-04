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

func EncodeWafregionalWebAclAssociation(r WafregionalWebAclAssociation) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeWafregionalWebAclAssociation_Id(r.Spec.ForProvider, ctyVal)
	EncodeWafregionalWebAclAssociation_ResourceArn(r.Spec.ForProvider, ctyVal)
	EncodeWafregionalWebAclAssociation_WebAclId(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeWafregionalWebAclAssociation_Id(p WafregionalWebAclAssociationParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeWafregionalWebAclAssociation_ResourceArn(p WafregionalWebAclAssociationParameters, vals map[string]cty.Value) {
	vals["resource_arn"] = cty.StringVal(p.ResourceArn)
}

func EncodeWafregionalWebAclAssociation_WebAclId(p WafregionalWebAclAssociationParameters, vals map[string]cty.Value) {
	vals["web_acl_id"] = cty.StringVal(p.WebAclId)
}