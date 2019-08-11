// +build !ignore_autogenerated

//
// DISCLAIMER
//
// Copyright 2018 ArangoDB GmbH, Cologne, Germany
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Copyright holder is ArangoDB GmbH, Cologne, Germany
//

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha

import (
	time "time"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Action) DeepCopyInto(out *Action) {
	*out = *in
	in.CreationTime.DeepCopyInto(&out.CreationTime)
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Action.
func (in *Action) DeepCopy() *Action {
	if in == nil {
		return nil
	}
	out := new(Action)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArangoBackup) DeepCopyInto(out *ArangoBackup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArangoBackup.
func (in *ArangoBackup) DeepCopy() *ArangoBackup {
	if in == nil {
		return nil
	}
	out := new(ArangoBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ArangoBackup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArangoBackupDetails) DeepCopyInto(out *ArangoBackupDetails) {
	*out = *in
	out.CreationTimestamp = in.CreationTimestamp
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArangoBackupDetails.
func (in *ArangoBackupDetails) DeepCopy() *ArangoBackupDetails {
	if in == nil {
		return nil
	}
	out := new(ArangoBackupDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArangoBackupList) DeepCopyInto(out *ArangoBackupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ArangoBackup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArangoBackupList.
func (in *ArangoBackupList) DeepCopy() *ArangoBackupList {
	if in == nil {
		return nil
	}
	out := new(ArangoBackupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ArangoBackupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArangoBackupSpec) DeepCopyInto(out *ArangoBackupSpec) {
	*out = *in
	out.Deployment = in.Deployment
	if in.Download != nil {
		in, out := &in.Download, &out.Download
		*out = new(ArangoBackupSpecOperation)
		**out = **in
	}
	if in.Upload != nil {
		in, out := &in.Upload, &out.Upload
		*out = new(ArangoBackupSpecOperation)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArangoBackupSpec.
func (in *ArangoBackupSpec) DeepCopy() *ArangoBackupSpec {
	if in == nil {
		return nil
	}
	out := new(ArangoBackupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArangoBackupSpecDeployment) DeepCopyInto(out *ArangoBackupSpecDeployment) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArangoBackupSpecDeployment.
func (in *ArangoBackupSpecDeployment) DeepCopy() *ArangoBackupSpecDeployment {
	if in == nil {
		return nil
	}
	out := new(ArangoBackupSpecDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArangoBackupSpecOperation) DeepCopyInto(out *ArangoBackupSpecOperation) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArangoBackupSpecOperation.
func (in *ArangoBackupSpecOperation) DeepCopy() *ArangoBackupSpecOperation {
	if in == nil {
		return nil
	}
	out := new(ArangoBackupSpecOperation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArangoBackupState) DeepCopyInto(out *ArangoBackupState) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArangoBackupState.
func (in *ArangoBackupState) DeepCopy() *ArangoBackupState {
	if in == nil {
		return nil
	}
	out := new(ArangoBackupState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArangoBackupStatus) DeepCopyInto(out *ArangoBackupStatus) {
	*out = *in
	out.State = in.State
	if in.Details != nil {
		in, out := &in.Details, &out.Details
		*out = new(ArangoBackupDetails)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArangoBackupStatus.
func (in *ArangoBackupStatus) DeepCopy() *ArangoBackupStatus {
	if in == nil {
		return nil
	}
	out := new(ArangoBackupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArangoDeployment) DeepCopyInto(out *ArangoDeployment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArangoDeployment.
func (in *ArangoDeployment) DeepCopy() *ArangoDeployment {
	if in == nil {
		return nil
	}
	out := new(ArangoDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ArangoDeployment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArangoDeploymentList) DeepCopyInto(out *ArangoDeploymentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ArangoDeployment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArangoDeploymentList.
func (in *ArangoDeploymentList) DeepCopy() *ArangoDeploymentList {
	if in == nil {
		return nil
	}
	out := new(ArangoDeploymentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ArangoDeploymentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthenticationSpec) DeepCopyInto(out *AuthenticationSpec) {
	*out = *in
	if in.JWTSecretName != nil {
		in, out := &in.JWTSecretName, &out.JWTSecretName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthenticationSpec.
func (in *AuthenticationSpec) DeepCopy() *AuthenticationSpec {
	if in == nil {
		return nil
	}
	out := new(AuthenticationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BootstrapSpec) DeepCopyInto(out *BootstrapSpec) {
	*out = *in
	if in.PasswordSecretNames != nil {
		in, out := &in.PasswordSecretNames, &out.PasswordSecretNames
		*out = make(PasswordSecretNameList, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BootstrapSpec.
func (in *BootstrapSpec) DeepCopy() *BootstrapSpec {
	if in == nil {
		return nil
	}
	out := new(BootstrapSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChaosSpec) DeepCopyInto(out *ChaosSpec) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Interval != nil {
		in, out := &in.Interval, &out.Interval
		*out = new(time.Duration)
		**out = **in
	}
	if in.KillPodProbability != nil {
		in, out := &in.KillPodProbability, &out.KillPodProbability
		*out = new(Percent)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChaosSpec.
func (in *ChaosSpec) DeepCopy() *ChaosSpec {
	if in == nil {
		return nil
	}
	out := new(ChaosSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in ConditionList) DeepCopyInto(out *ConditionList) {
	{
		in := &in
		*out = make(ConditionList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConditionList.
func (in ConditionList) DeepCopy() ConditionList {
	if in == nil {
		return nil
	}
	out := new(ConditionList)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentSpec) DeepCopyInto(out *DeploymentSpec) {
	*out = *in
	if in.Mode != nil {
		in, out := &in.Mode, &out.Mode
		*out = new(DeploymentMode)
		**out = **in
	}
	if in.Environment != nil {
		in, out := &in.Environment, &out.Environment
		*out = new(Environment)
		**out = **in
	}
	if in.StorageEngine != nil {
		in, out := &in.StorageEngine, &out.StorageEngine
		*out = new(StorageEngine)
		**out = **in
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(string)
		**out = **in
	}
	if in.ImagePullPolicy != nil {
		in, out := &in.ImagePullPolicy, &out.ImagePullPolicy
		*out = new(v1.PullPolicy)
		**out = **in
	}
	if in.DowntimeAllowed != nil {
		in, out := &in.DowntimeAllowed, &out.DowntimeAllowed
		*out = new(bool)
		**out = **in
	}
	if in.DisableIPv6 != nil {
		in, out := &in.DisableIPv6, &out.DisableIPv6
		*out = new(bool)
		**out = **in
	}
	in.ExternalAccess.DeepCopyInto(&out.ExternalAccess)
	in.RocksDB.DeepCopyInto(&out.RocksDB)
	in.Authentication.DeepCopyInto(&out.Authentication)
	in.TLS.DeepCopyInto(&out.TLS)
	in.Sync.DeepCopyInto(&out.Sync)
	in.License.DeepCopyInto(&out.License)
	in.Metrics.DeepCopyInto(&out.Metrics)
	in.Single.DeepCopyInto(&out.Single)
	in.Agents.DeepCopyInto(&out.Agents)
	in.DBServers.DeepCopyInto(&out.DBServers)
	in.Coordinators.DeepCopyInto(&out.Coordinators)
	in.SyncMasters.DeepCopyInto(&out.SyncMasters)
	in.SyncWorkers.DeepCopyInto(&out.SyncWorkers)
	in.Chaos.DeepCopyInto(&out.Chaos)
	in.Bootstrap.DeepCopyInto(&out.Bootstrap)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentSpec.
func (in *DeploymentSpec) DeepCopy() *DeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(DeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentStatus) DeepCopyInto(out *DeploymentStatus) {
	*out = *in
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = make(ImageInfoList, len(*in))
		copy(*out, *in)
	}
	if in.CurrentImage != nil {
		in, out := &in.CurrentImage, &out.CurrentImage
		*out = new(ImageInfo)
		**out = **in
	}
	in.Members.DeepCopyInto(&out.Members)
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(ConditionList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Plan != nil {
		in, out := &in.Plan, &out.Plan
		*out = make(Plan, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AcceptedSpec != nil {
		in, out := &in.AcceptedSpec, &out.AcceptedSpec
		*out = new(DeploymentSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.SecretHashes != nil {
		in, out := &in.SecretHashes, &out.SecretHashes
		*out = new(SecretHashes)
		**out = **in
	}
	if in.ForceStatusReload != nil {
		in, out := &in.ForceStatusReload, &out.ForceStatusReload
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentStatus.
func (in *DeploymentStatus) DeepCopy() *DeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(DeploymentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentStatusMembers) DeepCopyInto(out *DeploymentStatusMembers) {
	*out = *in
	if in.Single != nil {
		in, out := &in.Single, &out.Single
		*out = make(MemberStatusList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Agents != nil {
		in, out := &in.Agents, &out.Agents
		*out = make(MemberStatusList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DBServers != nil {
		in, out := &in.DBServers, &out.DBServers
		*out = make(MemberStatusList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Coordinators != nil {
		in, out := &in.Coordinators, &out.Coordinators
		*out = make(MemberStatusList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SyncMasters != nil {
		in, out := &in.SyncMasters, &out.SyncMasters
		*out = make(MemberStatusList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SyncWorkers != nil {
		in, out := &in.SyncWorkers, &out.SyncWorkers
		*out = make(MemberStatusList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentStatusMembers.
func (in *DeploymentStatusMembers) DeepCopy() *DeploymentStatusMembers {
	if in == nil {
		return nil
	}
	out := new(DeploymentStatusMembers)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalAccessSpec) DeepCopyInto(out *ExternalAccessSpec) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(ExternalAccessType)
		**out = **in
	}
	if in.NodePort != nil {
		in, out := &in.NodePort, &out.NodePort
		*out = new(int)
		**out = **in
	}
	if in.LoadBalancerIP != nil {
		in, out := &in.LoadBalancerIP, &out.LoadBalancerIP
		*out = new(string)
		**out = **in
	}
	if in.LoadBalancerSourceRanges != nil {
		in, out := &in.LoadBalancerSourceRanges, &out.LoadBalancerSourceRanges
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AdvertisedEndpoint != nil {
		in, out := &in.AdvertisedEndpoint, &out.AdvertisedEndpoint
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalAccessSpec.
func (in *ExternalAccessSpec) DeepCopy() *ExternalAccessSpec {
	if in == nil {
		return nil
	}
	out := new(ExternalAccessSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageInfo) DeepCopyInto(out *ImageInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageInfo.
func (in *ImageInfo) DeepCopy() *ImageInfo {
	if in == nil {
		return nil
	}
	out := new(ImageInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in ImageInfoList) DeepCopyInto(out *ImageInfoList) {
	{
		in := &in
		*out = make(ImageInfoList, len(*in))
		copy(*out, *in)
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageInfoList.
func (in ImageInfoList) DeepCopy() ImageInfoList {
	if in == nil {
		return nil
	}
	out := new(ImageInfoList)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicenseSpec) DeepCopyInto(out *LicenseSpec) {
	*out = *in
	if in.SecretName != nil {
		in, out := &in.SecretName, &out.SecretName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseSpec.
func (in *LicenseSpec) DeepCopy() *LicenseSpec {
	if in == nil {
		return nil
	}
	out := new(LicenseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemberStatus) DeepCopyInto(out *MemberStatus) {
	*out = *in
	in.CreatedAt.DeepCopyInto(&out.CreatedAt)
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(ConditionList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RecentTerminations != nil {
		in, out := &in.RecentTerminations, &out.RecentTerminations
		*out = make([]metav1.Time, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SideCarSpecs != nil {
		in, out := &in.SideCarSpecs, &out.SideCarSpecs
		*out = make(map[string]v1.Container, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemberStatus.
func (in *MemberStatus) DeepCopy() *MemberStatus {
	if in == nil {
		return nil
	}
	out := new(MemberStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in MemberStatusList) DeepCopyInto(out *MemberStatusList) {
	{
		in := &in
		*out = make(MemberStatusList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemberStatusList.
func (in MemberStatusList) DeepCopy() MemberStatusList {
	if in == nil {
		return nil
	}
	out := new(MemberStatusList)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricsAuthenticationSpec) DeepCopyInto(out *MetricsAuthenticationSpec) {
	*out = *in
	if in.JWTTokenSecretName != nil {
		in, out := &in.JWTTokenSecretName, &out.JWTTokenSecretName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricsAuthenticationSpec.
func (in *MetricsAuthenticationSpec) DeepCopy() *MetricsAuthenticationSpec {
	if in == nil {
		return nil
	}
	out := new(MetricsAuthenticationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricsSpec) DeepCopyInto(out *MetricsSpec) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(string)
		**out = **in
	}
	in.Authentication.DeepCopyInto(&out.Authentication)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricsSpec.
func (in *MetricsSpec) DeepCopy() *MetricsSpec {
	if in == nil {
		return nil
	}
	out := new(MetricsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringSpec) DeepCopyInto(out *MonitoringSpec) {
	*out = *in
	if in.TokenSecretName != nil {
		in, out := &in.TokenSecretName, &out.TokenSecretName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringSpec.
func (in *MonitoringSpec) DeepCopy() *MonitoringSpec {
	if in == nil {
		return nil
	}
	out := new(MonitoringSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in PasswordSecretNameList) DeepCopyInto(out *PasswordSecretNameList) {
	{
		in := &in
		*out = make(PasswordSecretNameList, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PasswordSecretNameList.
func (in PasswordSecretNameList) DeepCopy() PasswordSecretNameList {
	if in == nil {
		return nil
	}
	out := new(PasswordSecretNameList)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Plan) DeepCopyInto(out *Plan) {
	{
		in := &in
		*out = make(Plan, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Plan.
func (in Plan) DeepCopy() Plan {
	if in == nil {
		return nil
	}
	out := new(Plan)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RocksDBEncryptionSpec) DeepCopyInto(out *RocksDBEncryptionSpec) {
	*out = *in
	if in.KeySecretName != nil {
		in, out := &in.KeySecretName, &out.KeySecretName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RocksDBEncryptionSpec.
func (in *RocksDBEncryptionSpec) DeepCopy() *RocksDBEncryptionSpec {
	if in == nil {
		return nil
	}
	out := new(RocksDBEncryptionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RocksDBSpec) DeepCopyInto(out *RocksDBSpec) {
	*out = *in
	in.Encryption.DeepCopyInto(&out.Encryption)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RocksDBSpec.
func (in *RocksDBSpec) DeepCopy() *RocksDBSpec {
	if in == nil {
		return nil
	}
	out := new(RocksDBSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretHashes) DeepCopyInto(out *SecretHashes) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretHashes.
func (in *SecretHashes) DeepCopy() *SecretHashes {
	if in == nil {
		return nil
	}
	out := new(SecretHashes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerGroupProbesSpec) DeepCopyInto(out *ServerGroupProbesSpec) {
	*out = *in
	if in.LivenessProbeDisabled != nil {
		in, out := &in.LivenessProbeDisabled, &out.LivenessProbeDisabled
		*out = new(bool)
		**out = **in
	}
	if in.ReadinessProbeDisabled != nil {
		in, out := &in.ReadinessProbeDisabled, &out.ReadinessProbeDisabled
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerGroupProbesSpec.
func (in *ServerGroupProbesSpec) DeepCopy() *ServerGroupProbesSpec {
	if in == nil {
		return nil
	}
	out := new(ServerGroupProbesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerGroupSpec) DeepCopyInto(out *ServerGroupSpec) {
	*out = *in
	if in.Count != nil {
		in, out := &in.Count, &out.Count
		*out = new(int)
		**out = **in
	}
	if in.MinCount != nil {
		in, out := &in.MinCount, &out.MinCount
		*out = new(int)
		**out = **in
	}
	if in.MaxCount != nil {
		in, out := &in.MaxCount, &out.MaxCount
		*out = new(int)
		**out = **in
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.StorageClassName != nil {
		in, out := &in.StorageClassName, &out.StorageClassName
		*out = new(string)
		**out = **in
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ServiceAccountName != nil {
		in, out := &in.ServiceAccountName, &out.ServiceAccountName
		*out = new(string)
		**out = **in
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Probes != nil {
		in, out := &in.Probes, &out.Probes
		*out = new(ServerGroupProbesSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.VolumeClaimTemplate != nil {
		in, out := &in.VolumeClaimTemplate, &out.VolumeClaimTemplate
		*out = new(v1.PersistentVolumeClaim)
		(*in).DeepCopyInto(*out)
	}
	if in.Sidecars != nil {
		in, out := &in.Sidecars, &out.Sidecars
		*out = make([]v1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerGroupSpec.
func (in *ServerGroupSpec) DeepCopy() *ServerGroupSpec {
	if in == nil {
		return nil
	}
	out := new(ServerGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SyncAuthenticationSpec) DeepCopyInto(out *SyncAuthenticationSpec) {
	*out = *in
	if in.JWTSecretName != nil {
		in, out := &in.JWTSecretName, &out.JWTSecretName
		*out = new(string)
		**out = **in
	}
	if in.ClientCASecretName != nil {
		in, out := &in.ClientCASecretName, &out.ClientCASecretName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SyncAuthenticationSpec.
func (in *SyncAuthenticationSpec) DeepCopy() *SyncAuthenticationSpec {
	if in == nil {
		return nil
	}
	out := new(SyncAuthenticationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SyncExternalAccessSpec) DeepCopyInto(out *SyncExternalAccessSpec) {
	*out = *in
	in.ExternalAccessSpec.DeepCopyInto(&out.ExternalAccessSpec)
	if in.MasterEndpoint != nil {
		in, out := &in.MasterEndpoint, &out.MasterEndpoint
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AccessPackageSecretNames != nil {
		in, out := &in.AccessPackageSecretNames, &out.AccessPackageSecretNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SyncExternalAccessSpec.
func (in *SyncExternalAccessSpec) DeepCopy() *SyncExternalAccessSpec {
	if in == nil {
		return nil
	}
	out := new(SyncExternalAccessSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SyncSpec) DeepCopyInto(out *SyncSpec) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	in.ExternalAccess.DeepCopyInto(&out.ExternalAccess)
	in.Authentication.DeepCopyInto(&out.Authentication)
	in.TLS.DeepCopyInto(&out.TLS)
	in.Monitoring.DeepCopyInto(&out.Monitoring)
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SyncSpec.
func (in *SyncSpec) DeepCopy() *SyncSpec {
	if in == nil {
		return nil
	}
	out := new(SyncSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSSpec) DeepCopyInto(out *TLSSpec) {
	*out = *in
	if in.CASecretName != nil {
		in, out := &in.CASecretName, &out.CASecretName
		*out = new(string)
		**out = **in
	}
	if in.AltNames != nil {
		in, out := &in.AltNames, &out.AltNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TTL != nil {
		in, out := &in.TTL, &out.TTL
		*out = new(Duration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSSpec.
func (in *TLSSpec) DeepCopy() *TLSSpec {
	if in == nil {
		return nil
	}
	out := new(TLSSpec)
	in.DeepCopyInto(out)
	return out
}
