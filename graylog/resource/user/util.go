package user

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/convert"
)

const (
	keyUsername      = "username"
	keyPermissions   = "permissions"
	keyClientAddress = "client_address"
	keyExternal      = "external"
	keyLastActivity  = "last_activity"
	keyUserID        = "user_id"
	keySessionActive = "session_active"
	keyReadOnly      = "read_only"
)

func getDataFromResourceData(d *schema.ResourceData) (map[string]interface{}, error) {
	data, err := convert.GetFromResourceData(d, Resource())
	if err != nil {
		return nil, err
	}

	delete(data, keyClientAddress)
	delete(data, keyExternal)
	delete(data, keyLastActivity)
	delete(data, keySessionActive)
	delete(data, keyReadOnly)
	delete(data, keyUserID)

	return data, nil
}

func setDataToResourceData(d *schema.ResourceData, data map[string]interface{}) error {
	if err := convert.SetResourceData(d, Resource(), data); err != nil {
		return err
	}
	d.SetId(data[keyUsername].(string))
	return nil
}
