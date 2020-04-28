package condition

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/suzuki-shunsuke/terraform-provider-graylog/graylog/convert"
	"github.com/suzuki-shunsuke/terraform-provider-graylog/graylog/util"
)

const (
	keyID               = "id"
	keyStreamID         = "stream_id"
	keyAlertConditionID = "alert_condition_id"
	keyParameters       = "parameters"
	keyInGrace          = "in_grace"
)

func getDataFromResourceData(d *schema.ResourceData) (map[string]interface{}, error) {
	data, err := convert.GetFromResourceData(d, Resource())
	if err != nil {
		return nil, err
	}
	if err := convert.ConvertJSONToData(data, keyParameters); err != nil {
		return nil, err
	}
	util.RenameKey(data, keyAlertConditionID, keyID)

	return data, nil
}

func setDataToResourceData(d *schema.ResourceData, data map[string]interface{}) error {
	if err := convert.ConvertDataToJSON(data, keyParameters); err != nil {
		return err
	}
	util.RenameKey(data, keyID, keyAlertConditionID)

	if err := convert.SetResourceData(d, Resource(), data); err != nil {
		return err
	}

	d.SetId(d.Get(keyStreamID).(string) + "/" + d.Get(keyAlertConditionID).(string))
	return nil
}
