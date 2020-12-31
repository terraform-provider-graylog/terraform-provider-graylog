package sidecar

import (
	"errors"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/convert"
)

const (
	keyNodeID   = "node_id"
	keyNodeName = "node_name"
	keySidecars = "sidecars"
)

var errNodeIDNotFound = errors.New("node_id isn't found")

func setDataToResourceData(d *schema.ResourceData, data map[string]interface{}) error {
	id, ok := data[keyNodeID]
	if !ok {
		return errNodeIDNotFound
	}

	if err := convert.SetResourceData(d, DataSource(), data); err != nil {
		return err
	}

	d.SetId(id.(string))
	return nil
}
