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
func (in *VpcEndpointRouteTableAssociation) DeepCopyInto(out *VpcEndpointRouteTableAssociation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcEndpointRouteTableAssociation.
func (in *VpcEndpointRouteTableAssociation) DeepCopy() *VpcEndpointRouteTableAssociation {
	if in == nil {
		return nil
	}
	out := new(VpcEndpointRouteTableAssociation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VpcEndpointRouteTableAssociation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcEndpointRouteTableAssociationList) DeepCopyInto(out *VpcEndpointRouteTableAssociationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VpcEndpointRouteTableAssociation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcEndpointRouteTableAssociationList.
func (in *VpcEndpointRouteTableAssociationList) DeepCopy() *VpcEndpointRouteTableAssociationList {
	if in == nil {
		return nil
	}
	out := new(VpcEndpointRouteTableAssociationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VpcEndpointRouteTableAssociationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcEndpointRouteTableAssociationObservation) DeepCopyInto(out *VpcEndpointRouteTableAssociationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcEndpointRouteTableAssociationObservation.
func (in *VpcEndpointRouteTableAssociationObservation) DeepCopy() *VpcEndpointRouteTableAssociationObservation {
	if in == nil {
		return nil
	}
	out := new(VpcEndpointRouteTableAssociationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcEndpointRouteTableAssociationParameters) DeepCopyInto(out *VpcEndpointRouteTableAssociationParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcEndpointRouteTableAssociationParameters.
func (in *VpcEndpointRouteTableAssociationParameters) DeepCopy() *VpcEndpointRouteTableAssociationParameters {
	if in == nil {
		return nil
	}
	out := new(VpcEndpointRouteTableAssociationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcEndpointRouteTableAssociationSpec) DeepCopyInto(out *VpcEndpointRouteTableAssociationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcEndpointRouteTableAssociationSpec.
func (in *VpcEndpointRouteTableAssociationSpec) DeepCopy() *VpcEndpointRouteTableAssociationSpec {
	if in == nil {
		return nil
	}
	out := new(VpcEndpointRouteTableAssociationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcEndpointRouteTableAssociationStatus) DeepCopyInto(out *VpcEndpointRouteTableAssociationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcEndpointRouteTableAssociationStatus.
func (in *VpcEndpointRouteTableAssociationStatus) DeepCopy() *VpcEndpointRouteTableAssociationStatus {
	if in == nil {
		return nil
	}
	out := new(VpcEndpointRouteTableAssociationStatus)
	in.DeepCopyInto(out)
	return out
}
