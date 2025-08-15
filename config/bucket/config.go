package bucket

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("minio_s3_bucket", func(r *config.Resource) {
		r.ShortGroup = "bucket"
	})
	p.AddResourceConfigurator("minio_s3_bucket_notification", func(r *config.Resource) {
		r.ShortGroup = "bucket"
		r.Kind = "Notification"
		r.References["bucket"] = config.Reference{
			TerraformName: "minio_s3_bucket",
		}
	})
	p.AddResourceConfigurator("minio_s3_bucket_policy", func(r *config.Resource) {
		r.ShortGroup = "bucket"
		r.Kind = "Policy"
		r.References["bucket"] = config.Reference{
			TerraformName: "minio_s3_bucket",
		}
	})
	p.AddResourceConfigurator("minio_s3_bucket_replication", func(r *config.Resource) {
		r.ShortGroup = "bucket"
		r.Kind = "Replication"
		r.References["bucket"] = config.Reference{
			TerraformName: "minio_s3_bucket",
		}
	})
	p.AddResourceConfigurator("minio_s3_bucket_retention", func(r *config.Resource) {
		r.ShortGroup = "bucket"
		r.Kind = "Retention"
		r.References["bucket"] = config.Reference{
			TerraformName: "minio_s3_bucket",
		}
	})
	p.AddResourceConfigurator("minio_s3_bucket_server_side_encryption", func(r *config.Resource) {
		r.ShortGroup = "bucket"
		r.Kind = "ServerSideEncryption"
		r.References["bucket"] = config.Reference{
			TerraformName: "minio_s3_bucket",
		}
	})
}
