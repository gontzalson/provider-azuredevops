package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	"github.com/gontzalson/provider-azuredevops/config/project"
)

const (
	resourcePrefix = "azuredevops"
	modulePath     = "github.com/gontzalson/provider-azuredevops"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("azuredevops.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		project.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

// ExternalNameConfigured returns the list of all resources whose external
// name is configured manually.
func ExternalNameConfigured() []string {
	// Return the list of all resources whose external name is configured manually
	return []string{
		// Project resources
		"azuredevops_project$",
		"azuredevops_project_features$",
		"azuredevops_project_pipeline_settings$",
	}
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() ujconfig.ResourceOption {
	return func(r *ujconfig.Resource) {
		if name, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = name
		}
	}
}

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]ujconfig.ExternalName{
	// Project resources - ID-based external names
	"azuredevops_project":                    ujconfig.IdentifierFromProvider,
	"azuredevops_project_features":           ujconfig.IdentifierFromProvider,
	"azuredevops_project_pipeline_settings":  ujconfig.IdentifierFromProvider,
}

// GetProviderNamespaced returns the namespaced provider configuration
// This is used for multi-tenancy scenarios where resources are scoped to namespaces
func GetProviderNamespaced() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("azuredevops.m.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		),
		ujconfig.WithExampleManifestConfiguration(ujconfig.ExampleManifestConfiguration{
			ManagedResourceNamespace: "crossplane-system",
		}))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		project.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
