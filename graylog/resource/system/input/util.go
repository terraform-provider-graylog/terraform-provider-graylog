package input

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/suzuki-shunsuke/go-dataeq/dataeq"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/convert"
)

const (
	keyID            = "id"
	keyAttributes    = "attributes"
	keyCreatedAt     = "created_at"
	keyCreatorUserID = "creator_user_id"
)

func getDataFromResourceData(d *schema.ResourceData) (map[string]interface{}, error) {
	data, err := convert.GetFromResourceData(d, Resource())
	if err != nil {
		return nil, err
	}

	attrS := d.Get(keyAttributes).(string)
	attr, err := dataeq.JSON.ConvertByte([]byte(attrS))
	if err != nil {
		return nil, fmt.Errorf("failed to parse the 'attributes'. 'attributes' must be a JSON string: %w", err)
	}
	data[keyAttributes] = attr

	delete(data, keyCreatedAt)
	delete(data, keyCreatorUserID)

	return data, nil
}

func setDataToResourceData(d *schema.ResourceData, data map[string]interface{}) error {
	attrS, err := json.Marshal(data[keyAttributes])
	if err != nil {
		return fmt.Errorf("failed to marshal the 'attributes' as JSON: %w", err)
	}
	data[keyAttributes] = string(attrS)

	if err := convert.SetResourceData(d, Resource(), data); err != nil {
		return err
	}

	d.SetId(data[keyID].(string))
	return nil
}
