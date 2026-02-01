package feed

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuredevops_feed", func(r *config.Resource) {
		r.ShortGroup = "feed"
		r.ExternalName = config.IdentifierFromProvider
	})
	
	p.AddResourceConfigurator("azuredevops_feed_permission", func(r *config.Resource) {
		r.ShortGroup = "feed"
		r.ExternalName = config.IdentifierFromProvider
	})
}
