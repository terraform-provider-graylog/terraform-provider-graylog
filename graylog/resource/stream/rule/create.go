package rule

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/suzuki-shunsuke/terraform-provider-graylog/graylog/client"
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
	d.Set(keyRuleID, ruleID)
	d.SetId(streamID + "/" + ruleID)
	return nil
}
