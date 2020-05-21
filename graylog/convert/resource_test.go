package convert

import (
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/stretchr/testify/require"
)

func TestGetResource(t *testing.T) {
	data := []struct {
		title string
		data  map[string]interface{}
		rsc   *schema.Resource
		isErr bool
		exp   interface{}
	}{
		{
			title: "normal",
			data: map[string]interface{}{
				"name": "foo",
			},
			rsc: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"name": {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
			exp: map[string]interface{}{
				"name": "foo",
			},
		},
	}

	for _, d := range data {
		d := d
		t.Run(d.title, func(t *testing.T) {
			a, err := GetResource(d.data, d.rsc)
			if d.isErr {
				require.NotNil(t, err)
				return
			}
			require.Nil(t, err)
			require.Equal(t, d.exp, a)
		})
	}
}
