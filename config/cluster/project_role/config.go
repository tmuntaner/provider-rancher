package project_role

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "projectRole"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("rancher2_project_role_template_binding", func(r *config.Resource) {
		r.ShortGroup = shortGroup

		r.MarkAsRequired(
			"project_id",
			"role_template_id",
			"name",
		)

		r.References["project_id"] = config.Reference{
			TerraformName: "rancher2_project",
		}
	})
}
