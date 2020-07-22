---
page_title: "Graylog: graylog_dashboard_widget_positions"
---

# Resource: graylog_dashboard_widget_positions

* [Example](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/examples/v0.12/dashboard.tf)
* [Source Code](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/graylog/resource/dashboard/position/resource.go)

## Argument Reference

## Required Argument

name | type
--- | ---
dashboard_id | string
positions | JSON string

### Optional Argument

Nothing.

### Attributes Reference

Nothing.

### Import

`graylog_dashboard_widget_positions` can be imported using the Dashboard id, e.g.

```console
$ terraform import graylog_dashboard_widget_positions.test 5c4acaefc9e77bbbbbbbbbbb
```
