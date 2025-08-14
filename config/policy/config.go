package policy

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("minio_iam_policy", func(r *config.Resource) {
		r.ShortGroup = "policy"
	})
}
