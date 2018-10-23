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

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotPolicy) DeepCopyInto(out *SnapshotPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotPolicy.
func (in *SnapshotPolicy) DeepCopy() *SnapshotPolicy {
	if in == nil {
		return nil
	}
	out := new(SnapshotPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SnapshotPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotPolicyList) DeepCopyInto(out *SnapshotPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SnapshotPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotPolicyList.
func (in *SnapshotPolicyList) DeepCopy() *SnapshotPolicyList {
	if in == nil {
		return nil
	}
	out := new(SnapshotPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SnapshotPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotPolicySpec) DeepCopyInto(out *SnapshotPolicySpec) {
	*out = *in
	if in.PVCNames != nil {
		in, out := &in.PVCNames, &out.PVCNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Period != nil {
		in, out := &in.Period, &out.Period
		*out = new(int32)
		**out = **in
	}
	if in.Retention != nil {
		in, out := &in.Retention, &out.Retention
		*out = new(int32)
		**out = **in
	}
	out.Strategy = in.Strategy
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotPolicySpec.
func (in *SnapshotPolicySpec) DeepCopy() *SnapshotPolicySpec {
	if in == nil {
		return nil
	}
	out := new(SnapshotPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotPolicyStatus) DeepCopyInto(out *SnapshotPolicyStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotPolicyStatus.
func (in *SnapshotPolicyStatus) DeepCopy() *SnapshotPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(SnapshotPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Strategy) DeepCopyInto(out *Strategy) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Strategy.
func (in *Strategy) DeepCopy() *Strategy {
	if in == nil {
		return nil
	}
	out := new(Strategy)
	in.DeepCopyInto(out)
	return out
}
