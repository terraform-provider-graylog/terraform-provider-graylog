package alarmcallback

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const schemaVersion = 1

var stateUpgraders = []schema.StateUpgrader{
	stateUpgraderV1,
}

func alarmCallbackResourceV0() *schema.Resource {
	return &schema.Resource{}
}

var stateUpgraderV1 = schema.StateUpgrader{
	Version: 0,
	Type:    alarmCallbackResourceV0().CoreConfigSchema().ImpliedType(),
	Upgrade: func(_ context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
		attrTypes := []string{
			"http_configuration",
			"email_configuration",
			"slack_configuration",
		}

		generalAttrTypes := []string{
			"general_int_configuration",
			"general_bool_configuration",
			"general_float_configuration",
			"general_string_configuration",
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
			return nil, fmt.Errorf("failed to marshal attributes '%s' as JSON: %w", keyConfiguration, err)
		}
		rawState[keyConfiguration] = string(b)

		streamID := rawState[keyStreamID].(string)
		alarmCallbackID := rawState[keyID].(string)

		rawState[keyAlarmCallbackID] = alarmCallbackID
		rawState[keyID] = streamID + "/" + alarmCallbackID

		return rawState, nil
	},
}
