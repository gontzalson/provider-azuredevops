// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	features "github.com/gontzalson/provider-azuredevops/internal/controller/namespaced/project/features"
	pipelinesettings "github.com/gontzalson/provider-azuredevops/internal/controller/namespaced/project/pipelinesettings"
	project "github.com/gontzalson/provider-azuredevops/internal/controller/namespaced/project/project"
	providerconfig "github.com/gontzalson/provider-azuredevops/internal/controller/namespaced/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		features.Setup,
		pipelinesettings.Setup,
		project.Setup,
		providerconfig.Setup,
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
		features.SetupGated,
		pipelinesettings.SetupGated,
		project.SetupGated,
		providerconfig.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
