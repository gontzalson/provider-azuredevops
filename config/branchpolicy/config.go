package branchpolicy

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuredevops_branch_policy_auto_reviewers", func(r *config.Resource) {
		r.ShortGroup = "branchpolicy"
		r.ExternalName = config.IdentifierFromProvider
	})
	
	p.AddResourceConfigurator("azuredevops_branch_policy_build_validation", func(r *config.Resource) {
		r.ShortGroup = "branchpolicy"
		r.ExternalName = config.IdentifierFromProvider
	})
	
	p.AddResourceConfigurator("azuredevops_branch_policy_comment_resolution", func(r *config.Resource) {
		r.ShortGroup = "branchpolicy"
		r.ExternalName = config.IdentifierFromProvider
	})
	
	p.AddResourceConfigurator("azuredevops_branch_policy_merge_types", func(r *config.Resource) {
		r.ShortGroup = "branchpolicy"
		r.ExternalName = config.IdentifierFromProvider
	})
	
	p.AddResourceConfigurator("azuredevops_branch_policy_min_reviewers", func(r *config.Resource) {
		r.ShortGroup = "branchpolicy"
		r.ExternalName = config.IdentifierFromProvider
	})
	
	p.AddResourceConfigurator("azuredevops_branch_policy_status_check", func(r *config.Resource) {
		r.ShortGroup = "branchpolicy"
		r.ExternalName = config.IdentifierFromProvider
	})
	
	p.AddResourceConfigurator("azuredevops_branch_policy_work_item_linking", func(r *config.Resource) {
		r.ShortGroup = "branchpolicy"
		r.ExternalName = config.IdentifierFromProvider
	})
}
