package graylog

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProvider(t *testing.T) {
	require.NotNil(t, Provider())
}
