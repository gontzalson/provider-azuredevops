#!/bin/bash
# Generate config files for all Azure DevOps resources

set -e

cd "$(dirname "$0")"

# Create directories for each resource group
mkdir -p agent build git branchpolicy repositorypolicy serviceendpoint check \
         group team feed environment variablegroup wiki workitem permissions other

# Agent resources
cat > agent/config.go <<'EOF'
package agent

import "github.com/crossplane/upjet/pkg/config"

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
EOF

# Build resources
cat > build/config.go <<'EOF'
package build

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuredevops_build_definition", func(r *config.Resource) {
		r.ShortGroup = "build"
		r.ExternalName = config.IdentifierFromProvider
	})
	
	p.AddResourceConfigurator("azuredevops_build_definition_permissions", func(r *config.Resource) {
		r.ShortGroup = "build"
		r.ExternalName = config.IdentifierFromProvider
	})
	
	p.AddResourceConfigurator("azuredevops_build_folder", func(r *config.Resource) {
		r.ShortGroup = "build"
		r.ExternalName = config.IdentifierFromProvider
	})
	
	p.AddResourceConfigurator("azuredevops_build_folder_permissions", func(r *config.Resource) {
		r.ShortGroup = "build"
		r.ExternalName = config.IdentifierFromProvider
	})
}
EOF

# Git resources
cat > git/config.go <<'EOF'
package git

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuredevops_git_repository", func(r *config.Resource) {
		r.ShortGroup = "git"
		r.ExternalName = config.IdentifierFromProvider
	})
	
	p.AddResourceConfigurator("azuredevops_git_repository_branch", func(r *config.Resource) {
		r.ShortGroup = "git"
		r.ExternalName = config.IdentifierFromProvider
	})
	
	p.AddResourceConfigurator("azuredevops_git_repository_file", func(r *config.Resource) {
		r.ShortGroup = "git"
		r.ExternalName = config.IdentifierFromProvider
	})
	
	p.AddResourceConfigurator("azuredevops_git_permissions", func(r *config.Resource) {
		r.ShortGroup = "git"
		r.ExternalName = config.IdentifierFromProvider
	})
}
EOF

# Branch policy resources
cat > branchpolicy/config.go <<'EOF'
package branchpolicy

import "github.com/crossplane/upjet/pkg/config"

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
EOF

# Repository policy resources
cat > repositorypolicy/config.go <<'EOF'
package repositorypolicy

import "github.com/crossplane/upjet/pkg/config"

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
EOF

# Service endpoint resources (28 resources!)
cat > serviceendpoint/config.go <<'EOF'
package serviceendpoint

import "github.com/crossplane/upjet/pkg/config"

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
EOF

# Check resources
cat > check/config.go <<'EOF'
package check

import "github.com/crossplane/upjet/pkg/config"

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
EOF

# Group resources
cat > group/config.go <<'EOF'
package group

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuredevops_group", func(r *config.Resource) {
		r.ShortGroup = "group"
		r.ExternalName = config.IdentifierFromProvider
	})
	
	p.AddResourceConfigurator("azuredevops_group_entitlement", func(r *config.Resource) {
		r.ShortGroup = "group"
		r.ExternalName = config.IdentifierFromProvider
	})
	
	p.AddResourceConfigurator("azuredevops_group_membership", func(r *config.Resource) {
		r.ShortGroup = "group"
		r.ExternalName = config.IdentifierFromProvider
	})
}
EOF

# Team resources
cat > team/config.go <<'EOF'
package team

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuredevops_team", func(r *config.Resource) {
		r.ShortGroup = "team"
		r.ExternalName = config.IdentifierFromProvider
	})
	
	p.AddResourceConfigurator("azuredevops_team_administrators", func(r *config.Resource) {
		r.ShortGroup = "team"
		r.ExternalName = config.IdentifierFromProvider
	})
	
	p.AddResourceConfigurator("azuredevops_team_members", func(r *config.Resource) {
		r.ShortGroup = "team"
		r.ExternalName = config.IdentifierFromProvider
	})
}
EOF

# Feed resources
cat > feed/config.go <<'EOF'
package feed

import "github.com/crossplane/upjet/pkg/config"

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
EOF

# Environment resources
cat > environment/config.go <<'EOF'
package environment

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuredevops_environment", func(r *config.Resource) {
		r.ShortGroup = "environment"
		r.ExternalName = config.IdentifierFromProvider
	})
	
	p.AddResourceConfigurator("azuredevops_environment_resource_kubernetes", func(r *config.Resource) {
		r.ShortGroup = "environment"
		r.ExternalName = config.IdentifierFromProvider
	})
}
EOF

# Variable group resources
cat > variablegroup/config.go <<'EOF'
package variablegroup

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuredevops_variable_group", func(r *config.Resource) {
		r.ShortGroup = "variablegroup"
		r.ExternalName = config.IdentifierFromProvider
	})
	
	p.AddResourceConfigurator("azuredevops_variable_group_permissions", func(r *config.Resource) {
		r.ShortGroup = "variablegroup"
		r.ExternalName = config.IdentifierFromProvider
	})
}
EOF

# Wiki resources
cat > wiki/config.go <<'EOF'
package wiki

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuredevops_wiki", func(r *config.Resource) {
		r.ShortGroup = "wiki"
		r.ExternalName = config.IdentifierFromProvider
	})
	
	p.AddResourceConfigurator("azuredevops_wiki_page", func(r *config.Resource) {
		r.ShortGroup = "wiki"
		r.ExternalName = config.IdentifierFromProvider
	})
}
EOF

# Workitem resources
cat > workitem/config.go <<'EOF'
package workitem

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuredevops_workitem", func(r *config.Resource) {
		r.ShortGroup = "workitem"
		r.ExternalName = config.IdentifierFromProvider
	})
	
	p.AddResourceConfigurator("azuredevops_workitemquery_permissions", func(r *config.Resource) {
		r.ShortGroup = "workitem"
		r.ExternalName = config.IdentifierFromProvider
	})
}
EOF

# Permissions resources
cat > permissions/config.go <<'EOF'
package permissions

import "github.com/crossplane/upjet/pkg/config"

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
EOF

# Other resources
cat > other/config.go <<'EOF'
package other

import "github.com/crossplane/upjet/pkg/config"

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
EOF

echo "✅ All config files generated successfully!"
