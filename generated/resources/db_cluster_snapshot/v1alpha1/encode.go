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

func EncodeDbClusterSnapshot(r DbClusterSnapshot) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDbClusterSnapshot_DbClusterSnapshotIdentifier(r.Spec.ForProvider, ctyVal)
	EncodeDbClusterSnapshot_DbClusterIdentifier(r.Spec.ForProvider, ctyVal)
	EncodeDbClusterSnapshot_Id(r.Spec.ForProvider, ctyVal)
	EncodeDbClusterSnapshot_Tags(r.Spec.ForProvider, ctyVal)
	EncodeDbClusterSnapshot_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeDbClusterSnapshot_LicenseModel(r.Status.AtProvider, ctyVal)
	EncodeDbClusterSnapshot_SnapshotType(r.Status.AtProvider, ctyVal)
	EncodeDbClusterSnapshot_AllocatedStorage(r.Status.AtProvider, ctyVal)
	EncodeDbClusterSnapshot_AvailabilityZones(r.Status.AtProvider, ctyVal)
	EncodeDbClusterSnapshot_SourceDbClusterSnapshotArn(r.Status.AtProvider, ctyVal)
	EncodeDbClusterSnapshot_VpcId(r.Status.AtProvider, ctyVal)
	EncodeDbClusterSnapshot_Port(r.Status.AtProvider, ctyVal)
	EncodeDbClusterSnapshot_Status(r.Status.AtProvider, ctyVal)
	EncodeDbClusterSnapshot_StorageEncrypted(r.Status.AtProvider, ctyVal)
	EncodeDbClusterSnapshot_DbClusterSnapshotArn(r.Status.AtProvider, ctyVal)
	EncodeDbClusterSnapshot_Engine(r.Status.AtProvider, ctyVal)
	EncodeDbClusterSnapshot_EngineVersion(r.Status.AtProvider, ctyVal)
	EncodeDbClusterSnapshot_KmsKeyId(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeDbClusterSnapshot_DbClusterSnapshotIdentifier(p DbClusterSnapshotParameters, vals map[string]cty.Value) {
	vals["db_cluster_snapshot_identifier"] = cty.StringVal(p.DbClusterSnapshotIdentifier)
}

func EncodeDbClusterSnapshot_DbClusterIdentifier(p DbClusterSnapshotParameters, vals map[string]cty.Value) {
	vals["db_cluster_identifier"] = cty.StringVal(p.DbClusterIdentifier)
}

func EncodeDbClusterSnapshot_Id(p DbClusterSnapshotParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeDbClusterSnapshot_Tags(p DbClusterSnapshotParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeDbClusterSnapshot_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeDbClusterSnapshot_Timeouts_Create(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeDbClusterSnapshot_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeDbClusterSnapshot_LicenseModel(p DbClusterSnapshotObservation, vals map[string]cty.Value) {
	vals["license_model"] = cty.StringVal(p.LicenseModel)
}

func EncodeDbClusterSnapshot_SnapshotType(p DbClusterSnapshotObservation, vals map[string]cty.Value) {
	vals["snapshot_type"] = cty.StringVal(p.SnapshotType)
}

func EncodeDbClusterSnapshot_AllocatedStorage(p DbClusterSnapshotObservation, vals map[string]cty.Value) {
	vals["allocated_storage"] = cty.NumberIntVal(p.AllocatedStorage)
}

func EncodeDbClusterSnapshot_AvailabilityZones(p DbClusterSnapshotObservation, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.AvailabilityZones {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["availability_zones"] = cty.ListVal(colVals)
}

func EncodeDbClusterSnapshot_SourceDbClusterSnapshotArn(p DbClusterSnapshotObservation, vals map[string]cty.Value) {
	vals["source_db_cluster_snapshot_arn"] = cty.StringVal(p.SourceDbClusterSnapshotArn)
}

func EncodeDbClusterSnapshot_VpcId(p DbClusterSnapshotObservation, vals map[string]cty.Value) {
	vals["vpc_id"] = cty.StringVal(p.VpcId)
}

func EncodeDbClusterSnapshot_Port(p DbClusterSnapshotObservation, vals map[string]cty.Value) {
	vals["port"] = cty.NumberIntVal(p.Port)
}

func EncodeDbClusterSnapshot_Status(p DbClusterSnapshotObservation, vals map[string]cty.Value) {
	vals["status"] = cty.StringVal(p.Status)
}

func EncodeDbClusterSnapshot_StorageEncrypted(p DbClusterSnapshotObservation, vals map[string]cty.Value) {
	vals["storage_encrypted"] = cty.BoolVal(p.StorageEncrypted)
}

func EncodeDbClusterSnapshot_DbClusterSnapshotArn(p DbClusterSnapshotObservation, vals map[string]cty.Value) {
	vals["db_cluster_snapshot_arn"] = cty.StringVal(p.DbClusterSnapshotArn)
}

func EncodeDbClusterSnapshot_Engine(p DbClusterSnapshotObservation, vals map[string]cty.Value) {
	vals["engine"] = cty.StringVal(p.Engine)
}

func EncodeDbClusterSnapshot_EngineVersion(p DbClusterSnapshotObservation, vals map[string]cty.Value) {
	vals["engine_version"] = cty.StringVal(p.EngineVersion)
}

func EncodeDbClusterSnapshot_KmsKeyId(p DbClusterSnapshotObservation, vals map[string]cty.Value) {
	vals["kms_key_id"] = cty.StringVal(p.KmsKeyId)
}