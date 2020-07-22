---
page_title: "Graylog: graylog_stream_rule"
---

# Resource: graylog_stream_rule

* [Example](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/examples/v0.12/stream_rule.tf)
* [Source Code](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/graylog/resource/stream/rule/resource.go)

## Argument Reference

### Required Argument

name | type
--- | ---
field | string
description | string
type | int
stream_id | string

### Optional Argument

name | default | type
--- | --- | ---
value |  | string
inverted | | bool

## Attributes Reference

name | type
--- | ---
rule_id | string

## Import

`graylog_stream_rule` can be imported using `<stream id/stream rule id>`, e.g.

```console
$ terraform import graylog_stream_rule.test 5bb1b4b5c9e77bbbbbbbbbbb/5c4acaefc9e77bbbbbbbbbbb
```
