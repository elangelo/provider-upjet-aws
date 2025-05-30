//go:build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Application) DeepCopyInto(out *Application) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Application.
func (in *Application) DeepCopy() *Application {
	if in == nil {
		return nil
	}
	out := new(Application)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Application) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationInitParameters) DeepCopyInto(out *ApplicationInitParameters) {
	*out = *in
	if in.Architecture != nil {
		in, out := &in.Architecture, &out.Architecture
		*out = new(string)
		**out = **in
	}
	if in.AutoStartConfiguration != nil {
		in, out := &in.AutoStartConfiguration, &out.AutoStartConfiguration
		*out = make([]AutoStartConfigurationInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AutoStopConfiguration != nil {
		in, out := &in.AutoStopConfiguration, &out.AutoStopConfiguration
		*out = make([]AutoStopConfigurationInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ImageConfiguration != nil {
		in, out := &in.ImageConfiguration, &out.ImageConfiguration
		*out = make([]ImageConfigurationInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.InitialCapacity != nil {
		in, out := &in.InitialCapacity, &out.InitialCapacity
		*out = make([]InitialCapacityInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.InteractiveConfiguration != nil {
		in, out := &in.InteractiveConfiguration, &out.InteractiveConfiguration
		*out = make([]InteractiveConfigurationInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MaximumCapacity != nil {
		in, out := &in.MaximumCapacity, &out.MaximumCapacity
		*out = make([]MaximumCapacityInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NetworkConfiguration != nil {
		in, out := &in.NetworkConfiguration, &out.NetworkConfiguration
		*out = make([]NetworkConfigurationInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ReleaseLabel != nil {
		in, out := &in.ReleaseLabel, &out.ReleaseLabel
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationInitParameters.
func (in *ApplicationInitParameters) DeepCopy() *ApplicationInitParameters {
	if in == nil {
		return nil
	}
	out := new(ApplicationInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationList) DeepCopyInto(out *ApplicationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Application, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationList.
func (in *ApplicationList) DeepCopy() *ApplicationList {
	if in == nil {
		return nil
	}
	out := new(ApplicationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApplicationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationObservation) DeepCopyInto(out *ApplicationObservation) {
	*out = *in
	if in.Architecture != nil {
		in, out := &in.Architecture, &out.Architecture
		*out = new(string)
		**out = **in
	}
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.AutoStartConfiguration != nil {
		in, out := &in.AutoStartConfiguration, &out.AutoStartConfiguration
		*out = make([]AutoStartConfigurationObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AutoStopConfiguration != nil {
		in, out := &in.AutoStopConfiguration, &out.AutoStopConfiguration
		*out = make([]AutoStopConfigurationObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.ImageConfiguration != nil {
		in, out := &in.ImageConfiguration, &out.ImageConfiguration
		*out = make([]ImageConfigurationObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.InitialCapacity != nil {
		in, out := &in.InitialCapacity, &out.InitialCapacity
		*out = make([]InitialCapacityObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.InteractiveConfiguration != nil {
		in, out := &in.InteractiveConfiguration, &out.InteractiveConfiguration
		*out = make([]InteractiveConfigurationObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MaximumCapacity != nil {
		in, out := &in.MaximumCapacity, &out.MaximumCapacity
		*out = make([]MaximumCapacityObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NetworkConfiguration != nil {
		in, out := &in.NetworkConfiguration, &out.NetworkConfiguration
		*out = make([]NetworkConfigurationObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ReleaseLabel != nil {
		in, out := &in.ReleaseLabel, &out.ReleaseLabel
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationObservation.
func (in *ApplicationObservation) DeepCopy() *ApplicationObservation {
	if in == nil {
		return nil
	}
	out := new(ApplicationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationParameters) DeepCopyInto(out *ApplicationParameters) {
	*out = *in
	if in.Architecture != nil {
		in, out := &in.Architecture, &out.Architecture
		*out = new(string)
		**out = **in
	}
	if in.AutoStartConfiguration != nil {
		in, out := &in.AutoStartConfiguration, &out.AutoStartConfiguration
		*out = make([]AutoStartConfigurationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AutoStopConfiguration != nil {
		in, out := &in.AutoStopConfiguration, &out.AutoStopConfiguration
		*out = make([]AutoStopConfigurationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ImageConfiguration != nil {
		in, out := &in.ImageConfiguration, &out.ImageConfiguration
		*out = make([]ImageConfigurationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.InitialCapacity != nil {
		in, out := &in.InitialCapacity, &out.InitialCapacity
		*out = make([]InitialCapacityParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.InteractiveConfiguration != nil {
		in, out := &in.InteractiveConfiguration, &out.InteractiveConfiguration
		*out = make([]InteractiveConfigurationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MaximumCapacity != nil {
		in, out := &in.MaximumCapacity, &out.MaximumCapacity
		*out = make([]MaximumCapacityParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NetworkConfiguration != nil {
		in, out := &in.NetworkConfiguration, &out.NetworkConfiguration
		*out = make([]NetworkConfigurationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.ReleaseLabel != nil {
		in, out := &in.ReleaseLabel, &out.ReleaseLabel
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationParameters.
func (in *ApplicationParameters) DeepCopy() *ApplicationParameters {
	if in == nil {
		return nil
	}
	out := new(ApplicationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSpec) DeepCopyInto(out *ApplicationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSpec.
func (in *ApplicationSpec) DeepCopy() *ApplicationSpec {
	if in == nil {
		return nil
	}
	out := new(ApplicationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationStatus) DeepCopyInto(out *ApplicationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationStatus.
func (in *ApplicationStatus) DeepCopy() *ApplicationStatus {
	if in == nil {
		return nil
	}
	out := new(ApplicationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoStartConfigurationInitParameters) DeepCopyInto(out *AutoStartConfigurationInitParameters) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoStartConfigurationInitParameters.
func (in *AutoStartConfigurationInitParameters) DeepCopy() *AutoStartConfigurationInitParameters {
	if in == nil {
		return nil
	}
	out := new(AutoStartConfigurationInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoStartConfigurationObservation) DeepCopyInto(out *AutoStartConfigurationObservation) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoStartConfigurationObservation.
func (in *AutoStartConfigurationObservation) DeepCopy() *AutoStartConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(AutoStartConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoStartConfigurationParameters) DeepCopyInto(out *AutoStartConfigurationParameters) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoStartConfigurationParameters.
func (in *AutoStartConfigurationParameters) DeepCopy() *AutoStartConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(AutoStartConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoStopConfigurationInitParameters) DeepCopyInto(out *AutoStopConfigurationInitParameters) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.IdleTimeoutMinutes != nil {
		in, out := &in.IdleTimeoutMinutes, &out.IdleTimeoutMinutes
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoStopConfigurationInitParameters.
func (in *AutoStopConfigurationInitParameters) DeepCopy() *AutoStopConfigurationInitParameters {
	if in == nil {
		return nil
	}
	out := new(AutoStopConfigurationInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoStopConfigurationObservation) DeepCopyInto(out *AutoStopConfigurationObservation) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.IdleTimeoutMinutes != nil {
		in, out := &in.IdleTimeoutMinutes, &out.IdleTimeoutMinutes
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoStopConfigurationObservation.
func (in *AutoStopConfigurationObservation) DeepCopy() *AutoStopConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(AutoStopConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoStopConfigurationParameters) DeepCopyInto(out *AutoStopConfigurationParameters) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.IdleTimeoutMinutes != nil {
		in, out := &in.IdleTimeoutMinutes, &out.IdleTimeoutMinutes
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoStopConfigurationParameters.
func (in *AutoStopConfigurationParameters) DeepCopy() *AutoStopConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(AutoStopConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageConfigurationInitParameters) DeepCopyInto(out *ImageConfigurationInitParameters) {
	*out = *in
	if in.ImageURI != nil {
		in, out := &in.ImageURI, &out.ImageURI
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageConfigurationInitParameters.
func (in *ImageConfigurationInitParameters) DeepCopy() *ImageConfigurationInitParameters {
	if in == nil {
		return nil
	}
	out := new(ImageConfigurationInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageConfigurationObservation) DeepCopyInto(out *ImageConfigurationObservation) {
	*out = *in
	if in.ImageURI != nil {
		in, out := &in.ImageURI, &out.ImageURI
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageConfigurationObservation.
func (in *ImageConfigurationObservation) DeepCopy() *ImageConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(ImageConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageConfigurationParameters) DeepCopyInto(out *ImageConfigurationParameters) {
	*out = *in
	if in.ImageURI != nil {
		in, out := &in.ImageURI, &out.ImageURI
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageConfigurationParameters.
func (in *ImageConfigurationParameters) DeepCopy() *ImageConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(ImageConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InitialCapacityConfigInitParameters) DeepCopyInto(out *InitialCapacityConfigInitParameters) {
	*out = *in
	if in.WorkerConfiguration != nil {
		in, out := &in.WorkerConfiguration, &out.WorkerConfiguration
		*out = make([]WorkerConfigurationInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.WorkerCount != nil {
		in, out := &in.WorkerCount, &out.WorkerCount
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InitialCapacityConfigInitParameters.
func (in *InitialCapacityConfigInitParameters) DeepCopy() *InitialCapacityConfigInitParameters {
	if in == nil {
		return nil
	}
	out := new(InitialCapacityConfigInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InitialCapacityConfigObservation) DeepCopyInto(out *InitialCapacityConfigObservation) {
	*out = *in
	if in.WorkerConfiguration != nil {
		in, out := &in.WorkerConfiguration, &out.WorkerConfiguration
		*out = make([]WorkerConfigurationObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.WorkerCount != nil {
		in, out := &in.WorkerCount, &out.WorkerCount
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InitialCapacityConfigObservation.
func (in *InitialCapacityConfigObservation) DeepCopy() *InitialCapacityConfigObservation {
	if in == nil {
		return nil
	}
	out := new(InitialCapacityConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InitialCapacityConfigParameters) DeepCopyInto(out *InitialCapacityConfigParameters) {
	*out = *in
	if in.WorkerConfiguration != nil {
		in, out := &in.WorkerConfiguration, &out.WorkerConfiguration
		*out = make([]WorkerConfigurationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.WorkerCount != nil {
		in, out := &in.WorkerCount, &out.WorkerCount
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InitialCapacityConfigParameters.
func (in *InitialCapacityConfigParameters) DeepCopy() *InitialCapacityConfigParameters {
	if in == nil {
		return nil
	}
	out := new(InitialCapacityConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InitialCapacityInitParameters) DeepCopyInto(out *InitialCapacityInitParameters) {
	*out = *in
	if in.InitialCapacityConfig != nil {
		in, out := &in.InitialCapacityConfig, &out.InitialCapacityConfig
		*out = make([]InitialCapacityConfigInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.InitialCapacityType != nil {
		in, out := &in.InitialCapacityType, &out.InitialCapacityType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InitialCapacityInitParameters.
func (in *InitialCapacityInitParameters) DeepCopy() *InitialCapacityInitParameters {
	if in == nil {
		return nil
	}
	out := new(InitialCapacityInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InitialCapacityObservation) DeepCopyInto(out *InitialCapacityObservation) {
	*out = *in
	if in.InitialCapacityConfig != nil {
		in, out := &in.InitialCapacityConfig, &out.InitialCapacityConfig
		*out = make([]InitialCapacityConfigObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.InitialCapacityType != nil {
		in, out := &in.InitialCapacityType, &out.InitialCapacityType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InitialCapacityObservation.
func (in *InitialCapacityObservation) DeepCopy() *InitialCapacityObservation {
	if in == nil {
		return nil
	}
	out := new(InitialCapacityObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InitialCapacityParameters) DeepCopyInto(out *InitialCapacityParameters) {
	*out = *in
	if in.InitialCapacityConfig != nil {
		in, out := &in.InitialCapacityConfig, &out.InitialCapacityConfig
		*out = make([]InitialCapacityConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.InitialCapacityType != nil {
		in, out := &in.InitialCapacityType, &out.InitialCapacityType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InitialCapacityParameters.
func (in *InitialCapacityParameters) DeepCopy() *InitialCapacityParameters {
	if in == nil {
		return nil
	}
	out := new(InitialCapacityParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InteractiveConfigurationInitParameters) DeepCopyInto(out *InteractiveConfigurationInitParameters) {
	*out = *in
	if in.LivyEndpointEnabled != nil {
		in, out := &in.LivyEndpointEnabled, &out.LivyEndpointEnabled
		*out = new(bool)
		**out = **in
	}
	if in.StudioEnabled != nil {
		in, out := &in.StudioEnabled, &out.StudioEnabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InteractiveConfigurationInitParameters.
func (in *InteractiveConfigurationInitParameters) DeepCopy() *InteractiveConfigurationInitParameters {
	if in == nil {
		return nil
	}
	out := new(InteractiveConfigurationInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InteractiveConfigurationObservation) DeepCopyInto(out *InteractiveConfigurationObservation) {
	*out = *in
	if in.LivyEndpointEnabled != nil {
		in, out := &in.LivyEndpointEnabled, &out.LivyEndpointEnabled
		*out = new(bool)
		**out = **in
	}
	if in.StudioEnabled != nil {
		in, out := &in.StudioEnabled, &out.StudioEnabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InteractiveConfigurationObservation.
func (in *InteractiveConfigurationObservation) DeepCopy() *InteractiveConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(InteractiveConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InteractiveConfigurationParameters) DeepCopyInto(out *InteractiveConfigurationParameters) {
	*out = *in
	if in.LivyEndpointEnabled != nil {
		in, out := &in.LivyEndpointEnabled, &out.LivyEndpointEnabled
		*out = new(bool)
		**out = **in
	}
	if in.StudioEnabled != nil {
		in, out := &in.StudioEnabled, &out.StudioEnabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InteractiveConfigurationParameters.
func (in *InteractiveConfigurationParameters) DeepCopy() *InteractiveConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(InteractiveConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaximumCapacityInitParameters) DeepCopyInto(out *MaximumCapacityInitParameters) {
	*out = *in
	if in.CPU != nil {
		in, out := &in.CPU, &out.CPU
		*out = new(string)
		**out = **in
	}
	if in.Disk != nil {
		in, out := &in.Disk, &out.Disk
		*out = new(string)
		**out = **in
	}
	if in.Memory != nil {
		in, out := &in.Memory, &out.Memory
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaximumCapacityInitParameters.
func (in *MaximumCapacityInitParameters) DeepCopy() *MaximumCapacityInitParameters {
	if in == nil {
		return nil
	}
	out := new(MaximumCapacityInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaximumCapacityObservation) DeepCopyInto(out *MaximumCapacityObservation) {
	*out = *in
	if in.CPU != nil {
		in, out := &in.CPU, &out.CPU
		*out = new(string)
		**out = **in
	}
	if in.Disk != nil {
		in, out := &in.Disk, &out.Disk
		*out = new(string)
		**out = **in
	}
	if in.Memory != nil {
		in, out := &in.Memory, &out.Memory
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaximumCapacityObservation.
func (in *MaximumCapacityObservation) DeepCopy() *MaximumCapacityObservation {
	if in == nil {
		return nil
	}
	out := new(MaximumCapacityObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaximumCapacityParameters) DeepCopyInto(out *MaximumCapacityParameters) {
	*out = *in
	if in.CPU != nil {
		in, out := &in.CPU, &out.CPU
		*out = new(string)
		**out = **in
	}
	if in.Disk != nil {
		in, out := &in.Disk, &out.Disk
		*out = new(string)
		**out = **in
	}
	if in.Memory != nil {
		in, out := &in.Memory, &out.Memory
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaximumCapacityParameters.
func (in *MaximumCapacityParameters) DeepCopy() *MaximumCapacityParameters {
	if in == nil {
		return nil
	}
	out := new(MaximumCapacityParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkConfigurationInitParameters) DeepCopyInto(out *NetworkConfigurationInitParameters) {
	*out = *in
	if in.SecurityGroupIds != nil {
		in, out := &in.SecurityGroupIds, &out.SecurityGroupIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SubnetIds != nil {
		in, out := &in.SubnetIds, &out.SubnetIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkConfigurationInitParameters.
func (in *NetworkConfigurationInitParameters) DeepCopy() *NetworkConfigurationInitParameters {
	if in == nil {
		return nil
	}
	out := new(NetworkConfigurationInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkConfigurationObservation) DeepCopyInto(out *NetworkConfigurationObservation) {
	*out = *in
	if in.SecurityGroupIds != nil {
		in, out := &in.SecurityGroupIds, &out.SecurityGroupIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SubnetIds != nil {
		in, out := &in.SubnetIds, &out.SubnetIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkConfigurationObservation.
func (in *NetworkConfigurationObservation) DeepCopy() *NetworkConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(NetworkConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkConfigurationParameters) DeepCopyInto(out *NetworkConfigurationParameters) {
	*out = *in
	if in.SecurityGroupIds != nil {
		in, out := &in.SecurityGroupIds, &out.SecurityGroupIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SubnetIds != nil {
		in, out := &in.SubnetIds, &out.SubnetIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkConfigurationParameters.
func (in *NetworkConfigurationParameters) DeepCopy() *NetworkConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(NetworkConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerConfigurationInitParameters) DeepCopyInto(out *WorkerConfigurationInitParameters) {
	*out = *in
	if in.CPU != nil {
		in, out := &in.CPU, &out.CPU
		*out = new(string)
		**out = **in
	}
	if in.Disk != nil {
		in, out := &in.Disk, &out.Disk
		*out = new(string)
		**out = **in
	}
	if in.Memory != nil {
		in, out := &in.Memory, &out.Memory
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerConfigurationInitParameters.
func (in *WorkerConfigurationInitParameters) DeepCopy() *WorkerConfigurationInitParameters {
	if in == nil {
		return nil
	}
	out := new(WorkerConfigurationInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerConfigurationObservation) DeepCopyInto(out *WorkerConfigurationObservation) {
	*out = *in
	if in.CPU != nil {
		in, out := &in.CPU, &out.CPU
		*out = new(string)
		**out = **in
	}
	if in.Disk != nil {
		in, out := &in.Disk, &out.Disk
		*out = new(string)
		**out = **in
	}
	if in.Memory != nil {
		in, out := &in.Memory, &out.Memory
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerConfigurationObservation.
func (in *WorkerConfigurationObservation) DeepCopy() *WorkerConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(WorkerConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerConfigurationParameters) DeepCopyInto(out *WorkerConfigurationParameters) {
	*out = *in
	if in.CPU != nil {
		in, out := &in.CPU, &out.CPU
		*out = new(string)
		**out = **in
	}
	if in.Disk != nil {
		in, out := &in.Disk, &out.Disk
		*out = new(string)
		**out = **in
	}
	if in.Memory != nil {
		in, out := &in.Memory, &out.Memory
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerConfigurationParameters.
func (in *WorkerConfigurationParameters) DeepCopy() *WorkerConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(WorkerConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}
