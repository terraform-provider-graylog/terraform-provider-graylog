# graylog_stream Data Source

* [Example](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/examples/v0.12/stream.tf)
* [Source Code](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/graylog/datasource/stream/data_source.go)

## Argument Reference

One of `stream_id` or `title` must be set.
If `title` is specified, the title must be unique in all streams.

## Attributes Reference

* `title` - The title of the Stream. The data type is `string`.
* `index_set_id` - The id of the Index Set which the Stream is associated with. The data type is `string`.
* `disabled` - The data type is `bool`.
* `matching_type` - The data type is `string`.
* `description` - The data type is `string`.
* `remove_matches_from_default_stream` - The data type is `bool`.
* `is_default` - The data type is `bool`.
* `creator_user_id` - The user id who created the Stream. The data type is `string`.
* `created_at` - The date time when the Stream is created. The data type is `string`.
* `stream_id` - The id of the Stream. The data type is `string`.
