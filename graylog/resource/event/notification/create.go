package notification

import (
	"context"
	"fmt"

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

	ds, _, err := cl.EventNotification.Create(ctx, data)
	if err != nil {
		return fmt.Errorf("failed to create a event notification: %w", err)
	}
	d.SetId(ds[keyID].(string))
	return nil
}
