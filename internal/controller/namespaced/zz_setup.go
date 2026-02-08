// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	appv2 "github.com/tmuntaner/provider-rancher/internal/controller/namespaced/app/appv2"
	catalogv2 "github.com/tmuntaner/provider-rancher/internal/controller/namespaced/catalog/catalogv2"
	cluster "github.com/tmuntaner/provider-rancher/internal/controller/namespaced/cluster/cluster"
	namespace "github.com/tmuntaner/provider-rancher/internal/controller/namespaced/namespace/namespace"
	project "github.com/tmuntaner/provider-rancher/internal/controller/namespaced/project/project"
	projectroletemplatebinding "github.com/tmuntaner/provider-rancher/internal/controller/namespaced/project/projectroletemplatebinding"
	providerconfig "github.com/tmuntaner/provider-rancher/internal/controller/namespaced/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		appv2.Setup,
		catalogv2.Setup,
		cluster.Setup,
		namespace.Setup,
		project.Setup,
		projectroletemplatebinding.Setup,
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
		appv2.SetupGated,
		catalogv2.SetupGated,
		cluster.SetupGated,
		namespace.SetupGated,
		project.SetupGated,
		projectroletemplatebinding.SetupGated,
		providerconfig.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
