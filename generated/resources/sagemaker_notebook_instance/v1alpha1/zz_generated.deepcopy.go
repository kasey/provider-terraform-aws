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
func (in *SagemakerNotebookInstance) DeepCopyInto(out *SagemakerNotebookInstance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SagemakerNotebookInstance.
func (in *SagemakerNotebookInstance) DeepCopy() *SagemakerNotebookInstance {
	if in == nil {
		return nil
	}
	out := new(SagemakerNotebookInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SagemakerNotebookInstance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SagemakerNotebookInstanceList) DeepCopyInto(out *SagemakerNotebookInstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SagemakerNotebookInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SagemakerNotebookInstanceList.
func (in *SagemakerNotebookInstanceList) DeepCopy() *SagemakerNotebookInstanceList {
	if in == nil {
		return nil
	}
	out := new(SagemakerNotebookInstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SagemakerNotebookInstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SagemakerNotebookInstanceObservation) DeepCopyInto(out *SagemakerNotebookInstanceObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SagemakerNotebookInstanceObservation.
func (in *SagemakerNotebookInstanceObservation) DeepCopy() *SagemakerNotebookInstanceObservation {
	if in == nil {
		return nil
	}
	out := new(SagemakerNotebookInstanceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SagemakerNotebookInstanceParameters) DeepCopyInto(out *SagemakerNotebookInstanceParameters) {
	*out = *in
	if in.SecurityGroups != nil {
		in, out := &in.SecurityGroups, &out.SecurityGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SagemakerNotebookInstanceParameters.
func (in *SagemakerNotebookInstanceParameters) DeepCopy() *SagemakerNotebookInstanceParameters {
	if in == nil {
		return nil
	}
	out := new(SagemakerNotebookInstanceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SagemakerNotebookInstanceSpec) DeepCopyInto(out *SagemakerNotebookInstanceSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SagemakerNotebookInstanceSpec.
func (in *SagemakerNotebookInstanceSpec) DeepCopy() *SagemakerNotebookInstanceSpec {
	if in == nil {
		return nil
	}
	out := new(SagemakerNotebookInstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SagemakerNotebookInstanceStatus) DeepCopyInto(out *SagemakerNotebookInstanceStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SagemakerNotebookInstanceStatus.
func (in *SagemakerNotebookInstanceStatus) DeepCopy() *SagemakerNotebookInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(SagemakerNotebookInstanceStatus)
	in.DeepCopyInto(out)
	return out
}
