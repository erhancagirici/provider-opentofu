/*
SPDX-FileCopyrightText: 2025 Upbound Inc. <https://upbound.io>

SPDX-License-Identifier: Apache-2.0
*/

package v1beta1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/v2/apis/common/v1"
	sharedv1beta1 "github.com/upbound/provider-opentofu/apis/shared/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// A Var represents a tofu configuration variable.
type Var = sharedv1beta1.Var

// A VarFileSource specifies the source of a Terraform vars file.
// +kubebuilder:validation:Enum=ConfigMapKey;SecretKey
type VarFileSource = sharedv1beta1.VarFileSource

// Vars file sources.
const (
	VarFileSourceConfigMapKey = sharedv1beta1.VarFileSourceConfigMapKey
	VarFileSourceSecretKey    = sharedv1beta1.VarFileSourceSecretKey
)

// A FileFormat specifies the format of a Terraform file.
// +kubebuilder:validation:Enum=HCL;JSON
type FileFormat = sharedv1beta1.FileFormat

// Vars file formats.
var (
	FileFormatHCL  = sharedv1beta1.FileFormatHCL
	FileFormatJSON = sharedv1beta1.FileFormatJSON
)

// A VarFile is a file containing many Terraform variables.
type VarFile = sharedv1beta1.VarFile

// An EnvVar specifies an environment variable to be set for the workspace.
type EnvVar = sharedv1beta1.EnvVar

// A KeyReference references a key within a Secret or a ConfigMap.
type KeyReference = sharedv1beta1.KeyReference

// A ModuleSource represents the source of a Terraform module.
// +kubebuilder:validation:Enum=Remote;Inline
type ModuleSource = sharedv1beta1.ModuleSource

// Module sources.
const (
	ModuleSourceRemote = sharedv1beta1.ModuleSourceRemote
	ModuleSourceInline = sharedv1beta1.ModuleSourceInline
)

// WorkspaceParameters are the configurable fields of a Workspace.
type WorkspaceParameters = sharedv1beta1.WorkspaceParameters

// WorkspaceObservation are the observable fields of a Workspace.
type WorkspaceObservation = sharedv1beta1.WorkspaceObservation

// A WorkspaceSpec defines the desired state of a Workspace.
type WorkspaceSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       WorkspaceParameters `json:"forProvider"`
}

// A WorkspaceStatus represents the observed state of a Workspace.
type WorkspaceStatus = sharedv1beta1.WorkspaceStatus

// +kubebuilder:object:root=true

// A Workspace of OpenTofu Configuration.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opentofu}
type Workspace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WorkspaceSpec   `json:"spec"`
	Status WorkspaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WorkspaceList contains a list of Workspace
type WorkspaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Workspace `json:"items"`
}
