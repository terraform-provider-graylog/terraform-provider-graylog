package graylog

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProvider(t *testing.T) {
	t.Parallel()
	require.NotNil(t, Provider())
}
