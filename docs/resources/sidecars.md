---
page_title: "Graylog: graylog_sidecars"
---

# Resource: graylog_sidecars

* [Example](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/examples/v0.12/sidecar.tf)
* [Source Code](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/graylog/resource/sidecar/resource.go)

Manages to assign Sidecars's configuration to Sidecars.
Due to the Graylog API's restriction, we have to manage all assignments by one Terraform resource,
which means we shouldn't use this resource only once.

Good

```hcl
resource "graylog_sidecars" "all" {
  sidecars {
    # ...
  }
}
```

NG

```hcl
resource "graylog_sidecars" "foo" {
  sidecars {
    # ...
  }
}

resource "graylog_sidecars" "bar" {
  sidecars {
    # ...
  }
}
```

## Argument Reference

### Required Argument

name | type
--- | ---
sidecars | []object (set)
sidecars[].node_id | string
sidecars[].assignments | []object (set)
sidecars[].assignments[].collector_id | string
sidecars[].assignments[].configuration_id | string

### Optional Argument

None.

## Attrs Reference

None.

## Import

Unlike other resources, the given ID is ignored so please specify any string as ID.

e.g.

```console
$ terraform import graylog_sidecars.all all
```
