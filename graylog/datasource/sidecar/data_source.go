package sidecar

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func DataSource() *schema.Resource {
	return &schema.Resource{
		Read: read,

		Schema: map[string]*schema.Schema{
			"node_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"node_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
