package indexset

import (
	"context"
	"errors"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/suzuki-shunsuke/terraform-provider-graylog/graylog/client"
)

func read(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := client.New(m)
	if err != nil {
		return err
	}

	if id, ok := d.GetOk("index_set_id"); ok {
		if _, ok := d.GetOk("title"); ok {
			return errors.New("only one of index_set_id or title or index_prefix must be set")
		}
		if _, ok := d.GetOk("index_prefix"); ok {
			return errors.New("only one of index_set_id or title or index_prefix must be set")
		}
		data, _, err := cl.IndexSet.Get(ctx, id.(string))
		if err != nil {
			return err
		}
		return setDataToResourceData(d, data)
	}

	if t, ok := d.GetOk("title"); ok {
		if _, ok := d.GetOk("index_prefix"); ok {
			return errors.New("only one of index_set_id or title or index_prefix must be set")
		}
		title := t.(string)
		indexSets, _, err := cl.IndexSet.Gets(ctx, nil)
		if err != nil {
			return err
		}
		cnt := 0
		var data map[string]interface{}
		for _, is := range indexSets["index_sets"].([]interface{}) {
			a := is.(map[string]interface{})
			if a["title"].(string) == title {
				data = a
				cnt++
				if cnt > 1 {
					return errors.New("title isn't unique")
				}
			}
		}
		if cnt == 0 {
			return errors.New("matched index set is not found")
		}
		return setDataToResourceData(d, data)
	}

	if p, ok := d.GetOk("index_prefix"); ok {
		prefix := p.(string)

		indexSets, _, err := cl.IndexSet.Gets(ctx, nil)
		if err != nil {
			return err
		}
		for _, is := range indexSets["index_sets"].([]interface{}) {
			data := is.(map[string]interface{})
			if data["index_prefix"].(string) == prefix {
				return setDataToResourceData(d, data)
			}
		}
		return errors.New("matched index prefix is not found")
	}
	return errors.New("one of index_set_id or title or prefix must be set")
}
