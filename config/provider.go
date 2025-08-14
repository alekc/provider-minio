/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"
	"github.com/alekc/provider-minio/config/bucket"
	"github.com/alekc/provider-minio/config/iam"
	"github.com/alekc/provider-minio/config/policy"

	ujconfig "github.com/crossplane/upjet/pkg/config"
)

const (
	resourcePrefix = "minio"
	modulePath     = "github.com/alekc/provider-minio"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("minio.alekc.dev"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		bucket.Configure,
		iam.Configure,
		policy.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
