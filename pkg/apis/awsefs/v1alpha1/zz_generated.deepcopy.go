// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedVolume) DeepCopyInto(out *SharedVolume) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedVolume.
func (in *SharedVolume) DeepCopy() *SharedVolume {
	if in == nil {
		return nil
	}
	out := new(SharedVolume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SharedVolume) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedVolumeList) DeepCopyInto(out *SharedVolumeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SharedVolume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedVolumeList.
func (in *SharedVolumeList) DeepCopy() *SharedVolumeList {
	if in == nil {
		return nil
	}
	out := new(SharedVolumeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SharedVolumeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedVolumeSpec) DeepCopyInto(out *SharedVolumeSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedVolumeSpec.
func (in *SharedVolumeSpec) DeepCopy() *SharedVolumeSpec {
	if in == nil {
		return nil
	}
	out := new(SharedVolumeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedVolumeStatus) DeepCopyInto(out *SharedVolumeStatus) {
	*out = *in
	in.ClaimRef.DeepCopyInto(&out.ClaimRef)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedVolumeStatus.
func (in *SharedVolumeStatus) DeepCopy() *SharedVolumeStatus {
	if in == nil {
		return nil
	}
	out := new(SharedVolumeStatus)
	in.DeepCopyInto(out)
	return out
}