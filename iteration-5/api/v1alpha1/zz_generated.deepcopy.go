// +build !ignore_autogenerated

/*


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
func (in *WateringAlarm) DeepCopyInto(out *WateringAlarm) {
*out = *in
out.TypeMeta = in.TypeMeta
in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
out.Spec = in.Spec
in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WateringAlarm.
func (in *WateringAlarm) DeepCopy() *WateringAlarm {
	if in == nil { return nil }
	out := new(WateringAlarm)
	in.DeepCopyInto(out)
	return out
}


// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WateringAlarm) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WateringAlarmList) DeepCopyInto(out *WateringAlarmList) {
*out = *in
out.TypeMeta = in.TypeMeta
in.ListMeta.DeepCopyInto(&out.ListMeta)
if in.Items != nil {
in, out := &in.Items, &out.Items
*out = make([]WateringAlarm, len(*in))
for i := range *in {
(*in)[i].DeepCopyInto(&(*out)[i])
}
}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WateringAlarmList.
func (in *WateringAlarmList) DeepCopy() *WateringAlarmList {
	if in == nil { return nil }
	out := new(WateringAlarmList)
	in.DeepCopyInto(out)
	return out
}


// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WateringAlarmList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WateringAlarmSpec) DeepCopyInto(out *WateringAlarmSpec) {
*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WateringAlarmSpec.
func (in *WateringAlarmSpec) DeepCopy() *WateringAlarmSpec {
	if in == nil { return nil }
	out := new(WateringAlarmSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WateringAlarmStatus) DeepCopyInto(out *WateringAlarmStatus) {
*out = *in
if in.LastWateringDate != nil {
in, out := &in.LastWateringDate, &out.LastWateringDate
*out = new(invalid type)
**out = **in
}
if in.NextWateringDate != nil {
in, out := &in.NextWateringDate, &out.NextWateringDate
*out = new(invalid type)
**out = **in
}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WateringAlarmStatus.
func (in *WateringAlarmStatus) DeepCopy() *WateringAlarmStatus {
	if in == nil { return nil }
	out := new(WateringAlarmStatus)
	in.DeepCopyInto(out)
	return out
}

