package position

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/convert"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/util"
)

const (
	keyDashboardID = "dashboard_id"
	keyPositions   = "positions"
	keyWidgetID    = "widget_id"
	keyID          = "id"
)

func getDataFromResourceData(d *schema.ResourceData) (map[string]interface{}, error) {
	data, err := convert.GetFromResourceData(d, Resource())
	if err != nil {
		return nil, err
	}

	if err := convert.JSONToData(data, keyPositions); err != nil {
		return nil, err
	}

	p := data[keyPositions].(map[string]interface{})

	data[keyPositions] = convert.MapToList(p, keyID)

	return data, nil
}

func setDataToResourceData(d *schema.ResourceData, data map[string]interface{}) error {
	if err := convert.DataToJSON(data, keyPositions); err != nil {
		return err
	}

	v, ok := util.RenameKey(data, keyID, keyDashboardID)
	if err := convert.SetResourceData(d, Resource(), data); err != nil {
		return err
	}
	if ok {
		d.SetId(v.(string))
	}

	return nil
}
