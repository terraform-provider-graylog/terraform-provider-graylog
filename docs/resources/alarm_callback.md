# Resource: graylog_alarm_callback

* [Example](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/examples/v0.12/alarm_callback.tf)
* [Source Code](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/graylog/resource/stream/alarmcallback/resource.go)

## Argument Reference

* `type` - (Required) Alarm Callback type. The data type is `string`.
* `title` - (Required) Alarm Callback title. The data type is `string`.
* `stream_id` - (Required, Forces new resource) Stream id which the Alarm Callback is associated with. The data type is `string`.
* `configuration` - (Required) configuration of Alarm Callback. The data type is `JSON string`.

### configuration

`configuration` is a JSON string whose type is `object`.
The data structure of JSON is different per AlarmCallback `type`.
Please see the [example](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/examples/v0.12/alarm_callback.tf).

## Attributes Reference

* `alarmcallback_id` - Alarm Callback id. The data type is `string`.

## Import

`graylog_alarm_callback` can be imported using the User `<stream id>/<alarm callback id>`, e.g.

```console
$ terraform import graylog_alarm_callback.test 5bb1b4b5c9e77bbbbbbbbbbb/5c4acaefc9e77bbbbbbbbbbb
```
