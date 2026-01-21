package iam

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("minio_iam_policy", func(r *config.Resource) {})
	p.AddResourceConfigurator("minio_iam_user", func(r *config.Resource) {
		r.ExternalName.OmittedFields = []string{"name"}
	})
	p.AddResourceConfigurator("minio_iam_service_account", func(r *config.Resource) {
		r.References["target_user"] = config.Reference{
			TerraformName: "minio_iam_user",
		}
		if s, ok := r.TerraformResource.Schema["access_key"]; ok {
			s.Sensitive = true
		}
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if a, ok := attr["access_key"].(string); ok {
				conn["access_key"] = []byte(a)
			}
			if a, ok := attr["secret_key"].(string); ok {
				conn["secret_key"] = []byte(a)
			}
			return conn, nil
		}
	})
	p.AddResourceConfigurator("minio_iam_user_policy_attachment", func(r *config.Resource) {
		r.References["policy_name"] = config.Reference{
			TerraformName: "minio_iam_policy",
		}
		r.References["user_name"] = config.Reference{
			TerraformName: "minio_iam_user",
		}
	})
}
