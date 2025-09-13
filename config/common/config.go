package common

import "github.com/crossplane/upjet/v2/pkg/config"

const bucketShortGroup = "bucket"

// Configure configures cluster-scoped resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	configureBucket(p)
	configureIAM(p)
	p.AddResourceConfigurator("minio_ilm_policy", func(r *config.Resource) {
		r.References["bucket"] = config.Reference{
			TerraformName: "minio_s3_bucket",
		}
	})
	p.AddResourceConfigurator("minio_ilm_tier", func(r *config.Resource) {

	})
	p.AddResourceConfigurator("minio_kms_key", func(r *config.Resource) {
	})
}

func configureBucket(p *config.Provider) {
	p.AddResourceConfigurator("minio_s3_bucket", func(r *config.Resource) {
		r.ShortGroup = bucketShortGroup
	})
	p.AddResourceConfigurator("minio_s3_bucket_notification", func(r *config.Resource) {
		r.ShortGroup = bucketShortGroup
		r.Kind = "Notification"
		r.References[bucketShortGroup] = config.Reference{
			TerraformName: "minio_s3_bucket",
		}
	})
	p.AddResourceConfigurator("minio_s3_bucket_policy", func(r *config.Resource) {
		r.ShortGroup = bucketShortGroup
		r.Kind = "Policy"
		r.References[bucketShortGroup] = config.Reference{
			TerraformName: "minio_s3_bucket",
		}
	})
	p.AddResourceConfigurator("minio_s3_bucket_replication", func(r *config.Resource) {
		r.ShortGroup = bucketShortGroup
		r.Kind = "Replication"
		r.References[bucketShortGroup] = config.Reference{
			TerraformName: "minio_s3_bucket",
		}
	})
	p.AddResourceConfigurator("minio_s3_bucket_retention", func(r *config.Resource) {
		r.ShortGroup = bucketShortGroup
		r.Kind = "Retention"
		r.References[bucketShortGroup] = config.Reference{
			TerraformName: "minio_s3_bucket",
		}
	})
	p.AddResourceConfigurator("minio_s3_bucket_server_side_encryption", func(r *config.Resource) {
		r.ShortGroup = bucketShortGroup
		r.Kind = "ServerSideEncryption"
		r.References[bucketShortGroup] = config.Reference{
			TerraformName: "minio_s3_bucket",
		}
	})
	p.AddResourceConfigurator("minio_s3_bucket_versioning", func(r *config.Resource) {
		r.ShortGroup = bucketShortGroup
		r.Kind = "Versioning"
		r.References[bucketShortGroup] = config.Reference{
			TerraformName: "minio_s3_bucket",
		}
	})
	p.AddResourceConfigurator("minio_s3_object", func(r *config.Resource) {
		r.ShortGroup = bucketShortGroup
		r.Kind = "Object"
		r.References[bucketShortGroup] = config.Reference{
			TerraformName: "minio_s3_bucket",
			RefFieldName:  "bucket_name",
		}
	})
}
func configureIAM(p *config.Provider) {
	p.AddResourceConfigurator("minio_iam_user", func(r *config.Resource) {
	})
	p.AddResourceConfigurator("minio_iam_user_policy_attachment", func(r *config.Resource) {
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
	p.AddResourceConfigurator("minio_iam_group_policy_attachment", func(r *config.Resource) {
		r.References["group_name"] = config.Reference{
			TerraformName: "minio_iam_group",
		}
		r.References["policy_name"] = config.Reference{
			TerraformName: "minio_iam_policy",
		}
	})
	p.AddResourceConfigurator("minio_iam_group_user_attachment", func(r *config.Resource) {
		r.References["group_name"] = config.Reference{
			TerraformName: "minio_iam_group",
		}
		r.References["user_name"] = config.Reference{
			TerraformName: "minio_iam_user",
		}
	})
	p.AddResourceConfigurator("minio_iam_ldap_group_policy_attachment", func(r *config.Resource) {
		r.Kind = "LDAPGroupPolicyAttachment"
		r.References["policy_name"] = config.Reference{
			TerraformName: "minio_iam_policy",
		}
	})
	p.AddResourceConfigurator("minio_iam_ldap_user_policy_attachment", func(r *config.Resource) {
		r.Kind = "LDAPUserPolicyAttachment"
		r.References["policy_name"] = config.Reference{
			TerraformName: "minio_iam_policy",
		}
	})
	p.AddResourceConfigurator("minio_iam_policy", func(r *config.Resource) {
	})
	p.AddResourceConfigurator("minio_iam_service_account", func(r *config.Resource) {
		r.References["target_user"] = config.Reference{
			TerraformName: "minio_iam_user",
		}
	})
}
