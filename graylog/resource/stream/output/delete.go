package output

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/client"
)

func destroy(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := client.New(m)
	if err != nil {
		return err
	}
	streamID := d.Get(keyStreamID).(string)
	for _, outputID := range d.Get(keyOutputIDs).(*schema.Set).List() {
		if _, err := cl.StreamOutput.Delete(ctx, streamID, outputID.(string)); err != nil {
			return err
		}
	}
	return nil
}
