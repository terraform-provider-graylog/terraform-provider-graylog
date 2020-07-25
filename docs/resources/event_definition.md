# Resource: graylog_event_definition

* [Example](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/examples/v0.12/event_definition.tf)
* [Source Code](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/graylog/resource/event/definition/resource.go)

## Argument Reference

* `title` - (Required) the title of the Event Definition. The data type is `string`.
* `config` - (Required) the configuration of the Event Definition. The data type is `JSON string`.
* `notifications[].notification_id` - (Required) the notification id. The data type is `string`.
* `priority` - (Required) the priority of the Event Definition. The data type is `int`. 1 (Low), 2 (Normal), 3 (High)
* `notification_settings` - (Required) the settings of the Event Definition. The data type is `object`.
* `description` - (Optional) the description of the Event Definition. The data type is `string`.
* `alert` - (Optional) The data type is `string`.
* `field_spec` - (Optional) The data type is `JSON string`.
* `key_spec` - (Optional) The data type is `[]string`. The default value is `[]`.
* `notification_settings.grace_period_ms` - (Optional) The data type is `int`.
* `notification_settings.backlog_size` - (Optional) The data type is `int`.
* `notifications` - (Optional) The data type is `[]object`. The default value is `[]`.

### config

The format of `config` depends on the Event Notification type.
Please see the [example](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/examples/v0.12/event_definition.tf).
Using the [Graylog's API browser](https://docs.graylog.org/en/3.1/pages/configuration/rest_api.html) you can check the format of `config`.

## Attribute Reference

None.

## Import

`graylog_event_definition` can be imported using the Event Definition id, e.g.

```console
$ terraform import graylog_event_definition.test 5c4acaefc9e77bbbbbbbbbbb
```
