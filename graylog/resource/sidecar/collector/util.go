package collector

import (
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/convert"
)

const (
	keyID = "id"
)

func getDataFromResourceData(d *schema.ResourceData) (map[string]interface{}, error) {
	return convert.GetFromResourceData(d, Resource())
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
