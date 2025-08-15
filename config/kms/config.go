package kms

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("minio_kms_key", func(r *config.Resource) {
	})
}
