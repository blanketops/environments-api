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
// Route Spec (Kubernetes-facing envelope)
// ─────────────────────────────────────────────────────────────
//

// RouteSpec is a Kubernetes-native envelope around the canonical
// BlanketOps Route contract.
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
type RouteSpec struct {

	// Contract is the canonical BlanketOps Route specification.
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
// Route Status (Kubernetes-facing envelope)
// ─────────────────────────────────────────────────────────────
//

// RouteStatus represents observed state owned by the controller.
//
// This mirrors the contract status but remains opaque to Kubernetes.
type RouteStatus struct {

	// Contract is the canonical BlanketOps Route status.
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
// Route CRD
// ─────────────────────────────────────────────────────────────
//

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// +kubebuilder:printcolumn:name="Phase",type="string",JSONPath=".status.contract.phase",description="The operational lifecycle phase of the network path"
// +kubebuilder:printcolumn:name="Host",type="string",JSONPath=".spec.contract.host",description="The target domain address to match"
// +kubebuilder:printcolumn:name="Path",type="string",JSONPath=".spec.contract.path",description="The HTTP path prefix endpoint rule matching string"
// +kubebuilder:printcolumn:name="Service-Unit",type="string",JSONPath=".spec.contract.serviceUnitRef.name",description="The underlying target service unit CR being exposed"
// +kubebuilder:printcolumn:name="Enabled",type="boolean",JSONPath=".spec.contract.enabled",description="Whether this network trajectory rule is actively serving traffic"
// +kubebuilder:printcolumn:name="Endpoint",type="string",JSONPath=".status.contract.resolvedAddress",description="The load balancer host target address string output",priority=1
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.contract.message",description="Failure diagnosis log description detailed summary",priority=1
// +kubebuilder:printcolumn:name="Last-Updated",type="date",JSONPath=".status.contract.lastUpdated",description="The last time the Route status was updated",priority=1

// Route declares an application ingress trajectory mapping pattern ruleset.s
//
// Route represents a declarative request to Route an artifact.
//
// This resource is a Kubernetes-native envelope around a
// transport-agnostic BlanketOps Route contract.
type Route struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RouteSpec   `json:"spec,omitempty"`
	Status RouteStatus `json:"status,omitempty"`
}

//
// ─────────────────────────────────────────────────────────────
// RouteList
// ─────────────────────────────────────────────────────────────
//

// +kubebuilder:object:root=true

// RouteList contains a list of Route resources.
type RouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Route `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Route{}, &RouteList{})
}
