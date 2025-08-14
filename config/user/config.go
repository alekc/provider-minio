package user

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("minio_iam_user", func(r *config.Resource) {
		r.ShortGroup = "user"
	})
	p.AddResourceConfigurator("minio_iam_user_policy_attachment", func(r *config.Resource) {
		r.ShortGroup = "user"
		r.References["user_name"] = config.Reference{
			TerraformName: "minio_iam_user",
		}
		r.References["policy_name"] = config.Reference{
			TerraformName: "minio_iam_policy",
		}
	})
}
