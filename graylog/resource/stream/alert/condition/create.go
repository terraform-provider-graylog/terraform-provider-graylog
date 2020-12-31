package condition

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
	delete(data, keyID)
	delete(data, keyInGrace)

	streamID := data[keyStreamID].(string)
	delete(data, keyStreamID)
	delete(data, keyAlertConditionID)
	ac, _, err := cl.AlertCondition.Create(ctx, streamID, data)
	if err != nil {
		return fmt.Errorf("failed to create a alert condition (stream id: %s): %w", streamID, err)
	}
	acID := ac[keyAlertConditionID].(string)
	if err := d.Set(keyAlertConditionID, acID); err != nil {
		return err
	}
	d.SetId(streamID + "/" + acID)
	return nil
}
