package other

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuredevops_pipeline_authorization", func(r *config.Resource) {
		r.ShortGroup = "other"
		r.ExternalName = config.IdentifierFromProvider
	})
	
	p.AddResourceConfigurator("azuredevops_resource_authorization", func(r *config.Resource) {
		r.ShortGroup = "other"
		r.ExternalName = config.IdentifierFromProvider
	})
	
	p.AddResourceConfigurator("azuredevops_securityrole_assignment", func(r *config.Resource) {
		r.ShortGroup = "other"
		r.ExternalName = config.IdentifierFromProvider
	})
	
	p.AddResourceConfigurator("azuredevops_servicehook_storage_queue_pipelines", func(r *config.Resource) {
		r.ShortGroup = "other"
		r.ExternalName = config.IdentifierFromProvider
	})
	
	p.AddResourceConfigurator("azuredevops_user_entitlement", func(r *config.Resource) {
		r.ShortGroup = "other"
		r.ExternalName = config.IdentifierFromProvider
	})
}
