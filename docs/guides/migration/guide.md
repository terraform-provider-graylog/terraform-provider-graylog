# Migration Guide

In this page, we describe the overview of how to migrate Terraform provider from [go-graylog v11.3.0](https://github.com/suzuki-shunsuke/go-graylog/tree/v11.3.0) to this provider.  
Unfortunately there are some breaking changes between go-graylog and terraform-provider-graylog, so
we have to fix Terraform configuration files (.tf) manually.

For detail of the changes of each resources, please see [here](migration/detail.md) also.

## Migration steps

1. Backup Terraform configuration (.tf) and State (.tfstate)
1. Replace the binary of provider
1. Fix the configuration (.tf)

## What is changed

* The default value of `api_version` is changed from `v2` to `v3`
* Some attribute names are changed
* Some attribute types are changed to JSON string
* Resource ID becomes to be same as the string which is used for `terraform import`

### api_version

The default value of `api_version` is changed from `v2` to `v3`.
Currently, we support only `v3`.
If you use `v2`, please use go-graylog instead of this provider.

### Some attribute names are changed

For example, at `go-graylog` AlarmCallback has the following attributes.

https://github.com/suzuki-shunsuke/go-graylog/blob/v11.3.0/docs/resources/alarm_callback.md

* `http_configuration`
* `email_configuration`
* etc

These attributes are unified to `configuration`.

### Some attribute types are changed to JSON string

For example, at `go-graylog` AlarmCallback's `http_configuration` is the following structure.

```tf
http_configuration {
  url = "https://example.com"
}
```

At this provider, this is changed to JSON string.

```tf
configuration = <<EOF
{
  "url": "http://example.com"
}
EOF
```

By [jsonencode](https://www.terraform.io/docs/configuration/functions/jsonencode.html), you can migrate more easily.

```tf
configuration = jsonencode({
  url = "https://example.com"
})
```

### Resource ID

For example, at `go-graylog` `graylog_alarmcallback`'s `id` is `<AlarmCallback's id>`.
On the other hand, at `terraform-provider-graylog` `graylog_alarmcallback`'s `id` is `<Stream's id>/<AlarmCallback's id>`, which is same as the argument of `terraform import`.

In case of the following AlarmCallback,

```json
{
  "id": "5ea2bc0a2ab79c001274e26f",
  "stream_id": "5ea26bb42ab79c0012521287",
  ...
}
```

provider | attribute | value
--- | ---- | ---
go-graylog | id | `5ea2bc0a2ab79c001274e26f`
terraform-provider-graylog | id | `5ea26bb42ab79c0012521287/5ea2bc0a2ab79c001274e26f`
terraform-provider-graylog | alarmcallback_id | `5ea2bc0a2ab79c001274e26f`

So please fix `graylog_alarm_callback.<name>.id` to `graylog_alarm_callback.<name>.alarmcallback_id`.

## State Migration

The schema of Terraform State is also changed, but Terraform supports the state migration.

https://www.terraform.io/docs/extend/resources/state-migration.html

We implement `StateUpgrades`, so state is migrated automatically.
If you encounter the trouble about state migration, please [create an issue](https://github.com/terraform-provider-graylog/terraform-provider-graylog/issues/new).
