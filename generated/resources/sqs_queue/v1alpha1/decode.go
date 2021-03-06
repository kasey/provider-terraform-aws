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
	r, ok := mr.(*SqsQueue)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeSqsQueue(r, ctyValue)
}

func DecodeSqsQueue(prev *SqsQueue, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeSqsQueue_MessageRetentionSeconds(&new.Spec.ForProvider, valMap)
	DecodeSqsQueue_Name(&new.Spec.ForProvider, valMap)
	DecodeSqsQueue_NamePrefix(&new.Spec.ForProvider, valMap)
	DecodeSqsQueue_ReceiveWaitTimeSeconds(&new.Spec.ForProvider, valMap)
	DecodeSqsQueue_DelaySeconds(&new.Spec.ForProvider, valMap)
	DecodeSqsQueue_MaxMessageSize(&new.Spec.ForProvider, valMap)
	DecodeSqsQueue_VisibilityTimeoutSeconds(&new.Spec.ForProvider, valMap)
	DecodeSqsQueue_FifoQueue(&new.Spec.ForProvider, valMap)
	DecodeSqsQueue_RedrivePolicy(&new.Spec.ForProvider, valMap)
	DecodeSqsQueue_KmsDataKeyReusePeriodSeconds(&new.Spec.ForProvider, valMap)
	DecodeSqsQueue_KmsMasterKeyId(&new.Spec.ForProvider, valMap)
	DecodeSqsQueue_Policy(&new.Spec.ForProvider, valMap)
	DecodeSqsQueue_Tags(&new.Spec.ForProvider, valMap)
	DecodeSqsQueue_ContentBasedDeduplication(&new.Spec.ForProvider, valMap)
	DecodeSqsQueue_Arn(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeSqsQueue_MessageRetentionSeconds(p *SqsQueueParameters, vals map[string]cty.Value) {
	p.MessageRetentionSeconds = ctwhy.ValueAsInt64(vals["message_retention_seconds"])
}

//primitiveTypeDecodeTemplate
func DecodeSqsQueue_Name(p *SqsQueueParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//primitiveTypeDecodeTemplate
func DecodeSqsQueue_NamePrefix(p *SqsQueueParameters, vals map[string]cty.Value) {
	p.NamePrefix = ctwhy.ValueAsString(vals["name_prefix"])
}

//primitiveTypeDecodeTemplate
func DecodeSqsQueue_ReceiveWaitTimeSeconds(p *SqsQueueParameters, vals map[string]cty.Value) {
	p.ReceiveWaitTimeSeconds = ctwhy.ValueAsInt64(vals["receive_wait_time_seconds"])
}

//primitiveTypeDecodeTemplate
func DecodeSqsQueue_DelaySeconds(p *SqsQueueParameters, vals map[string]cty.Value) {
	p.DelaySeconds = ctwhy.ValueAsInt64(vals["delay_seconds"])
}

//primitiveTypeDecodeTemplate
func DecodeSqsQueue_MaxMessageSize(p *SqsQueueParameters, vals map[string]cty.Value) {
	p.MaxMessageSize = ctwhy.ValueAsInt64(vals["max_message_size"])
}

//primitiveTypeDecodeTemplate
func DecodeSqsQueue_VisibilityTimeoutSeconds(p *SqsQueueParameters, vals map[string]cty.Value) {
	p.VisibilityTimeoutSeconds = ctwhy.ValueAsInt64(vals["visibility_timeout_seconds"])
}

//primitiveTypeDecodeTemplate
func DecodeSqsQueue_FifoQueue(p *SqsQueueParameters, vals map[string]cty.Value) {
	p.FifoQueue = ctwhy.ValueAsBool(vals["fifo_queue"])
}

//primitiveTypeDecodeTemplate
func DecodeSqsQueue_RedrivePolicy(p *SqsQueueParameters, vals map[string]cty.Value) {
	p.RedrivePolicy = ctwhy.ValueAsString(vals["redrive_policy"])
}

//primitiveTypeDecodeTemplate
func DecodeSqsQueue_KmsDataKeyReusePeriodSeconds(p *SqsQueueParameters, vals map[string]cty.Value) {
	p.KmsDataKeyReusePeriodSeconds = ctwhy.ValueAsInt64(vals["kms_data_key_reuse_period_seconds"])
}

//primitiveTypeDecodeTemplate
func DecodeSqsQueue_KmsMasterKeyId(p *SqsQueueParameters, vals map[string]cty.Value) {
	p.KmsMasterKeyId = ctwhy.ValueAsString(vals["kms_master_key_id"])
}

//primitiveTypeDecodeTemplate
func DecodeSqsQueue_Policy(p *SqsQueueParameters, vals map[string]cty.Value) {
	p.Policy = ctwhy.ValueAsString(vals["policy"])
}

//primitiveMapTypeDecodeTemplate
func DecodeSqsQueue_Tags(p *SqsQueueParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

//primitiveTypeDecodeTemplate
func DecodeSqsQueue_ContentBasedDeduplication(p *SqsQueueParameters, vals map[string]cty.Value) {
	p.ContentBasedDeduplication = ctwhy.ValueAsBool(vals["content_based_deduplication"])
}

//primitiveTypeDecodeTemplate
func DecodeSqsQueue_Arn(p *SqsQueueObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}