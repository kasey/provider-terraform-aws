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
func (in *ServicequotasServiceQuota) DeepCopyInto(out *ServicequotasServiceQuota) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicequotasServiceQuota.
func (in *ServicequotasServiceQuota) DeepCopy() *ServicequotasServiceQuota {
	if in == nil {
		return nil
	}
	out := new(ServicequotasServiceQuota)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServicequotasServiceQuota) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicequotasServiceQuotaList) DeepCopyInto(out *ServicequotasServiceQuotaList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServicequotasServiceQuota, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicequotasServiceQuotaList.
func (in *ServicequotasServiceQuotaList) DeepCopy() *ServicequotasServiceQuotaList {
	if in == nil {
		return nil
	}
	out := new(ServicequotasServiceQuotaList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServicequotasServiceQuotaList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicequotasServiceQuotaObservation) DeepCopyInto(out *ServicequotasServiceQuotaObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicequotasServiceQuotaObservation.
func (in *ServicequotasServiceQuotaObservation) DeepCopy() *ServicequotasServiceQuotaObservation {
	if in == nil {
		return nil
	}
	out := new(ServicequotasServiceQuotaObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicequotasServiceQuotaParameters) DeepCopyInto(out *ServicequotasServiceQuotaParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicequotasServiceQuotaParameters.
func (in *ServicequotasServiceQuotaParameters) DeepCopy() *ServicequotasServiceQuotaParameters {
	if in == nil {
		return nil
	}
	out := new(ServicequotasServiceQuotaParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicequotasServiceQuotaSpec) DeepCopyInto(out *ServicequotasServiceQuotaSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicequotasServiceQuotaSpec.
func (in *ServicequotasServiceQuotaSpec) DeepCopy() *ServicequotasServiceQuotaSpec {
	if in == nil {
		return nil
	}
	out := new(ServicequotasServiceQuotaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicequotasServiceQuotaStatus) DeepCopyInto(out *ServicequotasServiceQuotaStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicequotasServiceQuotaStatus.
func (in *ServicequotasServiceQuotaStatus) DeepCopy() *ServicequotasServiceQuotaStatus {
	if in == nil {
		return nil
	}
	out := new(ServicequotasServiceQuotaStatus)
	in.DeepCopyInto(out)
	return out
}
