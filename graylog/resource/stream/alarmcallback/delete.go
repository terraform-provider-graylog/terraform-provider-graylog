package alarmcallback

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/suzuki-shunsuke/terraform-provider-graylog/graylog/client"
)

func destroy(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := client.New(m)
	if err != nil {
		return err
	}
	if _, err := cl.AlarmCallback.Delete(ctx, d.Get(keyStreamID).(string), d.Get(keyAlarmCallbackID).(string)); err != nil {
		return err
	}
	return nil
}
