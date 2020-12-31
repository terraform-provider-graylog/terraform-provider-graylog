package rule

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

	delete(data, keyRuleID)

	streamID := data[keyStreamID].(string)
	delete(data, keyStreamID)

	rule, _, err := cl.StreamRule.Create(ctx, streamID, data)
	if err != nil {
		return fmt.Errorf("failed to create a stream rule (stream id: %s): %w", streamID, err)
	}
	ruleID := rule[keyStreamRuleID].(string)
	if err := d.Set(keyRuleID, ruleID); err != nil {
		return err
	}
	d.SetId(streamID + "/" + ruleID)
	return nil
}
