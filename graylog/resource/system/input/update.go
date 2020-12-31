package input

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
	data, err := getDataFromResourceData(d)
	if err != nil {
		return err
	}

	delete(data, "created_at")
	delete(data, "creator_user_id")

	if _, _, err := cl.Input.Update(ctx, d.Id(), data); err != nil {
		return fmt.Errorf("failed to update a input %s: %w", d.Id(), err)
	}
	return nil
}
