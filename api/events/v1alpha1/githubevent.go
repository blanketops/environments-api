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
// GitHubEvent Spec (Kubernetes-facing envelope)
// ─────────────────────────────────────────────────────────────
//

// GitHubEventSpec is a Kubernetes-native envelope around the canonical
// BlanketOps GitHubEvent contract.
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
type GitHubEventSpec struct {

	// Contract is the canonical BlanketOps GitHubEvent specification.
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
// GitHubEvent Status (Kubernetes-facing envelope)
// ─────────────────────────────────────────────────────────────
//

// GitHubEventStatus represents observed state owned by the controller.
//
// This mirrors the contract status but remains opaque to Kubernetes.
type GitHubEventStatus struct {

	// Contract is the canonical BlanketOps GitHubEvent status.
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
// GitHubEvent CRD
// ─────────────────────────────────────────────────────────────
//

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// +kubebuilder:printcolumn:name="Event-Type",type="string",JSONPath=".spec.contract.eventType",description="The provider event action (e.g., PUSH, PULL_REQUEST)"
// +kubebuilder:printcolumn:name="Repository",type="string",JSONPath=".spec.contract.repository",description="The source repository origin owner/name"
// +kubebuilder:printcolumn:name="Accepted",type="boolean",JSONPath=".status.contract.accepted",description="True if webhook signature validation passed"
// +kubebuilder:printcolumn:name="Triggered",type="boolean",JSONPath=".status.contract.triggered",description="True if an automation build trigger kicked off"
// +kubebuilder:printcolumn:name="Commit-SHA",type="string",JSONPath=".spec.contract.commitSha",description="The target commit hex identifier",priority=1
// +kubebuilder:printcolumn:name="Triggered-Ref",type="string",JSONPath=".status.contract.triggeredRef",description="Reference to the downstream resource invoked",priority=1
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.contract.reason",description="Rejection or execution bypass message detail",priority=1
// +kubebuilder:printcolumn:name="Last-Updated",type="date",JSONPath=".status.contract.lastUpdated",description="The last time the GitHubEvent status was updated",priority=1

// GitHubEvent represents an immutable log of an external Git platform action.
// This resource is a Kubernetes-native envelope around a transport-agnostic BlanketOps GitHubEvent contract.
// GitHubEvent represents a declarative request to GitHubEvent an artifact.
//
// This resource is a Kubernetes-native envelope around a
// transport-agnostic BlanketOps GitHubEvent contract.
type GitHubEvent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GitHubEventSpec   `json:"spec,omitempty"`
	Status            GitHubEventStatus `json:"status,omitempty"`
}

//
// ─────────────────────────────────────────────────────────────
// GitHubEventList
// ─────────────────────────────────────────────────────────────
//

// +kubebuilder:object:root=true

// GitHubEventList contains a list of GitHubEvent resources.
type GitHubEventList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GitHubEvent `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GitHubEvent{}, &GitHubEventList{})
}
