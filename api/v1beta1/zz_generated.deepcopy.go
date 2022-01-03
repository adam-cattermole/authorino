// +build !ignore_autogenerated

/*
Copyright 2020 Red Hat, Inc.

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

package v1beta1

import (
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthConfig) DeepCopyInto(out *AuthConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthConfig.
func (in *AuthConfig) DeepCopy() *AuthConfig {
	if in == nil {
		return nil
	}
	out := new(AuthConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AuthConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthConfigList) DeepCopyInto(out *AuthConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AuthConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthConfigList.
func (in *AuthConfigList) DeepCopy() *AuthConfigList {
	if in == nil {
		return nil
	}
	out := new(AuthConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AuthConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthConfigSpec) DeepCopyInto(out *AuthConfigSpec) {
	*out = *in
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Patterns != nil {
		in, out := &in.Patterns, &out.Patterns
		*out = make(map[string]JSONPatternExpressions, len(*in))
		for key, val := range *in {
			var outVal []JSONPatternExpression
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(JSONPatternExpressions, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]JSONPattern, len(*in))
		copy(*out, *in)
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = make([]*Identity, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Identity)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = make([]*Metadata, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Metadata)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Authorization != nil {
		in, out := &in.Authorization, &out.Authorization
		*out = make([]*Authorization, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Authorization)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Response != nil {
		in, out := &in.Response, &out.Response
		*out = make([]*Response, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Response)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.DenyWith != nil {
		in, out := &in.DenyWith, &out.DenyWith
		*out = new(DenyWith)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthConfigSpec.
func (in *AuthConfigSpec) DeepCopy() *AuthConfigSpec {
	if in == nil {
		return nil
	}
	out := new(AuthConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthConfigStatus) DeepCopyInto(out *AuthConfigStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthConfigStatus.
func (in *AuthConfigStatus) DeepCopy() *AuthConfigStatus {
	if in == nil {
		return nil
	}
	out := new(AuthConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Authorization) DeepCopyInto(out *Authorization) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]JSONPattern, len(*in))
		copy(*out, *in)
	}
	if in.OPA != nil {
		in, out := &in.OPA, &out.OPA
		*out = new(Authorization_OPA)
		(*in).DeepCopyInto(*out)
	}
	if in.JSON != nil {
		in, out := &in.JSON, &out.JSON
		*out = new(Authorization_JSONPatternMatching)
		(*in).DeepCopyInto(*out)
	}
	if in.KubernetesAuthz != nil {
		in, out := &in.KubernetesAuthz, &out.KubernetesAuthz
		*out = new(Authorization_KubernetesAuthz)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Authorization.
func (in *Authorization) DeepCopy() *Authorization {
	if in == nil {
		return nil
	}
	out := new(Authorization)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Authorization_JSONPatternMatching) DeepCopyInto(out *Authorization_JSONPatternMatching) {
	*out = *in
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]JSONPattern, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Authorization_JSONPatternMatching.
func (in *Authorization_JSONPatternMatching) DeepCopy() *Authorization_JSONPatternMatching {
	if in == nil {
		return nil
	}
	out := new(Authorization_JSONPatternMatching)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Authorization_KubernetesAuthz) DeepCopyInto(out *Authorization_KubernetesAuthz) {
	*out = *in
	out.User = in.User
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ResourceAttributes != nil {
		in, out := &in.ResourceAttributes, &out.ResourceAttributes
		*out = new(Authorization_KubernetesAuthz_ResourceAttributes)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Authorization_KubernetesAuthz.
func (in *Authorization_KubernetesAuthz) DeepCopy() *Authorization_KubernetesAuthz {
	if in == nil {
		return nil
	}
	out := new(Authorization_KubernetesAuthz)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Authorization_KubernetesAuthz_Attribute) DeepCopyInto(out *Authorization_KubernetesAuthz_Attribute) {
	*out = *in
	out.ValueFrom = in.ValueFrom
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Authorization_KubernetesAuthz_Attribute.
func (in *Authorization_KubernetesAuthz_Attribute) DeepCopy() *Authorization_KubernetesAuthz_Attribute {
	if in == nil {
		return nil
	}
	out := new(Authorization_KubernetesAuthz_Attribute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Authorization_KubernetesAuthz_ResourceAttributes) DeepCopyInto(out *Authorization_KubernetesAuthz_ResourceAttributes) {
	*out = *in
	out.Namespace = in.Namespace
	out.Group = in.Group
	out.Resource = in.Resource
	out.Name = in.Name
	out.SubResource = in.SubResource
	out.Verb = in.Verb
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Authorization_KubernetesAuthz_ResourceAttributes.
func (in *Authorization_KubernetesAuthz_ResourceAttributes) DeepCopy() *Authorization_KubernetesAuthz_ResourceAttributes {
	if in == nil {
		return nil
	}
	out := new(Authorization_KubernetesAuthz_ResourceAttributes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Authorization_OPA) DeepCopyInto(out *Authorization_OPA) {
	*out = *in
	in.ExternalRegistry.DeepCopyInto(&out.ExternalRegistry)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Authorization_OPA.
func (in *Authorization_OPA) DeepCopy() *Authorization_OPA {
	if in == nil {
		return nil
	}
	out := new(Authorization_OPA)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Credentials) DeepCopyInto(out *Credentials) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Credentials.
func (in *Credentials) DeepCopy() *Credentials {
	if in == nil {
		return nil
	}
	out := new(Credentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DenyWith) DeepCopyInto(out *DenyWith) {
	*out = *in
	if in.Unauthenticated != nil {
		in, out := &in.Unauthenticated, &out.Unauthenticated
		*out = new(DenyWithSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Unauthorized != nil {
		in, out := &in.Unauthorized, &out.Unauthorized
		*out = new(DenyWithSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DenyWith.
func (in *DenyWith) DeepCopy() *DenyWith {
	if in == nil {
		return nil
	}
	out := new(DenyWith)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DenyWithSpec) DeepCopyInto(out *DenyWithSpec) {
	*out = *in
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make([]JsonProperty, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DenyWithSpec.
func (in *DenyWithSpec) DeepCopy() *DenyWithSpec {
	if in == nil {
		return nil
	}
	out := new(DenyWithSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalRegistry) DeepCopyInto(out *ExternalRegistry) {
	*out = *in
	if in.SharedSecret != nil {
		in, out := &in.SharedSecret, &out.SharedSecret
		*out = new(SecretKeyReference)
		**out = **in
	}
	out.Credentials = in.Credentials
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalRegistry.
func (in *ExternalRegistry) DeepCopy() *ExternalRegistry {
	if in == nil {
		return nil
	}
	out := new(ExternalRegistry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Identity) DeepCopyInto(out *Identity) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]JSONPattern, len(*in))
		copy(*out, *in)
	}
	out.Credentials = in.Credentials
	if in.ExtendedProperties != nil {
		in, out := &in.ExtendedProperties, &out.ExtendedProperties
		*out = make([]JsonProperty, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.OAuth2 != nil {
		in, out := &in.OAuth2, &out.OAuth2
		*out = new(Identity_OAuth2Config)
		(*in).DeepCopyInto(*out)
	}
	if in.Oidc != nil {
		in, out := &in.Oidc, &out.Oidc
		*out = new(Identity_OidcConfig)
		**out = **in
	}
	if in.APIKey != nil {
		in, out := &in.APIKey, &out.APIKey
		*out = new(Identity_APIKey)
		(*in).DeepCopyInto(*out)
	}
	if in.KubernetesAuth != nil {
		in, out := &in.KubernetesAuth, &out.KubernetesAuth
		*out = new(Identity_KubernetesAuth)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Identity.
func (in *Identity) DeepCopy() *Identity {
	if in == nil {
		return nil
	}
	out := new(Identity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Identity_APIKey) DeepCopyInto(out *Identity_APIKey) {
	*out = *in
	if in.LabelSelectors != nil {
		in, out := &in.LabelSelectors, &out.LabelSelectors
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Identity_APIKey.
func (in *Identity_APIKey) DeepCopy() *Identity_APIKey {
	if in == nil {
		return nil
	}
	out := new(Identity_APIKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Identity_KubernetesAuth) DeepCopyInto(out *Identity_KubernetesAuth) {
	*out = *in
	if in.Audiences != nil {
		in, out := &in.Audiences, &out.Audiences
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Identity_KubernetesAuth.
func (in *Identity_KubernetesAuth) DeepCopy() *Identity_KubernetesAuth {
	if in == nil {
		return nil
	}
	out := new(Identity_KubernetesAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Identity_OAuth2Config) DeepCopyInto(out *Identity_OAuth2Config) {
	*out = *in
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Identity_OAuth2Config.
func (in *Identity_OAuth2Config) DeepCopy() *Identity_OAuth2Config {
	if in == nil {
		return nil
	}
	out := new(Identity_OAuth2Config)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Identity_OidcConfig) DeepCopyInto(out *Identity_OidcConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Identity_OidcConfig.
func (in *Identity_OidcConfig) DeepCopy() *Identity_OidcConfig {
	if in == nil {
		return nil
	}
	out := new(Identity_OidcConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JSONPattern) DeepCopyInto(out *JSONPattern) {
	*out = *in
	out.JSONPatternRef = in.JSONPatternRef
	out.JSONPatternExpression = in.JSONPatternExpression
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JSONPattern.
func (in *JSONPattern) DeepCopy() *JSONPattern {
	if in == nil {
		return nil
	}
	out := new(JSONPattern)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JSONPatternExpression) DeepCopyInto(out *JSONPatternExpression) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JSONPatternExpression.
func (in *JSONPatternExpression) DeepCopy() *JSONPatternExpression {
	if in == nil {
		return nil
	}
	out := new(JSONPatternExpression)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in JSONPatternExpressions) DeepCopyInto(out *JSONPatternExpressions) {
	{
		in := &in
		*out = make(JSONPatternExpressions, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JSONPatternExpressions.
func (in JSONPatternExpressions) DeepCopy() JSONPatternExpressions {
	if in == nil {
		return nil
	}
	out := new(JSONPatternExpressions)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JSONPatternRef) DeepCopyInto(out *JSONPatternRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JSONPatternRef.
func (in *JSONPatternRef) DeepCopy() *JSONPatternRef {
	if in == nil {
		return nil
	}
	out := new(JSONPatternRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JsonProperty) DeepCopyInto(out *JsonProperty) {
	*out = *in
	in.Value.DeepCopyInto(&out.Value)
	out.ValueFrom = in.ValueFrom
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JsonProperty.
func (in *JsonProperty) DeepCopy() *JsonProperty {
	if in == nil {
		return nil
	}
	out := new(JsonProperty)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Metadata) DeepCopyInto(out *Metadata) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]JSONPattern, len(*in))
		copy(*out, *in)
	}
	if in.UserInfo != nil {
		in, out := &in.UserInfo, &out.UserInfo
		*out = new(Metadata_UserInfo)
		**out = **in
	}
	if in.UMA != nil {
		in, out := &in.UMA, &out.UMA
		*out = new(Metadata_UMA)
		(*in).DeepCopyInto(*out)
	}
	if in.GenericHTTP != nil {
		in, out := &in.GenericHTTP, &out.GenericHTTP
		*out = new(Metadata_GenericHTTP)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Metadata.
func (in *Metadata) DeepCopy() *Metadata {
	if in == nil {
		return nil
	}
	out := new(Metadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Metadata_GenericHTTP) DeepCopyInto(out *Metadata_GenericHTTP) {
	*out = *in
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make([]JsonProperty, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make([]JsonProperty, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SharedSecret != nil {
		in, out := &in.SharedSecret, &out.SharedSecret
		*out = new(SecretKeyReference)
		**out = **in
	}
	out.Credentials = in.Credentials
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Metadata_GenericHTTP.
func (in *Metadata_GenericHTTP) DeepCopy() *Metadata_GenericHTTP {
	if in == nil {
		return nil
	}
	out := new(Metadata_GenericHTTP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Metadata_UMA) DeepCopyInto(out *Metadata_UMA) {
	*out = *in
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Metadata_UMA.
func (in *Metadata_UMA) DeepCopy() *Metadata_UMA {
	if in == nil {
		return nil
	}
	out := new(Metadata_UMA)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Metadata_UserInfo) DeepCopyInto(out *Metadata_UserInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Metadata_UserInfo.
func (in *Metadata_UserInfo) DeepCopy() *Metadata_UserInfo {
	if in == nil {
		return nil
	}
	out := new(Metadata_UserInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Response) DeepCopyInto(out *Response) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]JSONPattern, len(*in))
		copy(*out, *in)
	}
	if in.Wristband != nil {
		in, out := &in.Wristband, &out.Wristband
		*out = new(Response_Wristband)
		(*in).DeepCopyInto(*out)
	}
	if in.JSON != nil {
		in, out := &in.JSON, &out.JSON
		*out = new(Response_DynamicJSON)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Response.
func (in *Response) DeepCopy() *Response {
	if in == nil {
		return nil
	}
	out := new(Response)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Response_DynamicJSON) DeepCopyInto(out *Response_DynamicJSON) {
	*out = *in
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make([]JsonProperty, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Response_DynamicJSON.
func (in *Response_DynamicJSON) DeepCopy() *Response_DynamicJSON {
	if in == nil {
		return nil
	}
	out := new(Response_DynamicJSON)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Response_Wristband) DeepCopyInto(out *Response_Wristband) {
	*out = *in
	if in.CustomClaims != nil {
		in, out := &in.CustomClaims, &out.CustomClaims
		*out = make([]JsonProperty, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TokenDuration != nil {
		in, out := &in.TokenDuration, &out.TokenDuration
		*out = new(int64)
		**out = **in
	}
	if in.SigningKeyRefs != nil {
		in, out := &in.SigningKeyRefs, &out.SigningKeyRefs
		*out = make([]*SigningKeyRef, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(SigningKeyRef)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Response_Wristband.
func (in *Response_Wristband) DeepCopy() *Response_Wristband {
	if in == nil {
		return nil
	}
	out := new(Response_Wristband)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretKeyReference) DeepCopyInto(out *SecretKeyReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretKeyReference.
func (in *SecretKeyReference) DeepCopy() *SecretKeyReference {
	if in == nil {
		return nil
	}
	out := new(SecretKeyReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SigningKeyRef) DeepCopyInto(out *SigningKeyRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SigningKeyRef.
func (in *SigningKeyRef) DeepCopy() *SigningKeyRef {
	if in == nil {
		return nil
	}
	out := new(SigningKeyRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValueFromAuthJSON) DeepCopyInto(out *ValueFromAuthJSON) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValueFromAuthJSON.
func (in *ValueFromAuthJSON) DeepCopy() *ValueFromAuthJSON {
	if in == nil {
		return nil
	}
	out := new(ValueFromAuthJSON)
	in.DeepCopyInto(out)
	return out
}
