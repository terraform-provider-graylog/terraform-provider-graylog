package staticfield

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func update(d *schema.ResourceData, m interface{}) error {
	return create(d, m)
}
