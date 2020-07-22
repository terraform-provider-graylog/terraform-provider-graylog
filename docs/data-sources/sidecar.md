---
page_title: "Graylog: graylog_sidecar"
---

# graylog_sidecar Data Source

* [Source Code](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/graylog/datasource/sidecar/data_source.go)

## Example Usage

```tf
data "graylog_sidecar" "test" {
  node_name = "test"
}
```

## Argument Reference

One of `node_id` or `node_name` must be set.

## Attributes Reference

* `node_id` - The data type is `string`.
* `node_name` - The data type is `string`.
