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
// Environment Spec (Kubernetes-facing envelope)
// ─────────────────────────────────────────────────────────────
//

// EnvironmentSpec is a Kubernetes-native envelope around the canonical
// BlanketOps Environments Environment contract.
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
type EnvironmentSpec struct {

	// Contract is the canonical BlanketOps Environments Environment specification.
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
// Environment Status (Kubernetes-facing envelope)
// ─────────────────────────────────────────────────────────────
//

// EnvironmentStatus represents observed state owned by the controller.
//
// This mirrors the contract status but remains opaque to Kubernetes.
type EnvironmentStatus struct {

	// Contract is the canonical BlanketOps Environments Environment status.
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
// Environment CRD
// ─────────────────────────────────────────────────────────────
//

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// +kubebuilder:printcolumn:name="Phase",type="string",JSONPath=".status.contract.phase",description="The aggregate health phase of the delivery chain"
// +kubebuilder:printcolumn:name="Application",type="string",JSONPath=".spec.contract.applicationName",description="The human-readable application identifier"
// +kubebuilder:printcolumn:name="Type",type="string",JSONPath=".spec.contract.environmentType",description="Classification of the environment (e.g., Staging, Production)"
// +kubebuilder:printcolumn:name="Version",type="string",JSONPath=".spec.contract.version",description="The current semantic version deployed"
// +kubebuilder:printcolumn:name="Branch",type="string",JSONPath=".spec.contract.branch",description="The source Git branch tracked by this environment",priority=1
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.contract.message",description="Aggregated failure diagnostics summary",priority=1
// +kubebuilder:printcolumn:name="Last-Updated",type="date",JSONPath=".status.contract.lastUpdated",description="The last time the environment status was updated",priority=1

// Environment tracks a versioned, isolated execution context where applications run.
// Environment represents a declarative request to Environment an artifact.
// This resource is a Kubernetes-native envelope around a transport-agnostic BlanketOps Environments Environment contract.
type Environment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EnvironmentSpec   `json:"spec,omitempty"`
	Status EnvironmentStatus `json:"status,omitempty"`
}

//
// ─────────────────────────────────────────────────────────────
// EnvironmentList
// ─────────────────────────────────────────────────────────────
//

// +kubebuilder:object:root=true

// EnvironmentList contains a list of Environment resources.
type EnvironmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Environment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Environment{}, &EnvironmentList{})
}
