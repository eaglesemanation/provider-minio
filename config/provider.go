/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	"github.com/eaglesemanation/provider-minio/config/accesskey"
	"github.com/eaglesemanation/provider-minio/config/iam"
	"github.com/eaglesemanation/provider-minio/config/s3"
)

const (
	resourcePrefix = "minio"
	modulePath     = "github.com/eaglesemanation/provider-minio"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("minio.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		accesskey.Configure,
		s3.Configure,
		iam.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
