package sidecar

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
	if _, err := cl.SidecarConfiguration.Assign(ctx, data); err != nil {
		return fmt.Errorf("failed to update sidecars's assignments: %w", err)
	}
	return nil
}
