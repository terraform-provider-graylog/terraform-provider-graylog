---
page_title: "Graylog: graylog_pipeline"
---

# Resource: graylog_pipeline

* [Example](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/examples/v0.12/pipeline.tf)
* [Source Code](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/graylog/resource/system/pipeline/pipeline/resource.go)

## Argument Reference

### Required Argument

* `source` - (Required) The source of the Pipeline. The data type is `string`.
* `description` - (Required) The description of the Pipeline. The data type is `string`.

## Attributes Reference

Nothing.

## Import

`graylog_pipeline` can be imported using the Pipeline id, e.g.

```console
$ terraform import graylog_pipeline.test 5c4acaefc9e77bbbbbbbbbbb
```
