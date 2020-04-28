package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/suzuki-shunsuke/terraform-provider-graylog/graylog/config"
)

func Configure(d *schema.ResourceData) (interface{}, error) {
	cfg := &config.Config{
		Endpoint:     d.Get("web_endpoint_uri").(string),
		AuthName:     d.Get("auth_name").(string),
		AuthPassword: d.Get("auth_password").(string),
		XRequestedBy: d.Get("x_requested_by").(string),
		APIVersion:   d.Get("api_version").(string),
	}

	if err := cfg.LoadAndValidate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func SchemaMap() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"web_endpoint_uri": {
			Type:        schema.TypeString,
			Required:    true,
			DefaultFunc: schema.MultiEnvDefaultFunc([]string{"GRAYLOG_WEB_ENDPOINT_URI"}, nil),
		},
		"auth_name": {
			Type:     schema.TypeString,
			Required: true,
			DefaultFunc: schema.MultiEnvDefaultFunc([]string{
				"GRAYLOG_AUTH_NAME"}, nil),
		},
		"auth_password": {
			Type:     schema.TypeString,
			Required: true,
			DefaultFunc: schema.MultiEnvDefaultFunc([]string{
				"GRAYLOG_AUTH_PASSWORD"}, nil),
		},
		"x_requested_by": {
			Type:     schema.TypeString,
			Optional: true,
			DefaultFunc: schema.MultiEnvDefaultFunc([]string{
				"GRAYLOG_X_REQUESTED_BY"}, "terraform-provider-graylog"),
		},
		"api_version": {
			Type:     schema.TypeString,
			Optional: true,
			DefaultFunc: schema.MultiEnvDefaultFunc([]string{
				"GRAYLOG_API_VERSION"}, "v3"),
		},
	}
}
