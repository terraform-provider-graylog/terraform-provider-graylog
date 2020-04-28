package extractor

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

	data, err := getDataFromResourceData(d)
	if err != nil {
		return err
	}

	inputID := data[keyInputID].(string)
	delete(data, keyInputID)
	delete(data, keyID)

	ac, _, err := cl.Extractor.Create(ctx, inputID, data)
	if err != nil {
		return fmt.Errorf("failed to create a extractor (input id: %s): %w", inputID, err)
	}
	acID := ac[keyExtractorID].(string)
	d.Set(keyExtractorID, acID)
	d.SetId(inputID + "/" + acID)
	return nil
}
