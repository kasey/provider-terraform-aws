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
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import runtimev1alpha1 "github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"

// GetCondition of this ShieldProtection.
func (mg *ShieldProtection) GetCondition(ct runtimev1alpha1.ConditionType) runtimev1alpha1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this ShieldProtection.
func (mg *ShieldProtection) GetDeletionPolicy() runtimev1alpha1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this ShieldProtection.
func (mg *ShieldProtection) GetProviderConfigReference() *runtimev1alpha1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this ShieldProtection.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *ShieldProtection) GetProviderReference() *runtimev1alpha1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this ShieldProtection.
func (mg *ShieldProtection) GetWriteConnectionSecretToReference() *runtimev1alpha1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this ShieldProtection.
func (mg *ShieldProtection) SetConditions(c ...runtimev1alpha1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this ShieldProtection.
func (mg *ShieldProtection) SetDeletionPolicy(r runtimev1alpha1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this ShieldProtection.
func (mg *ShieldProtection) SetProviderConfigReference(r *runtimev1alpha1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this ShieldProtection.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *ShieldProtection) SetProviderReference(r *runtimev1alpha1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this ShieldProtection.
func (mg *ShieldProtection) SetWriteConnectionSecretToReference(r *runtimev1alpha1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
