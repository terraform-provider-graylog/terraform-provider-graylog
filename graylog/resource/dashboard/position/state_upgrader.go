package position

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/convert"
)

const schemaVersion = 1

var stateUpgraders = []schema.StateUpgrader{
	stateUpgraderV1,
}

func positionResourceV0() *schema.Resource {
	return &schema.Resource{}
}

var stateUpgraderV1 = schema.StateUpgrader{
	Version: 0,
	Type:    positionResourceV0().CoreConfigSchema().ImpliedType(),
	Upgrade: func(_ context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
		if a, ok := rawState[keyPositions]; ok {
			b, err := json.Marshal(convert.ListToMap(a.([]interface{}), keyWidgetID))
			if err != nil {
				return nil, fmt.Errorf("failed to marshal 'positions' as JSON: %w", err)
			}
			rawState[keyPositions] = string(b)
		}

		return rawState, nil
	},
}
