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

package v1alpha1func EncodeElasticsearchDomain(r ElasticsearchDomain) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeElasticsearchDomain_Id(r.Spec.ForProvider, ctyVal)
	EncodeElasticsearchDomain_Tags(r.Spec.ForProvider, ctyVal)
	EncodeElasticsearchDomain_DomainName(r.Spec.ForProvider, ctyVal)
	EncodeElasticsearchDomain_ElasticsearchVersion(r.Spec.ForProvider, ctyVal)
	EncodeElasticsearchDomain_AccessPolicies(r.Spec.ForProvider, ctyVal)
	EncodeElasticsearchDomain_AdvancedOptions(r.Spec.ForProvider, ctyVal)
	EncodeElasticsearchDomain_ClusterConfig(r.Spec.ForProvider.ClusterConfig, ctyVal)
	EncodeElasticsearchDomain_DomainEndpointOptions(r.Spec.ForProvider.DomainEndpointOptions, ctyVal)
	EncodeElasticsearchDomain_EbsOptions(r.Spec.ForProvider.EbsOptions, ctyVal)
	EncodeElasticsearchDomain_NodeToNodeEncryption(r.Spec.ForProvider.NodeToNodeEncryption, ctyVal)
	EncodeElasticsearchDomain_SnapshotOptions(r.Spec.ForProvider.SnapshotOptions, ctyVal)
	EncodeElasticsearchDomain_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeElasticsearchDomain_AdvancedSecurityOptions(r.Spec.ForProvider.AdvancedSecurityOptions, ctyVal)
	EncodeElasticsearchDomain_CognitoOptions(r.Spec.ForProvider.CognitoOptions, ctyVal)
	EncodeElasticsearchDomain_EncryptAtRest(r.Spec.ForProvider.EncryptAtRest, ctyVal)
	EncodeElasticsearchDomain_LogPublishingOptions(r.Spec.ForProvider.LogPublishingOptions, ctyVal)
	EncodeElasticsearchDomain_VpcOptions(r.Spec.ForProvider.VpcOptions, ctyVal)
	EncodeElasticsearchDomain_KibanaEndpoint(r.Status.AtProvider, ctyVal)
	EncodeElasticsearchDomain_Arn(r.Status.AtProvider, ctyVal)
	EncodeElasticsearchDomain_DomainId(r.Status.AtProvider, ctyVal)
	EncodeElasticsearchDomain_Endpoint(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeElasticsearchDomain_Id(p *ElasticsearchDomainParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeElasticsearchDomain_Tags(p *ElasticsearchDomainParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeElasticsearchDomain_DomainName(p *ElasticsearchDomainParameters, vals map[string]cty.Value) {
	vals["domain_name"] = cty.StringVal(p.DomainName)
}

func EncodeElasticsearchDomain_ElasticsearchVersion(p *ElasticsearchDomainParameters, vals map[string]cty.Value) {
	vals["elasticsearch_version"] = cty.StringVal(p.ElasticsearchVersion)
}

func EncodeElasticsearchDomain_AccessPolicies(p *ElasticsearchDomainParameters, vals map[string]cty.Value) {
	vals["access_policies"] = cty.StringVal(p.AccessPolicies)
}

func EncodeElasticsearchDomain_AdvancedOptions(p *ElasticsearchDomainParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.AdvancedOptions {
		mVals[key] = cty.StringVal(value)
	}
	vals["advanced_options"] = cty.MapVal(mVals)
}

func EncodeElasticsearchDomain_ClusterConfig(p *ClusterConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.ClusterConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeElasticsearchDomain_ClusterConfig_ZoneAwarenessEnabled(v, ctyVal)
		EncodeElasticsearchDomain_ClusterConfig_DedicatedMasterCount(v, ctyVal)
		EncodeElasticsearchDomain_ClusterConfig_InstanceCount(v, ctyVal)
		EncodeElasticsearchDomain_ClusterConfig_InstanceType(v, ctyVal)
		EncodeElasticsearchDomain_ClusterConfig_WarmCount(v, ctyVal)
		EncodeElasticsearchDomain_ClusterConfig_WarmEnabled(v, ctyVal)
		EncodeElasticsearchDomain_ClusterConfig_WarmType(v, ctyVal)
		EncodeElasticsearchDomain_ClusterConfig_DedicatedMasterEnabled(v, ctyVal)
		EncodeElasticsearchDomain_ClusterConfig_DedicatedMasterType(v, ctyVal)
		EncodeElasticsearchDomain_ClusterConfig_ZoneAwarenessConfig(v.ZoneAwarenessConfig, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["cluster_config"] = cty.ListVal(valsForCollection)
}

func EncodeElasticsearchDomain_ClusterConfig_ZoneAwarenessEnabled(p *ClusterConfig, vals map[string]cty.Value) {
	vals["zone_awareness_enabled"] = cty.BoolVal(p.ZoneAwarenessEnabled)
}

func EncodeElasticsearchDomain_ClusterConfig_DedicatedMasterCount(p *ClusterConfig, vals map[string]cty.Value) {
	vals["dedicated_master_count"] = cty.IntVal(p.DedicatedMasterCount)
}

func EncodeElasticsearchDomain_ClusterConfig_InstanceCount(p *ClusterConfig, vals map[string]cty.Value) {
	vals["instance_count"] = cty.IntVal(p.InstanceCount)
}

func EncodeElasticsearchDomain_ClusterConfig_InstanceType(p *ClusterConfig, vals map[string]cty.Value) {
	vals["instance_type"] = cty.StringVal(p.InstanceType)
}

func EncodeElasticsearchDomain_ClusterConfig_WarmCount(p *ClusterConfig, vals map[string]cty.Value) {
	vals["warm_count"] = cty.IntVal(p.WarmCount)
}

func EncodeElasticsearchDomain_ClusterConfig_WarmEnabled(p *ClusterConfig, vals map[string]cty.Value) {
	vals["warm_enabled"] = cty.BoolVal(p.WarmEnabled)
}

func EncodeElasticsearchDomain_ClusterConfig_WarmType(p *ClusterConfig, vals map[string]cty.Value) {
	vals["warm_type"] = cty.StringVal(p.WarmType)
}

func EncodeElasticsearchDomain_ClusterConfig_DedicatedMasterEnabled(p *ClusterConfig, vals map[string]cty.Value) {
	vals["dedicated_master_enabled"] = cty.BoolVal(p.DedicatedMasterEnabled)
}

func EncodeElasticsearchDomain_ClusterConfig_DedicatedMasterType(p *ClusterConfig, vals map[string]cty.Value) {
	vals["dedicated_master_type"] = cty.StringVal(p.DedicatedMasterType)
}

func EncodeElasticsearchDomain_ClusterConfig_ZoneAwarenessConfig(p *ZoneAwarenessConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.ZoneAwarenessConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeElasticsearchDomain_ClusterConfig_ZoneAwarenessConfig_AvailabilityZoneCount(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["zone_awareness_config"] = cty.ListVal(valsForCollection)
}

func EncodeElasticsearchDomain_ClusterConfig_ZoneAwarenessConfig_AvailabilityZoneCount(p *ZoneAwarenessConfig, vals map[string]cty.Value) {
	vals["availability_zone_count"] = cty.IntVal(p.AvailabilityZoneCount)
}

func EncodeElasticsearchDomain_DomainEndpointOptions(p *DomainEndpointOptions, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.DomainEndpointOptions {
		ctyVal = make(map[string]cty.Value)
		EncodeElasticsearchDomain_DomainEndpointOptions_TlsSecurityPolicy(v, ctyVal)
		EncodeElasticsearchDomain_DomainEndpointOptions_EnforceHttps(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["domain_endpoint_options"] = cty.ListVal(valsForCollection)
}

func EncodeElasticsearchDomain_DomainEndpointOptions_TlsSecurityPolicy(p *DomainEndpointOptions, vals map[string]cty.Value) {
	vals["tls_security_policy"] = cty.StringVal(p.TlsSecurityPolicy)
}

func EncodeElasticsearchDomain_DomainEndpointOptions_EnforceHttps(p *DomainEndpointOptions, vals map[string]cty.Value) {
	vals["enforce_https"] = cty.BoolVal(p.EnforceHttps)
}

func EncodeElasticsearchDomain_EbsOptions(p *EbsOptions, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.EbsOptions {
		ctyVal = make(map[string]cty.Value)
		EncodeElasticsearchDomain_EbsOptions_EbsEnabled(v, ctyVal)
		EncodeElasticsearchDomain_EbsOptions_Iops(v, ctyVal)
		EncodeElasticsearchDomain_EbsOptions_VolumeSize(v, ctyVal)
		EncodeElasticsearchDomain_EbsOptions_VolumeType(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["ebs_options"] = cty.ListVal(valsForCollection)
}

func EncodeElasticsearchDomain_EbsOptions_EbsEnabled(p *EbsOptions, vals map[string]cty.Value) {
	vals["ebs_enabled"] = cty.BoolVal(p.EbsEnabled)
}

func EncodeElasticsearchDomain_EbsOptions_Iops(p *EbsOptions, vals map[string]cty.Value) {
	vals["iops"] = cty.IntVal(p.Iops)
}

func EncodeElasticsearchDomain_EbsOptions_VolumeSize(p *EbsOptions, vals map[string]cty.Value) {
	vals["volume_size"] = cty.IntVal(p.VolumeSize)
}

func EncodeElasticsearchDomain_EbsOptions_VolumeType(p *EbsOptions, vals map[string]cty.Value) {
	vals["volume_type"] = cty.StringVal(p.VolumeType)
}

func EncodeElasticsearchDomain_NodeToNodeEncryption(p *NodeToNodeEncryption, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.NodeToNodeEncryption {
		ctyVal = make(map[string]cty.Value)
		EncodeElasticsearchDomain_NodeToNodeEncryption_Enabled(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["node_to_node_encryption"] = cty.ListVal(valsForCollection)
}

func EncodeElasticsearchDomain_NodeToNodeEncryption_Enabled(p *NodeToNodeEncryption, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeElasticsearchDomain_SnapshotOptions(p *SnapshotOptions, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SnapshotOptions {
		ctyVal = make(map[string]cty.Value)
		EncodeElasticsearchDomain_SnapshotOptions_AutomatedSnapshotStartHour(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["snapshot_options"] = cty.ListVal(valsForCollection)
}

func EncodeElasticsearchDomain_SnapshotOptions_AutomatedSnapshotStartHour(p *SnapshotOptions, vals map[string]cty.Value) {
	vals["automated_snapshot_start_hour"] = cty.IntVal(p.AutomatedSnapshotStartHour)
}

func EncodeElasticsearchDomain_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	ctyVal = make(map[string]cty.Value)
	EncodeElasticsearchDomain_Timeouts_Update(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeElasticsearchDomain_Timeouts_Update(p *Timeouts, vals map[string]cty.Value) {
	vals["update"] = cty.StringVal(p.Update)
}

func EncodeElasticsearchDomain_AdvancedSecurityOptions(p *AdvancedSecurityOptions, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AdvancedSecurityOptions {
		ctyVal = make(map[string]cty.Value)
		EncodeElasticsearchDomain_AdvancedSecurityOptions_Enabled(v, ctyVal)
		EncodeElasticsearchDomain_AdvancedSecurityOptions_InternalUserDatabaseEnabled(v, ctyVal)
		EncodeElasticsearchDomain_AdvancedSecurityOptions_MasterUserOptions(v.MasterUserOptions, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["advanced_security_options"] = cty.ListVal(valsForCollection)
}

func EncodeElasticsearchDomain_AdvancedSecurityOptions_Enabled(p *AdvancedSecurityOptions, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeElasticsearchDomain_AdvancedSecurityOptions_InternalUserDatabaseEnabled(p *AdvancedSecurityOptions, vals map[string]cty.Value) {
	vals["internal_user_database_enabled"] = cty.BoolVal(p.InternalUserDatabaseEnabled)
}

func EncodeElasticsearchDomain_AdvancedSecurityOptions_MasterUserOptions(p *MasterUserOptions, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.MasterUserOptions {
		ctyVal = make(map[string]cty.Value)
		EncodeElasticsearchDomain_AdvancedSecurityOptions_MasterUserOptions_MasterUserArn(v, ctyVal)
		EncodeElasticsearchDomain_AdvancedSecurityOptions_MasterUserOptions_MasterUserName(v, ctyVal)
		EncodeElasticsearchDomain_AdvancedSecurityOptions_MasterUserOptions_MasterUserPassword(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["master_user_options"] = cty.ListVal(valsForCollection)
}

func EncodeElasticsearchDomain_AdvancedSecurityOptions_MasterUserOptions_MasterUserArn(p *MasterUserOptions, vals map[string]cty.Value) {
	vals["master_user_arn"] = cty.StringVal(p.MasterUserArn)
}

func EncodeElasticsearchDomain_AdvancedSecurityOptions_MasterUserOptions_MasterUserName(p *MasterUserOptions, vals map[string]cty.Value) {
	vals["master_user_name"] = cty.StringVal(p.MasterUserName)
}

func EncodeElasticsearchDomain_AdvancedSecurityOptions_MasterUserOptions_MasterUserPassword(p *MasterUserOptions, vals map[string]cty.Value) {
	vals["master_user_password"] = cty.StringVal(p.MasterUserPassword)
}

func EncodeElasticsearchDomain_CognitoOptions(p *CognitoOptions, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.CognitoOptions {
		ctyVal = make(map[string]cty.Value)
		EncodeElasticsearchDomain_CognitoOptions_Enabled(v, ctyVal)
		EncodeElasticsearchDomain_CognitoOptions_IdentityPoolId(v, ctyVal)
		EncodeElasticsearchDomain_CognitoOptions_RoleArn(v, ctyVal)
		EncodeElasticsearchDomain_CognitoOptions_UserPoolId(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["cognito_options"] = cty.ListVal(valsForCollection)
}

func EncodeElasticsearchDomain_CognitoOptions_Enabled(p *CognitoOptions, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeElasticsearchDomain_CognitoOptions_IdentityPoolId(p *CognitoOptions, vals map[string]cty.Value) {
	vals["identity_pool_id"] = cty.StringVal(p.IdentityPoolId)
}

func EncodeElasticsearchDomain_CognitoOptions_RoleArn(p *CognitoOptions, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeElasticsearchDomain_CognitoOptions_UserPoolId(p *CognitoOptions, vals map[string]cty.Value) {
	vals["user_pool_id"] = cty.StringVal(p.UserPoolId)
}

func EncodeElasticsearchDomain_EncryptAtRest(p *EncryptAtRest, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.EncryptAtRest {
		ctyVal = make(map[string]cty.Value)
		EncodeElasticsearchDomain_EncryptAtRest_KmsKeyId(v, ctyVal)
		EncodeElasticsearchDomain_EncryptAtRest_Enabled(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["encrypt_at_rest"] = cty.ListVal(valsForCollection)
}

func EncodeElasticsearchDomain_EncryptAtRest_KmsKeyId(p *EncryptAtRest, vals map[string]cty.Value) {
	vals["kms_key_id"] = cty.StringVal(p.KmsKeyId)
}

func EncodeElasticsearchDomain_EncryptAtRest_Enabled(p *EncryptAtRest, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeElasticsearchDomain_LogPublishingOptions(p *LogPublishingOptions, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.LogPublishingOptions {
		ctyVal = make(map[string]cty.Value)
		EncodeElasticsearchDomain_LogPublishingOptions_CloudwatchLogGroupArn(v, ctyVal)
		EncodeElasticsearchDomain_LogPublishingOptions_Enabled(v, ctyVal)
		EncodeElasticsearchDomain_LogPublishingOptions_LogType(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["log_publishing_options"] = cty.SetVal(valsForCollection)
}

func EncodeElasticsearchDomain_LogPublishingOptions_CloudwatchLogGroupArn(p *LogPublishingOptions, vals map[string]cty.Value) {
	vals["cloudwatch_log_group_arn"] = cty.StringVal(p.CloudwatchLogGroupArn)
}

func EncodeElasticsearchDomain_LogPublishingOptions_Enabled(p *LogPublishingOptions, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeElasticsearchDomain_LogPublishingOptions_LogType(p *LogPublishingOptions, vals map[string]cty.Value) {
	vals["log_type"] = cty.StringVal(p.LogType)
}

func EncodeElasticsearchDomain_VpcOptions(p *VpcOptions, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.VpcOptions {
		ctyVal = make(map[string]cty.Value)
		EncodeElasticsearchDomain_VpcOptions_VpcId(v, ctyVal)
		EncodeElasticsearchDomain_VpcOptions_AvailabilityZones(v, ctyVal)
		EncodeElasticsearchDomain_VpcOptions_SecurityGroupIds(v, ctyVal)
		EncodeElasticsearchDomain_VpcOptions_SubnetIds(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["vpc_options"] = cty.ListVal(valsForCollection)
}

func EncodeElasticsearchDomain_VpcOptions_VpcId(p *VpcOptions, vals map[string]cty.Value) {
	vals["vpc_id"] = cty.StringVal(p.VpcId)
}

func EncodeElasticsearchDomain_VpcOptions_AvailabilityZones(p *VpcOptions, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.AvailabilityZones {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["availability_zones"] = cty.SetVal(colVals)
}

func EncodeElasticsearchDomain_VpcOptions_SecurityGroupIds(p *VpcOptions, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SecurityGroupIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["security_group_ids"] = cty.SetVal(colVals)
}

func EncodeElasticsearchDomain_VpcOptions_SubnetIds(p *VpcOptions, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SubnetIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["subnet_ids"] = cty.SetVal(colVals)
}

func EncodeElasticsearchDomain_KibanaEndpoint(p *ElasticsearchDomainObservation, vals map[string]cty.Value) {
	vals["kibana_endpoint"] = cty.StringVal(p.KibanaEndpoint)
}

func EncodeElasticsearchDomain_Arn(p *ElasticsearchDomainObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeElasticsearchDomain_DomainId(p *ElasticsearchDomainObservation, vals map[string]cty.Value) {
	vals["domain_id"] = cty.StringVal(p.DomainId)
}

func EncodeElasticsearchDomain_Endpoint(p *ElasticsearchDomainObservation, vals map[string]cty.Value) {
	vals["endpoint"] = cty.StringVal(p.Endpoint)
}