---
page_title: "Graylog: graylog_input"
---

# Resource: graylog_input

* [Example](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/examples/v0.12/input.tf)
* [Source Code](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/graylog/resource/system/input/resource.go)

## Argument Reference

### Required Argument

name | type
--- | ---
title | string
type | string
attributes | JSON string

### Optional Argument

name | default | type
--- | --- | ---
global | false | bool
node | "" | string

## Attributes Reference

name | type
--- | ---
created_at | string
creator_user_id | string

## Import

`graylog_input` can be imported using the Input id, e.g.

```console
$ terraform import graylog_input.test 5c4acaefc9e77bbbbbbbbbbb
```
