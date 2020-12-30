# Resource: graylog_stream_rule

* [Example](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/examples/stream_rule.tf)
* [Source Code](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/graylog/resource/stream/rule/resource.go)

## Argument Reference

* `field` - (Required) The data type is `string`.
* `description` - (Required) The data type is `string`.
* `type` - (Required) The data type is `int`.
* `stream_id` - (Required, Forces new resource) The data type is `string`.
* `value` - (Optional) The data type is `string`.
* `inverted` - (Optional) The data type is `bool`.

## Attributes Reference

* `rule_id` - The data type is `string`.

## Import

`graylog_stream_rule` can be imported using `<stream id/stream rule id>`, e.g.

```console
$ terraform import graylog_stream_rule.test 5bb1b4b5c9e77bbbbbbbbbbb/5c4acaefc9e77bbbbbbbbbbb
```
