package condition

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
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

	delete(data, keyID)
	delete(data, keyInGrace)
	delete(data, keyStreamID)
	delete(data, keyAlertConditionID)

	sID := d.Get(keyStreamID).(string)
	aID := d.Get(keyAlertConditionID).(string)
	if _, err := cl.AlertCondition.Update(ctx, sID, aID, data); err != nil {
		return fmt.Errorf(
			"failed to update a alert condition (stream id: %s, alert condition id: %s): %w", sID, aID, err)
	}
	return nil
}
