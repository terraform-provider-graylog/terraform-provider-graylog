package dashboard

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
	delete(data, keyCreatedAt)
	if _, err := cl.Dashboard.Update(ctx, d.Id(), data); err != nil {
		return fmt.Errorf("failed to update a dashboard %s: %w", d.Id(), err)
	}
	return nil
}
