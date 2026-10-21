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
// GitRepository Spec (Kubernetes-facing envelope)
// ─────────────────────────────────────────────────────────────
//

// GitRepositorySpec is a Kubernetes-native envelope around the canonical
// BlanketOps GitRepository contract.
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
type GitRepositorySpec struct {

	// Contract is the canonical BlanketOps GitRepository specification.
	//
	// This field is intentionally opaque to Kubernetes and schema generation.
	// It is preserved verbatim by the API server.
	//
	// +kubebuilder:validation:Required
	// +kubebuilder:pruning:PreserveUnknownFields
	Contract runtime.RawExtension `json:"contract"`
}

// ─────────────────────────────────────────────────────────────
// GitRepository Status (Kubernetes-facing envelope)
// ─────────────────────────────────────────────────────────────
//
// GitRepositoryStatus represents observed state owned by the controller.
//
// This mirrors the contract status but remains opaque to Kubernetes.
type GitRepositoryStatus struct {

	// Contract is the canonical BlanketOps GitRepository status.
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
// GitRepository CRD
// ─────────────────────────────────────────────────────────────
//

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// +kubebuilder:printcolumn:name="Ready",type="boolean",JSONPath=".status.contract.ready",description="True if provider webhooks are registered and active"
// +kubebuilder:printcolumn:name="Provider",type="string",JSONPath=".spec.contract.provider",description="The code hosting platform (e.g., GITHUB, GITLAB)"
// +kubebuilder:printcolumn:name="Repository",type="string",JSONPath=".spec.contract.repository.name",description="The target repository identifier"
// +kubebuilder:printcolumn:name="Owner",type="string",JSONPath=".spec.contract.repository.owner",description="The user or organization account on the provider"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.contract.reason",description="Deregistration or connection failure summary detail",priority=1
// +kubebuilder:printcolumn:name="Last-Updated",type="date",JSONPath=".status.contract.lastUpdated",description="The last time the GitRepository status was updated",priority=1

// GitRepository declares an external source system observed by the delivery controller.
// This resource is a Kubernetes-native envelope around a transport-agnostic BlanketOps GitRepository contract.
// GitRepository represents a declarative request to GitRepository an artifact.
//
// This resource is a Kubernetes-native envelope around a
// transport-agnostic BlanketOps GitRepository contract.
type GitRepository struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GitRepositorySpec   `json:"spec,omitempty"`
	Status            GitRepositoryStatus `json:"status,omitempty"`
}

//
// ─────────────────────────────────────────────────────────────
// GitRepositoryList
// ─────────────────────────────────────────────────────────────
//

// +kubebuilder:object:root=true

// GitRepositoryList contains a list of GitRepository resources.
type GitRepositoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GitRepository `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GitRepository{}, &GitRepositoryList{})
}
