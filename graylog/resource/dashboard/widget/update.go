package widget

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/client"
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

	delete(data, keyDashboardID)
	delete(data, keyWidgetID)
	delete(data, keyCreatorUserID)
	delete(data, keyID)

	dsID := d.Get(keyDashboardID).(string)
	wID := d.Get(keyWidgetID).(string)

	if _, err := cl.DashboardWidget.Update(ctx, dsID, wID, data); err != nil {
		return fmt.Errorf(
			"faield to update a dashboard widget(dashboard id: %s widget id: %s): %w", dsID, wID, err)
	}
	return nil
}
