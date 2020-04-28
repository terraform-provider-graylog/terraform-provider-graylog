package indexset

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

	is, _, err := cl.IndexSet.Create(ctx, data)
	if err != nil {
		return fmt.Errorf("failed to create a index set: %w", err)
	}
	d.SetId(is[keyID].(string))
	return nil
}
