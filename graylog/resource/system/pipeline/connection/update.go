package connection

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func update(d *schema.ResourceData, m interface{}) error {
	return create(d, m)
}
