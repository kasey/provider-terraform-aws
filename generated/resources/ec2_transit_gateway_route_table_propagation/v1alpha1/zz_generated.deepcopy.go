// +build !ignore_autogenerated

/*
Copyright 2020 The Crossplane Authors.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ec2TransitGatewayRouteTablePropagation) DeepCopyInto(out *Ec2TransitGatewayRouteTablePropagation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ec2TransitGatewayRouteTablePropagation.
func (in *Ec2TransitGatewayRouteTablePropagation) DeepCopy() *Ec2TransitGatewayRouteTablePropagation {
	if in == nil {
		return nil
	}
	out := new(Ec2TransitGatewayRouteTablePropagation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Ec2TransitGatewayRouteTablePropagation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ec2TransitGatewayRouteTablePropagationList) DeepCopyInto(out *Ec2TransitGatewayRouteTablePropagationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Ec2TransitGatewayRouteTablePropagation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ec2TransitGatewayRouteTablePropagationList.
func (in *Ec2TransitGatewayRouteTablePropagationList) DeepCopy() *Ec2TransitGatewayRouteTablePropagationList {
	if in == nil {
		return nil
	}
	out := new(Ec2TransitGatewayRouteTablePropagationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Ec2TransitGatewayRouteTablePropagationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ec2TransitGatewayRouteTablePropagationObservation) DeepCopyInto(out *Ec2TransitGatewayRouteTablePropagationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ec2TransitGatewayRouteTablePropagationObservation.
func (in *Ec2TransitGatewayRouteTablePropagationObservation) DeepCopy() *Ec2TransitGatewayRouteTablePropagationObservation {
	if in == nil {
		return nil
	}
	out := new(Ec2TransitGatewayRouteTablePropagationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ec2TransitGatewayRouteTablePropagationParameters) DeepCopyInto(out *Ec2TransitGatewayRouteTablePropagationParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ec2TransitGatewayRouteTablePropagationParameters.
func (in *Ec2TransitGatewayRouteTablePropagationParameters) DeepCopy() *Ec2TransitGatewayRouteTablePropagationParameters {
	if in == nil {
		return nil
	}
	out := new(Ec2TransitGatewayRouteTablePropagationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ec2TransitGatewayRouteTablePropagationSpec) DeepCopyInto(out *Ec2TransitGatewayRouteTablePropagationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ec2TransitGatewayRouteTablePropagationSpec.
func (in *Ec2TransitGatewayRouteTablePropagationSpec) DeepCopy() *Ec2TransitGatewayRouteTablePropagationSpec {
	if in == nil {
		return nil
	}
	out := new(Ec2TransitGatewayRouteTablePropagationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ec2TransitGatewayRouteTablePropagationStatus) DeepCopyInto(out *Ec2TransitGatewayRouteTablePropagationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ec2TransitGatewayRouteTablePropagationStatus.
func (in *Ec2TransitGatewayRouteTablePropagationStatus) DeepCopy() *Ec2TransitGatewayRouteTablePropagationStatus {
	if in == nil {
		return nil
	}
	out := new(Ec2TransitGatewayRouteTablePropagationStatus)
	in.DeepCopyInto(out)
	return out
}
