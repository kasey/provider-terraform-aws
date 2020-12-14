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
	r, ok := mr.(*KinesisStream)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeKinesisStream(r, ctyValue)
}

func DecodeKinesisStream(prev *KinesisStream, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeKinesisStream_Arn(&new.Spec.ForProvider, valMap)
	DecodeKinesisStream_EnforceConsumerDeletion(&new.Spec.ForProvider, valMap)
	DecodeKinesisStream_Id(&new.Spec.ForProvider, valMap)
	DecodeKinesisStream_Name(&new.Spec.ForProvider, valMap)
	DecodeKinesisStream_Tags(&new.Spec.ForProvider, valMap)
	DecodeKinesisStream_EncryptionType(&new.Spec.ForProvider, valMap)
	DecodeKinesisStream_KmsKeyId(&new.Spec.ForProvider, valMap)
	DecodeKinesisStream_RetentionPeriod(&new.Spec.ForProvider, valMap)
	DecodeKinesisStream_ShardCount(&new.Spec.ForProvider, valMap)
	DecodeKinesisStream_ShardLevelMetrics(&new.Spec.ForProvider, valMap)
	DecodeKinesisStream_Timeouts(&new.Spec.ForProvider.Timeouts, valMap)

	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodeKinesisStream_Arn(p *KinesisStreamParameters, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}

func DecodeKinesisStream_EnforceConsumerDeletion(p *KinesisStreamParameters, vals map[string]cty.Value) {
	p.EnforceConsumerDeletion = ctwhy.ValueAsBool(vals["enforce_consumer_deletion"])
}

func DecodeKinesisStream_Id(p *KinesisStreamParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

func DecodeKinesisStream_Name(p *KinesisStreamParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

func DecodeKinesisStream_Tags(p *KinesisStreamParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

func DecodeKinesisStream_EncryptionType(p *KinesisStreamParameters, vals map[string]cty.Value) {
	p.EncryptionType = ctwhy.ValueAsString(vals["encryption_type"])
}

func DecodeKinesisStream_KmsKeyId(p *KinesisStreamParameters, vals map[string]cty.Value) {
	p.KmsKeyId = ctwhy.ValueAsString(vals["kms_key_id"])
}

func DecodeKinesisStream_RetentionPeriod(p *KinesisStreamParameters, vals map[string]cty.Value) {
	p.RetentionPeriod = ctwhy.ValueAsInt64(vals["retention_period"])
}

func DecodeKinesisStream_ShardCount(p *KinesisStreamParameters, vals map[string]cty.Value) {
	p.ShardCount = ctwhy.ValueAsInt64(vals["shard_count"])
}

func DecodeKinesisStream_ShardLevelMetrics(p *KinesisStreamParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["shard_level_metrics"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.ShardLevelMetrics = goVals
}

func DecodeKinesisStream_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	valMap := vals["timeouts"].AsValueMap()
	DecodeKinesisStream_Timeouts_Create(p, valMap)
	DecodeKinesisStream_Timeouts_Delete(p, valMap)
	DecodeKinesisStream_Timeouts_Update(p, valMap)
}

func DecodeKinesisStream_Timeouts_Create(p *Timeouts, vals map[string]cty.Value) {
	p.Create = ctwhy.ValueAsString(vals["create"])
}

func DecodeKinesisStream_Timeouts_Delete(p *Timeouts, vals map[string]cty.Value) {
	p.Delete = ctwhy.ValueAsString(vals["delete"])
}

func DecodeKinesisStream_Timeouts_Update(p *Timeouts, vals map[string]cty.Value) {
	p.Update = ctwhy.ValueAsString(vals["update"])
}