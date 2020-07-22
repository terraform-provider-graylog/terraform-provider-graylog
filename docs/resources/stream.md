---
page_title: "Graylog: graylog_stream"
---

# Resource: graylog_stream

* [Example](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/examples/v0.12/stream.tf)
* [Source Code](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/graylog/resource/stream/resource.go)

## Argument Reference

### Required Argument

name | type
--- | ---
title | string
index_set_id | string

### Optional Argument

name | default | type
--- | --- | ---
disabled | | bool
matching_type | | string
description | | string
remove_matches_from_default_stream | | bool
is_default | | bool

## Attributes Reference

name | type | etc
--- | --- | ---
creator_user_id | string | computed
created_at | string | computed

## Import

`graylog_stream` can be imported using the Stream id, e.g.

```console
$ terraform import graylog_stream.test 5c4acaefc9e77bbbbbbbbbbb
```
