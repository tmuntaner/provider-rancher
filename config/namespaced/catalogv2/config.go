package catalogv2

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "catalogV2"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("rancher2_catalog_v2", func(r *config.Resource) {
		r.ShortGroup = shortGroup

		r.MarkAsRequired(
			"name",
			"cluster_id",
			"project_id",
			"target_namespace",
			"template_name",
		)

		r.References["cluster_id"] = config.Reference{
			TerraformName: "rancher2_cluster",
		}

		r.References["project_id"] = config.Reference{
			TerraformName: "rancher2_project",
		}

		r.References["target_namespace"] = config.Reference{
			TerraformName: "rancher2_namespace",
		}
	})
}
