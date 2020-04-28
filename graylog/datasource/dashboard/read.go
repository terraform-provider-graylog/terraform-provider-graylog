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
		if _, ok := d.GetOk("title"); ok {
			return errors.New("both dashboard_id and title must not be set")
		}
		data, _, err := cl.Dashboard.Get(ctx, id.(string))
		if err != nil {
			return err
		}
		return setDataToResourceData(d, data)
	}

	if t, ok := d.GetOk("title"); ok {
		title := t.(string)
		dashboards, _, err := cl.Dashboard.Gets(ctx)
		if err != nil {
			return err
		}
		cnt := 0
		var data map[string]interface{}
		for _, a := range dashboards["dashboards"].([]interface{}) {
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
	return errors.New("one of dashboard_id or title must be set")
}
