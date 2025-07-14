package iam

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("minio_iam_policy", func(r *config.Resource) {})
	p.AddResourceConfigurator("minio_iam_user", func(r *config.Resource) {
		r.ExternalName.OmittedFields = []string{"name"}
	})
	p.AddResourceConfigurator("minio_iam_user_policy_attachment", func(r *config.Resource) {
		r.References["policy_name"] = config.Reference{
			TerraformName:     "minio_iam_policy",
			RefFieldName:      "PolicyRef",
			SelectorFieldName: "PolicySelector",
		}
		r.References["user_name"] = config.Reference{
			TerraformName:     "minio_iam_user",
			RefFieldName:      "UserRef",
			SelectorFieldName: "UserSelector",
		}
	})
}
