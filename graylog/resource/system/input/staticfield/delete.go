package staticfield

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/suzuki-shunsuke/terraform-provider-graylog/graylog/client"
)

func destroy(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := client.New(m)
	if err != nil {
		return err
	}

	data, _, err := getDataFromResourceData(d)
	if err != nil {
		return err
	}

	fields := data[KeyFields].(map[string]interface{})
	inputID := data[KeyInputID].(string)

	for k := range fields {
		if _, err := cl.InputStaticField.Delete(ctx, inputID, k); err != nil {
			return fmt.Errorf(
				"failed to delete a input static field (input id: %s, key: %s): %w",
				inputID, k, err)
		}
	}

	return nil
}
