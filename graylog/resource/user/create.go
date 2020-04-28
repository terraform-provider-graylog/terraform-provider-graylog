package user

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/suzuki-shunsuke/terraform-provider-graylog/graylog/client"
)

func create(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := client.New(m)
	if err != nil {
		return err
	}
	data, err := getDataFromResourceData(d)
	if err != nil {
		return err
	}
	if _, ok := data[keyPermissions]; !ok {
		data[keyPermissions] = []string{}
	}

	if _, err := cl.User.Create(ctx, data); err != nil {
		return err
	}
	d.SetId(data[keyUsername].(string))
	return nil
}
