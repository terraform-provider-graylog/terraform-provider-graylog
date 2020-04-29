package indexset

import (
	"github.com/hashicorp/terraform/helper/schema"
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
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"title": {
				Type:     schema.TypeString,
				Required: true,
			},
			"index_prefix": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"rotation_strategy_class": {
				Type:     schema.TypeString,
				Required: true,
			},
			"rotation_strategy": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: util.SchemaDiffSuppressJSONString,
				ValidateFunc:     util.ValidateIsJSON,
			},
			"retention_strategy_class": {
				Type:     schema.TypeString,
				Required: true,
			},
			"retention_strategy": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: util.SchemaDiffSuppressJSONString,
				ValidateFunc:     util.ValidateIsJSON,
			},
			"index_analyzer": {
				Type:     schema.TypeString,
				Required: true,
			},
			// >= 1
			"shards": {
				Type:     schema.TypeInt,
				Required: true,
			},
			// >= 1
			"index_optimization_max_num_segments": {
				Type:     schema.TypeInt,
				Required: true,
			},

			// Optional
			"creation_date": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"replicas": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"index_optimization_disabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"writable": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"default": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			// field_type_refresh_interval was added from Graylog API v3
			"field_type_refresh_interval": {
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}
