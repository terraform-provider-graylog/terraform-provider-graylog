---
page_title: "Graylog: graylog_event_notification"
---

# Resource: graylog_event_notification

* [Example](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/examples/v0.12/event_notification.tf)
* [Source Code](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/graylog/resource/event/notification/resource.go)

## Argument Reference

* `title` - (Required) The title of the Event Notification. The data type is `string`.
* `config` - (Required) the configuration of the Event Notification. The data type is `JSON string`.
* `description` - (Optional) the description of the Event Notification. The data type is `string`.

### config

`config` is a JSON string.
The format of `config` depends on the Event Notification type.
Please see the [example](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/examples/v0.12/event_notification.tf).
Using the [Graylog's API browser](https://docs.graylog.org/en/latest/pages/configuration/rest_api.html) you can check the format of `config`.

## Attributes Reference

None.

## Import

`graylog_event_notification` can be imported using the Event Notification id, e.g.

```console
$ terraform import graylog_event_notification.test 5c4acaefc9e77bbbbbbbbbbb
```
