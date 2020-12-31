package stream

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/convert"
)

const (
	keyID            = "id"
	keyStreamID      = "stream_id"
	keyDisabled      = "disabled"
	keyCreatorUserID = "creator_user_id"
	keyCreatedAt     = "created_at"
	keyIsDefault     = "is_default"
)

func getDataFromResourceData(d *schema.ResourceData) (map[string]interface{}, error) {
	data, err := convert.GetFromResourceData(d, Resource())
	if err != nil {
		return nil, err
	}
	delete(data, keyCreatedAt)
	delete(data, keyCreatorUserID)
	delete(data, keyIsDefault)
	return data, nil
}

func setDataToResourceData(d *schema.ResourceData, data map[string]interface{}) error {
	if err := convert.SetResourceData(d, Resource(), data); err != nil {
		return err
	}

	d.SetId(data[keyID].(string))
	return nil
}
