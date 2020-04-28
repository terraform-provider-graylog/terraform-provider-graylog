package position

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
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			// Required
			"dashboard_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"positions": {
				Required:         true,
				Type:             schema.TypeString,
				DiffSuppressFunc: util.SchemaDiffSuppressJSONString,
				ValidateFunc:     util.ValidateIsMapJSON,
			},
		},
	}
}
