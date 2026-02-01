package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	"github.com/gontzalson/provider-azuredevops/config/agent"
	"github.com/gontzalson/provider-azuredevops/config/branchpolicy"
	"github.com/gontzalson/provider-azuredevops/config/build"
	"github.com/gontzalson/provider-azuredevops/config/check"
	"github.com/gontzalson/provider-azuredevops/config/environment"
	"github.com/gontzalson/provider-azuredevops/config/feed"
	"github.com/gontzalson/provider-azuredevops/config/git"
	"github.com/gontzalson/provider-azuredevops/config/group"
	"github.com/gontzalson/provider-azuredevops/config/other"
	"github.com/gontzalson/provider-azuredevops/config/permissions"
	"github.com/gontzalson/provider-azuredevops/config/project"
	"github.com/gontzalson/provider-azuredevops/config/repositorypolicy"
	"github.com/gontzalson/provider-azuredevops/config/serviceendpoint"
	"github.com/gontzalson/provider-azuredevops/config/team"
	"github.com/gontzalson/provider-azuredevops/config/variablegroup"
	"github.com/gontzalson/provider-azuredevops/config/wiki"
	"github.com/gontzalson/provider-azuredevops/config/workitem"
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
		agent.Configure,
		branchpolicy.Configure,
		build.Configure,
		check.Configure,
		environment.Configure,
		feed.Configure,
		git.Configure,
		group.Configure,
		other.Configure,
		permissions.Configure,
		project.Configure,
		repositorypolicy.Configure,
		serviceendpoint.Configure,
		team.Configure,
		variablegroup.Configure,
		wiki.Configure,
		workitem.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
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
		agent.Configure,
		branchpolicy.Configure,
		build.Configure,
		check.Configure,
		environment.Configure,
		feed.Configure,
		git.Configure,
		group.Configure,
		other.Configure,
		permissions.Configure,
		project.Configure,
		repositorypolicy.Configure,
		serviceendpoint.Configure,
		team.Configure,
		variablegroup.Configure,
		wiki.Configure,
		workitem.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
