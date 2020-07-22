---
page_title: "Graylog: graylog_role"
---

# Resource: graylog_role

* [Example](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/examples/v0.12/role.tf)
* [Source Code](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/graylog/resource/role/resource.go)

## Argument Reference

### Required Argument

name | type
--- | ---
name | string
permissions | []string

### Optional Argument

name | default | type
--- | --- | --- |
description | `""` | string

## Attributes Reference

name | type
--- | ---
read_only | bool

## Import

`graylog_role` can be imported using the Role name, e.g.

```
$ terraform import graylog_role.foo foo
```
