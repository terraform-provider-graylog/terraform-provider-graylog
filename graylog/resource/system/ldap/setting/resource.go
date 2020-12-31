package setting

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Resource() *schema.Resource {
	return &schema.Resource{
		Create: create,
		Read:   read,
		Update: update,
		Delete: destroy,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			// required
			"system_username": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ldap_uri": {
				Type:     schema.TypeString,
				Required: true,
			},
			"search_base": {
				Type:     schema.TypeString,
				Required: true,
			},
			"search_pattern": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name_attribute": {
				Type:     schema.TypeString,
				Required: true,
			},
			"default_group": {
				Type:     schema.TypeString,
				Required: true,
			},

			// optional
			// system_password is required to create the resource
			"system_password": {
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"use_start_tls": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"trust_all_certificates": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"active_directory": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"group_search_base": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"group_id_attribute": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"group_search_pattern": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"group_mapping": {
				Type:     schema.TypeMap,
				Optional: true,
			},
			"additional_default_groups": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"system_password_set": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
		},
	}
}
