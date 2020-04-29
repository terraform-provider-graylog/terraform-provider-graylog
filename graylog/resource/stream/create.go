package stream

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

	disabled := data[keyDisabled].(bool)
	delete(data, keyDisabled)

	stream, _, err := cl.Stream.Create(ctx, data)
	if err != nil {
		return fmt.Errorf("failed to create a stream: %w", err)
	}
	id := stream[keyStreamID].(string)
	d.SetId(id)

	// resume if needed
	if !disabled {
		if _, err := cl.Stream.Resume(ctx, id); err != nil {
			return fmt.Errorf("failed to resume a stream %s: %w", id, err)
		}
	}

	return nil
}
