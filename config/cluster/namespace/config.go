package namespace

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "namespace"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("rancher2_namespace", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Namespace"

		r.MarkAsRequired(
			"name",
			"project_id",
		)

		r.References["project_id"] = config.Reference{
			TerraformName: "rancher2_project",
		}
	})
}
