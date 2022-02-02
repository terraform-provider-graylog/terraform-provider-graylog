# Resource: graylog_sidecar_collector

* [Example](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/examples/sidecar_collector.tf)
* [Source Code](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/graylog/resource/sidecar/collector/resource.go)

## Argument Reference

* `name` - (Required) The Sidecar Collector name. The data type is `string`.
* `service_type` - (Required) The data type is `string`.
* `node_operating_system` - (Required) The data type is `string`.
* `executable_path` - (Required) The data type is `string`.
* `execute_parameters` - (Optional) The data type is `string`.
* `validation_parameters` - (Optional) The data type is `string`.
* `default_template` - (Optional) The data type is `string`.

## Attributes Reference

None.

## Import

`graylog_sidecar_collector` can be imported using the Collector id, e.g.

```console
$ terraform import graylog_sidecar_collector.test 5c4acaefc9e77bbbbbbbbbbb
```

