# Resource: graylog_role

* [Example](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/examples/role.tf)
* [Source Code](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/graylog/resource/role/resource.go)

## Argument Reference

* `name` - (Required) The Role name. The data type is `string`.
* `permissions` - (Required) The permissions of the Role. The data type is `[]string`.
* `description` - (Optional) description of the Role. The data type is `string`.

## Attributes Reference

* `read_only` - The data type is `bool`.

## Import

`graylog_role` can be imported using the Role name, e.g.

```
$ terraform import graylog_role.foo foo
```
