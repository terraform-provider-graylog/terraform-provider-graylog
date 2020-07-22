---
page_title: "Graylog: graylog_grok_pattern"
---

# Resource: graylog_grok_pattern

* [Source Code](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/graylog/resource/system/grok/resource.go)

## Note

Note that currently this resource doesn't support content packs.

And if you use the sequence `%{`, you have to escape it.

https://github.com/hashicorp/hcl2/blob/57bd5f374f26cdb7ae1b1c92fd6eb71335b9805b/hcl/hclsyntax/spec.md#template-literals

> The interpolation and directive introductions are escaped by doubling their leading characters.
> The ${ sequence is escaped as $${ and the %{ sequence is escaped as %%{.

## Example Usage

[Example](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/examples/v0.12/grok_pattern.tf)

```hcl
resource "graylog_grok_pattern" "datestamp" {
  name = "DATESTAMP"
  pattern = "%%{DATE}[- ]%%{TIME}"
}
```

## Argument Reference

* `name` - (Required) the name of the Grok Pattern. The data type is `string`.
* `pattern` - (Required) the pattern of the Grok Pattern. The data type is `string`.

## Attributes Reference

Nothing.

## Import

`graylog_grok_pattern` can be imported using the Grok id, e.g.

```console
$ terraform import graylog_grok_pattern.test 5c4acaefc9e77bbbbbbbbbbb
```
