---
page_title: "Graylog: graylog_input"
---

# Resource: graylog_input

* [Example](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/examples/v0.12/input.tf)
* [Source Code](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/graylog/resource/system/input/resource.go)

## Argument Reference

* `title` - (Required) the title of the Input. The data type is `string`.
* `type` - (Required) the type of the Input. The data type is `string`.
* `attributes` - (Required) the attributes of the Input. The data type is `JSON string`.
* `global` - (Optional) The data type is `bool`. The default value is `false`.
* `node` - (Optional) The data type is `string`.

## Attributes Reference

* `created_at` - The date time when the Index Set is created. The data type is `string`.
* `creator_user_id` - The user id who created the Input. The data type is `string`.

## Import

`graylog_input` can be imported using the Input id, e.g.

```console
$ terraform import graylog_input.test 5c4acaefc9e77bbbbbbbbbbb
```
