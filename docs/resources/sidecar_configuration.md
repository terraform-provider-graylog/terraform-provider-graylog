---
page_title: "Graylog: graylog_sidecar_configuration"
---

# Resource: graylog_sidecar_configuration

* [Example](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/examples/v0.12/sidecar_configuration.tf)
* [Source Code](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/graylog/resource/sidecar/configuration/resource.go)

## Argument Reference

### Required Argument

name | type
--- | ---
name | string
color | string
collector_id | string
template | string

### Optional Argument

None.

## Attrs Reference

None.

## Import

`graylog_sidecar_configuration` can be imported using the Collector id, e.g.

```console
$ terraform import graylog_sidecar_configuration.test 5c4acaefc9e77bbbbbbbbbbb
```
