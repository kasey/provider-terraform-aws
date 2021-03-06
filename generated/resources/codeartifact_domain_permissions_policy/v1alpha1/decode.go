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
	r, ok := mr.(*CodeartifactDomainPermissionsPolicy)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeCodeartifactDomainPermissionsPolicy(r, ctyValue)
}

func DecodeCodeartifactDomainPermissionsPolicy(prev *CodeartifactDomainPermissionsPolicy, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeCodeartifactDomainPermissionsPolicy_PolicyRevision(&new.Spec.ForProvider, valMap)
	DecodeCodeartifactDomainPermissionsPolicy_Domain(&new.Spec.ForProvider, valMap)
	DecodeCodeartifactDomainPermissionsPolicy_DomainOwner(&new.Spec.ForProvider, valMap)
	DecodeCodeartifactDomainPermissionsPolicy_PolicyDocument(&new.Spec.ForProvider, valMap)
	DecodeCodeartifactDomainPermissionsPolicy_ResourceArn(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeCodeartifactDomainPermissionsPolicy_PolicyRevision(p *CodeartifactDomainPermissionsPolicyParameters, vals map[string]cty.Value) {
	p.PolicyRevision = ctwhy.ValueAsString(vals["policy_revision"])
}

//primitiveTypeDecodeTemplate
func DecodeCodeartifactDomainPermissionsPolicy_Domain(p *CodeartifactDomainPermissionsPolicyParameters, vals map[string]cty.Value) {
	p.Domain = ctwhy.ValueAsString(vals["domain"])
}

//primitiveTypeDecodeTemplate
func DecodeCodeartifactDomainPermissionsPolicy_DomainOwner(p *CodeartifactDomainPermissionsPolicyParameters, vals map[string]cty.Value) {
	p.DomainOwner = ctwhy.ValueAsString(vals["domain_owner"])
}

//primitiveTypeDecodeTemplate
func DecodeCodeartifactDomainPermissionsPolicy_PolicyDocument(p *CodeartifactDomainPermissionsPolicyParameters, vals map[string]cty.Value) {
	p.PolicyDocument = ctwhy.ValueAsString(vals["policy_document"])
}

//primitiveTypeDecodeTemplate
func DecodeCodeartifactDomainPermissionsPolicy_ResourceArn(p *CodeartifactDomainPermissionsPolicyObservation, vals map[string]cty.Value) {
	p.ResourceArn = ctwhy.ValueAsString(vals["resource_arn"])
}