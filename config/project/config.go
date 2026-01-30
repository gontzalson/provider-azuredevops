package project

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuredevops_project", func(r *config.Resource) {
		// We use the ID as the external name
		r.ExternalName = config.IdentifierFromProvider
		
		// Set the API version
		r.Version = "v1alpha1"
		
		// Project is the root resource, no references needed
		r.ShortGroup = "project"
	})
	
	p.AddResourceConfigurator("azuredevops_project_features", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.Version = "v1alpha1"
		r.ShortGroup = "project"
		
		// Reference to project
		r.References["project_id"] = config.Reference{
			TerraformName: "azuredevops_project",
		}
	})
	
	p.AddResourceConfigurator("azuredevops_project_pipeline_settings", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.Version = "v1alpha1"
		r.ShortGroup = "project"
		
		// Reference to project
		r.References["project_id"] = config.Reference{
			TerraformName: "azuredevops_project",
		}
	})
}
