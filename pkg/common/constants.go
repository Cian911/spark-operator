/*
Copyright 2024 The Kubeflow authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package common

const (
	ErrorCodePodAlreadyExists = "code=409"
)

const (
	SparkApplicationFinalizerName          = "sparkoperator.k8s.io/finalizer"
	ScheduledSparkApplicationFinalizerName = "sparkoperator.k8s.io/finalizer"
)

const (
	RSAKeySize = 2048
)

const (
	CAKeyPem      = "ca-key.pem"
	CACertPem     = "ca-cert.pem"
	ServerKeyPem  = "server-key.pem"
	ServerCertPem = "server-cert.pem"
)

// Kubernetes volume types.
const (
	VolumeTypeEmptyDir              = "emptyDir"
	VolumeTypeHostPath              = "hostPath"
	VolumeTypeNFS                   = "nfs"
	VolumeTypePersistentVolumeClaim = "persistentVolumeClaim"
)

// Kubernetes memory types
const (
	MemoryTypeLimit = "limits.memory"
)

const (
	// Epsilon is a small number used to compare 64 bit floating point numbers.
	Epsilon = 1e-9
)
