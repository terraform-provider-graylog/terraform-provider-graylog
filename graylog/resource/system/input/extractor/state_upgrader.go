package extractor

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

func extractorResourceV0() *schema.Resource {
	return &schema.Resource{}
}

var stateUpgraderV1 = schema.StateUpgrader{
	Version: 0,
	Type:    extractorResourceV0().CoreConfigSchema().ImpliedType(),
	Upgrade: func(_ context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
		inputID := rawState[keyInputID].(string)
		extractorID := rawState[keyID].(string)

		rawState[keyExtractorID] = extractorID
		rawState[keyID] = inputID + "/" + extractorID

		attrTypes := []string{
			"grok_type_extractor_config",
			"json_type_extractor_config",
			"regex_type_extractor_config",
		}

		generalAttrTypes := []string{
			"general_int_extractor_config",
			"general_bool_extractor_config",
			"general_float_extractor_config",
			"general_string_extractor_config",
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
			return nil, fmt.Errorf("failed to marshal attributes '%s' as JSON: %w", keyConfig, err)
		}
		rawState[keyExtractorConfig] = string(b)

		if a, ok := rawState[keyConverters]; ok {
			list := a.([]interface{})
			for i, e := range list {
				if err := convert.OneSizeListToJSON(e.(map[string]interface{}), keyConfig); err != nil {
					return nil, err
				}
				list[i] = e
			}
			rawState[keyConverters] = list
		}

		return rawState, nil
	},
}
