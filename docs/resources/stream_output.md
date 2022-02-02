# Resource: graylog_stream_output

* [Example](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/examples/stream_output.tf)
* [Source Code](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/graylog/resource/stream/output/resource.go)

## Argument Reference

* `stream_id` - (Required, Forces new resource) Stream id which the Stream Output is associated with. The data type is `string`.
* `output_ids` - (Required) The data type is `[]string`.

### Note

This resource treats the stream id as the resource id.
So please make the stream id unique in all `graylog_stream_output` resources.

## Attributes Reference

None.

## Import

`graylog_stream_output` can be imported using the Stream id, e.g.

```console
$ terraform import graylog_stream_output.test 5c4acaefc9e77bbbbbbbbbbb
```
