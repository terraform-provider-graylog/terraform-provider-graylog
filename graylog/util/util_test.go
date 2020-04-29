package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRenameKey(t *testing.T) {
	data := []struct {
		title  string
		data   map[string]interface{}
		oldKey string
		newKey string
		exp    interface{}
		expF   bool
	}{
		{
			title: "normal",
			data: map[string]interface{}{
				"foo": "bar",
			},
			oldKey: "foo",
			newKey: "bar",
			exp:    "bar",
			expF:   true,
		},
		{
			title: "not found",
			data: map[string]interface{}{
				"foo": "bar",
			},
			oldKey: "bar",
			newKey: "foo",
		},
	}
	for _, d := range data {
		t.Run(d.title, func(t *testing.T) {
			value, f := RenameKey(d.data, d.oldKey, d.newKey)
			if d.expF {
				require.True(t, f)
				require.Equal(t, d.exp, value)
				return
			}
			require.False(t, f)
		})
	}
}
