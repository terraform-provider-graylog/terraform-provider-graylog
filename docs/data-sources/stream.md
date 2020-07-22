---
page_title: "Graylog: graylog_stream"
---

# graylog_stream Data Source

* [Example](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/examples/v0.12/stream.tf)
* [Source Code](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/graylog/datasource/stream/data_source.go)

## Required Argument

One of `stream_id` or `title` must be set.
If `title` is specified, the title must be unique in all streams.

## Attributes

name | type
--- | ---
title | string
stream_id | string
index_set_id | string
disabled | bool
matching_type | string
remove_matches_from_default_stream | bool
is_default | bool
creator_user_id | string
created_at | string
