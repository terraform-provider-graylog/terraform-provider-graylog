package condition

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/client"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/util"
)

func read(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := client.New(m)
	if err != nil {
		return err
	}
	data, resp, err := cl.AlertCondition.Get(ctx, d.Get(keyStreamID).(string), d.Get(keyAlertConditionID).(string))
	if err != nil {
		return util.HandleGetResourceError(d, resp, err)
	}
	return setDataToResourceData(d, data)
}
