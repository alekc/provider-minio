/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/crossplane/upjet/pkg/config"

//func ExternalNameConfig() config.ExternalName {
//	return config.ExternalName{
//		SetIdentifierArgumentFn: config.NopSetIdentifierArgument,
//		GetExternalNameFn:       config.IDAsExternalName,
//		GetIDFn:                 BuildIdentifyingPropertiesLookupIDFn(lookupConfig),
//		DisableNameInitializer:  true,
//	}
//}

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	"minio_accesskey":                  config.IdentifierFromProvider,
	"minio_iam_group":                  config.IdentifierFromProvider,
	"minio_s3_bucket":                  CustomParameterAsIdentifier("bucket", []string{"bucket_prefix"}),
	"minio_iam_user":                   CustomParameterAsIdentifier("name", []string{}),
	"minio_iam_user_policy_attachment": config.IdentifierFromProvider,
	"minio_iam_policy":                 config.IdentifierFromProvider,
}

// CustomParameterAsIdentifier follows the same logic as ParameterAsIdentifier, but
// keeps the parameter in the spec. This is useful when you want to be able to use the same bucket name in different
// providers

func CustomParameterAsIdentifier(param string, omittedFields []string) config.ExternalName {
	e := config.NameAsIdentifier
	e.SetIdentifierArgumentFn = func(base map[string]any, name string) {
		base[param] = name
	}
	e.OmittedFields = omittedFields
	e.IdentifierFields = []string{param}
	return e
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
