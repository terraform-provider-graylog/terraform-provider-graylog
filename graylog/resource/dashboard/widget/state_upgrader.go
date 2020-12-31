package widget

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

func widgetResourceV0() *schema.Resource {
	return &schema.Resource{}
}

var stateUpgraderV1 = schema.StateUpgrader{
	Version: 0,
	Type:    widgetResourceV0().CoreConfigSchema().ImpliedType(),
	Upgrade: func(_ context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
		dashboardID := rawState[keyDashboardID].(string)
		widgetID := rawState[keyID].(string)

		rawState[keyWidgetID] = widgetID
		rawState[keyID] = dashboardID + "/" + widgetID

		attrTypes := []string{
			"stream_search_result_count_configuration",
			"quick_values_configuration",
			"quick_values_histogram_configuration",
			"search_result_chart_configuration",
			"field_chart_configuration",
			"stats_count_configuration",
		}

		if a, ok := rawState["json_configuration"]; ok {
			if a != "" {
				rawState[keyConfig] = a
				return rawState, nil
			}
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

		b, err := json.Marshal(attrs)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal attributes '%s' as JSON: %w", keyConfig, err)
		}
		rawState[keyConfig] = string(b)

		return rawState, nil
	},
}
