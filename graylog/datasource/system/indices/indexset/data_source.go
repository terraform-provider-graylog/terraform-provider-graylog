package indexset

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSource() *schema.Resource {
	return &schema.Resource{
		Read: read,

		Schema: map[string]*schema.Schema{
			"index_set_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"title": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"index_prefix": {
				Type:     schema.TypeString,
				Optional: true,
			},

			// computed
			"rotation_strategy_class": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"rotation_strategy": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"retention_strategy_class": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"retention_strategy": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"index_analyzer": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"shards": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"index_optimization_max_num_segments": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			"creation_date": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"replicas": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"index_optimization_disabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"writable": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"default": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			// field_type_refresh_interval is added from Graylog API v3
			"field_type_refresh_interval": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}
