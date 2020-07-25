# Resource: graylog_extractor

* [Example](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/examples/v0.12/extractor.tf)
* [Source Code](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/graylog/resource/system/input/extractor/resource.go)

## Argument Reference

### Required Argument

* `input_id` - (Required) the Input id which the Extractor is associated with. The data type is `string`.
* `type` - (Required) the type of the Extractor. The data type is `string`.
* `title` - (Required) the title of the Extractor. The data type is `string`.
* `cursor_strategy` - (Required) The data type is `string`.
* `source_field` - (Required) The data type is `string`.
* `condition_type` - (Required) the condition type of the Extractor. The data type is `string`.
* `extractor_config` - (Required) The data type is `JSON string`.
* `converters[].type` - (Required) the type of the converter. The data type is `string`.
* `converters[].config` - (Required) the configuration of the converter. The data type is `JSON string`.
* `converters` - (Optional) The data type is `[]object`. The default value is `[]`.
* `target_field` - (Optional) The data type is `string`.
* `condition_value` - (Optional) The data type is `string`.
* `order` - (Optional) The data type is `int`.

### Attributes Reference

* `extractor_id` - The id of the extractor. The data type is `string`.

## Import

`graylog_extractor` can be imported using the User `<input id>/<extractor id>`, e.g.

```console
$ terraform import graylog_extractor.test 5bb1b4b5c9e77bbbbbbbbbbb/5c4acaefc9e77bbbbbbbbbbb
```
