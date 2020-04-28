package condition

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
	streamID := d.Get(keyStreamID).(string)
	aID := d.Get(keyAlertConditionID).(string)
	if _, err := cl.AlertCondition.Delete(ctx, streamID, aID); err != nil {
		return fmt.Errorf(
			"failed to delete a alert condition (stream id: %s, alert condition id: %s): %w",
			streamID, aID, err)
	}
	return nil
}
