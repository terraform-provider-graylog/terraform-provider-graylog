package widget

import (
	"context"
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

	dsID := data[keyDashboardID].(string)
	delete(data, keyID)
	delete(data, keyDashboardID)
	delete(data, keyWidgetID)
	delete(data, keyCreatorUserID)

	dw, _, err := cl.DashboardWidget.Create(ctx, dsID, data)
	if err != nil {
		return fmt.Errorf("failed to create a dashboard widget(dashboard id: %s): %w", dsID, err)
	}
	widgetID := dw[keyWidgetID].(string)
	if err := d.Set(keyWidgetID, widgetID); err != nil {
		return err
	}
	d.SetId(dsID + "/" + widgetID)
	return nil
}
