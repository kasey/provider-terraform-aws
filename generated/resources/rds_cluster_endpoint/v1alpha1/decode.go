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
	r, ok := mr.(*RdsClusterEndpoint)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeRdsClusterEndpoint(r, ctyValue)
}

func DecodeRdsClusterEndpoint(prev *RdsClusterEndpoint, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeRdsClusterEndpoint_ClusterEndpointIdentifier(&new.Spec.ForProvider, valMap)
	DecodeRdsClusterEndpoint_CustomEndpointType(&new.Spec.ForProvider, valMap)
	DecodeRdsClusterEndpoint_Id(&new.Spec.ForProvider, valMap)
	DecodeRdsClusterEndpoint_StaticMembers(&new.Spec.ForProvider, valMap)
	DecodeRdsClusterEndpoint_ClusterIdentifier(&new.Spec.ForProvider, valMap)
	DecodeRdsClusterEndpoint_ExcludedMembers(&new.Spec.ForProvider, valMap)
	DecodeRdsClusterEndpoint_Tags(&new.Spec.ForProvider, valMap)
	DecodeRdsClusterEndpoint_Arn(&new.Status.AtProvider, valMap)
	DecodeRdsClusterEndpoint_Endpoint(&new.Status.AtProvider, valMap)
	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodeRdsClusterEndpoint_ClusterEndpointIdentifier(p *RdsClusterEndpointParameters, vals map[string]cty.Value) {
	p.ClusterEndpointIdentifier = ctwhy.ValueAsString(vals["cluster_endpoint_identifier"])
}

func DecodeRdsClusterEndpoint_CustomEndpointType(p *RdsClusterEndpointParameters, vals map[string]cty.Value) {
	p.CustomEndpointType = ctwhy.ValueAsString(vals["custom_endpoint_type"])
}

func DecodeRdsClusterEndpoint_Id(p *RdsClusterEndpointParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

func DecodeRdsClusterEndpoint_StaticMembers(p *RdsClusterEndpointParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["static_members"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.StaticMembers = goVals
}

func DecodeRdsClusterEndpoint_ClusterIdentifier(p *RdsClusterEndpointParameters, vals map[string]cty.Value) {
	p.ClusterIdentifier = ctwhy.ValueAsString(vals["cluster_identifier"])
}

func DecodeRdsClusterEndpoint_ExcludedMembers(p *RdsClusterEndpointParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["excluded_members"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.ExcludedMembers = goVals
}

func DecodeRdsClusterEndpoint_Tags(p *RdsClusterEndpointParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

func DecodeRdsClusterEndpoint_Arn(p *RdsClusterEndpointObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}

func DecodeRdsClusterEndpoint_Endpoint(p *RdsClusterEndpointObservation, vals map[string]cty.Value) {
	p.Endpoint = ctwhy.ValueAsString(vals["endpoint"])
}