package output

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func Resource() *schema.Resource {
	return &schema.Resource{
		Create: create,
		Read:   read,
		Update: update,
		Delete: destroy,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"stream_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"output_ids": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}
