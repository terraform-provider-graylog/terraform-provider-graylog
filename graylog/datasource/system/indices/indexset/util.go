package indexset

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/convert"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/resource/system/indices/indexset"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/util"
)

func setDataToResourceData(d *schema.ResourceData, data map[string]interface{}) error {
	id, ok := util.RenameKey(data, "id", "index_set_id")
	if err := convert.DataToJSON(data, "rotation_strategy", "retention_strategy"); err != nil {
		return err
	}

	if err := convert.SetResourceData(d, indexset.Resource(), data); err != nil {
		return err
	}

	if ok {
		d.SetId(id.(string))
	}
	return nil
}
