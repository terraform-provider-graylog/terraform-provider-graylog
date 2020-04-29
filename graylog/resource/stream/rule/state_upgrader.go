package rule

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/util"
)

const schemaVersion = 1

var stateUpgraders = []schema.StateUpgrader{
	stateUpgraderV1,
}

var stateUpgraderV1 = schema.StateUpgrader{
	Version: 0,
	Type:    util.UpgraderType(),
	Upgrade: func(rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
		streamID := rawState[keyStreamID].(string)
		ruleID := rawState[keyID].(string)

		rawState[keyRuleID] = ruleID
		rawState[keyID] = streamID + "/" + ruleID

		return rawState, nil
	},
}
