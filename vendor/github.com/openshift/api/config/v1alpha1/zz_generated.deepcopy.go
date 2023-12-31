//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Backup) DeepCopyInto(out *Backup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Backup.
func (in *Backup) DeepCopy() *Backup {
	if in == nil {
		return nil
	}
	out := new(Backup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Backup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupList) DeepCopyInto(out *BackupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Backup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupList.
func (in *BackupList) DeepCopy() *BackupList {
	if in == nil {
		return nil
	}
	out := new(BackupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupSpec) DeepCopyInto(out *BackupSpec) {
	*out = *in
	in.EtcdBackupSpec.DeepCopyInto(&out.EtcdBackupSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupSpec.
func (in *BackupSpec) DeepCopy() *BackupSpec {
	if in == nil {
		return nil
	}
	out := new(BackupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupStatus) DeepCopyInto(out *BackupStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupStatus.
func (in *BackupStatus) DeepCopy() *BackupStatus {
	if in == nil {
		return nil
	}
	out := new(BackupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EtcdBackupSpec) DeepCopyInto(out *EtcdBackupSpec) {
	*out = *in
	in.RetentionPolicy.DeepCopyInto(&out.RetentionPolicy)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EtcdBackupSpec.
func (in *EtcdBackupSpec) DeepCopy() *EtcdBackupSpec {
	if in == nil {
		return nil
	}
	out := new(EtcdBackupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatherConfig) DeepCopyInto(out *GatherConfig) {
	*out = *in
	if in.DisabledGatherers != nil {
		in, out := &in.DisabledGatherers, &out.DisabledGatherers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatherConfig.
func (in *GatherConfig) DeepCopy() *GatherConfig {
	if in == nil {
		return nil
	}
	out := new(GatherConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InsightsDataGather) DeepCopyInto(out *InsightsDataGather) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InsightsDataGather.
func (in *InsightsDataGather) DeepCopy() *InsightsDataGather {
	if in == nil {
		return nil
	}
	out := new(InsightsDataGather)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InsightsDataGather) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InsightsDataGatherList) DeepCopyInto(out *InsightsDataGatherList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]InsightsDataGather, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InsightsDataGatherList.
func (in *InsightsDataGatherList) DeepCopy() *InsightsDataGatherList {
	if in == nil {
		return nil
	}
	out := new(InsightsDataGatherList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InsightsDataGatherList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InsightsDataGatherSpec) DeepCopyInto(out *InsightsDataGatherSpec) {
	*out = *in
	in.GatherConfig.DeepCopyInto(&out.GatherConfig)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InsightsDataGatherSpec.
func (in *InsightsDataGatherSpec) DeepCopy() *InsightsDataGatherSpec {
	if in == nil {
		return nil
	}
	out := new(InsightsDataGatherSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InsightsDataGatherStatus) DeepCopyInto(out *InsightsDataGatherStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InsightsDataGatherStatus.
func (in *InsightsDataGatherStatus) DeepCopy() *InsightsDataGatherStatus {
	if in == nil {
		return nil
	}
	out := new(InsightsDataGatherStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetentionNumberConfig) DeepCopyInto(out *RetentionNumberConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetentionNumberConfig.
func (in *RetentionNumberConfig) DeepCopy() *RetentionNumberConfig {
	if in == nil {
		return nil
	}
	out := new(RetentionNumberConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetentionPolicy) DeepCopyInto(out *RetentionPolicy) {
	*out = *in
	if in.RetentionNumber != nil {
		in, out := &in.RetentionNumber, &out.RetentionNumber
		*out = new(RetentionNumberConfig)
		**out = **in
	}
	if in.RetentionSize != nil {
		in, out := &in.RetentionSize, &out.RetentionSize
		*out = new(RetentionSizeConfig)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetentionPolicy.
func (in *RetentionPolicy) DeepCopy() *RetentionPolicy {
	if in == nil {
		return nil
	}
	out := new(RetentionPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetentionSizeConfig) DeepCopyInto(out *RetentionSizeConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetentionSizeConfig.
func (in *RetentionSizeConfig) DeepCopy() *RetentionSizeConfig {
	if in == nil {
		return nil
	}
	out := new(RetentionSizeConfig)
	in.DeepCopyInto(out)
	return out
}
