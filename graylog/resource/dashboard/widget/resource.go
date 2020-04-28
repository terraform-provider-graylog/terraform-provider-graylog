package widget

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/util"
)

func Resource() *schema.Resource {
	return &schema.Resource{
		Create: create,
		Read:   read,
		Update: update,
		Delete: destroy,

		SchemaVersion: 1,
		StateUpgraders: []schema.StateUpgrader{
			stateUpgraderV1,
		},

		Importer: &schema.ResourceImporter{
			State: util.GenStateFunc("dashboard_id", "widget_id"),
		},

		Schema: map[string]*schema.Schema{
			// Required
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"dashboard_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},

			"config": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: util.SchemaDiffSuppressJSONString,
				ValidateFunc:     util.ValidateIsJSON,
			},

			// Optional
			"cache_time": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"creator_user_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"widget_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}
