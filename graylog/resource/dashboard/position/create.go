package position

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func create(d *schema.ResourceData, m interface{}) error {
	if err := update(d, m); err != nil {
		return err
	}
	d.SetId(d.Get(keyDashboardID).(string))
	return nil
}
