package setting

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/convert"
)

const (
	ldapSettingID        = "ldap_setting_id" // https://github.com/suzuki-shunsuke/go-graylog/blob/a26bd268c99483896b456e972309aa5ed0d2c88c/graylog/terraform/resource_ldap_setting.go#L14
	keySystemPasswordSet = "system_password_set"
)

func getDataFromResourceData(d *schema.ResourceData) (map[string]interface{}, error) {
	data, err := convert.GetFromResourceData(d, Resource())
	if err != nil {
		return nil, err
	}
	return data, nil
}

func setDataToResourceData(d *schema.ResourceData, data map[string]interface{}) error {
	if err := convert.SetResourceData(d, Resource(), data); err != nil {
		return err
	}

	d.SetId(ldapSettingID)
	return nil
}
