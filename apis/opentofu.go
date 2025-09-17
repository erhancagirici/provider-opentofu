/*
SPDX-FileCopyrightText: 2025 Upbound Inc. <https://upbound.io>

SPDX-License-Identifier: Apache-2.0
*/

// Package apis contains Kubernetes API for the OpenTofu provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	v1beta1Cluster "github.com/upbound/provider-opentofu/apis/cluster/v1beta1"
	v1beta1Namespaced "github.com/upbound/provider-opentofu/apis/namespaced/v1beta1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		v1beta1Cluster.SchemeBuilder.AddToScheme,
		v1beta1Namespaced.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}
