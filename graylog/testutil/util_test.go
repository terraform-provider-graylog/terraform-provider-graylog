package testutil

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/stretchr/testify/require"
	"github.com/suzuki-shunsuke/flute/flute"
)

func TestHeader(t *testing.T) {
	require.NotNil(t, Header())
}

func TestSetEnv(t *testing.T) {
	require.Nil(t, SetEnv())
	require.Equal(t, "http://example.com/api", os.Getenv("GRAYLOG_WEB_ENDPOINT_URI"))
	require.Equal(t, "admin", os.Getenv("GRAYLOG_AUTH_NAME"))
	require.Equal(t, "admin", os.Getenv("GRAYLOG_AUTH_PASSWORD"))
}

func TestSetHTTPClient(t *testing.T) {
	routes := []flute.Route{}
	SetHTTPClient(t, routes...)
}

func TestSingleResourceProviders(t *testing.T) {
	require.NotNil(t, SingleResourceProviders("graylog_dashboard", &schema.Resource{}))
}

func TestSingleDataSourceProviders(t *testing.T) {
	require.NotNil(t, SingleDataSourceProviders("graylog_dashboard", &schema.Resource{}))
}

func TestEqualMapKeys(t *testing.T) {
	data := []struct {
		title string
		data  map[string]interface{}
		keys  []string
		isErr bool
	}{
		{
			title: "normal",
			data: map[string]interface{}{
				"foo": "bar",
			},
			keys: []string{"foo"},
		},
		{
			title: "unexpected key",
			data: map[string]interface{}{
				"foo":   "bar",
				"hello": "hello",
			},
			keys:  []string{"foo"},
			isErr: true,
		},
		{
			title: "key is required",
			data: map[string]interface{}{
				"foo":   "bar",
				"hello": "hello",
			},
			keys:  []string{"foo", "hello", "bar"},
			isErr: true,
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.title, func(t *testing.T) {
			err := EqualMapKeys(d.data, d.keys...)
			if d.isErr {
				require.NotNil(t, err)
				return
			}
			require.Nil(t, err)
		})
	}
}
