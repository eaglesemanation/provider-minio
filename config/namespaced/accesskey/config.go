package accesskey

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("minio_accesskey", func(r *config.Resource) {
		r.ShortGroup = ""
		r.Kind = "AccessKey"
		r.References["user"] = config.Reference{
			TerraformName:     "minio_iam_user",
			RefFieldName:      "UserRef",
			SelectorFieldName: "UserSelector",
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
}
