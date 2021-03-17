// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v2alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApisixRoute) DeepCopyInto(out *ApisixRoute) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(ApisixRouteSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApisixRoute.
func (in *ApisixRoute) DeepCopy() *ApisixRoute {
	if in == nil {
		return nil
	}
	out := new(ApisixRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApisixRoute) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApisixRouteHTTP) DeepCopyInto(out *ApisixRouteHTTP) {
	*out = *in
	if in.Match != nil {
		in, out := &in.Match, &out.Match
		*out = new(ApisixRouteHTTPMatch)
		(*in).DeepCopyInto(*out)
	}
	if in.Backend != nil {
		in, out := &in.Backend, &out.Backend
		*out = new(ApisixRouteHTTPBackend)
		**out = **in
	}
	if in.Plugins != nil {
		in, out := &in.Plugins, &out.Plugins
		*out = make([]*ApisixRouteHTTPPlugin, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ApisixRouteHTTPPlugin)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApisixRouteHTTP.
func (in *ApisixRouteHTTP) DeepCopy() *ApisixRouteHTTP {
	if in == nil {
		return nil
	}
	out := new(ApisixRouteHTTP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApisixRouteHTTPBackend) DeepCopyInto(out *ApisixRouteHTTPBackend) {
	*out = *in
	out.ServicePort = in.ServicePort
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApisixRouteHTTPBackend.
func (in *ApisixRouteHTTPBackend) DeepCopy() *ApisixRouteHTTPBackend {
	if in == nil {
		return nil
	}
	out := new(ApisixRouteHTTPBackend)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApisixRouteHTTPMatch) DeepCopyInto(out *ApisixRouteHTTPMatch) {
	*out = *in
	if in.Paths != nil {
		in, out := &in.Paths, &out.Paths
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Methods != nil {
		in, out := &in.Methods, &out.Methods
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RemoteAddrs != nil {
		in, out := &in.RemoteAddrs, &out.RemoteAddrs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NginxVars != nil {
		in, out := &in.NginxVars, &out.NginxVars
		*out = make([]ApisixRouteHTTPMatchNginxVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApisixRouteHTTPMatch.
func (in *ApisixRouteHTTPMatch) DeepCopy() *ApisixRouteHTTPMatch {
	if in == nil {
		return nil
	}
	out := new(ApisixRouteHTTPMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApisixRouteHTTPMatchNginxVar) DeepCopyInto(out *ApisixRouteHTTPMatchNginxVar) {
	*out = *in
	if in.Set != nil {
		in, out := &in.Set, &out.Set
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApisixRouteHTTPMatchNginxVar.
func (in *ApisixRouteHTTPMatchNginxVar) DeepCopy() *ApisixRouteHTTPMatchNginxVar {
	if in == nil {
		return nil
	}
	out := new(ApisixRouteHTTPMatchNginxVar)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApisixRouteHTTPPlugin) DeepCopyInto(out *ApisixRouteHTTPPlugin) {
	*out = *in
	in.Config.DeepCopyInto(&out.Config)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApisixRouteHTTPPlugin.
func (in *ApisixRouteHTTPPlugin) DeepCopy() *ApisixRouteHTTPPlugin {
	if in == nil {
		return nil
	}
	out := new(ApisixRouteHTTPPlugin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApisixRouteList) DeepCopyInto(out *ApisixRouteList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ApisixRoute, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApisixRouteList.
func (in *ApisixRouteList) DeepCopy() *ApisixRouteList {
	if in == nil {
		return nil
	}
	out := new(ApisixRouteList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApisixRouteList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApisixRouteSpec) DeepCopyInto(out *ApisixRouteSpec) {
	*out = *in
	if in.HTTP != nil {
		in, out := &in.HTTP, &out.HTTP
		*out = make([]*ApisixRouteHTTP, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ApisixRouteHTTP)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApisixRouteSpec.
func (in *ApisixRouteSpec) DeepCopy() *ApisixRouteSpec {
	if in == nil {
		return nil
	}
	out := new(ApisixRouteSpec)
	in.DeepCopyInto(out)
	return out
}
