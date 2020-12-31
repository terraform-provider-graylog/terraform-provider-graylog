package input

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/util"
)

func Resource() *schema.Resource {
	return &schema.Resource{
		Create: create,
		Read:   read,
		Update: update,
		Delete: destroy,

		SchemaVersion:  schemaVersion,
		StateUpgraders: stateUpgraders,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			// required
			"title": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"attributes": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: util.SchemaDiffSuppressJSONString,
				ValidateFunc:     util.ValidateIsJSON,
			},

			"global": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"node": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"created_at": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"creator_user_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			// "context_pack": &schema.Schema{
			// 	Type:     schema.TypeString,
			// 	Optional: true,
			// },
		},
	}
}
