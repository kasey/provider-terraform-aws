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
func (in *Apigatewayv2RouteResponse) DeepCopyInto(out *Apigatewayv2RouteResponse) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Apigatewayv2RouteResponse.
func (in *Apigatewayv2RouteResponse) DeepCopy() *Apigatewayv2RouteResponse {
	if in == nil {
		return nil
	}
	out := new(Apigatewayv2RouteResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Apigatewayv2RouteResponse) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Apigatewayv2RouteResponseList) DeepCopyInto(out *Apigatewayv2RouteResponseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Apigatewayv2RouteResponse, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Apigatewayv2RouteResponseList.
func (in *Apigatewayv2RouteResponseList) DeepCopy() *Apigatewayv2RouteResponseList {
	if in == nil {
		return nil
	}
	out := new(Apigatewayv2RouteResponseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Apigatewayv2RouteResponseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Apigatewayv2RouteResponseObservation) DeepCopyInto(out *Apigatewayv2RouteResponseObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Apigatewayv2RouteResponseObservation.
func (in *Apigatewayv2RouteResponseObservation) DeepCopy() *Apigatewayv2RouteResponseObservation {
	if in == nil {
		return nil
	}
	out := new(Apigatewayv2RouteResponseObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Apigatewayv2RouteResponseParameters) DeepCopyInto(out *Apigatewayv2RouteResponseParameters) {
	*out = *in
	if in.ResponseModels != nil {
		in, out := &in.ResponseModels, &out.ResponseModels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Apigatewayv2RouteResponseParameters.
func (in *Apigatewayv2RouteResponseParameters) DeepCopy() *Apigatewayv2RouteResponseParameters {
	if in == nil {
		return nil
	}
	out := new(Apigatewayv2RouteResponseParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Apigatewayv2RouteResponseSpec) DeepCopyInto(out *Apigatewayv2RouteResponseSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Apigatewayv2RouteResponseSpec.
func (in *Apigatewayv2RouteResponseSpec) DeepCopy() *Apigatewayv2RouteResponseSpec {
	if in == nil {
		return nil
	}
	out := new(Apigatewayv2RouteResponseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Apigatewayv2RouteResponseStatus) DeepCopyInto(out *Apigatewayv2RouteResponseStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Apigatewayv2RouteResponseStatus.
func (in *Apigatewayv2RouteResponseStatus) DeepCopy() *Apigatewayv2RouteResponseStatus {
	if in == nil {
		return nil
	}
	out := new(Apigatewayv2RouteResponseStatus)
	in.DeepCopyInto(out)
	return out
}
