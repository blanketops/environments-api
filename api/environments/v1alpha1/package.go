/*
Copyright 2026 The BlanketOps Authors.
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

// +k8s:openapi-gen=true
package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

//
// ─────────────────────────────────────────────────────────────
// Package Spec (Kubernetes-facing envelope)
// ─────────────────────────────────────────────────────────────
//

// PackageSpec is a Kubernetes-native envelope around the canonical
// BlanketOps Environment Package contract.
//
// IMPORTANT:
// - Kubernetes does NOT understand the contents of `Contract`
// - Kubernetes does NOT validate the contents of `Contract`
// - Kubernetes ONLY stores and round-trips this field
//
// Ownership boundaries:
// - API server: envelope + metadata
// - Controller: lifecycle orchestration
// - Contract layer: semantic meaning
type PackageSpec struct {

	// Contract is the canonical BlanketOps Environments Package specification.
	//
	// This field is intentionally opaque to Kubernetes and schema generation.
	// It is preserved verbatim by the API server.
	//
	// +kubebuilder:validation:Required
	// +kubebuilder:pruning:PreserveUnknownFields
	Contract runtime.RawExtension `json:"contract"`
}

//
// ─────────────────────────────────────────────────────────────
// Package Status (Kubernetes-facing envelope)
// ─────────────────────────────────────────────────────────────
//

// PackageStatus represents observed state owned by the controller.
//
// This mirrors the contract status but remains opaque to Kubernetes.
type PackageStatus struct {

	// Contract is the canonical BlanketOps Environments Package status.
	//
	// This field is opaque and preserved verbatim.
	//
	// +optional
	// +kubebuilder:pruning:PreserveUnknownFields
	Contract runtime.RawExtension `json:"contract,omitempty"`

	// Conditions follows standard Kubernetes condition conventions.
	// These are intended for kubectl, UIs, and ecosystem tooling.
	//
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

//
// ─────────────────────────────────────────────────────────────
// Package CRD
// ─────────────────────────────────────────────────────────────
//

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// +kubebuilder:printcolumn:name="Phase",type="string",JSONPath=".status.contract.phase",description="The current reconciliation phase of the package"
// +kubebuilder:printcolumn:name="Package-Name",type="string",JSONPath=".spec.contract.name",description="The logical name of the packaged application"
// +kubebuilder:printcolumn:name="Version",type="string",JSONPath=".spec.contract.version",description="The version identifier of the package contents"
// +kubebuilder:printcolumn:name="Enabled",type="boolean",JSONPath=".spec.contract.enabled",description="Whether this package is actively reconciling"
// +kubebuilder:printcolumn:name="Diff-Applied",type="boolean",JSONPath=".status.contract.diffApplied",description="Whether the last sync resulted in changes",priority=1
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.contract.message",description="Reconciliation failure or diagnostics summary",priority=1
// +kubebuilder:printcolumn:name="Last-Updated",type="date",JSONPath=".status.contract.lastUpdated",description="The last time the package status was updated",priority=1

// Package represents a versioned delivery artifact tracking manifest state.
// Package represents a declarative request to Package an artifact.
// This resource is a Kubernetes-native envelope around a transport-agnostic BlanketOps Environments Package contract.
type Package struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PackageSpec   `json:"spec,omitempty"`
	Status PackageStatus `json:"status,omitempty"`
}

//
// ─────────────────────────────────────────────────────────────
// PackageList
// ─────────────────────────────────────────────────────────────
//

// +kubebuilder:object:root=true

// PackageList contains a list of Package resources.
type PackageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Package `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Package{}, &PackageList{})
}
