# Resource: graylog_ldap_setting

* [Example](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/examples/v0.12/ldap_setting.tf)
* [Source Code](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/graylog/resource/system/ldap/setting/resource.go)

## Argument Reference

* `system_username` - (Required) The data type is `string`.
* `ldap_uri` - (Required) The data type is `string`.
* `search_base` - (Required) The data type is `string`.
* `search_pattern` - (Required) The data type is `string`.
* `display_name_attribute` - (Required) The data type is `string`.
* `default_group` - (Required) The data type is `string`.
* `description` - (Optional) The data type is `string`.
* `enabled` - (Optional) The data type is `bool`.
* `use_start_tls` - (Optional) The data type is `bool`.
* `trust_all_certificates` - (Optional) The data type is `bool`.
* `active_directory` - (Optional) The data type is `bool`.
* `group_search_base` - (Optional) The data type is `string`.
* `group_id_attribute` - (Optional) The data type is `string`.
* `group_search_pattern` - (Optional) The data type is `string`.
* `group_mapping` - (Optional) The data type is `map[string]string`.
* `system_password` - (Optional) The data type is `string`.

Note that `system_passoword` is optional as Terraform schema but is required to create a LDAP setting.
If we make `system_password` required as Terrafrom schema, we have to store `system_password` in the Terraform state file, which some users wouldn't want it.

## Attributes Reference

* `system_password_set` - The data type is `bool`.

## Import

Unlike other resources, LDAP settings has no id,
so when you import the LDAP settings, please specify some string as id.

```console
$ terraform import graylog_ldap_setting.foo bar
```
