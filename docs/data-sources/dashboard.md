---
page_title: "Graylog: graylog_dashboard"
---

# graylog_dashboard Data Source

* [Source Code](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/graylog/datasource/dashboard/resource.go)

## Example Usage

```tf
data "graylog_dashboard" "test" {
  title = "test"
}
```

## Argument Reference

One of `dashboard_id` or `title` must be set.

## Attributes Reference

* `title` - The data type is `string`
* `dashboard_id` - The data type is `string`
* `description` - The data type is `string`
