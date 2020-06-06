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
	data, _, err := cl.Sidecar.GetAll(ctx)
	if err != nil {
		return fmt.Errorf("failed to get all sidecars to destroy: %w", err)
	}
	sidecars := data[keySidecars].([]interface{})
	for i, sidecar := range sidecars {
		nodeID := sidecar.(map[string]interface{})[keyNodeID].(string)
		sidecars[i] = map[string]interface{}{
			keyNodeID:      nodeID,
			keyAssignments: []interface{}{},
		}
	}

	if _, err := cl.SidecarConfiguration.Assign(ctx, map[string]interface{}{
		"nodes": sidecars,
	}); err != nil {
		return fmt.Errorf("failed to delete congiuration assignments to sidecars: %w", err)
	}
	return nil
}
