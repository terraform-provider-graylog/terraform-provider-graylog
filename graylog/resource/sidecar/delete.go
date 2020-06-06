package sidecar

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/client"
)

func destroy(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := client.New(m)
	if err != nil {
		return err
	}
	if _, err := cl.SidecarConfiguration.Assign(ctx, map[string]interface{}{
		"nodes": []interface{}{},
	}); err != nil {
		return fmt.Errorf("failed to delete congiuration assignments to sidecars: %w", err)
	}
	return nil
}
