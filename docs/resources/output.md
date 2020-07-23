# Resource: graylog_output

* [Example](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/examples/v0.12/output.tf)
* [Source Code](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/graylog/resource/system/output/resource.go)

## Argument Reference

* `title` - (Required) The title of the Output. The data type is `string`.
* `type` - (Required) The type of the Output. The data type is `string`.
* `configuration` - (Required) The configuration of the Output. The data type is `JSON string`.

`configuration` is a JSON string.
The format of `configuration` depends on the output type.
Please see the [example](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/examples/v0.12/output.tf).
Using the [Graylog's API browser](https://docs.graylog.org/en/3.1/pages/configuration/rest_api.html) you can check the format of `configuration`.

## Attributes Reference

None.

## Import

`graylog_output` can be imported using the Output id, e.g.

```console
$ terraform import graylog_output.test 5c4acaefc9e77bbbbbbbbbbb
```
