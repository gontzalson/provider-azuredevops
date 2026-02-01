package permissions

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuredevops_area_permissions", func(r *config.Resource) {
		r.ShortGroup = "permissions"
		r.ExternalName = config.IdentifierFromProvider
	})
	
	p.AddResourceConfigurator("azuredevops_iteration_permissions", func(r *config.Resource) {
		r.ShortGroup = "permissions"
		r.ExternalName = config.IdentifierFromProvider
	})
	
	p.AddResourceConfigurator("azuredevops_library_permissions", func(r *config.Resource) {
		r.ShortGroup = "permissions"
		r.ExternalName = config.IdentifierFromProvider
	})
	
	p.AddResourceConfigurator("azuredevops_project_permissions", func(r *config.Resource) {
		r.ShortGroup = "permissions"
		r.ExternalName = config.IdentifierFromProvider
	})
	
	p.AddResourceConfigurator("azuredevops_servicehook_permissions", func(r *config.Resource) {
		r.ShortGroup = "permissions"
		r.ExternalName = config.IdentifierFromProvider
	})
	
	p.AddResourceConfigurator("azuredevops_tagging_permissions", func(r *config.Resource) {
		r.ShortGroup = "permissions"
		r.ExternalName = config.IdentifierFromProvider
	})
}
