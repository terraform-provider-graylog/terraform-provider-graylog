package provider

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSchemaMap(t *testing.T) {
	require.NotNil(t, SchemaMap())
}
