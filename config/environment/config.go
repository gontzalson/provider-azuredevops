package environment

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuredevops_environment", func(r *config.Resource) {
		r.ShortGroup = "environment"
		r.ExternalName = config.IdentifierFromProvider
	})
	
	p.AddResourceConfigurator("azuredevops_environment_resource_kubernetes", func(r *config.Resource) {
		r.ShortGroup = "environment"
		r.ExternalName = config.IdentifierFromProvider
	})
}
