package position

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/suzuki-shunsuke/terraform-provider-graylog/graylog/convert"
	"github.com/suzuki-shunsuke/terraform-provider-graylog/graylog/util"
)

var stateUpgraderV1 = schema.StateUpgrader{
	Version: 0,
	Type:    util.UpgraderType(),
	Upgrade: func(rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
		if a, ok := rawState[keyPositions]; ok {
			b, err := json.Marshal(convert.ConvertListToMap(a.([]interface{}), keyWidgetID))
			if err != nil {
				return nil, fmt.Errorf("failed to marshal 'positions' as JSON: %w", err)
			}
			rawState[keyPositions] = string(b)
		}

		return rawState, nil
	},
}
