package output

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/client"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/convert"
)

func update(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := client.New(m)
	if err != nil {
		return err
	}

	streamID := d.Get(keyStreamID).(string)
	oldV, newV := d.GetChange(keyOutputIDs)

	oldVSet := oldV.(*schema.Set)
	newVSet := newV.(*schema.Set)

	for _, v := range oldVSet.Difference(newVSet).List() {
		oID := v.(string)
		if _, err := cl.StreamOutput.Delete(ctx, streamID, oID); err != nil {
			return fmt.Errorf(
				"failed to delete a relation of stream %s and output %s: %w", streamID, oID, err)
		}
	}

	newList := newVSet.Difference(oldVSet).List()
	if len(newList) != 0 {
		if _, err := cl.StreamOutput.AssociateOutputsWithStream(ctx, streamID, convert.InterfaceListToStringList(newList)); err != nil {
			return fmt.Errorf("failed to associate outputs to stream %s: %w", streamID, err)
		}
	}
	d.SetId(streamID)

	return nil
}
