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

func EncodeDirectoryServiceConditionalForwarder(r DirectoryServiceConditionalForwarder) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDirectoryServiceConditionalForwarder_DirectoryId(r.Spec.ForProvider, ctyVal)
	EncodeDirectoryServiceConditionalForwarder_DnsIps(r.Spec.ForProvider, ctyVal)
	EncodeDirectoryServiceConditionalForwarder_Id(r.Spec.ForProvider, ctyVal)
	EncodeDirectoryServiceConditionalForwarder_RemoteDomainName(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeDirectoryServiceConditionalForwarder_DirectoryId(p DirectoryServiceConditionalForwarderParameters, vals map[string]cty.Value) {
	vals["directory_id"] = cty.StringVal(p.DirectoryId)
}

func EncodeDirectoryServiceConditionalForwarder_DnsIps(p DirectoryServiceConditionalForwarderParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.DnsIps {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["dns_ips"] = cty.ListVal(colVals)
}

func EncodeDirectoryServiceConditionalForwarder_Id(p DirectoryServiceConditionalForwarderParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeDirectoryServiceConditionalForwarder_RemoteDomainName(p DirectoryServiceConditionalForwarderParameters, vals map[string]cty.Value) {
	vals["remote_domain_name"] = cty.StringVal(p.RemoteDomainName)
}