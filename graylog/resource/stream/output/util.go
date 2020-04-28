package output

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/convert"
)

const (
	keyStreamID  = "stream_id"
	keyOutputIDs = "output_ids"
	keyOutputs   = "outputs"
	keyID        = "id"
)

func setDataToResourceData(d *schema.ResourceData, data map[string]interface{}) error {
	list := data[keyOutputs].([]interface{})
	outputIDs := make([]interface{}, len(list))
	for i, a := range list {
		elem := a.(map[string]interface{})
		outputIDs[i] = elem[keyID]
	}

	if err := convert.SetResourceData(d, Resource(), map[string]interface{}{
		keyStreamID:  d.Id(),
		keyOutputIDs: outputIDs,
	}); err != nil {
		return err
	}

	return nil
}
