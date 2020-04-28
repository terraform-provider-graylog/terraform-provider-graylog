package staticfield

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/suzuki-shunsuke/terraform-provider-graylog/graylog/convert"
)

const (
	KeyInputID      = "input_id"
	KeyStaticFields = "static_fields"
	KeyFields       = "fields"
	KeyID           = "id"
)

func getDataFromResourceData(d *schema.ResourceData) (map[string]interface{}, string, error) {
	data, err := convert.GetFromResourceData(d, Resource())
	if err != nil {
		return nil, "", err
	}
	return data, d.Id(), nil
}

func setDataToResourceData(d *schema.ResourceData, data map[string]interface{}) error {
	if err := convert.SetResourceData(d, Resource(), data); err != nil {
		return err
	}

	d.SetId(data[KeyInputID].(string))
	return nil
}
