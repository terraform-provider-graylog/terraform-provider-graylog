package convert

import (
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/stretchr/testify/require"
)

func TestGetSchema(t *testing.T) {
	data := []struct {
		title string
		data  interface{}
		sc    *schema.Schema
		isErr bool
		exp   interface{}
	}{
		{
			title: "normal",
			data:  "foo",
			sc: &schema.Schema{
				Type: schema.TypeString,
			},
			exp: "foo",
		},
		{
			title: "list",
			data:  []interface{}{"foo"},
			sc: &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			exp: []interface{}{"foo"},
		},
		{
			title: "set",
			data: schema.NewSet(schema.HashSchema(&schema.Schema{
				Type: schema.TypeString,
			}), []interface{}{"foo"}),
			sc: &schema.Schema{
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			exp: []interface{}{"foo"},
		},
		{
			title: "min max items 1",
			data: []interface{}{
				map[string]interface{}{
					"name": "foo",
				},
			},
			sc: &schema.Schema{
				Type:     schema.TypeList,
				MinItems: 1,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString,
						},
					},
				},
			},
			exp: map[string]interface{}{
				"name": "foo",
			},
		},
	}

	for _, d := range data {
		t.Run(d.title, func(t *testing.T) {
			a, err := GetSchema(d.data, d.sc)
			if d.isErr {
				require.NotNil(t, err)
				return
			}
			require.Nil(t, err)
			require.Equal(t, d.exp, a)
		})
	}
}

func TestSetSchema(t *testing.T) {
	data := []struct {
		title string
		data  interface{}
		sc    *schema.Schema
		isErr bool
		exp   interface{}
	}{
		{
			title: "normal",
			data:  "foo",
			sc: &schema.Schema{
				Type: schema.TypeString,
			},
			exp: "foo",
		},
		{
			title: "list of map",
			data: []interface{}{
				map[string]interface{}{
					"name": "foo",
					"id":   "001",
				},
			},
			sc: &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString,
						},
					},
				},
			},
			exp: []interface{}{
				map[string]interface{}{
					"name": "foo",
				},
			},
		},
	}

	for _, d := range data {
		t.Run(d.title, func(t *testing.T) {
			a, err := SetSchema(d.data, d.sc)
			if d.isErr {
				require.NotNil(t, err)
				return
			}
			require.Nil(t, err)
			require.Equal(t, d.exp, a)
		})
	}
}
