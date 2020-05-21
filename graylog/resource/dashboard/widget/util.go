package widget

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/convert"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/util"
)

const (
	keyDashboardID   = "dashboard_id"
	keyWidgetID      = "widget_id"
	keyConfig        = "config"
	keyID            = "id"
	keyCreatorUserID = "creator_user_id"
)

func getDataFromResourceData(d *schema.ResourceData) (map[string]interface{}, error) {
	data, err := convert.GetFromResourceData(d, Resource())
	if err != nil {
		return nil, err
	}

	if err := convert.JSONToData(data, keyConfig); err != nil {
		return nil, err
	}
	util.RenameKey(data, keyWidgetID, keyID)

	return data, nil
}

func setDataToResourceData(d *schema.ResourceData, data map[string]interface{}) error {
	if err := convert.DataToJSON(data, keyConfig); err != nil {
		return err
	}
	util.RenameKey(data, keyID, keyWidgetID)

	if err := convert.SetResourceData(d, Resource(), data); err != nil {
		return err
	}

	d.SetId(d.Get(keyDashboardID).(string) + "/" + d.Get(keyWidgetID).(string))
	return nil
}
