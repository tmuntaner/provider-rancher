package app

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "app"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("rancher2_app", func(r *config.Resource) {
		r.ShortGroup = shortGroup

		r.MarkAsRequired(
			"catalog_name",
			"name",
			"project_id",
			"target_namespace",
			"template_name",
		)

		r.References["catalog_name"] = config.Reference{
			TerraformName: "rancher2_catalog_v2",
		}
	})
}
