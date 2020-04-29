package dashboard

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func DataSource() *schema.Resource {
	return &schema.Resource{
		Read: read,

		Schema: map[string]*schema.Schema{
			"dashboard_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"title": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"description": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},

			"created_at": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}
