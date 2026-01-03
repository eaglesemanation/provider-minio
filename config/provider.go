/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	accesskeyCluster "github.com/eaglesemanation/provider-minio/config/cluster/accesskey"
	iamCluster "github.com/eaglesemanation/provider-minio/config/cluster/iam"
	s3Cluster "github.com/eaglesemanation/provider-minio/config/cluster/s3"

	accesskeyNamespaced "github.com/eaglesemanation/provider-minio/config/namespaced/accesskey"
	iamNamespaced "github.com/eaglesemanation/provider-minio/config/namespaced/iam"
	s3Namespaced "github.com/eaglesemanation/provider-minio/config/namespaced/s3"
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
		accesskeyCluster.Configure,
		s3Cluster.Configure,
		iamCluster.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

// GetProviderNamespaced returns provider configuration
func GetProviderNamespaced() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("minio.m.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		accesskeyNamespaced.Configure,
		s3Namespaced.Configure,
		iamNamespaced.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
