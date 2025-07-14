package s3

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("minio_s3_bucket", func(r *config.Resource) {
		r.ExternalName.SetIdentifierArgumentFn = func(base map[string]any, externalName string) {
			base["bucket"] = externalName
		}
		r.ExternalName.OmittedFields = []string{
			"bucket",
		}
	})
}
