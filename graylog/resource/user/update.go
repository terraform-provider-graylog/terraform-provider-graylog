package user

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/client"
)

func update(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := client.New(m)
	if err != nil {
		return err
	}
	oldName, newName := d.GetChange(keyUsername)
	data, err := getDataFromResourceData(d)
	if err != nil {
		return err
	}
	if _, ok := data[keyPermissions]; !ok {
		data[keyPermissions] = []string{}
	}
	if _, err := cl.User.Update(ctx, oldName.(string), data); err != nil {
		return err
	}
	d.SetId(newName.(string))
	return nil
}
