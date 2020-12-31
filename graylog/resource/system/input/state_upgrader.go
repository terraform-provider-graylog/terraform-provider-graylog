package input

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/convert"
)

const schemaVersion = 1

var stateUpgraders = []schema.StateUpgrader{
	stateUpgraderV1,
}

func resourceV0() *schema.Resource {
	return &schema.Resource{}
}

var stateUpgraderV1 = schema.StateUpgrader{
	Version: 0,
	Type:    resourceV0().CoreConfigSchema().ImpliedType(),
	Upgrade: func(_ context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
		if err := convert.OneSizeListToJSON(rawState, keyAttributes); err != nil {
			return nil, err
		}
		return rawState, nil
	},
}
