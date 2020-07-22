---
page_title: "Graylog: graylog_dashboard"
---

# Resource: graylog_dashboard

* [Example](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/examples/v0.12/dashboard.tf)
* [Source Code](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/graylog/resource/dashboard/resource.go)

## Argument Reference

### Required Argument

name | type
--- | ---
title | string
description | string

### Optional Argument

None

## Attrs Reference

name | type
--- | ---
created_at | string

## Import

`graylog_dashboard` can be imported using the Dashboard id, e.g.

```console
$ terraform import graylog_dashboard.test 5c4acaefc9e77bbbbbbbbbbb
```
