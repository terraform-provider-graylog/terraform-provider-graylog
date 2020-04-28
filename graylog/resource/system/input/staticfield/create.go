package staticfield

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/suzuki-shunsuke/terraform-provider-graylog/graylog/client"
)

func create(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := client.New(m)
	if err != nil {
		return err
	}
	data, _, err := getDataFromResourceData(d)
	if err != nil {
		return err
	}

	newFields := data[KeyFields].(map[string]interface{})
	inputID := data[KeyInputID].(string)

	input, _, err := cl.Input.Get(ctx, inputID)
	if err != nil {
		return err
	}
	oldFields := input[KeyStaticFields].(map[string]interface{})
	if oldFields == nil {
		oldFields = map[string]interface{}{}
	}

	for k, v := range newFields {
		if oldV, ok := oldFields[k]; ok {
			if v == oldV {
				continue
			}
		}
		value := v.(string)
		if _, err := cl.InputStaticField.Create(ctx, inputID, k, value); err != nil {
			return fmt.Errorf(
				"failed to create a input static field (input id: %s, key: %s, value: %s): %w",
				inputID, k, value, err)
		}
	}

	for k := range oldFields {
		if _, ok := newFields[k]; ok {
			continue
		}
		if _, err := cl.InputStaticField.Delete(ctx, inputID, k); err != nil {
			return fmt.Errorf(
				"failed to delete a input static field (input id: %s, key: %s): %w",
				inputID, k, err)
		}
	}

	d.SetId(inputID)
	return nil
}
