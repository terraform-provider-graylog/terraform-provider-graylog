package dashboard

import (
	"errors"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/convert"
)

const (
	keyID          = "id"
	keyCreatedAt   = "created_at"
	keyDashboardID = "dashboard_id"
)

func getDataFromResourceData(d *schema.ResourceData) (map[string]interface{}, error) {
	data, err := convert.GetFromResourceData(d, Resource())
	if err != nil {
		return nil, err
	}
	delete(data, keyCreatedAt)
	return data, nil
}

func setDataToResourceData(d *schema.ResourceData, data map[string]interface{}) error {
	if err := convert.SetResourceData(d, Resource(), data); err != nil {
		return err
	}

	a, ok := data[keyID]
	if !ok {
		return errors.New("failed to set id. 'id' isn't found")
	}
	dID, ok := a.(string)
	if !ok {
		return fmt.Errorf("'id' should be string: %v", a)
	}

	d.SetId(dID)
	return nil
}
