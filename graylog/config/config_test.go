package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfig_LoadAndValidate(t *testing.T) {
	t.Parallel()
	cfg := Config{
		Endpoint:     "http://example.com:9000/api",
		AuthName:     "xxx",
		AuthPassword: "token",
		XRequestedBy: "terraform-provider-graylog",
		APIVersion:   "v3",
	}
	require.Nil(t, cfg.LoadAndValidate())
}
