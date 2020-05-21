package collector

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
	if _, err := cl.Collector.Delete(ctx, d.Id()); err != nil {
		return fmt.Errorf("failed to delete a collector %s: %w", d.Id(), err)
	}
	return nil
}
