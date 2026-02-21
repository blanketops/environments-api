// Code generated manually. DO NOT EDIT via generator.

package openapi

import (
	common "k8s.io/kube-openapi/pkg/common"
	spec "k8s.io/kube-openapi/pkg/validation/spec"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/ntlaletsi70/blanketops-environments-api/api/environments/v1alpha1.Environment":       schemaEnvironment(ref),
		"github.com/ntlaletsi70/blanketops-environments-api/api/environments/v1alpha1.EnvironmentList":   schemaEnvironmentList(ref),
		"github.com/ntlaletsi70/blanketops-environments-api/api/environments/v1alpha1.EnvironmentSpec":   schemaEnvironmentSpec(ref),
		"github.com/ntlaletsi70/blanketops-environments-api/api/environments/v1alpha1.EnvironmentStatus": schemaEnvironmentStatus(ref),
	}
}

func schemaEnvironment(_ common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type:        []string{"object"},
				Description: "Environment is a BlanketOps environment resource.",
			},
		},
	}
}

func schemaEnvironmentList(_ common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type:        []string{"object"},
				Description: "EnvironmentList contains a list of Environment resources.",
			},
		},
	}
}

func schemaEnvironmentSpec(_ common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type:        []string{"object"},
				Description: "EnvironmentSpec is the spec for an Environment.",
			},
		},
	}
}

func schemaEnvironmentStatus(_ common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type:        []string{"object"},
				Description: "EnvironmentStatus is the status for an Environment.",
			},
		},
	}
}
