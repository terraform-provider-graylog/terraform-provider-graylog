package role

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/client"
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

	if _, _, err := cl.Role.Create(ctx, data); err != nil {
		return fmt.Errorf("failed to create a role: %w", err)
	}
	d.SetId(data[keyName].(string))
	return nil
}
