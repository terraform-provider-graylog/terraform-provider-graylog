package extractor

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
			StateContext: util.GenStateFunc(keyInputID, keyExtractorID),
		},

		Schema: map[string]*schema.Schema{
			"input_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"title": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cursor_strategy": {
				Type:     schema.TypeString,
				Required: true,
			},
			"source_field": {
				Type:     schema.TypeString,
				Required: true,
			},
			"condition_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"order": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"condition_value": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"target_field": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"extractor_config": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: util.SchemaDiffSuppressJSONString,
				ValidateFunc:     util.ValidateIsJSON,
			},

			"converters": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"config": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: util.SchemaDiffSuppressJSONString,
							ValidateFunc:     util.ValidateIsJSON,
						},
					},
				},
			},

			"extractor_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}
