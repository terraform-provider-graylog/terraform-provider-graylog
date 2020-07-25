# Resource: graylog_dashboard_widget_positions

* [Example](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/examples/v0.12/dashboard.tf)
* [Source Code](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/graylog/resource/dashboard/position/resource.go)

## Argument Reference

* `dashboard_id` - (Required, Forces new resource) id of the dashboard which widgets are associated with. The data type is `string`.
* `positions` - (Required) positions of widgets. The data type is `JSON string`.

### Attributes Reference

Nothing.

### Import

`graylog_dashboard_widget_positions` can be imported using the Dashboard id, e.g.

```console
$ terraform import graylog_dashboard_widget_positions.test 5c4acaefc9e77bbbbbbbbbbb
```
