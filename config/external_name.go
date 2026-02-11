/*
Copyright 2022 Upbound Inc.
*/

package config

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	"minio_accesskey":                        config.IdentifierFromProvider,
	"minio_iam_group":                        config.IdentifierFromProvider,
	"minio_iam_group_membership":             config.IdentifierFromProvider,
	"minio_iam_group_policy":                 config.IdentifierFromProvider,
	"minio_iam_group_policy_attachment":      config.IdentifierFromProvider,
	"minio_iam_group_user_attachment":        config.IdentifierFromProvider,
	"minio_iam_ldap_group_policy_attachment": config.IdentifierFromProvider,
	"minio_iam_ldap_user_policy_attachment":  config.IdentifierFromProvider,
	"minio_iam_user":                         config.IdentifierFromProvider,
	"minio_iam_user_policy_attachment":       config.IdentifierFromProvider,
	"minio_iam_policy":                       config.IdentifierFromProvider,
	"minio_iam_service_account":              config.IdentifierFromProvider,
	"minio_ilm_policy":                       config.IdentifierFromProvider,
	"minio_ilm_tier":                         config.IdentifierFromProvider,
	"minio_kms_key":                          config.IdentifierFromProvider,
	"minio_s3_bucket":                        config.ParameterAsIdentifier("bucket"),
	"minio_s3_bucket_notification":           config.IdentifierFromProvider,
	"minio_s3_bucket_policy":                 config.IdentifierFromProvider,
	"minio_s3_bucket_replication":            config.IdentifierFromProvider,
	"minio_s3_bucket_retention":              config.IdentifierFromProvider,
	"minio_s3_bucket_server_side_encryption": config.IdentifierFromProvider,
	"minio_s3_bucket_versioning":             config.IdentifierFromProvider,
	"minio_s3_object":                        config.IdentifierFromProvider,
}

// ExternalNameConfigsNamespaceOverride contains all external name configurations overrides for namespaced resources
var ExternalNameConfigsNamespaceOverride = map[string]config.ExternalName{
	"minio_iam_group":            config.NameAsIdentifier,
	"minio_iam_group_membership": config.NameAsIdentifier,
	"minio_iam_policy":           config.NameAsIdentifier,
	"minio_ilm_tier":             CustomParameterAsIdentifier("name", []string{"prefix"}),
	"minio_iam_service_account":  config.NameAsIdentifier,
	"minio_iam_user":             config.NameAsIdentifier,
}

// CustomParameterAsIdentifier follows the same logic as ParameterAsIdentifier, but
// keeps the parameter in the spec. This is useful when you want to be able to use the same bucket name in different
// providers

func CustomParameterAsIdentifier(param string, omittedFields []string) config.ExternalName {
	e := config.ExternalName{
		SetIdentifierArgumentFn: config.NopSetIdentifierArgument,
		GetExternalNameFn:       config.IDAsExternalName,
		GetIDFn:                 config.ExternalNameAsID,
		DisableNameInitializer:  true,
		OmittedFields:           omittedFields,
	}

	return e
}

// getExternalNameConfigs returns the map of external name configurations.
// If `namespaced` is true, it returns the default configs.
// Otherwise, it applies namespace overrides to the configs before returning.
func getExternalNameConfigs(namespaced bool) map[string]config.ExternalName {
	if !namespaced {
		return ExternalNameConfigs
	}
	// Create a copy to avoid mutating the global map
	configs := make(map[string]config.ExternalName, len(ExternalNameConfigs))
	for k, v := range ExternalNameConfigs {
		configs[k] = v
	}
	for k, v := range ExternalNameConfigsNamespaceOverride {
		configs[k] = v
	}
	return configs
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations(namespaced bool) config.ResourceOption {
	return func(r *config.Resource) {
		nameConfigs := getExternalNameConfigs(namespaced)
		if e, ok := nameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured(namespaced bool) []string {
	nameConfigs := getExternalNameConfigs(namespaced)
	l := make([]string, len(nameConfigs))
	i := 0
	for name := range nameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
