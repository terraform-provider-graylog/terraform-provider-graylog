package connection

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

	sID := data[keyStreamID].(string)
	if _, err := cl.PipelineConnection.ConnectPipelinesToStream(ctx, data); err != nil {
		return fmt.Errorf("failed to connect pipelines to a stream (stream id: %s): %w", sID, err)
	}
	d.SetId(sID)
	return nil
}
