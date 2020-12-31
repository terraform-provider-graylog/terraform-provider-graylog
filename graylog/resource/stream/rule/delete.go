package rule

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/client"
)

func destroy(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := client.New(m)
	if err != nil {
		return err
	}

	sID := d.Get(keyStreamID).(string)
	rID := d.Get(keyRuleID).(string)
	if _, err := cl.StreamRule.Delete(ctx, sID, rID); err != nil {
		return fmt.Errorf(
			"failed to delete a stream rule (stream id: %s, rule id: %s): %w", sID, rID, err)
	}
	return nil
}
