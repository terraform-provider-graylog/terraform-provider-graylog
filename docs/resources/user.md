# Resource: graylog_user

* [Example](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/examples/v0.12/user.tf)
* [Source Code](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/graylog/resource/user/resource.go)

## Argument Reference

* `username` - (Required, Forces new resource) The data type is `string`.
* `email` - (Required) The data type is `string`.
* `full_name` - (Required) The data type is `string`.
* `password` - (Optonal, Sensitive) The data type is `string`.
* `roles` - (Optional) The data type is `set of string`.
* `timezone` - (Optional, Computed) The data type is `string`.
* `session_timeout_ms` - (Optional) The data type is `int`. The default value is `3600000`.

### password

`password` is required to create a resource.
Once the user is created, `password` is optional.

## Attributes Reference

* `user_id` - The data type is `string`.
* `read_only` - The data type is `bool`.
* `external` - The data type is `bool`.
* `client_address` - The data type is `string`.
* `session_active` - The data type is `bool`.
* `last_activity` - The data type is `string`.
* `permissions` - The data type is `set of string`.

## Import

`graylog_user` can be imported using the User `username`, e.g.

```
$ terraform import graylog_user.foo foo
```
