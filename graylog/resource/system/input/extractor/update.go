package extractor

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/suzuki-shunsuke/terraform-provider-graylog/graylog/client"
)

func update(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := client.New(m)
	if err != nil {
		return err
	}

	data, err := getDataFromResourceData(d)
	if err != nil {
		return err
	}

	delete(data, keyInputID)
	delete(data, keyID)

	inputID := d.Get(keyInputID).(string)
	eID := d.Get(keyExtractorID).(string)
	if _, _, err := cl.Extractor.Update(ctx, inputID, eID, data); err != nil {
		return fmt.Errorf("failed to update a extractor (input id: %s, id: %s): %w", inputID, eID, err)
	}
	return nil
}
