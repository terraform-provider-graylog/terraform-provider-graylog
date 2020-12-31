package provider

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSchemaMap(t *testing.T) {
	t.Parallel()
	require.NotNil(t, SchemaMap())
}
