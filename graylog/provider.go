package graylog

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/suzuki-shunsuke/terraform-provider-graylog/graylog/provider"
)

// Provider returns a terraform resource provider for graylog.
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema:         provider.SchemaMap(),
		ResourcesMap:   resourceMap,
		DataSourcesMap: dataSourcesMap,
		ConfigureFunc:  provider.Configure,
	}
}
