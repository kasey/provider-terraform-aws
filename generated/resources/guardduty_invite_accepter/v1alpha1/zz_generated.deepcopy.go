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
func (in *GuarddutyInviteAccepter) DeepCopyInto(out *GuarddutyInviteAccepter) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GuarddutyInviteAccepter.
func (in *GuarddutyInviteAccepter) DeepCopy() *GuarddutyInviteAccepter {
	if in == nil {
		return nil
	}
	out := new(GuarddutyInviteAccepter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GuarddutyInviteAccepter) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GuarddutyInviteAccepterList) DeepCopyInto(out *GuarddutyInviteAccepterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GuarddutyInviteAccepter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GuarddutyInviteAccepterList.
func (in *GuarddutyInviteAccepterList) DeepCopy() *GuarddutyInviteAccepterList {
	if in == nil {
		return nil
	}
	out := new(GuarddutyInviteAccepterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GuarddutyInviteAccepterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GuarddutyInviteAccepterObservation) DeepCopyInto(out *GuarddutyInviteAccepterObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GuarddutyInviteAccepterObservation.
func (in *GuarddutyInviteAccepterObservation) DeepCopy() *GuarddutyInviteAccepterObservation {
	if in == nil {
		return nil
	}
	out := new(GuarddutyInviteAccepterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GuarddutyInviteAccepterParameters) DeepCopyInto(out *GuarddutyInviteAccepterParameters) {
	*out = *in
	out.Timeouts = in.Timeouts
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GuarddutyInviteAccepterParameters.
func (in *GuarddutyInviteAccepterParameters) DeepCopy() *GuarddutyInviteAccepterParameters {
	if in == nil {
		return nil
	}
	out := new(GuarddutyInviteAccepterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GuarddutyInviteAccepterSpec) DeepCopyInto(out *GuarddutyInviteAccepterSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GuarddutyInviteAccepterSpec.
func (in *GuarddutyInviteAccepterSpec) DeepCopy() *GuarddutyInviteAccepterSpec {
	if in == nil {
		return nil
	}
	out := new(GuarddutyInviteAccepterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GuarddutyInviteAccepterStatus) DeepCopyInto(out *GuarddutyInviteAccepterStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GuarddutyInviteAccepterStatus.
func (in *GuarddutyInviteAccepterStatus) DeepCopy() *GuarddutyInviteAccepterStatus {
	if in == nil {
		return nil
	}
	out := new(GuarddutyInviteAccepterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Timeouts) DeepCopyInto(out *Timeouts) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Timeouts.
func (in *Timeouts) DeepCopy() *Timeouts {
	if in == nil {
		return nil
	}
	out := new(Timeouts)
	in.DeepCopyInto(out)
	return out
}
