package definition

import (
	"errors"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/util"
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
			"title": {
				Type:     schema.TypeString,
				Required: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Required: true,
				ValidateFunc: util.WrapValidateFunc(func(v interface{}, k string) error {
					priority := v.(int)
					if priority < 1 || priority > 3 {
						return errors.New("'priority' should be either 1, 2, and 3")
					}
					return nil
				}),
			},
			"config": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: util.SchemaDiffSuppressJSONString,
				ValidateFunc:     util.ValidateIsJSON,
			},
			"notification_settings": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"grace_period_ms": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"backlog_size": {
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},

			"alert": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"field_spec": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "{}",
				DiffSuppressFunc: util.SchemaDiffSuppressJSONString,
				ValidateFunc:     util.ValidateIsJSON,
			},
			"notifications": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"notification_id": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"key_spec": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			// optional
			//	"storage": {
			//		Type:     schema.TypeString,
			//		Optional: true,
			//	},
		},
	}
}
