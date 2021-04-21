# Resource: graylog_dashboard_widget

* [Example](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/examples/dashboard.tf)
* [Source Code](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/graylog/resource/dashboard/widget/resource.go)

## Argument Reference

* `type` - (Required) widget type. The data type is `string`.
* `description` - (Required) description of the widget. The data type is `string`.
* `dashboard_id` - (Required, Forces new resource) id of the dashboard which the widget is associated with. The data type is `string`.
* `configuration` - (Required) configuration of the widget. The data type is `JSON string`.
* `cache_time` - (Optional) The data type is `int`.

## Attributes Reference

* `creator_user_id` - The user id who created the widget. The data type is `string`.

## Import

`graylog_dashboard_widget` can be imported using `<Dashboard id>/<Widget id>`, e.g.

```console
$ terraform import graylog_dashboard_widget.test 5c4acaefc9e77bbbbbbbbbbb/5c4acaefc9e77bbbbbbbbbbb
```
