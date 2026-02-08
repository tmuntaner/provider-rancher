package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
	appCluster "github.com/tmuntaner/provider-rancher/config/cluster/app"
	appV2Cluster "github.com/tmuntaner/provider-rancher/config/cluster/app_v2"
	catalogV2Cluster "github.com/tmuntaner/provider-rancher/config/cluster/catalog_v2"
	clusterCluster "github.com/tmuntaner/provider-rancher/config/cluster/cluster"
	namespaceCluster "github.com/tmuntaner/provider-rancher/config/cluster/namespace"
	projectCluster "github.com/tmuntaner/provider-rancher/config/cluster/project"
	projectRoleCluster "github.com/tmuntaner/provider-rancher/config/cluster/project_role"
	appNamespaced "github.com/tmuntaner/provider-rancher/config/namespaced/app"
	appV2Namespaced "github.com/tmuntaner/provider-rancher/config/namespaced/app_v2"
	catalogV2Namespaced "github.com/tmuntaner/provider-rancher/config/namespaced/catalog_v2"
	clusterNamespaced "github.com/tmuntaner/provider-rancher/config/namespaced/cluster"
	namespaceNamespaced "github.com/tmuntaner/provider-rancher/config/namespaced/namespace"
	projectNamespaced "github.com/tmuntaner/provider-rancher/config/namespaced/project"
	projectRoleNamespaced "github.com/tmuntaner/provider-rancher/config/namespaced/project_role"
)

const (
	resourcePrefix = "rancher"
	modulePath     = "github.com/tmuntaner/provider-rancher"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("rancher.openplatform.suse.com"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		appCluster.Configure,
		appV2Cluster.Configure,
		catalogV2Cluster.Configure,
		clusterCluster.Configure,
		namespaceCluster.Configure,
		projectCluster.Configure,
		projectRoleCluster.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

// GetProviderNamespaced returns the namespaced provider configuration
func GetProviderNamespaced() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("rancher.m.openplatform.suse.com"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		),
		ujconfig.WithExampleManifestConfiguration(ujconfig.ExampleManifestConfiguration{
			ManagedResourceNamespace: "crossplane-system",
		}))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		appNamespaced.Configure,
		appV2Namespaced.Configure,
		catalogV2Namespaced.Configure,
		clusterNamespaced.Configure,
		namespaceNamespaced.Configure,
		projectNamespaced.Configure,
		projectRoleNamespaced.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
