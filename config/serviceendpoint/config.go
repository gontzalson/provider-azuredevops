package serviceendpoint

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	endpoints := []string{
		"azuredevops_serviceendpoint_argocd",
		"azuredevops_serviceendpoint_artifactory",
		"azuredevops_serviceendpoint_aws",
		"azuredevops_serviceendpoint_azurecr",
		"azuredevops_serviceendpoint_azuredevops",
		"azuredevops_serviceendpoint_azurerm",
		"azuredevops_serviceendpoint_bitbucket",
		"azuredevops_serviceendpoint_dockerregistry",
		"azuredevops_serviceendpoint_externaltfs",
		"azuredevops_serviceendpoint_gcp_terraform",
		"azuredevops_serviceendpoint_generic",
		"azuredevops_serviceendpoint_generic_git",
		"azuredevops_serviceendpoint_github",
		"azuredevops_serviceendpoint_github_enterprise",
		"azuredevops_serviceendpoint_incomingwebhook",
		"azuredevops_serviceendpoint_jenkins",
		"azuredevops_serviceendpoint_jfrog_artifactory_v2",
		"azuredevops_serviceendpoint_jfrog_distribution_v2",
		"azuredevops_serviceendpoint_jfrog_platform_v2",
		"azuredevops_serviceendpoint_jfrog_xray_v2",
		"azuredevops_serviceendpoint_kubernetes",
		"azuredevops_serviceendpoint_maven",
		"azuredevops_serviceendpoint_nexus",
		"azuredevops_serviceendpoint_npm",
		"azuredevops_serviceendpoint_nuget",
		"azuredevops_serviceendpoint_octopusdeploy",
		"azuredevops_serviceendpoint_permissions",
		"azuredevops_serviceendpoint_runpipeline",
		"azuredevops_serviceendpoint_servicefabric",
		"azuredevops_serviceendpoint_sonarcloud",
		"azuredevops_serviceendpoint_sonarqube",
		"azuredevops_serviceendpoint_ssh",
	}
	
	for _, ep := range endpoints {
		endpoint := ep
		p.AddResourceConfigurator(endpoint, func(r *config.Resource) {
			r.ShortGroup = "serviceendpoint"
			r.ExternalName = config.IdentifierFromProvider
		})
	}
}
