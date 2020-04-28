package view

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/convert"
)

const (
	keyID            = "id"
	keyWidgetMapping = "widget_mapping"
	keyPositions     = "positions"
	keyState         = "state"
	keyWidgets       = "widgets"
	keyConfig        = "config"
	keyTimerange     = "timerange"
	keyWidgetID      = "widget_id"
)

func getDataFromResourceData(d *schema.ResourceData) (map[string]interface{}, error) {
	data, err := convert.GetFromResourceData(d, Resource())
	if err != nil {
		return nil, err
	}

	state := data[keyState].(map[string]interface{})

	if err := convert.ConvertJSONToData(state, keyWidgetMapping, keyPositions); err != nil {
		return nil, err
	}

	widgets := convert.ConvertListToMap(state[keyWidgets].([]interface{}), keyWidgetID)
	for k, a := range widgets {
		widget := a.(map[string]interface{})
		if err := convert.ConvertJSONToData(widget, keyConfig, keyTimerange); err != nil {
			return nil, err
		}
		widgets[k] = widget
	}
	state[keyWidgets] = widgets
	data[keyState] = state
	return data, nil
}

func setDataToResourceData(d *schema.ResourceData, data map[string]interface{}) error {
	state := data[keyState].(map[string]interface{})

	if err := convert.ConvertDataToJSON(state, keyWidgetMapping, keyPositions); err != nil {
		return err
	}

	widgets := convert.ConvertMapToList(state[keyWidgets].(map[string]interface{}), keyWidgetID)
	for i, a := range widgets {
		widget := a.(map[string]interface{})
		if err := convert.ConvertDataToJSON(widget, keyConfig, keyTimerange); err != nil {
			return err
		}
		widgets[i] = widget
	}

	state[keyWidgets] = widgets

	data[keyState] = state

	if err := convert.SetResourceData(d, Resource(), data); err != nil {
		return err
	}

	d.SetId(data[keyID].(string))
	return nil
}
