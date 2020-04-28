package definition

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

	ds, _, err := cl.EventDefinition.Create(ctx, data)
	if err != nil {
		return fmt.Errorf("failed to create a event definition: %w", err)
	}
	d.SetId(ds[keyID].(string))
	return nil
}
