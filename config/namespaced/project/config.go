package project

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "project"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("rancher2_project", func(r *config.Resource) {
		r.ShortGroup = shortGroup

		/*
			r.MarkAsRequired(
				"cluster_id",
				"name",
			)

			r.References["cluster_id"] = config.Reference{
				TerraformName: "rancher2_cluster_v2",
			}
		*/
	})
}
