# JSON string attributes

In this provider, the type of some attributes is not complex type but JSON string.

e.g,

```tf
  # graylog_extractor's extractor_config
  extractor_config = <<EOF
{
  "list_separator": ", ",
  "kv_separator": "=",
  "key_prefix": "visit_",
  "key_separator": "_",
  "replace_key_whitespace": false,
  "key_whitespace_replacement": "_"
}
EOF
```

Or using [jsonencode](https://www.terraform.io/docs/configuration/functions/jsonencode.html)


```tf
  extractor_config = jsonencode({
    list_separator             = ", "
    kv_separator               = "="
    key_prefix                 = "visit_"
    key_separator              = "_"
    replace_key_whitespace     = false
    key_whitespace_replacement = "_"
  })
```


By adopting JSON string, there are some merits.

For users

* users can use third party plugins freely
* users can use Graylog API response body without converting to HCL
* users don't take care whether this provider supports the attributes

For developers

* developers don't have to implement the attributes per types
* developer can support attributes which are complicated data structure easily

Even if the same attribute, data structure is different per type at Graylog.  
For example, Graylog supports various Input types, and data structure of the attribute `attributes` is diffeernt per Input type.
Furthermore, Graylog supports the plugin architecture.
So it is hard to support all types including third party plugins.

By using JSON string, we don't have to implement per types and we can support including third party plugins.

## Please tell us the format of JSON attributes

We provide some examples, but we don't provide the document about data structure of JSON attributes.

There are some ways to know the data structure.

* Refer examples which we provide
* Call [Graylog REST API](https://docs.graylog.org/en/latest/pages/configuration/rest_api.html)
* Read [Graylog source code](https://github.com/Graylog2/graylog2-server)
* Ask to [community](https://community.graylog.org/)
