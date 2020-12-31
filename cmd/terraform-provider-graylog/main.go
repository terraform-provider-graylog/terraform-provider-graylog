package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: graylog.Provider,
	})
}
