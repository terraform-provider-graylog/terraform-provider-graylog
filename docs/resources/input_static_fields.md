---
page_title: "Graylog: graylog_input_static_fields"
---

# Resource: graylog_input_static_fields

* [Example](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/examples/v0.12/input.tf)
* [Source Code](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/graylog/resource/system/input/staticfield/resource.go)

## Argument Reference

* `input_id` - (Required) id of the Input which the static fields are associated with. The data type is `string`.
* `fields` - (Optional) The data type is `map[string]string`.

## Attributes Reference

None.

## Import

`graylog_input_static_fields` can be imported using the Input id, e.g.

```console
$ terraform import graylog_input_static_fields.test 5c4acaefc9e77bbbbbbbbbbb
```
