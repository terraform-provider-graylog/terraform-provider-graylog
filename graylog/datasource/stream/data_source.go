package stream

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSource() *schema.Resource {
	return &schema.Resource{
		Read: read,

		Schema: map[string]*schema.Schema{
			"title": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"stream_id": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"index_set_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			// content_pack
			"matching_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"remove_matches_from_default_stream": {
				Type:     schema.TypeBool,
				Computed: true,
			},

			// attributes
			"creator_user_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			// outputs
			"created_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"disabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_default": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}
