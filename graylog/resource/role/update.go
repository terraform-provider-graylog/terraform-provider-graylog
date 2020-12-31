package role

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/client"
)

func update(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := client.New(m)
	if err != nil {
		return err
	}
	oldName, newName := d.GetChange(keyName)
	data, err := getDataFromResourceData(d)
	if err != nil {
		return err
	}
	o := oldName.(string)
	if _, _, err := cl.Role.Update(ctx, o, data); err != nil {
		return fmt.Errorf("failed to update a role %s: %w", o, err)
	}
	d.SetId(newName.(string))
	return nil
}
