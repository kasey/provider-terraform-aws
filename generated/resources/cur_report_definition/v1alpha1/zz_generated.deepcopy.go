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
func (in *CurReportDefinition) DeepCopyInto(out *CurReportDefinition) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CurReportDefinition.
func (in *CurReportDefinition) DeepCopy() *CurReportDefinition {
	if in == nil {
		return nil
	}
	out := new(CurReportDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CurReportDefinition) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CurReportDefinitionList) DeepCopyInto(out *CurReportDefinitionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CurReportDefinition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CurReportDefinitionList.
func (in *CurReportDefinitionList) DeepCopy() *CurReportDefinitionList {
	if in == nil {
		return nil
	}
	out := new(CurReportDefinitionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CurReportDefinitionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CurReportDefinitionObservation) DeepCopyInto(out *CurReportDefinitionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CurReportDefinitionObservation.
func (in *CurReportDefinitionObservation) DeepCopy() *CurReportDefinitionObservation {
	if in == nil {
		return nil
	}
	out := new(CurReportDefinitionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CurReportDefinitionParameters) DeepCopyInto(out *CurReportDefinitionParameters) {
	*out = *in
	if in.AdditionalSchemaElements != nil {
		in, out := &in.AdditionalSchemaElements, &out.AdditionalSchemaElements
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AdditionalArtifacts != nil {
		in, out := &in.AdditionalArtifacts, &out.AdditionalArtifacts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CurReportDefinitionParameters.
func (in *CurReportDefinitionParameters) DeepCopy() *CurReportDefinitionParameters {
	if in == nil {
		return nil
	}
	out := new(CurReportDefinitionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CurReportDefinitionSpec) DeepCopyInto(out *CurReportDefinitionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CurReportDefinitionSpec.
func (in *CurReportDefinitionSpec) DeepCopy() *CurReportDefinitionSpec {
	if in == nil {
		return nil
	}
	out := new(CurReportDefinitionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CurReportDefinitionStatus) DeepCopyInto(out *CurReportDefinitionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CurReportDefinitionStatus.
func (in *CurReportDefinitionStatus) DeepCopy() *CurReportDefinitionStatus {
	if in == nil {
		return nil
	}
	out := new(CurReportDefinitionStatus)
	in.DeepCopyInto(out)
	return out
}
