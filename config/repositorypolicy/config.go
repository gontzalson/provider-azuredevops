package repositorypolicy

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuredevops_repository_policy_author_email_pattern", func(r *config.Resource) {
		r.ShortGroup = "repositorypolicy"
		r.ExternalName = config.IdentifierFromProvider
	})
	
	p.AddResourceConfigurator("azuredevops_repository_policy_case_enforcement", func(r *config.Resource) {
		r.ShortGroup = "repositorypolicy"
		r.ExternalName = config.IdentifierFromProvider
	})
	
	p.AddResourceConfigurator("azuredevops_repository_policy_check_credentials", func(r *config.Resource) {
		r.ShortGroup = "repositorypolicy"
		r.ExternalName = config.IdentifierFromProvider
	})
	
	p.AddResourceConfigurator("azuredevops_repository_policy_file_path_pattern", func(r *config.Resource) {
		r.ShortGroup = "repositorypolicy"
		r.ExternalName = config.IdentifierFromProvider
	})
	
	p.AddResourceConfigurator("azuredevops_repository_policy_max_file_size", func(r *config.Resource) {
		r.ShortGroup = "repositorypolicy"
		r.ExternalName = config.IdentifierFromProvider
	})
	
	p.AddResourceConfigurator("azuredevops_repository_policy_max_path_length", func(r *config.Resource) {
		r.ShortGroup = "repositorypolicy"
		r.ExternalName = config.IdentifierFromProvider
	})
	
	p.AddResourceConfigurator("azuredevops_repository_policy_reserved_names", func(r *config.Resource) {
		r.ShortGroup = "repositorypolicy"
		r.ExternalName = config.IdentifierFromProvider
	})
}
