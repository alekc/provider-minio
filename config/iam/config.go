package iam

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("minio_iam_user", func(r *config.Resource) {
	})
	p.AddResourceConfigurator("minio_iam_user_policy_attachment", func(r *config.Resource) {
		r.Kind = "PolicyAttachment"
		r.References["user_name"] = config.Reference{
			TerraformName: "minio_iam_user",
		}
		r.References["policy_name"] = config.Reference{
			TerraformName: "minio_iam_policy",
		}
	})
	p.AddResourceConfigurator("minio_accesskey", func(r *config.Resource) {
		r.ShortGroup = "iam"
		r.References["user"] = config.Reference{
			TerraformName: "minio_iam_user",
		}
	})
	p.AddResourceConfigurator("minio_iam_group", func(r *config.Resource) {
	})
	p.AddResourceConfigurator("minio_iam_group_membership", func(r *config.Resource) {
		r.References["group"] = config.Reference{
			TerraformName: "minio_iam_group",
		}
		r.References["users"] = config.Reference{
			TerraformName: "minio_iam_user",
		}
	})
	p.AddResourceConfigurator("minio_iam_group_policy", func(r *config.Resource) {
		r.References["group"] = config.Reference{
			TerraformName: "minio_iam_group",
		}
	})
}
