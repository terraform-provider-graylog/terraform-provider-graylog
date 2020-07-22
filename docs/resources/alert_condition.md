---
page_title: "Graylog: graylog_alert_condition"
---

# Resource: graylog_alert_condition

* [Example](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/examples/v0.12/alert_condition.tf)
* [Source Code](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/graylog/resource/stream/alert/condition/resource.go)

## Argument Reference

* `type` - (Required) Alert Condition type. The data type is `string`.
* `title` - (Required) Alert Condition title. The data type is `string`.
* `stream_id` - (Required) Stream id which the Alert Condition is associated with. The data type is `string`.
* `parameters` - (Required) parameters of Alert Condition. The data type is `JSON string`.
* `in_grace` - (Opitonal) The data type is `bool`. The default value is `false`.

### parameters

`parameters` is a JSON string whose type is `object`.
The data structure of JSON is different per AlertCondition `type`.
Please see the [example](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/examples/v0.12/alert_condition.tf).

### Attributes Reference

* `alert_condition_id` - Alert Condition id. The data type is `string`.

## Import

`graylog_alert_condition` can be imported using the User `<stream id>/<alert condition id>`, e.g.

```console
$ terraform import graylog_alert_condition.test 5bb1b4b5c9e77bbbbbbbbbbb/5c4acaefc9e77bbbbbbbbbbb
```
