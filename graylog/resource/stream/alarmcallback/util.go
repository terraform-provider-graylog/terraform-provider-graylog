package alarmcallback

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/suzuki-shunsuke/terraform-provider-graylog/graylog/convert"
	"github.com/suzuki-shunsuke/terraform-provider-graylog/graylog/util"
)

const (
	keyAlarmCallbackID = "alarmcallback_id"
	keyConfiguration   = "configuration"
	keyID              = "id"
	keyStreamID        = "stream_id"
)

func getDataFromResourceData(d *schema.ResourceData) (map[string]interface{}, error) {
	data, err := convert.GetFromResourceData(d, Resource())
	if err != nil {
		return nil, err
	}
	if err := convert.ConvertJSONToData(data, keyConfiguration); err != nil {
		return nil, err
	}
	util.RenameKey(data, keyAlarmCallbackID, keyID)

	return data, nil
}

func setDataToResourceData(d *schema.ResourceData, data map[string]interface{}) error {
	if err := convert.ConvertDataToJSON(data, keyConfiguration); err != nil {
		return err
	}
	util.RenameKey(data, keyID, keyAlarmCallbackID)

	if err := convert.SetResourceData(d, Resource(), data); err != nil {
		return err
	}

	d.SetId(d.Get(keyStreamID).(string) + "/" + d.Get(keyAlarmCallbackID).(string))
	return nil
}
