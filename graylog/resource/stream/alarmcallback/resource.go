package alarmcallback

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/suzuki-shunsuke/terraform-provider-graylog/graylog/util"
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
			State: util.GenStateFunc("stream_id", "alarm_callback_id"),
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
			"configuration": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: util.SchemaDiffSuppressJSONString,
				ValidateFunc:     util.ValidateIsMapJSON,
			},
			"alarmcallback_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}
