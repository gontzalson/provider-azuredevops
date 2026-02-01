// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	pool "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/agent/pool"
	queue "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/agent/queue"
	policyautoreviewers "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/branchpolicy/policyautoreviewers"
	policybuildvalidation "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/branchpolicy/policybuildvalidation"
	policycommentresolution "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/branchpolicy/policycommentresolution"
	policymergetypes "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/branchpolicy/policymergetypes"
	policyminreviewers "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/branchpolicy/policyminreviewers"
	policystatuscheck "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/branchpolicy/policystatuscheck"
	policyworkitemlinking "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/branchpolicy/policyworkitemlinking"
	definition "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/build/definition"
	definitionpermissions "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/build/definitionpermissions"
	folder "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/build/folder"
	folderpermissions "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/build/folderpermissions"
	approval "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/check/approval"
	branchcontrol "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/check/branchcontrol"
	businesshours "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/check/businesshours"
	exclusivelock "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/check/exclusivelock"
	requiredtemplate "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/check/requiredtemplate"
	environment "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/environment/environment"
	resourcekubernetes "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/environment/resourcekubernetes"
	feed "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/feed/feed"
	permission "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/feed/permission"
	permissions "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/git/permissions"
	repository "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/git/repository"
	repositorybranch "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/git/repositorybranch"
	repositoryfile "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/git/repositoryfile"
	entitlement "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/group/entitlement"
	group "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/group/group"
	membership "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/group/membership"
	assignment "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/other/assignment"
	authorization "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/other/authorization"
	entitlementother "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/other/entitlement"
	storagequeuepipelines "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/other/storagequeuepipelines"
	permissionspermissions "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/permissions/permissions"
	features "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/project/features"
	pipelinesettings "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/project/pipelinesettings"
	project "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/project/project"
	providerconfig "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/providerconfig"
	policyauthoremailpattern "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/repositorypolicy/policyauthoremailpattern"
	policycaseenforcement "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/repositorypolicy/policycaseenforcement"
	policycheckcredentials "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/repositorypolicy/policycheckcredentials"
	policyfilepathpattern "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/repositorypolicy/policyfilepathpattern"
	policymaxfilesize "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/repositorypolicy/policymaxfilesize"
	policymaxpathlength "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/repositorypolicy/policymaxpathlength"
	policyreservednames "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/repositorypolicy/policyreservednames"
	argocd "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/serviceendpoint/argocd"
	artifactory "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/serviceendpoint/artifactory"
	aws "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/serviceendpoint/aws"
	azurecr "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/serviceendpoint/azurecr"
	azuredevops "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/serviceendpoint/azuredevops"
	azurerm "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/serviceendpoint/azurerm"
	bitbucket "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/serviceendpoint/bitbucket"
	dockerregistry "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/serviceendpoint/dockerregistry"
	externaltfs "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/serviceendpoint/externaltfs"
	gcpterraform "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/serviceendpoint/gcpterraform"
	generic "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/serviceendpoint/generic"
	genericgit "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/serviceendpoint/genericgit"
	github "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/serviceendpoint/github"
	githubenterprise "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/serviceendpoint/githubenterprise"
	incomingwebhook "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/serviceendpoint/incomingwebhook"
	jenkins "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/serviceendpoint/jenkins"
	jfrogartifactoryv2 "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/serviceendpoint/jfrogartifactoryv2"
	jfrogdistributionv2 "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/serviceendpoint/jfrogdistributionv2"
	jfrogplatformv2 "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/serviceendpoint/jfrogplatformv2"
	jfrogxrayv2 "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/serviceendpoint/jfrogxrayv2"
	kubernetes "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/serviceendpoint/kubernetes"
	maven "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/serviceendpoint/maven"
	nexus "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/serviceendpoint/nexus"
	npm "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/serviceendpoint/npm"
	nuget "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/serviceendpoint/nuget"
	octopusdeploy "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/serviceendpoint/octopusdeploy"
	permissionsserviceendpoint "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/serviceendpoint/permissions"
	runpipeline "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/serviceendpoint/runpipeline"
	servicefabric "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/serviceendpoint/servicefabric"
	sonarcloud "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/serviceendpoint/sonarcloud"
	sonarqube "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/serviceendpoint/sonarqube"
	ssh "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/serviceendpoint/ssh"
	administrators "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/team/administrators"
	members "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/team/members"
	team "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/team/team"
	groupvariablegroup "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/variablegroup/group"
	grouppermissions "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/variablegroup/grouppermissions"
	page "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/wiki/page"
	wiki "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/wiki/wiki"
	permissionsworkitem "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/workitem/permissions"
	workitem "github.com/gontzalson/provider-azuredevops/internal/controller/cluster/workitem/workitem"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		pool.Setup,
		pool.Setup,
		queue.Setup,
		policyautoreviewers.Setup,
		policybuildvalidation.Setup,
		policycommentresolution.Setup,
		policymergetypes.Setup,
		policyminreviewers.Setup,
		policystatuscheck.Setup,
		policyworkitemlinking.Setup,
		definition.Setup,
		definitionpermissions.Setup,
		folder.Setup,
		folderpermissions.Setup,
		approval.Setup,
		branchcontrol.Setup,
		businesshours.Setup,
		exclusivelock.Setup,
		requiredtemplate.Setup,
		environment.Setup,
		resourcekubernetes.Setup,
		feed.Setup,
		permission.Setup,
		permissions.Setup,
		repository.Setup,
		repositorybranch.Setup,
		repositoryfile.Setup,
		entitlement.Setup,
		group.Setup,
		membership.Setup,
		assignment.Setup,
		authorization.Setup,
		authorization.Setup,
		entitlementother.Setup,
		storagequeuepipelines.Setup,
		permissionspermissions.Setup,
		permissionspermissions.Setup,
		permissionspermissions.Setup,
		permissionspermissions.Setup,
		permissionspermissions.Setup,
		permissionspermissions.Setup,
		features.Setup,
		pipelinesettings.Setup,
		project.Setup,
		providerconfig.Setup,
		policyauthoremailpattern.Setup,
		policycaseenforcement.Setup,
		policycheckcredentials.Setup,
		policyfilepathpattern.Setup,
		policymaxfilesize.Setup,
		policymaxpathlength.Setup,
		policyreservednames.Setup,
		argocd.Setup,
		artifactory.Setup,
		aws.Setup,
		azurecr.Setup,
		azuredevops.Setup,
		azurerm.Setup,
		bitbucket.Setup,
		dockerregistry.Setup,
		externaltfs.Setup,
		gcpterraform.Setup,
		generic.Setup,
		genericgit.Setup,
		github.Setup,
		githubenterprise.Setup,
		incomingwebhook.Setup,
		jenkins.Setup,
		jfrogartifactoryv2.Setup,
		jfrogdistributionv2.Setup,
		jfrogplatformv2.Setup,
		jfrogxrayv2.Setup,
		kubernetes.Setup,
		maven.Setup,
		nexus.Setup,
		npm.Setup,
		nuget.Setup,
		octopusdeploy.Setup,
		permissionsserviceendpoint.Setup,
		runpipeline.Setup,
		servicefabric.Setup,
		sonarcloud.Setup,
		sonarqube.Setup,
		ssh.Setup,
		administrators.Setup,
		members.Setup,
		team.Setup,
		groupvariablegroup.Setup,
		grouppermissions.Setup,
		page.Setup,
		wiki.Setup,
		permissionsworkitem.Setup,
		workitem.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		pool.SetupGated,
		pool.SetupGated,
		queue.SetupGated,
		policyautoreviewers.SetupGated,
		policybuildvalidation.SetupGated,
		policycommentresolution.SetupGated,
		policymergetypes.SetupGated,
		policyminreviewers.SetupGated,
		policystatuscheck.SetupGated,
		policyworkitemlinking.SetupGated,
		definition.SetupGated,
		definitionpermissions.SetupGated,
		folder.SetupGated,
		folderpermissions.SetupGated,
		approval.SetupGated,
		branchcontrol.SetupGated,
		businesshours.SetupGated,
		exclusivelock.SetupGated,
		requiredtemplate.SetupGated,
		environment.SetupGated,
		resourcekubernetes.SetupGated,
		feed.SetupGated,
		permission.SetupGated,
		permissions.SetupGated,
		repository.SetupGated,
		repositorybranch.SetupGated,
		repositoryfile.SetupGated,
		entitlement.SetupGated,
		group.SetupGated,
		membership.SetupGated,
		assignment.SetupGated,
		authorization.SetupGated,
		authorization.SetupGated,
		entitlementother.SetupGated,
		storagequeuepipelines.SetupGated,
		permissionspermissions.SetupGated,
		permissionspermissions.SetupGated,
		permissionspermissions.SetupGated,
		permissionspermissions.SetupGated,
		permissionspermissions.SetupGated,
		permissionspermissions.SetupGated,
		features.SetupGated,
		pipelinesettings.SetupGated,
		project.SetupGated,
		providerconfig.SetupGated,
		policyauthoremailpattern.SetupGated,
		policycaseenforcement.SetupGated,
		policycheckcredentials.SetupGated,
		policyfilepathpattern.SetupGated,
		policymaxfilesize.SetupGated,
		policymaxpathlength.SetupGated,
		policyreservednames.SetupGated,
		argocd.SetupGated,
		artifactory.SetupGated,
		aws.SetupGated,
		azurecr.SetupGated,
		azuredevops.SetupGated,
		azurerm.SetupGated,
		bitbucket.SetupGated,
		dockerregistry.SetupGated,
		externaltfs.SetupGated,
		gcpterraform.SetupGated,
		generic.SetupGated,
		genericgit.SetupGated,
		github.SetupGated,
		githubenterprise.SetupGated,
		incomingwebhook.SetupGated,
		jenkins.SetupGated,
		jfrogartifactoryv2.SetupGated,
		jfrogdistributionv2.SetupGated,
		jfrogplatformv2.SetupGated,
		jfrogxrayv2.SetupGated,
		kubernetes.SetupGated,
		maven.SetupGated,
		nexus.SetupGated,
		npm.SetupGated,
		nuget.SetupGated,
		octopusdeploy.SetupGated,
		permissionsserviceendpoint.SetupGated,
		runpipeline.SetupGated,
		servicefabric.SetupGated,
		sonarcloud.SetupGated,
		sonarqube.SetupGated,
		ssh.SetupGated,
		administrators.SetupGated,
		members.SetupGated,
		team.SetupGated,
		groupvariablegroup.SetupGated,
		grouppermissions.SetupGated,
		page.SetupGated,
		wiki.SetupGated,
		permissionsworkitem.SetupGated,
		workitem.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
