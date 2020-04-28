package notification

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/client"
)

func update(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := client.New(m)
	if err != nil {
		return err
	}
	data, err := getDataFromResourceData(d)
	if err != nil {
		return err
	}
	data[keyID] = d.Id()
	if _, _, err := cl.EventNotification.Update(ctx, d.Id(), data); err != nil {
		return fmt.Errorf("failed to update a event notification %s: %w", d.Id(), err)
	}
	return nil
}
