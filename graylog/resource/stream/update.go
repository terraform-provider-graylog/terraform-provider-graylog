package stream

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

	disabled := data[keyDisabled].(bool)
	delete(data, keyDisabled)
	delete(data, keyCreatorUserID)
	delete(data, keyCreatedAt)

	if _, _, err := cl.Stream.Update(ctx, d.Id(), data); err != nil {
		return fmt.Errorf("failed to update a stream %s: %w", d.Id(), err)
	}

	if !d.HasChange(keyDisabled) {
		return nil
	}
	if disabled {
		if _, err := cl.Stream.Pause(ctx, d.Id()); err != nil {
			return fmt.Errorf("failed to pause a stream %s: %w", d.Id(), err)
		}
		return nil
	}
	if _, err := cl.Stream.Resume(ctx, d.Id()); err != nil {
		return fmt.Errorf("failed to resume a stream %s: %w", d.Id(), err)
	}

	return nil
}
