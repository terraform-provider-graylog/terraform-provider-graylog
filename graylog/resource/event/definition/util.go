package definition

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/suzuki-shunsuke/terraform-provider-graylog/graylog/convert"
)

const (
	keyID        = "id"
	keyConfig    = "config"
	keyFieldSpec = "field_spec"
)

func getDataFromResourceData(d *schema.ResourceData) (map[string]interface{}, error) {
	data, err := convert.GetFromResourceData(d, Resource())
	if err != nil {
		return nil, err
	}

	if err := convert.ConvertJSONToData(data, keyConfig, keyFieldSpec); err != nil {
		return nil, err
	}

	return data, nil
}

func setDataToResourceData(d *schema.ResourceData, data map[string]interface{}) error {
	if err := convert.ConvertDataToJSON(data, keyConfig, keyFieldSpec); err != nil {
		return err
	}

	if err := convert.SetResourceData(d, Resource(), data); err != nil {
		return err
	}

	d.SetId(data[keyID].(string))
	return nil
}
