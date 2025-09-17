/*
SPDX-FileCopyrightText: 2025 Upbound Inc. <https://upbound.io>

SPDX-License-Identifier: Apache-2.0
*/

package v1beta1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/v2/apis/common/v1"
)

// A ProviderConfigSpec defines the desired state of a ProviderConfig.
type ProviderConfigSpec struct {
	// Credentials required to authenticate to this provider.
	// +optional
	Credentials []ProviderCredentials `json:"credentials"`

	// Configuration that should be injected into all workspaces that use
	// this provider config, expressed as inline HCL. This can be used to
	// automatically inject Terraform provider configuration blocks.
	// +optional
	Configuration *string `json:"configuration,omitempty"`

	// Tofu backend file configuration content,
	// it has the contents of the backend block as top-level attributes,
	// without the need to wrap it in another opentofu or backend block.
	// More details at https://opentofu.org/docs/language/settings/backends/configuration/#file.
	// +optional
	BackendFile *string `json:"backendFile,omitempty"`

	// PluginCache enables tofu provider plugin caching mechanism
	// https://opentofu.org/docs/cli/config/config-file/#provider-plugin-cache
	// +optional
	// +kubebuilder:default=true
	PluginCache *bool `json:"pluginCache,omitempty"`
}

// ProviderCredentials required to authenticate.
type ProviderCredentials struct {
	// Filename (relative to main.tf) to which these provider credentials
	// should be written.
	Filename string `json:"filename"`

	// Source of the provider credentials.
	// +kubebuilder:validation:Enum=None;Secret;Environment;Filesystem
	Source xpv1.CredentialsSource `json:"source"`

	xpv1.CommonCredentialSelectors `json:",inline"`
}

// A ProviderConfigStatus reflects the observed state of a ProviderConfig.
type ProviderConfigStatus struct {
	xpv1.ProviderConfigStatus `json:",inline"`
}
