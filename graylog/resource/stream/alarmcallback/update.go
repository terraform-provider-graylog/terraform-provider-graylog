package alarmcallback

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
	delete(data, keyID)
	delete(data, keyStreamID)
	delete(data, keyAlarmCallbackID)

	sID := d.Get(keyStreamID).(string)
	aID := d.Get(keyAlarmCallbackID).(string)
	if _, err := cl.AlarmCallback.Update(ctx, sID, aID, data); err != nil {
		return fmt.Errorf(
			"failed to update a alarm callback (stream id: %s, alarm callback id: %s): %w", sID, aID, err)
	}
	return nil
}
