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
	r, ok := mr.(*S3BucketPublicAccessBlock)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeS3BucketPublicAccessBlock(r, ctyValue)
}

func DecodeS3BucketPublicAccessBlock(prev *S3BucketPublicAccessBlock, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeS3BucketPublicAccessBlock_BlockPublicAcls(&new.Spec.ForProvider, valMap)
	DecodeS3BucketPublicAccessBlock_BlockPublicPolicy(&new.Spec.ForProvider, valMap)
	DecodeS3BucketPublicAccessBlock_Bucket(&new.Spec.ForProvider, valMap)
	DecodeS3BucketPublicAccessBlock_IgnorePublicAcls(&new.Spec.ForProvider, valMap)
	DecodeS3BucketPublicAccessBlock_RestrictPublicBuckets(&new.Spec.ForProvider, valMap)

	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeS3BucketPublicAccessBlock_BlockPublicAcls(p *S3BucketPublicAccessBlockParameters, vals map[string]cty.Value) {
	p.BlockPublicAcls = ctwhy.ValueAsBool(vals["block_public_acls"])
}

//primitiveTypeDecodeTemplate
func DecodeS3BucketPublicAccessBlock_BlockPublicPolicy(p *S3BucketPublicAccessBlockParameters, vals map[string]cty.Value) {
	p.BlockPublicPolicy = ctwhy.ValueAsBool(vals["block_public_policy"])
}

//primitiveTypeDecodeTemplate
func DecodeS3BucketPublicAccessBlock_Bucket(p *S3BucketPublicAccessBlockParameters, vals map[string]cty.Value) {
	p.Bucket = ctwhy.ValueAsString(vals["bucket"])
}

//primitiveTypeDecodeTemplate
func DecodeS3BucketPublicAccessBlock_IgnorePublicAcls(p *S3BucketPublicAccessBlockParameters, vals map[string]cty.Value) {
	p.IgnorePublicAcls = ctwhy.ValueAsBool(vals["ignore_public_acls"])
}

//primitiveTypeDecodeTemplate
func DecodeS3BucketPublicAccessBlock_RestrictPublicBuckets(p *S3BucketPublicAccessBlockParameters, vals map[string]cty.Value) {
	p.RestrictPublicBuckets = ctwhy.ValueAsBool(vals["restrict_public_buckets"])
}