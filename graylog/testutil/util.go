package testutil

import (
	"errors"
	"net/http"
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/suzuki-shunsuke/flute/v2/flute"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/provider"
)

func Header() http.Header {
	return http.Header{
		"Content-Type":   []string{"application/json"},
		"X-Requested-By": []string{"terraform-provider-graylog"},
		"Authorization":  nil,
	}
}

func SetEnv() error {
	envs := map[string]string{
		"GRAYLOG_WEB_ENDPOINT_URI": "http://example.com/api",
		"GRAYLOG_AUTH_NAME":        "admin",
		"GRAYLOG_AUTH_PASSWORD":    "admin",
	}
	for k, v := range envs {
		if err := os.Setenv(k, v); err != nil {
			return err
		}
	}
	return nil
}

func SetHTTPClient(t *testing.T, routes ...flute.Route) {
	transport := flute.Transport{
		T: t,
		Services: []flute.Service{
			{
				Endpoint: "http://example.com",
				Routes:   routes,
			},
		},
	}

	http.DefaultClient = &http.Client{
		Transport: transport,
	}
}

func SingleResourceProviders(name string, rsc *schema.Resource) map[string]terraform.ResourceProvider {
	return map[string]terraform.ResourceProvider{
		"graylog": &schema.Provider{
			Schema: provider.SchemaMap(),
			ResourcesMap: map[string]*schema.Resource{
				name: rsc,
			},
			ConfigureFunc: provider.Configure,
		},
	}
}

func SingleDataSourceProviders(name string, rsc *schema.Resource) map[string]terraform.ResourceProvider {
	return map[string]terraform.ResourceProvider{
		"graylog": &schema.Provider{
			Schema: provider.SchemaMap(),
			DataSourcesMap: map[string]*schema.Resource{
				name: rsc,
			},
			ConfigureFunc: provider.Configure,
		},
	}
}

func EqualMapKeys(data map[string]interface{}, keys ...string) error {
	for _, k := range keys {
		if _, ok := data[k]; !ok {
			return errors.New("map should have a key: " + k)
		}
	}
	for d := range data {
		f := false
		for _, k := range keys {
			if d == k {
				f = true
				break
			}
		}
		if !f {
			return errors.New("map has an unexpected key: " + d)
		}
	}
	return nil
}
