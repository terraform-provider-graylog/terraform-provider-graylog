package output

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/client"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/util"
)

func read(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := client.New(m)
	if err != nil {
		return err
	}
	data, resp, err := cl.StreamOutput.GetOutputsOfStream(ctx, d.Id())
	if err != nil {
		return util.HandleGetResourceError(d, resp, err)
	}
	return setDataToResourceData(d, data)
}
