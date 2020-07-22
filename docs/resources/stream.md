---
page_title: "Graylog: graylog_stream"
---

# Resource: graylog_stream

* [Example](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/examples/v0.12/stream.tf)
* [Source Code](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/graylog/resource/stream/resource.go)

## Argument Reference

* `title` - (Required) The title of the Stream. The data type is `string`.
* `index_set_id` - (Required) The id of the Index Set which the Stream is associated with. The data type is `string`.
* `disabled` - (Optional) The data type is `bool`.
* `matching_type` - (Optional) The data type is `string`.
* `description` - (Optional) The data type is `string`.
* `remove_matches_from_default_stream` - (Optional) The data type is `bool`.
* `is_default` - (Optional) The data type is `bool`.

## Attributes Reference

* `creator_user_id` - The user id who created the Stream. The data type is `string`.
* `created_at` - The date time when the Stream is created. The data type is `string`.

## Import

`graylog_stream` can be imported using the Stream id, e.g.

```console
$ terraform import graylog_stream.test 5c4acaefc9e77bbbbbbbbbbb
```
