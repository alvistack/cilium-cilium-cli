//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Code generated by deepcopy-gen. DO NOT EDIT.

package store

import (
	loadbalancer "github.com/cilium/cilium/pkg/loadbalancer"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterService) DeepCopyInto(out *ClusterService) {
	*out = *in
	if in.Frontends != nil {
		in, out := &in.Frontends, &out.Frontends
		*out = make(map[string]PortConfiguration, len(*in))
		for key, val := range *in {
			var outVal map[string]*loadbalancer.L4Addr
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(PortConfiguration, len(*in))
				for key, val := range *in {
					var outVal *loadbalancer.L4Addr
					if val == nil {
						(*out)[key] = nil
					} else {
						in, out := &val, &outVal
						*out = new(loadbalancer.L4Addr)
						**out = **in
					}
					(*out)[key] = outVal
				}
			}
			(*out)[key] = outVal
		}
	}
	if in.Backends != nil {
		in, out := &in.Backends, &out.Backends
		*out = make(map[string]PortConfiguration, len(*in))
		for key, val := range *in {
			var outVal map[string]*loadbalancer.L4Addr
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(PortConfiguration, len(*in))
				for key, val := range *in {
					var outVal *loadbalancer.L4Addr
					if val == nil {
						(*out)[key] = nil
					} else {
						in, out := &val, &outVal
						*out = new(loadbalancer.L4Addr)
						**out = **in
					}
					(*out)[key] = outVal
				}
			}
			(*out)[key] = outVal
		}
	}
	if in.Hostnames != nil {
		in, out := &in.Hostnames, &out.Hostnames
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterService.
func (in *ClusterService) DeepCopy() *ClusterService {
	if in == nil {
		return nil
	}
	out := new(ClusterService)
	in.DeepCopyInto(out)
	return out
}