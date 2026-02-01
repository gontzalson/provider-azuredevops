package agent

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuredevops_agent_pool", func(r *config.Resource) {
		r.ShortGroup = "agent"
		r.ExternalName = config.IdentifierFromProvider
	})
	
	p.AddResourceConfigurator("azuredevops_agent_queue", func(r *config.Resource) {
		r.ShortGroup = "agent"
		r.ExternalName = config.IdentifierFromProvider
	})
	
	p.AddResourceConfigurator("azuredevops_elastic_pool", func(r *config.Resource) {
		r.ShortGroup = "agent"
		r.ExternalName = config.IdentifierFromProvider
	})
}
