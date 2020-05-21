package collector

import (
	"context"
	"errors"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/client"
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

	ds, _, err := cl.Collector.Create(ctx, data)
	if err != nil {
		return fmt.Errorf("failed to create a collector: %w", err)
	}

	a, ok := ds[keyID]
	if !ok {
		return errors.New("response body of Graylog API is unexpected. 'id' isn't found")
	}
	dID, ok := a.(string)
	if !ok {
		return fmt.Errorf(
			"response body of Graylog API is unexpected. 'id' should be string: %v", a)
	}

	d.SetId(dID)
	return nil
}
