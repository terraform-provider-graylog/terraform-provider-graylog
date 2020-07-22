---
page_title: "Graylog: graylog_index_set"
---

# graylog_index_set Data Source

* [Example](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/examples/v0.12/index_set.tf)
* [Source Code](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/graylog/datasource/system/indices/indexset/data_source.go)

## Required Argument

One of `index_set_id` or `title` or `index_prefix` must be set.

## Attributes

name | type
--- | ---
title | string
index_prefix | string
rotation_strategy_class | string
rotation_strategy | JSON string
retention_strategy_class | string
retention_strategy | JSON string
index_analyzer | string
shards | int
index_optimization_max_num_segments | int
description | string
replicas | int
index_optimization_disabled | bool
writable | bool
default | bool
creation_date | string
id | string
