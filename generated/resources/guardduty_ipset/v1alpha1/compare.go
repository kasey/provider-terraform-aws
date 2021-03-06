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
	k := kube.(*GuarddutyIpset)
	p := prov.(*GuarddutyIpset)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeGuarddutyIpset_Activate(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeGuarddutyIpset_DetectorId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeGuarddutyIpset_Format(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeGuarddutyIpset_Location(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeGuarddutyIpset_Name(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeGuarddutyIpset_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeGuarddutyIpset_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
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

//mergePrimitiveTemplateSpec
func MergeGuarddutyIpset_Activate(k *GuarddutyIpsetParameters, p *GuarddutyIpsetParameters, md *plugin.MergeDescription) bool {
	if k.Activate != p.Activate {
		p.Activate = k.Activate
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeGuarddutyIpset_DetectorId(k *GuarddutyIpsetParameters, p *GuarddutyIpsetParameters, md *plugin.MergeDescription) bool {
	if k.DetectorId != p.DetectorId {
		p.DetectorId = k.DetectorId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeGuarddutyIpset_Format(k *GuarddutyIpsetParameters, p *GuarddutyIpsetParameters, md *plugin.MergeDescription) bool {
	if k.Format != p.Format {
		p.Format = k.Format
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeGuarddutyIpset_Location(k *GuarddutyIpsetParameters, p *GuarddutyIpsetParameters, md *plugin.MergeDescription) bool {
	if k.Location != p.Location {
		p.Location = k.Location
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeGuarddutyIpset_Name(k *GuarddutyIpsetParameters, p *GuarddutyIpsetParameters, md *plugin.MergeDescription) bool {
	if k.Name != p.Name {
		p.Name = k.Name
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeGuarddutyIpset_Tags(k *GuarddutyIpsetParameters, p *GuarddutyIpsetParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeGuarddutyIpset_Arn(k *GuarddutyIpsetObservation, p *GuarddutyIpsetObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}