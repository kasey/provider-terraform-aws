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

package v1alpha1func EncodeEfsFileSystemPolicy(r EfsFileSystemPolicy) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeEfsFileSystemPolicy_Id(r.Spec.ForProvider, ctyVal)
	EncodeEfsFileSystemPolicy_Policy(r.Spec.ForProvider, ctyVal)
	EncodeEfsFileSystemPolicy_FileSystemId(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeEfsFileSystemPolicy_Id(p *EfsFileSystemPolicyParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeEfsFileSystemPolicy_Policy(p *EfsFileSystemPolicyParameters, vals map[string]cty.Value) {
	vals["policy"] = cty.StringVal(p.Policy)
}

func EncodeEfsFileSystemPolicy_FileSystemId(p *EfsFileSystemPolicyParameters, vals map[string]cty.Value) {
	vals["file_system_id"] = cty.StringVal(p.FileSystemId)
}