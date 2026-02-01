package variablegroup

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuredevops_variable_group", func(r *config.Resource) {
		r.ShortGroup = "variablegroup"
		r.ExternalName = config.IdentifierFromProvider
	})
	
	p.AddResourceConfigurator("azuredevops_variable_group_permissions", func(r *config.Resource) {
		r.ShortGroup = "variablegroup"
		r.ExternalName = config.IdentifierFromProvider
	})
}
