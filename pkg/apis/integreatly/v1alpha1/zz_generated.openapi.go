// +build !ignore_autogenerated

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/integr8ly/grafana-operator/pkg/apis/integreatly/v1alpha1.Grafana":                 schema_pkg_apis_integreatly_v1alpha1_Grafana(ref),
		"github.com/integr8ly/grafana-operator/pkg/apis/integreatly/v1alpha1.GrafanaDashboard":        schema_pkg_apis_integreatly_v1alpha1_GrafanaDashboard(ref),
		"github.com/integr8ly/grafana-operator/pkg/apis/integreatly/v1alpha1.GrafanaDataSource":       schema_pkg_apis_integreatly_v1alpha1_GrafanaDataSource(ref),
		"github.com/integr8ly/grafana-operator/pkg/apis/integreatly/v1alpha1.GrafanaDataSourceSpec":   schema_pkg_apis_integreatly_v1alpha1_GrafanaDataSourceSpec(ref),
		"github.com/integr8ly/grafana-operator/pkg/apis/integreatly/v1alpha1.GrafanaDataSourceStatus": schema_pkg_apis_integreatly_v1alpha1_GrafanaDataSourceStatus(ref),
	}
}

func schema_pkg_apis_integreatly_v1alpha1_Grafana(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Grafana is the Schema for the grafanas API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/integr8ly/grafana-operator/pkg/apis/integreatly/v1alpha1.GrafanaSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/integr8ly/grafana-operator/pkg/apis/integreatly/v1alpha1.GrafanaStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/integr8ly/grafana-operator/pkg/apis/integreatly/v1alpha1.GrafanaSpec", "github.com/integr8ly/grafana-operator/pkg/apis/integreatly/v1alpha1.GrafanaStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_integreatly_v1alpha1_GrafanaDashboard(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "GrafanaDashboard is the Schema for the grafanadashboards API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/integr8ly/grafana-operator/pkg/apis/integreatly/v1alpha1.GrafanaDashboardSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/integr8ly/grafana-operator/pkg/apis/integreatly/v1alpha1.GrafanaDashboardStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/integr8ly/grafana-operator/pkg/apis/integreatly/v1alpha1.GrafanaDashboardSpec", "github.com/integr8ly/grafana-operator/pkg/apis/integreatly/v1alpha1.GrafanaDashboardStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_integreatly_v1alpha1_GrafanaDataSource(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "GrafanaDataSource is the Schema for the grafanadatasources API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/integr8ly/grafana-operator/pkg/apis/integreatly/v1alpha1.GrafanaDataSourceSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/integr8ly/grafana-operator/pkg/apis/integreatly/v1alpha1.GrafanaDataSourceStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/integr8ly/grafana-operator/pkg/apis/integreatly/v1alpha1.GrafanaDataSourceSpec", "github.com/integr8ly/grafana-operator/pkg/apis/integreatly/v1alpha1.GrafanaDataSourceStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_integreatly_v1alpha1_GrafanaDataSourceSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "GrafanaDataSourceSpec defines the desired state of GrafanaDataSource",
				Type:        []string{"object"},
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_integreatly_v1alpha1_GrafanaDataSourceStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "GrafanaDataSourceStatus defines the observed state of GrafanaDataSource",
				Type:        []string{"object"},
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}