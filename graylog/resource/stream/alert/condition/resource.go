package condition

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

		SchemaVersion: 2,
		StateUpgraders: []schema.StateUpgrader{
			stateUpgraderV2,
		},

		Importer: &schema.ResourceImporter{
			State: util.GenStateFunc("stream_id", "alert_condition_id"),
		},

		Schema: map[string]*schema.Schema{
			// Required
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"title": {
				Type:     schema.TypeString,
				Required: true,
			},
			"stream_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"parameters": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: util.SchemaDiffSuppressJSONString,
				ValidateFunc:     util.ValidateIsMapJSON,
			},
			"alert_condition_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"in_grace": {
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}
