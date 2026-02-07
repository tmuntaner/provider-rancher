package cluster

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "cluster"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("rancher2_cluster", func(r *config.Resource) {
		r.ShortGroup = shortGroup

		r.MarkAsRequired(
			"name",
		)
	})
}
