---
page_title: "Graylog: graylog_sidecar_collector"
---

# Resource: graylog_sidecar_collector

* [Example](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/examples/v0.12/sidecar_collector.tf)
* [Source Code](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/graylog/resource/sidecar/collector/resource.go)

## Argument Reference

### Required Argument

name | type
--- | ---
name | string
service_type | string
node_operating_system | string
executable_path | string

### Optional Argument

name | type | default
--- | --- | ---
execute_parameters | string | ""
validation_parameters | string | ""
default_template | string | ""

## Attrs Reference

None.

## Import

`graylog_sidecar_collector` can be imported using the Collector id, e.g.

```console
$ terraform import graylog_sidecar_collector.test 5c4acaefc9e77bbbbbbbbbbb
```

