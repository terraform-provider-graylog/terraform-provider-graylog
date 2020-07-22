---
page_title: "Graylog: graylog_user"
---

# Resource: graylog_user

* [Example](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/examples/v0.12/user.tf)
* [Source Code](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/graylog/resource/user/resource.go)

## Argument Reference

### Required Argument

name | type | etc
--- | --- | ---
username | string | force_new
email | string |
full_name | string |

### Optional Argument

name | default | type | etc
--- | --- | --- | ---
password | string | sensitive
roles | `[]` | string set |
timezone | `""` | string | computed
session_timeout_ms | | int | computed

### password

`password` is required to create a resource.
Once the user is created, `password` is optional.

## Attributes Reference

name | type
--- | ---
user_id | string
external | bool
read_only | bool
client_address | string
session_active | bool
last_activity | string
permissions | string set

## Import

`graylog_user` can be imported using the User `username`, e.g.

```
$ terraform import graylog_user.foo foo
```
