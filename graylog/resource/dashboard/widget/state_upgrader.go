package widget

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/util"
)

var stateUpgraderV1 = schema.StateUpgrader{
	Version: 0,
	Type:    util.UpgraderType(),
	Upgrade: func(rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
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
