package stream

import (
	"github.com/hashicorp/terraform/helper/schema"
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
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			// content_pack
			"matching_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"remove_matches_from_default_stream": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
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
				Computed: true,
			},
			"is_default": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
		},
	}
}
