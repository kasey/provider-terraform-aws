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
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane-contrib/terraform-runtime/pkg/plugin"
)

//mergeManagedResourceEntrypointTemplate
type resourceMerger struct{}

func (r *resourceMerger) MergeResources(kube resource.Managed, prov resource.Managed) plugin.MergeDescription {
	k := kube.(*Wafv2IpSet)
	p := prov.(*Wafv2IpSet)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeWafv2IpSet_Addresses(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeWafv2IpSet_Description(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeWafv2IpSet_IpAddressVersion(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeWafv2IpSet_Scope(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeWafv2IpSet_Name(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeWafv2IpSet_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeWafv2IpSet_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeWafv2IpSet_LockToken(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	for key, v := range p.Annotations {
		if k.Annotations[key] != v {
			k.Annotations[key] = v
			md.AnnotationsUpdated = true
		}
	}
	md.AnyFieldUpdated = anyChildUpdated
	return *md
}

//mergePrimitiveContainerTemplateSpec
func MergeWafv2IpSet_Addresses(k *Wafv2IpSetParameters, p *Wafv2IpSetParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.Addresses, p.Addresses) {
		p.Addresses = k.Addresses
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeWafv2IpSet_Description(k *Wafv2IpSetParameters, p *Wafv2IpSetParameters, md *plugin.MergeDescription) bool {
	if k.Description != p.Description {
		p.Description = k.Description
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeWafv2IpSet_IpAddressVersion(k *Wafv2IpSetParameters, p *Wafv2IpSetParameters, md *plugin.MergeDescription) bool {
	if k.IpAddressVersion != p.IpAddressVersion {
		p.IpAddressVersion = k.IpAddressVersion
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeWafv2IpSet_Scope(k *Wafv2IpSetParameters, p *Wafv2IpSetParameters, md *plugin.MergeDescription) bool {
	if k.Scope != p.Scope {
		p.Scope = k.Scope
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeWafv2IpSet_Name(k *Wafv2IpSetParameters, p *Wafv2IpSetParameters, md *plugin.MergeDescription) bool {
	if k.Name != p.Name {
		p.Name = k.Name
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeWafv2IpSet_Tags(k *Wafv2IpSetParameters, p *Wafv2IpSetParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeWafv2IpSet_Arn(k *Wafv2IpSetObservation, p *Wafv2IpSetObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeWafv2IpSet_LockToken(k *Wafv2IpSetObservation, p *Wafv2IpSetObservation, md *plugin.MergeDescription) bool {
	if k.LockToken != p.LockToken {
		k.LockToken = p.LockToken
		md.StatusUpdated = true
		return true
	}
	return false
}