package ilm

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("minio_ilm_policy", func(r *config.Resource) {
		r.References["bucket"] = config.Reference{
			TerraformName: "minio_s3_bucket",
		}
	})
	p.AddResourceConfigurator("minio_ilm_tier", func(r *config.Resource) {

	})
}
