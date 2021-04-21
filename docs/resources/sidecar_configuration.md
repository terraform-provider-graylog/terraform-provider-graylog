# Resource: graylog_sidecar_configuration

* [Example](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/examples/sidecar_configuration.tf)
* [Source Code](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/graylog/resource/sidecar/configuration/resource.go)

## Argument Reference

* `name` - (Required) The name of the Sidecar Configuration. The data type is `string`.
* `color` - (Required) The data type is `string`.
* `collector_id` - (Required) The data type is `string`.
* `template` - (Required) The data type is `string`.

## Attributes Reference

None.

## Import

`graylog_sidecar_configuration` can be imported using the Collector id, e.g.

```console
$ terraform import graylog_sidecar_configuration.test 5c4acaefc9e77bbbbbbbbbbb
```
