---
page_title: "Graylog: graylog_index_set"
---

# Resource: graylog_index_set

* [Example](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/examples/v0.12/index_set.tf)
* [Source Code](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/graylog/resource/system/indices/indexset/resource.go)

## Argument Reference

### Required Argument

name | type | etc
--- | --- | ---
title | string |
index_prefix | string | `force new`
rotation_strategy_class | string |
rotation_strategy | JSON string |
retention_strategy_class | string |
retention_strategy | JSON string |
index_analyzer | string |
shards | int |
index_optimization_max_num_segments | int |

### Optional Argument

name | default | type
--- | --- | ---
description | "" | string
replicas | 0 | int
index_optimization_disabled | | bool
writable | | bool
default | | bool

## Attributes Reference

name | type | etc
--- | --- | ---
creation_date | computed | string |

## Import

`graylog_index_set` can be imported using the Index Set id, e.g.

```console
$ terraform import graylog_index_set.test 5c4acaefc9e77bbbbbbbbbbb
```
