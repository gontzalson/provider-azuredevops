package check

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuredevops_check_approval", func(r *config.Resource) {
		r.ShortGroup = "check"
		r.ExternalName = config.IdentifierFromProvider
	})
	
	p.AddResourceConfigurator("azuredevops_check_branch_control", func(r *config.Resource) {
		r.ShortGroup = "check"
		r.ExternalName = config.IdentifierFromProvider
	})
	
	p.AddResourceConfigurator("azuredevops_check_business_hours", func(r *config.Resource) {
		r.ShortGroup = "check"
		r.ExternalName = config.IdentifierFromProvider
	})
	
	p.AddResourceConfigurator("azuredevops_check_exclusive_lock", func(r *config.Resource) {
		r.ShortGroup = "check"
		r.ExternalName = config.IdentifierFromProvider
	})
	
	p.AddResourceConfigurator("azuredevops_check_required_template", func(r *config.Resource) {
		r.ShortGroup = "check"
		r.ExternalName = config.IdentifierFromProvider
	})
}
