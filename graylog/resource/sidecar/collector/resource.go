package collector

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
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"service_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"node_operating_system": {
				Type:     schema.TypeString,
				Required: true,
			},
			"executable_path": {
				Type:     schema.TypeString,
				Required: true,
			},
			"execute_parameters": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"validation_parameters": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"default_template": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
