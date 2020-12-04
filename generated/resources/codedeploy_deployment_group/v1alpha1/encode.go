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

func EncodeCodedeployDeploymentGroup(r CodedeployDeploymentGroup) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeCodedeployDeploymentGroup_AppName(r.Spec.ForProvider, ctyVal)
	EncodeCodedeployDeploymentGroup_AutoscalingGroups(r.Spec.ForProvider, ctyVal)
	EncodeCodedeployDeploymentGroup_DeploymentConfigName(r.Spec.ForProvider, ctyVal)
	EncodeCodedeployDeploymentGroup_DeploymentGroupName(r.Spec.ForProvider, ctyVal)
	EncodeCodedeployDeploymentGroup_Id(r.Spec.ForProvider, ctyVal)
	EncodeCodedeployDeploymentGroup_ServiceRoleArn(r.Spec.ForProvider, ctyVal)
	EncodeCodedeployDeploymentGroup_Ec2TagSet(r.Spec.ForProvider.Ec2TagSet, ctyVal)
	EncodeCodedeployDeploymentGroup_BlueGreenDeploymentConfig(r.Spec.ForProvider.BlueGreenDeploymentConfig, ctyVal)
	EncodeCodedeployDeploymentGroup_AutoRollbackConfiguration(r.Spec.ForProvider.AutoRollbackConfiguration, ctyVal)
	EncodeCodedeployDeploymentGroup_DeploymentStyle(r.Spec.ForProvider.DeploymentStyle, ctyVal)
	EncodeCodedeployDeploymentGroup_Ec2TagFilter(r.Spec.ForProvider.Ec2TagFilter, ctyVal)
	EncodeCodedeployDeploymentGroup_EcsService(r.Spec.ForProvider.EcsService, ctyVal)
	EncodeCodedeployDeploymentGroup_LoadBalancerInfo(r.Spec.ForProvider.LoadBalancerInfo, ctyVal)
	EncodeCodedeployDeploymentGroup_OnPremisesInstanceTagFilter(r.Spec.ForProvider.OnPremisesInstanceTagFilter, ctyVal)
	EncodeCodedeployDeploymentGroup_TriggerConfiguration(r.Spec.ForProvider.TriggerConfiguration, ctyVal)
	EncodeCodedeployDeploymentGroup_AlarmConfiguration(r.Spec.ForProvider.AlarmConfiguration, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeCodedeployDeploymentGroup_AppName(p CodedeployDeploymentGroupParameters, vals map[string]cty.Value) {
	vals["app_name"] = cty.StringVal(p.AppName)
}

func EncodeCodedeployDeploymentGroup_AutoscalingGroups(p CodedeployDeploymentGroupParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.AutoscalingGroups {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["autoscaling_groups"] = cty.SetVal(colVals)
}

func EncodeCodedeployDeploymentGroup_DeploymentConfigName(p CodedeployDeploymentGroupParameters, vals map[string]cty.Value) {
	vals["deployment_config_name"] = cty.StringVal(p.DeploymentConfigName)
}

func EncodeCodedeployDeploymentGroup_DeploymentGroupName(p CodedeployDeploymentGroupParameters, vals map[string]cty.Value) {
	vals["deployment_group_name"] = cty.StringVal(p.DeploymentGroupName)
}

func EncodeCodedeployDeploymentGroup_Id(p CodedeployDeploymentGroupParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeCodedeployDeploymentGroup_ServiceRoleArn(p CodedeployDeploymentGroupParameters, vals map[string]cty.Value) {
	vals["service_role_arn"] = cty.StringVal(p.ServiceRoleArn)
}

func EncodeCodedeployDeploymentGroup_Ec2TagSet(p Ec2TagSet, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeCodedeployDeploymentGroup_Ec2TagSet_Ec2TagFilter(p.Ec2TagFilter, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["ec2_tag_set"] = cty.SetVal(valsForCollection)
}

func EncodeCodedeployDeploymentGroup_Ec2TagSet_Ec2TagFilter(p Ec2TagFilter, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeCodedeployDeploymentGroup_Ec2TagSet_Ec2TagFilter_Value(p, ctyVal)
	EncodeCodedeployDeploymentGroup_Ec2TagSet_Ec2TagFilter_Key(p, ctyVal)
	EncodeCodedeployDeploymentGroup_Ec2TagSet_Ec2TagFilter_Type(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["ec2_tag_filter"] = cty.SetVal(valsForCollection)
}

func EncodeCodedeployDeploymentGroup_Ec2TagSet_Ec2TagFilter_Value(p Ec2TagFilter, vals map[string]cty.Value) {
	vals["value"] = cty.StringVal(p.Value)
}

func EncodeCodedeployDeploymentGroup_Ec2TagSet_Ec2TagFilter_Key(p Ec2TagFilter, vals map[string]cty.Value) {
	vals["key"] = cty.StringVal(p.Key)
}

func EncodeCodedeployDeploymentGroup_Ec2TagSet_Ec2TagFilter_Type(p Ec2TagFilter, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeCodedeployDeploymentGroup_BlueGreenDeploymentConfig(p BlueGreenDeploymentConfig, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeCodedeployDeploymentGroup_BlueGreenDeploymentConfig_DeploymentReadyOption(p.DeploymentReadyOption, ctyVal)
	EncodeCodedeployDeploymentGroup_BlueGreenDeploymentConfig_GreenFleetProvisioningOption(p.GreenFleetProvisioningOption, ctyVal)
	EncodeCodedeployDeploymentGroup_BlueGreenDeploymentConfig_TerminateBlueInstancesOnDeploymentSuccess(p.TerminateBlueInstancesOnDeploymentSuccess, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["blue_green_deployment_config"] = cty.ListVal(valsForCollection)
}

func EncodeCodedeployDeploymentGroup_BlueGreenDeploymentConfig_DeploymentReadyOption(p DeploymentReadyOption, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeCodedeployDeploymentGroup_BlueGreenDeploymentConfig_DeploymentReadyOption_ActionOnTimeout(p, ctyVal)
	EncodeCodedeployDeploymentGroup_BlueGreenDeploymentConfig_DeploymentReadyOption_WaitTimeInMinutes(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["deployment_ready_option"] = cty.ListVal(valsForCollection)
}

func EncodeCodedeployDeploymentGroup_BlueGreenDeploymentConfig_DeploymentReadyOption_ActionOnTimeout(p DeploymentReadyOption, vals map[string]cty.Value) {
	vals["action_on_timeout"] = cty.StringVal(p.ActionOnTimeout)
}

func EncodeCodedeployDeploymentGroup_BlueGreenDeploymentConfig_DeploymentReadyOption_WaitTimeInMinutes(p DeploymentReadyOption, vals map[string]cty.Value) {
	vals["wait_time_in_minutes"] = cty.NumberIntVal(p.WaitTimeInMinutes)
}

func EncodeCodedeployDeploymentGroup_BlueGreenDeploymentConfig_GreenFleetProvisioningOption(p GreenFleetProvisioningOption, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeCodedeployDeploymentGroup_BlueGreenDeploymentConfig_GreenFleetProvisioningOption_Action(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["green_fleet_provisioning_option"] = cty.ListVal(valsForCollection)
}

func EncodeCodedeployDeploymentGroup_BlueGreenDeploymentConfig_GreenFleetProvisioningOption_Action(p GreenFleetProvisioningOption, vals map[string]cty.Value) {
	vals["action"] = cty.StringVal(p.Action)
}

func EncodeCodedeployDeploymentGroup_BlueGreenDeploymentConfig_TerminateBlueInstancesOnDeploymentSuccess(p TerminateBlueInstancesOnDeploymentSuccess, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeCodedeployDeploymentGroup_BlueGreenDeploymentConfig_TerminateBlueInstancesOnDeploymentSuccess_TerminationWaitTimeInMinutes(p, ctyVal)
	EncodeCodedeployDeploymentGroup_BlueGreenDeploymentConfig_TerminateBlueInstancesOnDeploymentSuccess_Action(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["terminate_blue_instances_on_deployment_success"] = cty.ListVal(valsForCollection)
}

func EncodeCodedeployDeploymentGroup_BlueGreenDeploymentConfig_TerminateBlueInstancesOnDeploymentSuccess_TerminationWaitTimeInMinutes(p TerminateBlueInstancesOnDeploymentSuccess, vals map[string]cty.Value) {
	vals["termination_wait_time_in_minutes"] = cty.NumberIntVal(p.TerminationWaitTimeInMinutes)
}

func EncodeCodedeployDeploymentGroup_BlueGreenDeploymentConfig_TerminateBlueInstancesOnDeploymentSuccess_Action(p TerminateBlueInstancesOnDeploymentSuccess, vals map[string]cty.Value) {
	vals["action"] = cty.StringVal(p.Action)
}

func EncodeCodedeployDeploymentGroup_AutoRollbackConfiguration(p AutoRollbackConfiguration, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeCodedeployDeploymentGroup_AutoRollbackConfiguration_Enabled(p, ctyVal)
	EncodeCodedeployDeploymentGroup_AutoRollbackConfiguration_Events(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["auto_rollback_configuration"] = cty.ListVal(valsForCollection)
}

func EncodeCodedeployDeploymentGroup_AutoRollbackConfiguration_Enabled(p AutoRollbackConfiguration, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeCodedeployDeploymentGroup_AutoRollbackConfiguration_Events(p AutoRollbackConfiguration, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Events {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["events"] = cty.SetVal(colVals)
}

func EncodeCodedeployDeploymentGroup_DeploymentStyle(p DeploymentStyle, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeCodedeployDeploymentGroup_DeploymentStyle_DeploymentOption(p, ctyVal)
	EncodeCodedeployDeploymentGroup_DeploymentStyle_DeploymentType(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["deployment_style"] = cty.ListVal(valsForCollection)
}

func EncodeCodedeployDeploymentGroup_DeploymentStyle_DeploymentOption(p DeploymentStyle, vals map[string]cty.Value) {
	vals["deployment_option"] = cty.StringVal(p.DeploymentOption)
}

func EncodeCodedeployDeploymentGroup_DeploymentStyle_DeploymentType(p DeploymentStyle, vals map[string]cty.Value) {
	vals["deployment_type"] = cty.StringVal(p.DeploymentType)
}

func EncodeCodedeployDeploymentGroup_Ec2TagFilter(p Ec2TagFilter, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeCodedeployDeploymentGroup_Ec2TagFilter_Key(p, ctyVal)
	EncodeCodedeployDeploymentGroup_Ec2TagFilter_Type(p, ctyVal)
	EncodeCodedeployDeploymentGroup_Ec2TagFilter_Value(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["ec2_tag_filter"] = cty.SetVal(valsForCollection)
}

func EncodeCodedeployDeploymentGroup_Ec2TagFilter_Key(p Ec2TagFilter, vals map[string]cty.Value) {
	vals["key"] = cty.StringVal(p.Key)
}

func EncodeCodedeployDeploymentGroup_Ec2TagFilter_Type(p Ec2TagFilter, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeCodedeployDeploymentGroup_Ec2TagFilter_Value(p Ec2TagFilter, vals map[string]cty.Value) {
	vals["value"] = cty.StringVal(p.Value)
}

func EncodeCodedeployDeploymentGroup_EcsService(p EcsService, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeCodedeployDeploymentGroup_EcsService_ClusterName(p, ctyVal)
	EncodeCodedeployDeploymentGroup_EcsService_ServiceName(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["ecs_service"] = cty.ListVal(valsForCollection)
}

func EncodeCodedeployDeploymentGroup_EcsService_ClusterName(p EcsService, vals map[string]cty.Value) {
	vals["cluster_name"] = cty.StringVal(p.ClusterName)
}

func EncodeCodedeployDeploymentGroup_EcsService_ServiceName(p EcsService, vals map[string]cty.Value) {
	vals["service_name"] = cty.StringVal(p.ServiceName)
}

func EncodeCodedeployDeploymentGroup_LoadBalancerInfo(p LoadBalancerInfo, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeCodedeployDeploymentGroup_LoadBalancerInfo_ElbInfo(p.ElbInfo, ctyVal)
	EncodeCodedeployDeploymentGroup_LoadBalancerInfo_TargetGroupInfo(p.TargetGroupInfo, ctyVal)
	EncodeCodedeployDeploymentGroup_LoadBalancerInfo_TargetGroupPairInfo(p.TargetGroupPairInfo, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["load_balancer_info"] = cty.ListVal(valsForCollection)
}

func EncodeCodedeployDeploymentGroup_LoadBalancerInfo_ElbInfo(p ElbInfo, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeCodedeployDeploymentGroup_LoadBalancerInfo_ElbInfo_Name(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["elb_info"] = cty.SetVal(valsForCollection)
}

func EncodeCodedeployDeploymentGroup_LoadBalancerInfo_ElbInfo_Name(p ElbInfo, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeCodedeployDeploymentGroup_LoadBalancerInfo_TargetGroupInfo(p TargetGroupInfo, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeCodedeployDeploymentGroup_LoadBalancerInfo_TargetGroupInfo_Name(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["target_group_info"] = cty.SetVal(valsForCollection)
}

func EncodeCodedeployDeploymentGroup_LoadBalancerInfo_TargetGroupInfo_Name(p TargetGroupInfo, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeCodedeployDeploymentGroup_LoadBalancerInfo_TargetGroupPairInfo(p TargetGroupPairInfo, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeCodedeployDeploymentGroup_LoadBalancerInfo_TargetGroupPairInfo_ProdTrafficRoute(p.ProdTrafficRoute, ctyVal)
	EncodeCodedeployDeploymentGroup_LoadBalancerInfo_TargetGroupPairInfo_TargetGroup(p.TargetGroup, ctyVal)
	EncodeCodedeployDeploymentGroup_LoadBalancerInfo_TargetGroupPairInfo_TestTrafficRoute(p.TestTrafficRoute, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["target_group_pair_info"] = cty.ListVal(valsForCollection)
}

func EncodeCodedeployDeploymentGroup_LoadBalancerInfo_TargetGroupPairInfo_ProdTrafficRoute(p ProdTrafficRoute, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeCodedeployDeploymentGroup_LoadBalancerInfo_TargetGroupPairInfo_ProdTrafficRoute_ListenerArns(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["prod_traffic_route"] = cty.ListVal(valsForCollection)
}

func EncodeCodedeployDeploymentGroup_LoadBalancerInfo_TargetGroupPairInfo_ProdTrafficRoute_ListenerArns(p ProdTrafficRoute, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.ListenerArns {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["listener_arns"] = cty.SetVal(colVals)
}

func EncodeCodedeployDeploymentGroup_LoadBalancerInfo_TargetGroupPairInfo_TargetGroup(p []TargetGroup, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeCodedeployDeploymentGroup_LoadBalancerInfo_TargetGroupPairInfo_TargetGroup_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["target_group"] = cty.ListVal(valsForCollection)
}

func EncodeCodedeployDeploymentGroup_LoadBalancerInfo_TargetGroupPairInfo_TargetGroup_Name(p TargetGroup, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeCodedeployDeploymentGroup_LoadBalancerInfo_TargetGroupPairInfo_TestTrafficRoute(p TestTrafficRoute, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeCodedeployDeploymentGroup_LoadBalancerInfo_TargetGroupPairInfo_TestTrafficRoute_ListenerArns(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["test_traffic_route"] = cty.ListVal(valsForCollection)
}

func EncodeCodedeployDeploymentGroup_LoadBalancerInfo_TargetGroupPairInfo_TestTrafficRoute_ListenerArns(p TestTrafficRoute, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.ListenerArns {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["listener_arns"] = cty.SetVal(colVals)
}

func EncodeCodedeployDeploymentGroup_OnPremisesInstanceTagFilter(p OnPremisesInstanceTagFilter, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeCodedeployDeploymentGroup_OnPremisesInstanceTagFilter_Key(p, ctyVal)
	EncodeCodedeployDeploymentGroup_OnPremisesInstanceTagFilter_Type(p, ctyVal)
	EncodeCodedeployDeploymentGroup_OnPremisesInstanceTagFilter_Value(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["on_premises_instance_tag_filter"] = cty.SetVal(valsForCollection)
}

func EncodeCodedeployDeploymentGroup_OnPremisesInstanceTagFilter_Key(p OnPremisesInstanceTagFilter, vals map[string]cty.Value) {
	vals["key"] = cty.StringVal(p.Key)
}

func EncodeCodedeployDeploymentGroup_OnPremisesInstanceTagFilter_Type(p OnPremisesInstanceTagFilter, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeCodedeployDeploymentGroup_OnPremisesInstanceTagFilter_Value(p OnPremisesInstanceTagFilter, vals map[string]cty.Value) {
	vals["value"] = cty.StringVal(p.Value)
}

func EncodeCodedeployDeploymentGroup_TriggerConfiguration(p TriggerConfiguration, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeCodedeployDeploymentGroup_TriggerConfiguration_TriggerEvents(p, ctyVal)
	EncodeCodedeployDeploymentGroup_TriggerConfiguration_TriggerName(p, ctyVal)
	EncodeCodedeployDeploymentGroup_TriggerConfiguration_TriggerTargetArn(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["trigger_configuration"] = cty.SetVal(valsForCollection)
}

func EncodeCodedeployDeploymentGroup_TriggerConfiguration_TriggerEvents(p TriggerConfiguration, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.TriggerEvents {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["trigger_events"] = cty.SetVal(colVals)
}

func EncodeCodedeployDeploymentGroup_TriggerConfiguration_TriggerName(p TriggerConfiguration, vals map[string]cty.Value) {
	vals["trigger_name"] = cty.StringVal(p.TriggerName)
}

func EncodeCodedeployDeploymentGroup_TriggerConfiguration_TriggerTargetArn(p TriggerConfiguration, vals map[string]cty.Value) {
	vals["trigger_target_arn"] = cty.StringVal(p.TriggerTargetArn)
}

func EncodeCodedeployDeploymentGroup_AlarmConfiguration(p AlarmConfiguration, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeCodedeployDeploymentGroup_AlarmConfiguration_Alarms(p, ctyVal)
	EncodeCodedeployDeploymentGroup_AlarmConfiguration_Enabled(p, ctyVal)
	EncodeCodedeployDeploymentGroup_AlarmConfiguration_IgnorePollAlarmFailure(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["alarm_configuration"] = cty.ListVal(valsForCollection)
}

func EncodeCodedeployDeploymentGroup_AlarmConfiguration_Alarms(p AlarmConfiguration, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Alarms {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["alarms"] = cty.SetVal(colVals)
}

func EncodeCodedeployDeploymentGroup_AlarmConfiguration_Enabled(p AlarmConfiguration, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeCodedeployDeploymentGroup_AlarmConfiguration_IgnorePollAlarmFailure(p AlarmConfiguration, vals map[string]cty.Value) {
	vals["ignore_poll_alarm_failure"] = cty.BoolVal(p.IgnorePollAlarmFailure)
}