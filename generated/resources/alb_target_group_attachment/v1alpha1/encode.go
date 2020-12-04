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

func EncodeAlbTargetGroupAttachment(r AlbTargetGroupAttachment) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeAlbTargetGroupAttachment_Id(r.Spec.ForProvider, ctyVal)
	EncodeAlbTargetGroupAttachment_Port(r.Spec.ForProvider, ctyVal)
	EncodeAlbTargetGroupAttachment_TargetGroupArn(r.Spec.ForProvider, ctyVal)
	EncodeAlbTargetGroupAttachment_TargetId(r.Spec.ForProvider, ctyVal)
	EncodeAlbTargetGroupAttachment_AvailabilityZone(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeAlbTargetGroupAttachment_Id(p AlbTargetGroupAttachmentParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeAlbTargetGroupAttachment_Port(p AlbTargetGroupAttachmentParameters, vals map[string]cty.Value) {
	vals["port"] = cty.NumberIntVal(p.Port)
}

func EncodeAlbTargetGroupAttachment_TargetGroupArn(p AlbTargetGroupAttachmentParameters, vals map[string]cty.Value) {
	vals["target_group_arn"] = cty.StringVal(p.TargetGroupArn)
}

func EncodeAlbTargetGroupAttachment_TargetId(p AlbTargetGroupAttachmentParameters, vals map[string]cty.Value) {
	vals["target_id"] = cty.StringVal(p.TargetId)
}

func EncodeAlbTargetGroupAttachment_AvailabilityZone(p AlbTargetGroupAttachmentParameters, vals map[string]cty.Value) {
	vals["availability_zone"] = cty.StringVal(p.AvailabilityZone)
}