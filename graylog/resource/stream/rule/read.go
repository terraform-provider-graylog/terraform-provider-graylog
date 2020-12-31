package rule

import (
	"context"
	"fmt"

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

	sID := d.Get(keyStreamID).(string)
	rID := d.Get(keyRuleID).(string)
	data, resp, err := cl.StreamRule.Get(ctx, sID, rID)
	if err != nil {
		return util.HandleGetResourceError(
			d, resp, fmt.Errorf(
				"failed to get a stream rule (stream id: %s, rule id: %s): %w", sID, rID, err))
	}
	return setDataToResourceData(d, data)
}
