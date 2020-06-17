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

	if id, ok := d.GetOk("dashboard_id"); ok {
		return readByID(ctx, d, cl, id.(string))
	}

	if t, ok := d.GetOk("title"); ok {
		return readByTitle(ctx, d, cl, t.(string))
	}
	return errors.New("one of dashboard_id or title must be set")
}

func readByID(ctx context.Context, d *schema.ResourceData, cl *client.Client, id string) error {
	if _, ok := d.GetOk("title"); ok {
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
	cnt := 0
	var data map[string]interface{}
	key := "dashboards"
	if dashboards[key] == nil {
		// Graylog 3.2.2+
		key = "views"
	}
	for _, a := range dashboards[key].([]interface{}) {
		dashboard := a.(map[string]interface{})
		if dashboard["title"].(string) == title {
			data = dashboard
			cnt++
			if cnt > 1 {
				return errors.New("title isn't unique")
			}
		}
	}
	if cnt == 0 {
		return errors.New("matched dashboard is not found")
	}
	return setDataToResourceData(d, data)
}
