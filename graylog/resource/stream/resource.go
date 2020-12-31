package stream

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
			// Required
			"title": {
				Type:     schema.TypeString,
				Required: true,
			},
			"index_set_id": {
				Type:     schema.TypeString,
				Required: true,
				// Not ForceNew
			},

			// Optional
			// rules
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// content_pack
			"matching_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"remove_matches_from_default_stream": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			// attributes
			"creator_user_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			// outputs
			"created_at": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"disabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"is_default": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			// alert_conditions
			// alert_receivers
		},
	}
}
