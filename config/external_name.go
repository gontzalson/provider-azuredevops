package config

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Agent resources
	"azuredevops_agent_pool":   config.IdentifierFromProvider,
	"azuredevops_agent_queue":  config.IdentifierFromProvider,
	"azuredevops_elastic_pool": config.IdentifierFromProvider,

	// Build resources
	"azuredevops_build_definition":             config.IdentifierFromProvider,
	"azuredevops_build_definition_permissions": config.IdentifierFromProvider,
	"azuredevops_build_folder":                 config.IdentifierFromProvider,
	"azuredevops_build_folder_permissions":     config.IdentifierFromProvider,

	// Git resources
	"azuredevops_git_repository":        config.IdentifierFromProvider,
	"azuredevops_git_repository_branch": config.IdentifierFromProvider,
	"azuredevops_git_repository_file":   config.IdentifierFromProvider,
	"azuredevops_git_permissions":       config.IdentifierFromProvider,

	// Branch policy resources
	"azuredevops_branch_policy_auto_reviewers":      config.IdentifierFromProvider,
	"azuredevops_branch_policy_build_validation":    config.IdentifierFromProvider,
	"azuredevops_branch_policy_comment_resolution":  config.IdentifierFromProvider,
	"azuredevops_branch_policy_merge_types":         config.IdentifierFromProvider,
	"azuredevops_branch_policy_min_reviewers":       config.IdentifierFromProvider,
	"azuredevops_branch_policy_status_check":        config.IdentifierFromProvider,
	"azuredevops_branch_policy_work_item_linking":   config.IdentifierFromProvider,

	// Repository policy resources
	"azuredevops_repository_policy_author_email_pattern": config.IdentifierFromProvider,
	"azuredevops_repository_policy_case_enforcement":     config.IdentifierFromProvider,
	"azuredevops_repository_policy_check_credentials":    config.IdentifierFromProvider,
	"azuredevops_repository_policy_file_path_pattern":    config.IdentifierFromProvider,
	"azuredevops_repository_policy_max_file_size":        config.IdentifierFromProvider,
	"azuredevops_repository_policy_max_path_length":      config.IdentifierFromProvider,
	"azuredevops_repository_policy_reserved_names":       config.IdentifierFromProvider,

	// Service endpoint resources (28 resources)
	"azuredevops_serviceendpoint_argocd":               config.IdentifierFromProvider,
	"azuredevops_serviceendpoint_artifactory":          config.IdentifierFromProvider,
	"azuredevops_serviceendpoint_aws":                  config.IdentifierFromProvider,
	"azuredevops_serviceendpoint_azurecr":              config.IdentifierFromProvider,
	"azuredevops_serviceendpoint_azuredevops":          config.IdentifierFromProvider,
	"azuredevops_serviceendpoint_azurerm":              config.IdentifierFromProvider,
	"azuredevops_serviceendpoint_bitbucket":            config.IdentifierFromProvider,
	"azuredevops_serviceendpoint_dockerregistry":       config.IdentifierFromProvider,
	"azuredevops_serviceendpoint_externaltfs":          config.IdentifierFromProvider,
	"azuredevops_serviceendpoint_gcp_terraform":        config.IdentifierFromProvider,
	"azuredevops_serviceendpoint_generic":              config.IdentifierFromProvider,
	"azuredevops_serviceendpoint_generic_git":          config.IdentifierFromProvider,
	"azuredevops_serviceendpoint_github":               config.IdentifierFromProvider,
	"azuredevops_serviceendpoint_github_enterprise":    config.IdentifierFromProvider,
	"azuredevops_serviceendpoint_incomingwebhook":      config.IdentifierFromProvider,
	"azuredevops_serviceendpoint_jenkins":              config.IdentifierFromProvider,
	"azuredevops_serviceendpoint_jfrog_artifactory_v2": config.IdentifierFromProvider,
	"azuredevops_serviceendpoint_jfrog_distribution_v2": config.IdentifierFromProvider,
	"azuredevops_serviceendpoint_jfrog_platform_v2":    config.IdentifierFromProvider,
	"azuredevops_serviceendpoint_jfrog_xray_v2":        config.IdentifierFromProvider,
	"azuredevops_serviceendpoint_kubernetes":           config.IdentifierFromProvider,
	"azuredevops_serviceendpoint_maven":                config.IdentifierFromProvider,
	"azuredevops_serviceendpoint_nexus":                config.IdentifierFromProvider,
	"azuredevops_serviceendpoint_npm":                  config.IdentifierFromProvider,
	"azuredevops_serviceendpoint_nuget":                config.IdentifierFromProvider,
	"azuredevops_serviceendpoint_octopusdeploy":        config.IdentifierFromProvider,
	"azuredevops_serviceendpoint_permissions":          config.IdentifierFromProvider,
	"azuredevops_serviceendpoint_runpipeline":          config.IdentifierFromProvider,
	"azuredevops_serviceendpoint_servicefabric":        config.IdentifierFromProvider,
	"azuredevops_serviceendpoint_sonarcloud":           config.IdentifierFromProvider,
	"azuredevops_serviceendpoint_sonarqube":            config.IdentifierFromProvider,
	"azuredevops_serviceendpoint_ssh":                  config.IdentifierFromProvider,

	// Check resources
	"azuredevops_check_approval":          config.IdentifierFromProvider,
	"azuredevops_check_branch_control":    config.IdentifierFromProvider,
	"azuredevops_check_business_hours":    config.IdentifierFromProvider,
	"azuredevops_check_exclusive_lock":    config.IdentifierFromProvider,
	"azuredevops_check_required_template": config.IdentifierFromProvider,

	// Group resources
	"azuredevops_group":            config.IdentifierFromProvider,
	"azuredevops_group_entitlement": config.IdentifierFromProvider,
	"azuredevops_group_membership":  config.IdentifierFromProvider,

	// Team resources
	"azuredevops_team":               config.IdentifierFromProvider,
	"azuredevops_team_administrators": config.IdentifierFromProvider,
	"azuredevops_team_members":        config.IdentifierFromProvider,

	// Feed resources
	"azuredevops_feed":            config.IdentifierFromProvider,
	"azuredevops_feed_permission": config.IdentifierFromProvider,

	// Environment resources
	"azuredevops_environment":                     config.IdentifierFromProvider,
	"azuredevops_environment_resource_kubernetes": config.IdentifierFromProvider,

	// Variable group resources
	"azuredevops_variable_group":             config.IdentifierFromProvider,
	"azuredevops_variable_group_permissions": config.IdentifierFromProvider,

	// Wiki resources
	"azuredevops_wiki":      config.IdentifierFromProvider,
	"azuredevops_wiki_page": config.IdentifierFromProvider,

	// Workitem resources
	"azuredevops_workitem":                   config.IdentifierFromProvider,
	"azuredevops_workitemquery_permissions":  config.IdentifierFromProvider,

	// Permission resources
	"azuredevops_area_permissions":       config.IdentifierFromProvider,
	"azuredevops_iteration_permissions":  config.IdentifierFromProvider,
	"azuredevops_library_permissions":    config.IdentifierFromProvider,
	"azuredevops_project_permissions":    config.IdentifierFromProvider,
	"azuredevops_servicehook_permissions": config.IdentifierFromProvider,
	"azuredevops_tagging_permissions":    config.IdentifierFromProvider,

	// Project resources
	"azuredevops_project":                   config.IdentifierFromProvider,
	"azuredevops_project_features":          config.IdentifierFromProvider,
	"azuredevops_project_pipeline_settings": config.IdentifierFromProvider,

	// Other resources
	"azuredevops_pipeline_authorization":             config.IdentifierFromProvider,
	"azuredevops_resource_authorization":             config.IdentifierFromProvider,
	"azuredevops_securityrole_assignment":            config.IdentifierFromProvider,
	"azuredevops_servicehook_storage_queue_pipelines": config.IdentifierFromProvider,
	"azuredevops_user_entitlement":                   config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
