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
func (in *LbCookieStickinessPolicy) DeepCopyInto(out *LbCookieStickinessPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LbCookieStickinessPolicy.
func (in *LbCookieStickinessPolicy) DeepCopy() *LbCookieStickinessPolicy {
	if in == nil {
		return nil
	}
	out := new(LbCookieStickinessPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LbCookieStickinessPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LbCookieStickinessPolicyList) DeepCopyInto(out *LbCookieStickinessPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LbCookieStickinessPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LbCookieStickinessPolicyList.
func (in *LbCookieStickinessPolicyList) DeepCopy() *LbCookieStickinessPolicyList {
	if in == nil {
		return nil
	}
	out := new(LbCookieStickinessPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LbCookieStickinessPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LbCookieStickinessPolicyObservation) DeepCopyInto(out *LbCookieStickinessPolicyObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LbCookieStickinessPolicyObservation.
func (in *LbCookieStickinessPolicyObservation) DeepCopy() *LbCookieStickinessPolicyObservation {
	if in == nil {
		return nil
	}
	out := new(LbCookieStickinessPolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LbCookieStickinessPolicyParameters) DeepCopyInto(out *LbCookieStickinessPolicyParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LbCookieStickinessPolicyParameters.
func (in *LbCookieStickinessPolicyParameters) DeepCopy() *LbCookieStickinessPolicyParameters {
	if in == nil {
		return nil
	}
	out := new(LbCookieStickinessPolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LbCookieStickinessPolicySpec) DeepCopyInto(out *LbCookieStickinessPolicySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LbCookieStickinessPolicySpec.
func (in *LbCookieStickinessPolicySpec) DeepCopy() *LbCookieStickinessPolicySpec {
	if in == nil {
		return nil
	}
	out := new(LbCookieStickinessPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LbCookieStickinessPolicyStatus) DeepCopyInto(out *LbCookieStickinessPolicyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LbCookieStickinessPolicyStatus.
func (in *LbCookieStickinessPolicyStatus) DeepCopy() *LbCookieStickinessPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(LbCookieStickinessPolicyStatus)
	in.DeepCopyInto(out)
	return out
}
