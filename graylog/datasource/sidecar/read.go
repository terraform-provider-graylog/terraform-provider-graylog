package sidecar

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

	if id, ok := d.GetOk(keyNodeID); ok {
		if _, ok := d.GetOk(keyNodeName); ok {
			return errors.New("both node_id and node_name must not be set")
		}
		data, _, err := cl.Sidecar.Get(ctx, id.(string))
		if err != nil {
			return err
		}
		return setDataToResourceData(d, data)
	}

	if t, ok := d.GetOk(keyNodeName); ok {
		nodeName := t.(string)
		sidecars, _, err := cl.Sidecar.GetAll(ctx)
		if err != nil {
			return err
		}
		for _, a := range sidecars[keySidecars].([]interface{}) {
			sidecar := a.(map[string]interface{})
			if sidecar[keyNodeName].(string) == nodeName {
				return setDataToResourceData(d, sidecar)
			}
		}
		return errors.New("matched sidecar is not found")
	}
	return errors.New("one of node_id or node_name must be set")
}
