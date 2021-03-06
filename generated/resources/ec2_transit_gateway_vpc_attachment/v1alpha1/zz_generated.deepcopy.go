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
func (in *Ec2TransitGatewayVpcAttachment) DeepCopyInto(out *Ec2TransitGatewayVpcAttachment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ec2TransitGatewayVpcAttachment.
func (in *Ec2TransitGatewayVpcAttachment) DeepCopy() *Ec2TransitGatewayVpcAttachment {
	if in == nil {
		return nil
	}
	out := new(Ec2TransitGatewayVpcAttachment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Ec2TransitGatewayVpcAttachment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ec2TransitGatewayVpcAttachmentList) DeepCopyInto(out *Ec2TransitGatewayVpcAttachmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Ec2TransitGatewayVpcAttachment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ec2TransitGatewayVpcAttachmentList.
func (in *Ec2TransitGatewayVpcAttachmentList) DeepCopy() *Ec2TransitGatewayVpcAttachmentList {
	if in == nil {
		return nil
	}
	out := new(Ec2TransitGatewayVpcAttachmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Ec2TransitGatewayVpcAttachmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ec2TransitGatewayVpcAttachmentObservation) DeepCopyInto(out *Ec2TransitGatewayVpcAttachmentObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ec2TransitGatewayVpcAttachmentObservation.
func (in *Ec2TransitGatewayVpcAttachmentObservation) DeepCopy() *Ec2TransitGatewayVpcAttachmentObservation {
	if in == nil {
		return nil
	}
	out := new(Ec2TransitGatewayVpcAttachmentObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ec2TransitGatewayVpcAttachmentParameters) DeepCopyInto(out *Ec2TransitGatewayVpcAttachmentParameters) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SubnetIds != nil {
		in, out := &in.SubnetIds, &out.SubnetIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ec2TransitGatewayVpcAttachmentParameters.
func (in *Ec2TransitGatewayVpcAttachmentParameters) DeepCopy() *Ec2TransitGatewayVpcAttachmentParameters {
	if in == nil {
		return nil
	}
	out := new(Ec2TransitGatewayVpcAttachmentParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ec2TransitGatewayVpcAttachmentSpec) DeepCopyInto(out *Ec2TransitGatewayVpcAttachmentSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ec2TransitGatewayVpcAttachmentSpec.
func (in *Ec2TransitGatewayVpcAttachmentSpec) DeepCopy() *Ec2TransitGatewayVpcAttachmentSpec {
	if in == nil {
		return nil
	}
	out := new(Ec2TransitGatewayVpcAttachmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ec2TransitGatewayVpcAttachmentStatus) DeepCopyInto(out *Ec2TransitGatewayVpcAttachmentStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ec2TransitGatewayVpcAttachmentStatus.
func (in *Ec2TransitGatewayVpcAttachmentStatus) DeepCopy() *Ec2TransitGatewayVpcAttachmentStatus {
	if in == nil {
		return nil
	}
	out := new(Ec2TransitGatewayVpcAttachmentStatus)
	in.DeepCopyInto(out)
	return out
}
