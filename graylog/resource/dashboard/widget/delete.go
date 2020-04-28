package widget

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

	dsID := d.Get(keyDashboardID).(string)
	widgetID := d.Get(keyWidgetID).(string)
	if _, err := cl.DashboardWidget.Delete(ctx, dsID, widgetID); err != nil {
		return fmt.Errorf(
			"failed to delete a dashboard widget(dashboard id: %s widget id: %s): %w",
			dsID, widgetID, err)
	}
	return nil
}
