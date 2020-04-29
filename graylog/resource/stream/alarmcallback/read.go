package alarmcallback

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/client"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/util"
)

func read(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := client.New(m)
	if err != nil {
		return err
	}

	streamID := d.Get(keyStreamID).(string)
	aID := d.Get(keyAlarmCallbackID).(string)
	data, resp, err := cl.AlarmCallback.Get(ctx, streamID, aID)
	if err != nil {
		return util.HandleGetResourceError(
			d, resp, fmt.Errorf(
				"failed to get a alarm callback (stream id: %s, alarm callback id: %s): %w", streamID, aID, err))
	}
	return setDataToResourceData(d, data)
}
