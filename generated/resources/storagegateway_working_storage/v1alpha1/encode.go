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

package v1alpha1func EncodeStoragegatewayWorkingStorage(r StoragegatewayWorkingStorage) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeStoragegatewayWorkingStorage_DiskId(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewayWorkingStorage_GatewayArn(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewayWorkingStorage_Id(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeStoragegatewayWorkingStorage_DiskId(p *StoragegatewayWorkingStorageParameters, vals map[string]cty.Value) {
	vals["disk_id"] = cty.StringVal(p.DiskId)
}

func EncodeStoragegatewayWorkingStorage_GatewayArn(p *StoragegatewayWorkingStorageParameters, vals map[string]cty.Value) {
	vals["gateway_arn"] = cty.StringVal(p.GatewayArn)
}

func EncodeStoragegatewayWorkingStorage_Id(p *StoragegatewayWorkingStorageParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}