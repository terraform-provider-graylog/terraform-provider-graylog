package connection

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/suzuki-shunsuke/terraform-provider-graylog/graylog/client"
)

func destroy(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := client.New(m)
	if err != nil {
		return err
	}

	if _, err := cl.PipelineConnection.ConnectPipelinesToStream(ctx, map[string]interface{}{
		"stream_id":    d.Id(),
		"pipeline_ids": []string{},
	}); err != nil {
		return fmt.Errorf(
			"failed to delete connections between a stream and pipelines (stream id: %s): %w", d.Id(), err)
	}
	return nil
}
