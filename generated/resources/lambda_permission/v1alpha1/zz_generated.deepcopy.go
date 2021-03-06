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
func (in *LambdaPermission) DeepCopyInto(out *LambdaPermission) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LambdaPermission.
func (in *LambdaPermission) DeepCopy() *LambdaPermission {
	if in == nil {
		return nil
	}
	out := new(LambdaPermission)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LambdaPermission) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LambdaPermissionList) DeepCopyInto(out *LambdaPermissionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LambdaPermission, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LambdaPermissionList.
func (in *LambdaPermissionList) DeepCopy() *LambdaPermissionList {
	if in == nil {
		return nil
	}
	out := new(LambdaPermissionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LambdaPermissionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LambdaPermissionObservation) DeepCopyInto(out *LambdaPermissionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LambdaPermissionObservation.
func (in *LambdaPermissionObservation) DeepCopy() *LambdaPermissionObservation {
	if in == nil {
		return nil
	}
	out := new(LambdaPermissionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LambdaPermissionParameters) DeepCopyInto(out *LambdaPermissionParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LambdaPermissionParameters.
func (in *LambdaPermissionParameters) DeepCopy() *LambdaPermissionParameters {
	if in == nil {
		return nil
	}
	out := new(LambdaPermissionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LambdaPermissionSpec) DeepCopyInto(out *LambdaPermissionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LambdaPermissionSpec.
func (in *LambdaPermissionSpec) DeepCopy() *LambdaPermissionSpec {
	if in == nil {
		return nil
	}
	out := new(LambdaPermissionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LambdaPermissionStatus) DeepCopyInto(out *LambdaPermissionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LambdaPermissionStatus.
func (in *LambdaPermissionStatus) DeepCopy() *LambdaPermissionStatus {
	if in == nil {
		return nil
	}
	out := new(LambdaPermissionStatus)
	in.DeepCopyInto(out)
	return out
}
