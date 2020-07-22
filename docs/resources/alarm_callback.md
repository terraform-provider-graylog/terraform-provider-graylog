---
page_title: "Graylog: graylog_alarm_callback"
---

# Resource: graylog_alarm_callback

* [Example](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/examples/v0.12/alarm_callback.tf)
* [Source Code](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/graylog/resource/stream/alarmcallback/resource.go)

## Argument Reference

### Required Argument

name | type
--- | ---
type | string
title | string
stream_id | string
configuration | JSON string

### Optional Argument

None

### configuration

`configuration` is a JSON string whose type is `object`.
The data structure of JSON is different per AlarmCallback `type`.
Please see the [example](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/examples/v0.12/alarm_callback.tf).

## Attributes Reference

name | type
--- | ---
alarmcallback_id | string

## Import

`graylog_alarm_callback` can be imported using the User `<stream id>/<alarm callback id>`, e.g.

```console
$ terraform import graylog_alarm_callback.test 5bb1b4b5c9e77bbbbbbbbbbb/5c4acaefc9e77bbbbbbbbbbb
```
