package view

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/util"
)

var schemaState = &schema.Schema{
	Type:     schema.TypeList,
	Required: true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},

			// required
			// widgets widgetMapping widgetPositions
			"widgets": schemaWidgets,
			"widget_mapping": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: util.SchemaDiffSuppressJSONString,
				ValidateFunc:     util.ValidateIsJSON,
			},

			"positions": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: util.SchemaDiffSuppressJSONString,
				ValidateFunc:     util.ValidateIsJSON,
			},

			"titles": schemaTitles,

			// "selected_fields": {
			// 	Type:     schema.TypeList,
			// 	Optional: true,
			// 	Elem: &schema.Schema{
			// 		Type: schema.TypeString,
			// 	},
			// },
			// "static_message_list_id": null,
			// "formatting": null,
			// "display_mode_settings": {
			//   "positions": {}
			// }
		},
	},
}

var schemaTitles = &schema.Schema{
	Type:     schema.TypeList,
	Optional: true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"widget": &schema.Schema{
				Type:     schema.TypeMap,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	},
}

var schemaQuery = &schema.Schema{
	Type:     schema.TypeList,
	Optional: true,
	MinItems: 1,
	MaxItems: 1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"query_string": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	},
}

var schemaWidgets = &schema.Schema{
	Type:     schema.TypeList,
	Required: true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			// required
			// id type config
			"widget_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"config": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: util.SchemaDiffSuppressJSONString,
				ValidateFunc:     util.ValidateIsJSON,
			},
			// "filter": null,
			"timerange": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: util.SchemaDiffSuppressJSONString,
				ValidateFunc:     util.ValidateIsJSON,
			},
			"query": schemaQuery,
			// "streams": [],
		},
	},
}
