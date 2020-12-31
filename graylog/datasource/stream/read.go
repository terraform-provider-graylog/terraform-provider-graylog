package stream

import (
	"context"
	"errors"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/client"
)

func read(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := client.New(m)
	if err != nil {
		return err
	}

	if id, ok := d.GetOk("stream_id"); ok {
		if _, ok := d.GetOk("title"); ok {
			return errors.New("both stream_id and title must not be set")
		}
		data, _, err := cl.Stream.Get(ctx, id.(string))
		if err != nil {
			return err
		}
		return setDataToResourceData(d, data)
	}

	if t, ok := d.GetOk("title"); ok {
		title := t.(string)
		streams, _, err := cl.Stream.Gets(ctx)
		if err != nil {
			return err
		}
		cnt := 0
		var data map[string]interface{}
		for _, a := range streams["streams"].([]interface{}) {
			stream := a.(map[string]interface{})
			if stream["title"].(string) == title {
				data = stream
				cnt++
				if cnt > 1 {
					return errors.New("title isn't unique")
				}
			}
		}
		if cnt == 0 {
			return errors.New("matched stream is not found")
		}
		return setDataToResourceData(d, data)
	}
	return errors.New("one of stream_id or title must be set")
}
