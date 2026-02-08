package project

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "project"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("rancher2_project", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Project"

		r.MarkAsRequired(
			"cluster_id",
			"name",
		)

		r.References["cluster_id"] = config.Reference{
			TerraformName: "rancher2_cluster",
		}
	})

	p.AddResourceConfigurator("rancher2_project_role_template_binding", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ProjectRoleTemplateBinding"

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
