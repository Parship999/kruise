/*
Copyright 2021 The Kruise Authors.

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

package defaults

import (
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/kubernetes/pkg/apis/core/v1"
)

// SetDefaultPodSpec sets default pod spec
func SetDefaultPodSpec(in *corev1.PodSpec) {
	v1.SetDefaults_PodSpec(in)
	//default pod volumes
	SetDefaultPodVolumes(in.Volumes)
	for i := range in.InitContainers {
		a := &in.InitContainers[i]
		v1.SetDefaults_Container(a)
		for j := range a.Ports {
			b := &a.Ports[j]
			if b.Protocol == "" {
				b.Protocol = "TCP"
			}
		}
		for j := range a.Env {
			b := &a.Env[j]
			if b.ValueFrom != nil {
				if b.ValueFrom.FieldRef != nil {
					v1.SetDefaults_ObjectFieldSelector(b.ValueFrom.FieldRef)
				}
			}
		}
		v1.SetDefaults_ResourceList(&a.Resources.Limits)
		v1.SetDefaults_ResourceList(&a.Resources.Requests)
		if a.LivenessProbe != nil {
			v1.SetDefaults_Probe(a.LivenessProbe)
			if a.LivenessProbe.ProbeHandler.HTTPGet != nil {
				v1.SetDefaults_HTTPGetAction(a.LivenessProbe.ProbeHandler.HTTPGet)
			}
		}
		if a.ReadinessProbe != nil {
			v1.SetDefaults_Probe(a.ReadinessProbe)
			if a.ReadinessProbe.ProbeHandler.HTTPGet != nil {
				v1.SetDefaults_HTTPGetAction(a.ReadinessProbe.ProbeHandler.HTTPGet)
			}
		}
		if a.StartupProbe != nil {
			v1.SetDefaults_Probe(a.StartupProbe)
			if a.StartupProbe.ProbeHandler.HTTPGet != nil {
				v1.SetDefaults_HTTPGetAction(a.StartupProbe.ProbeHandler.HTTPGet)
			}
		}
		if a.Lifecycle != nil {
			if a.Lifecycle.PostStart != nil {
				if a.Lifecycle.PostStart.HTTPGet != nil {
					v1.SetDefaults_HTTPGetAction(a.Lifecycle.PostStart.HTTPGet)
				}
			}
			if a.Lifecycle.PreStop != nil {
				if a.Lifecycle.PreStop.HTTPGet != nil {
					v1.SetDefaults_HTTPGetAction(a.Lifecycle.PreStop.HTTPGet)
				}
			}
		}
	}
	for i := range in.Containers {
		a := &in.Containers[i]
		// For in-place update, we set default imagePullPolicy to Always
		if a.ImagePullPolicy == "" {
			a.ImagePullPolicy = corev1.PullAlways
		}
		v1.SetDefaults_Container(a)
		for j := range a.Ports {
			b := &a.Ports[j]
			if b.Protocol == "" {
				b.Protocol = "TCP"
			}
		}
		for j := range a.Env {
			b := &a.Env[j]
			if b.ValueFrom != nil {
				if b.ValueFrom.FieldRef != nil {
					v1.SetDefaults_ObjectFieldSelector(b.ValueFrom.FieldRef)
				}
			}
		}
		v1.SetDefaults_ResourceList(&a.Resources.Limits)
		v1.SetDefaults_ResourceList(&a.Resources.Requests)
		if a.LivenessProbe != nil {
			v1.SetDefaults_Probe(a.LivenessProbe)
			if a.LivenessProbe.ProbeHandler.HTTPGet != nil {
				v1.SetDefaults_HTTPGetAction(a.LivenessProbe.ProbeHandler.HTTPGet)
			}
		}
		if a.ReadinessProbe != nil {
			v1.SetDefaults_Probe(a.ReadinessProbe)
			if a.ReadinessProbe.ProbeHandler.HTTPGet != nil {
				v1.SetDefaults_HTTPGetAction(a.ReadinessProbe.ProbeHandler.HTTPGet)
			}
		}
		if a.StartupProbe != nil {
			v1.SetDefaults_Probe(a.StartupProbe)
			if a.StartupProbe.ProbeHandler.HTTPGet != nil {
				v1.SetDefaults_HTTPGetAction(a.StartupProbe.ProbeHandler.HTTPGet)
			}
		}
		if a.Lifecycle != nil {
			if a.Lifecycle.PostStart != nil {
				if a.Lifecycle.PostStart.HTTPGet != nil {
					v1.SetDefaults_HTTPGetAction(a.Lifecycle.PostStart.HTTPGet)
				}
			}
			if a.Lifecycle.PreStop != nil {
				if a.Lifecycle.PreStop.HTTPGet != nil {
					v1.SetDefaults_HTTPGetAction(a.Lifecycle.PreStop.HTTPGet)
				}
			}
		}
	}
	for i := range in.EphemeralContainers {
		a := &in.EphemeralContainers[i]
		for j := range a.EphemeralContainerCommon.Ports {
			b := &a.EphemeralContainerCommon.Ports[j]
			if b.Protocol == "" {
				b.Protocol = "TCP"
			}
		}
		for j := range a.EphemeralContainerCommon.Env {
			b := &a.EphemeralContainerCommon.Env[j]
			if b.ValueFrom != nil {
				if b.ValueFrom.FieldRef != nil {
					v1.SetDefaults_ObjectFieldSelector(b.ValueFrom.FieldRef)
				}
			}
		}
		v1.SetDefaults_ResourceList(&a.EphemeralContainerCommon.Resources.Limits)
		v1.SetDefaults_ResourceList(&a.EphemeralContainerCommon.Resources.Requests)
		if a.EphemeralContainerCommon.LivenessProbe != nil {
			v1.SetDefaults_Probe(a.EphemeralContainerCommon.LivenessProbe)
			if a.EphemeralContainerCommon.LivenessProbe.ProbeHandler.HTTPGet != nil {
				v1.SetDefaults_HTTPGetAction(a.EphemeralContainerCommon.LivenessProbe.ProbeHandler.HTTPGet)
			}
		}
		if a.EphemeralContainerCommon.ReadinessProbe != nil {
			v1.SetDefaults_Probe(a.EphemeralContainerCommon.ReadinessProbe)
			if a.EphemeralContainerCommon.ReadinessProbe.ProbeHandler.HTTPGet != nil {
				v1.SetDefaults_HTTPGetAction(a.EphemeralContainerCommon.ReadinessProbe.ProbeHandler.HTTPGet)
			}
		}
		if a.EphemeralContainerCommon.StartupProbe != nil {
			v1.SetDefaults_Probe(a.EphemeralContainerCommon.StartupProbe)
			if a.EphemeralContainerCommon.StartupProbe.ProbeHandler.HTTPGet != nil {
				v1.SetDefaults_HTTPGetAction(a.EphemeralContainerCommon.StartupProbe.ProbeHandler.HTTPGet)
			}
		}
		if a.EphemeralContainerCommon.Lifecycle != nil {
			if a.EphemeralContainerCommon.Lifecycle.PostStart != nil {
				if a.EphemeralContainerCommon.Lifecycle.PostStart.HTTPGet != nil {
					v1.SetDefaults_HTTPGetAction(a.EphemeralContainerCommon.Lifecycle.PostStart.HTTPGet)
				}
			}
			if a.EphemeralContainerCommon.Lifecycle.PreStop != nil {
				if a.EphemeralContainerCommon.Lifecycle.PreStop.HTTPGet != nil {
					v1.SetDefaults_HTTPGetAction(a.EphemeralContainerCommon.Lifecycle.PreStop.HTTPGet)
				}
			}
		}
	}
	v1.SetDefaults_ResourceList(&in.Overhead)
}

func SetDefaultPodVolumes(volumes []corev1.Volume) {
	for i := range volumes {
		a := &volumes[i]
		v1.SetDefaults_Volume(a)
		if a.VolumeSource.HostPath != nil {
			v1.SetDefaults_HostPathVolumeSource(a.VolumeSource.HostPath)
		}
		if a.VolumeSource.Secret != nil {
			v1.SetDefaults_SecretVolumeSource(a.VolumeSource.Secret)
		}
		if a.VolumeSource.DownwardAPI != nil {
			v1.SetDefaults_DownwardAPIVolumeSource(a.VolumeSource.DownwardAPI)
			for j := range a.VolumeSource.DownwardAPI.Items {
				b := &a.VolumeSource.DownwardAPI.Items[j]
				if b.FieldRef != nil {
					v1.SetDefaults_ObjectFieldSelector(b.FieldRef)
				}
			}
		}
		if a.VolumeSource.ConfigMap != nil {
			v1.SetDefaults_ConfigMapVolumeSource(a.VolumeSource.ConfigMap)
		}
		if a.VolumeSource.Projected != nil {
			v1.SetDefaults_ProjectedVolumeSource(a.VolumeSource.Projected)
			for j := range a.VolumeSource.Projected.Sources {
				b := &a.VolumeSource.Projected.Sources[j]
				if b.DownwardAPI != nil {
					for k := range b.DownwardAPI.Items {
						c := &b.DownwardAPI.Items[k]
						if c.FieldRef != nil {
							v1.SetDefaults_ObjectFieldSelector(c.FieldRef)
						}
					}
				}
				if b.ServiceAccountToken != nil {
					v1.SetDefaults_ServiceAccountTokenProjection(b.ServiceAccountToken)
				}
			}
		}
	}
}
