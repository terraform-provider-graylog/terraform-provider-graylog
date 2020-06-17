package dashboard

import (
	"context"
	"errors"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/client"
)

func read(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := client.New(m)
	if err != nil {
		return err
	}

	if id, ok := d.GetOk(keyDashboardID); ok {
		return readByID(ctx, d, cl, id.(string))
	}

	if t, ok := d.GetOk(keyTitle); ok {
		return readByTitle(ctx, d, cl, t.(string))
	}
	return errors.New("one of dashboard_id or title must be set")
}

func readByID(ctx context.Context, d *schema.ResourceData, cl *client.Client, id string) error {
	if _, ok := d.GetOk(keyTitle); ok {
		return errors.New("both dashboard_id and title must not be set")
	}
	data, _, err := cl.Dashboard.Get(ctx, id)
	if err != nil {
		return err
	}
	return setDataToResourceData(d, data)
}

func readByTitle(ctx context.Context, d *schema.ResourceData, cl *client.Client, title string) error {
	dashboards, _, err := cl.Dashboard.Gets(ctx)
	if err != nil {
		return err
	}
	ds, ok := dashboards[keyDashboards]
	if !ok {
		// Graylog 3.3.0+
		// Legacy Dashboard is migrated to view-based Dashboard
		// https://github.com/Graylog2/graylog2-server/pull/6924
		ds, ok = dashboards[keyViews]
		if !ok {
			return errors.New(`the response of Graylog API GET /api/dashboards is unexpected. The field "dashboards" and "views" aren't found`)
		}
	}
	var data map[string]interface{}
	for _, a := range ds.([]interface{}) {
		dashboard := a.(map[string]interface{})
		if dashboard[keyTitle].(string) == title {
			if data != nil {
				return errors.New("title isn't unique")
			}
			data = dashboard
		}
	}
	if data == nil {
		return errors.New("matched dashboard is not found")
	}
	return setDataToResourceData(d, data)
}
