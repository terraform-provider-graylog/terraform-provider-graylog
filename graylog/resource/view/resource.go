package view

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Resource() *schema.Resource {
	return &schema.Resource{
		Create: create,
		Read:   read,
		Update: update,
		Delete: destroy,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			// create
			// state, search_id, properties, summary, title, type, id, description, requires, owner, created_at

			// Required
			"title": {
				Type:     schema.TypeString,
				Required: true,
			},

			"type": {
				Type:     schema.TypeString,
				Required: true,
			},

			"search_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			"state": schemaState,

			// Optional
			// "properties": [],
			// "requires": {},

			"summary": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"owner": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"created_at": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}
