# Resource: graylog_dashboard

* [Example](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/examples/v0.12/dashboard.tf)
* [Source Code](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/graylog/resource/dashboard/resource.go)

## Argument Reference

* `title` - (Required) Dashboard title. The data type is `string`.
* `description` - (Required) Dashboard description. The data type is `string`.

## Attrs Reference

* `created_at` - The date time when the Dashboard is created. The data type is `string`.

## Import

`graylog_dashboard` can be imported using the Dashboard id, e.g.

```console
$ terraform import graylog_dashboard.test 5c4acaefc9e77bbbbbbbbbbb
```
