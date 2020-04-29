package output

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func create(d *schema.ResourceData, m interface{}) error {
	return update(d, m)
}
