//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

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
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MinecraftCluster) DeepCopyInto(out *MinecraftCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MinecraftCluster.
func (in *MinecraftCluster) DeepCopy() *MinecraftCluster {
	if in == nil {
		return nil
	}
	out := new(MinecraftCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MinecraftCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MinecraftClusterList) DeepCopyInto(out *MinecraftClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MinecraftCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MinecraftClusterList.
func (in *MinecraftClusterList) DeepCopy() *MinecraftClusterList {
	if in == nil {
		return nil
	}
	out := new(MinecraftClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MinecraftClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MinecraftClusterRef) DeepCopyInto(out *MinecraftClusterRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MinecraftClusterRef.
func (in *MinecraftClusterRef) DeepCopy() *MinecraftClusterRef {
	if in == nil {
		return nil
	}
	out := new(MinecraftClusterRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MinecraftClusterSpec) DeepCopyInto(out *MinecraftClusterSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MinecraftClusterSpec.
func (in *MinecraftClusterSpec) DeepCopy() *MinecraftClusterSpec {
	if in == nil {
		return nil
	}
	out := new(MinecraftClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MinecraftClusterStatus) DeepCopyInto(out *MinecraftClusterStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ServerPool != nil {
		in, out := &in.ServerPool, &out.ServerPool
		*out = make([]MinecraftClusterStatusServerEntry, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MinecraftClusterStatus.
func (in *MinecraftClusterStatus) DeepCopy() *MinecraftClusterStatus {
	if in == nil {
		return nil
	}
	out := new(MinecraftClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MinecraftClusterStatusServerEntry) DeepCopyInto(out *MinecraftClusterStatusServerEntry) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MinecraftClusterStatusServerEntry.
func (in *MinecraftClusterStatusServerEntry) DeepCopy() *MinecraftClusterStatusServerEntry {
	if in == nil {
		return nil
	}
	out := new(MinecraftClusterStatusServerEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MinecraftServer) DeepCopyInto(out *MinecraftServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MinecraftServer.
func (in *MinecraftServer) DeepCopy() *MinecraftServer {
	if in == nil {
		return nil
	}
	out := new(MinecraftServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MinecraftServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MinecraftServerList) DeepCopyInto(out *MinecraftServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MinecraftServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MinecraftServerList.
func (in *MinecraftServerList) DeepCopy() *MinecraftServerList {
	if in == nil {
		return nil
	}
	out := new(MinecraftServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MinecraftServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MinecraftServerPodOverridesSpec) DeepCopyInto(out *MinecraftServerPodOverridesSpec) {
	*out = *in
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]corev1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LivenessProbe != nil {
		in, out := &in.LivenessProbe, &out.LivenessProbe
		*out = new(MinecraftServerPodProbeSpec)
		**out = **in
	}
	if in.ReadinessProbe != nil {
		in, out := &in.ReadinessProbe, &out.ReadinessProbe
		*out = new(MinecraftServerPodProbeSpec)
		**out = **in
	}
	if in.TerminationGracePeriodSeconds != nil {
		in, out := &in.TerminationGracePeriodSeconds, &out.TerminationGracePeriodSeconds
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MinecraftServerPodOverridesSpec.
func (in *MinecraftServerPodOverridesSpec) DeepCopy() *MinecraftServerPodOverridesSpec {
	if in == nil {
		return nil
	}
	out := new(MinecraftServerPodOverridesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MinecraftServerPodProbeSpec) DeepCopyInto(out *MinecraftServerPodProbeSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MinecraftServerPodProbeSpec.
func (in *MinecraftServerPodProbeSpec) DeepCopy() *MinecraftServerPodProbeSpec {
	if in == nil {
		return nil
	}
	out := new(MinecraftServerPodProbeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MinecraftServerRconSpec) DeepCopyInto(out *MinecraftServerRconSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MinecraftServerRconSpec.
func (in *MinecraftServerRconSpec) DeepCopy() *MinecraftServerRconSpec {
	if in == nil {
		return nil
	}
	out := new(MinecraftServerRconSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MinecraftServerServiceSpec) DeepCopyInto(out *MinecraftServerServiceSpec) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MinecraftServerServiceSpec.
func (in *MinecraftServerServiceSpec) DeepCopy() *MinecraftServerServiceSpec {
	if in == nil {
		return nil
	}
	out := new(MinecraftServerServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MinecraftServerSpec) DeepCopyInto(out *MinecraftServerSpec) {
	*out = *in
	if in.ClusterRef != nil {
		in, out := &in.ClusterRef, &out.ClusterRef
		*out = new(MinecraftClusterRef)
		**out = **in
	}
	out.Version = in.Version
	if in.MaxPlayers != nil {
		in, out := &in.MaxPlayers, &out.MaxPlayers
		*out = new(int64)
		**out = **in
	}
	if in.Operators != nil {
		in, out := &in.Operators, &out.Operators
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.WhitelistedPlayers != nil {
		in, out := &in.WhitelistedPlayers, &out.WhitelistedPlayers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.World != nil {
		in, out := &in.World, &out.World
		*out = new(MinecraftServerWorldSpec)
		**out = **in
	}
	if in.Rcon != nil {
		in, out := &in.Rcon, &out.Rcon
		*out = new(MinecraftServerRconSpec)
		**out = **in
	}
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(MinecraftServerServiceSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.PodOverrides != nil {
		in, out := &in.PodOverrides, &out.PodOverrides
		*out = new(MinecraftServerPodOverridesSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(corev1.Affinity)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MinecraftServerSpec.
func (in *MinecraftServerSpec) DeepCopy() *MinecraftServerSpec {
	if in == nil {
		return nil
	}
	out := new(MinecraftServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MinecraftServerStatus) DeepCopyInto(out *MinecraftServerStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MinecraftServerStatus.
func (in *MinecraftServerStatus) DeepCopy() *MinecraftServerStatus {
	if in == nil {
		return nil
	}
	out := new(MinecraftServerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MinecraftServerVersionSpec) DeepCopyInto(out *MinecraftServerVersionSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MinecraftServerVersionSpec.
func (in *MinecraftServerVersionSpec) DeepCopy() *MinecraftServerVersionSpec {
	if in == nil {
		return nil
	}
	out := new(MinecraftServerVersionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MinecraftServerWorldSpec) DeepCopyInto(out *MinecraftServerWorldSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MinecraftServerWorldSpec.
func (in *MinecraftServerWorldSpec) DeepCopy() *MinecraftServerWorldSpec {
	if in == nil {
		return nil
	}
	out := new(MinecraftServerWorldSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxyDeployment) DeepCopyInto(out *ProxyDeployment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxyDeployment.
func (in *ProxyDeployment) DeepCopy() *ProxyDeployment {
	if in == nil {
		return nil
	}
	out := new(ProxyDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProxyDeployment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxyDeploymentList) DeepCopyInto(out *ProxyDeploymentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ProxyDeployment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxyDeploymentList.
func (in *ProxyDeploymentList) DeepCopy() *ProxyDeploymentList {
	if in == nil {
		return nil
	}
	out := new(ProxyDeploymentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProxyDeploymentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxyDeploymentPodOverridesSpec) DeepCopyInto(out *ProxyDeploymentPodOverridesSpec) {
	*out = *in
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]corev1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LivenessProbe != nil {
		in, out := &in.LivenessProbe, &out.LivenessProbe
		*out = new(ProxyDeploymentPodProbeSpec)
		**out = **in
	}
	if in.ReadinessProbe != nil {
		in, out := &in.ReadinessProbe, &out.ReadinessProbe
		*out = new(ProxyDeploymentPodProbeSpec)
		**out = **in
	}
	if in.TerminationGracePeriodSeconds != nil {
		in, out := &in.TerminationGracePeriodSeconds, &out.TerminationGracePeriodSeconds
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxyDeploymentPodOverridesSpec.
func (in *ProxyDeploymentPodOverridesSpec) DeepCopy() *ProxyDeploymentPodOverridesSpec {
	if in == nil {
		return nil
	}
	out := new(ProxyDeploymentPodOverridesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxyDeploymentPodProbeSpec) DeepCopyInto(out *ProxyDeploymentPodProbeSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxyDeploymentPodProbeSpec.
func (in *ProxyDeploymentPodProbeSpec) DeepCopy() *ProxyDeploymentPodProbeSpec {
	if in == nil {
		return nil
	}
	out := new(ProxyDeploymentPodProbeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxyDeploymentServiceSpec) DeepCopyInto(out *ProxyDeploymentServiceSpec) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxyDeploymentServiceSpec.
func (in *ProxyDeploymentServiceSpec) DeepCopy() *ProxyDeploymentServiceSpec {
	if in == nil {
		return nil
	}
	out := new(ProxyDeploymentServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxyDeploymentSpec) DeepCopyInto(out *ProxyDeploymentSpec) {
	*out = *in
	if in.ClusterRef != nil {
		in, out := &in.ClusterRef, &out.ClusterRef
		*out = new(MinecraftClusterRef)
		**out = **in
	}
	out.Version = in.Version
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(ProxyDeploymentServiceSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.PodOverrides != nil {
		in, out := &in.PodOverrides, &out.PodOverrides
		*out = new(MinecraftServerPodOverridesSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(corev1.Affinity)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxyDeploymentSpec.
func (in *ProxyDeploymentSpec) DeepCopy() *ProxyDeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(ProxyDeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxyDeploymentStatus) DeepCopyInto(out *ProxyDeploymentStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxyDeploymentStatus.
func (in *ProxyDeploymentStatus) DeepCopy() *ProxyDeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(ProxyDeploymentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxyDeploymentVersionSpec) DeepCopyInto(out *ProxyDeploymentVersionSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxyDeploymentVersionSpec.
func (in *ProxyDeploymentVersionSpec) DeepCopy() *ProxyDeploymentVersionSpec {
	if in == nil {
		return nil
	}
	out := new(ProxyDeploymentVersionSpec)
	in.DeepCopyInto(out)
	return out
}
