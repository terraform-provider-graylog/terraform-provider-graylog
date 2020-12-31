package rule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const schemaVersion = 1

var stateUpgraders = []schema.StateUpgrader{
	stateUpgraderV1,
}

func ruleResourceV0() *schema.Resource {
	return &schema.Resource{}
}

var stateUpgraderV1 = schema.StateUpgrader{
	Version: 0,
	Type:    ruleResourceV0().CoreConfigSchema().ImpliedType(),
	Upgrade: func(_ context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
		streamID := rawState[keyStreamID].(string)
		ruleID := rawState[keyID].(string)

		rawState[keyRuleID] = ruleID
		rawState[keyID] = streamID + "/" + ruleID

		return rawState, nil
	},
}
