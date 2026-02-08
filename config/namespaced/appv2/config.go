package appv2

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "appV2"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("rancher2_app_v2", func(r *config.Resource) {
		r.ShortGroup = shortGroup

		r.MarkAsRequired(
			"cluster_id",
			"name",
			"namespace",
			"repo_name",
			"chart_name",
		)
	})
}
