package bucket

import "github.com/crossplane/upjet/pkg/config"

const bucketShortGroup = "bucket"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
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
		}
	})
}
