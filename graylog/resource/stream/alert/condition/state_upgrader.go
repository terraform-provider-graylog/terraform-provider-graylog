package condition

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/suzuki-shunsuke/terraform-provider-graylog/graylog/util"
)

var stateUpgraderV2 = schema.StateUpgrader{
	Version: 1,
	Type:    util.UpgraderType(),
	Upgrade: func(rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
		attrTypes := []string{
			"field_content_value_parameters",
			"field_value_parameters",
			"message_count_parameters",
		}

		generalAttrTypes := []string{
			"general_int_parameters",
			"general_bool_parameters",
			"general_float_parameters",
			"general_string_parameters",
		}

		attrs := map[string]interface{}{}

		for _, a := range attrTypes {
			v, ok := rawState[a]
			if !ok || v == nil {
				continue
			}
			arr := v.([]interface{})
			if len(arr) == 0 {
				continue
			}
			for k, attr := range arr[0].(map[string]interface{}) {
				attrs[k] = attr
			}
			delete(rawState, a)
		}

		for _, a := range generalAttrTypes {
			v, ok := rawState[a]
			if !ok || v == nil {
				continue
			}
			for k, attr := range v.(map[string]interface{}) {
				attrs[k] = attr
			}
			delete(rawState, a)
		}

		b, err := json.Marshal(attrs)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal attributes '%s' as JSON: %w", keyParameters, err)
		}
		rawState[keyParameters] = string(b)

		streamID := rawState[keyStreamID].(string)
		alertConditionID := rawState[keyID].(string)

		rawState[keyAlertConditionID] = alertConditionID
		rawState[keyID] = streamID + "/" + alertConditionID

		return rawState, nil
	},
}
