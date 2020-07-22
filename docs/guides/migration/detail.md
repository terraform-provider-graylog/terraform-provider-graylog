# Migration Guide Detail

## No changes

* graylog_dashboard
* graylog_event_definition
* graylog_event_notification
* graylog_grok_pattern
* graylog_input_static_fields
* graylog_ldap_setting
* graylog_output
* graylog_pipeline
* graylog_pipeline_connection
* graylog_pipeline_rule
* graylog_role
* graylog_stream
* graylog_stream_output
* graylog_stream_rule
* graylog_user

## graylog_alarm_callback

Reference

* [go-graylog](https://github.com/suzuki-shunsuke/go-graylog/blob/v11.3.0/docs/resources/alarm_callback.md)

### Rename the following attributes to `configuration` and change to JSON string

* http_configuration
* email_configuration
* slack_configuration
* general_int_configuration
* general_bool_configuration
* general_float_configuration
* general_string_configuration

AS IS

```tf
http_configuration {
  url = "https://example.com"
}
```

TO BE

```tf
configuration = jsonencode({
  url = "https://example.com"
})
```

### Change the attribute `.id` to `.alarmcallback_id`

AS IS

```tf
graylog_alarm_callback.foo.id
```

TO BE

```tf
graylog_alarm_callback.foo.alarmcallback_id
```

## graylog_alert_condition

Reference

* [go-graylog](https://github.com/suzuki-shunsuke/go-graylog/blob/v11.3.0/docs/resources/alert_condition.md)

### Rename the following attributes to `parameters` and change to JSON string

* field_content_value_parameters
* field_value_parameters
* message_count_parameters
* general_int_parameters
* general_bool_parameters
* general_float_parameters
* general_string_parameters

AS IS

```tf
field_content_value_parameters {
  field                = "message"
  value                = "hoge hoge"
  backlog              = 2
  repeat_notifications = false
  query                = "*"
  grace                = 0
}
```

TO BE

```tf
parameters = jsonencode({
  field                = "message"
  value                = "hoge hoge"
  backlog              = 2
  repeat_notifications = false
  query                = "*"
  grace                = 0
})
```

### Change the attribute `.id` to `.alert_condition_id`

AS IS

```tf
graylog_alert_condition.foo.id
```

TO BE

```tf
graylog_alert_condition.foo.alert_condition_id
```

## graylog_dashboard_widget

Reference

* [go-graylog](https://github.com/suzuki-shunsuke/go-graylog/blob/v11.3.0/docs/resources/dashboard_widget.md)

### Rename the following attributes to `config` and change to JSON string

* json_configuration
* stream_search_result_count_configuration
* quick_values_configuration
* quick_values_histogram_configuration
* search_result_chart_configuration
* field_chart_configuration
* stats_count_configuration

AS IS

```tf
quick_values_histogram_configuration {
  timerange {
    type = "relative"
    range = 28800
  }
  stream_id = "5b3983000000000000000000"
  query = "status:200"
  field = "status"
  limit = 5
  sort_order = "desc"
  stacked_fields = ""
}
```

TO BE

```tf
config = jsonencode({
  timerange = {
    type = "relative"
    range = 28800
  }
  stream_id = "5b3983000000000000000000"
  query = "status:200"
  field = "status"
  limit = 5
  sort_order = "desc"
  stacked_fields = ""
})
```

### Change the attribute `.id` to `.widget_id`

AS IS

```tf
graylog_dashboard_widget.foo.id
```

TO BE

```tf
graylog_dashboard_widget.foo.widget_id
```

## graylog_dashboard_widget_positions

* Change attribute value to JSON string
  * positions

AS IS

```tf
positions {
  widget_id = graylog_dashboard_widget.test.id
  row       = 0
  col       = 0
  height    = 1
  width     = 1
}
positions {
  widget_id = graylog_dashboard_widget.test2.id
  row       = 0
  col       = 1
  height    = 2
  width     = 2
}
```

TO BE

```tf
positions = jsonencode({
  "${graylog_dashboard_widget.test.widget_id}" = {
    row    = 0
    col    = 0
    height = 1
    width  = 1
  }
  "${graylog_dashboard_widget.test2.widget_id}" = {
    row    = 0
    col    = 1
    height = 2
    width  = 2
  }
})
```

## graylog_extractor

* Change attribute value to JSON string
  * extractor_config
  * converters[].config

### Rename the following attributes to `extractor_config` and change to JSON string

* grok_type_extractor_config
* json_type_extractor_config
* regex_type_extractor_config
* general_int_extractor_config
* general_bool_extractor_config
* general_float_extractor_config
* general_string_extractor_config

AS IS

```tf
json_type_extractor_config {
  list_separator             = ", "
  kv_separator               = "="
  key_prefix                 = "visit_"
  key_separator              = "_"
  replace_key_whitespace     = false
  key_whitespace_replacement = "_"
}
```

TO BE

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

### Change the attribute `.id` to `.extractor_id`

AS IS

```tf
graylog_extractor.foo.id
```

TO BE

```tf
graylog_extractor.foo.extractor_id
```

## graylog_index_set

* Change attribute value to JSON string
  * retention_strategy
  * rotation_strategy

AS IS

```tf
retention_strategy {
  max_number_of_indices = 20
  type                  = "org.graylog2.indexer.retention.strategies.DeletionRetentionStrategyConfig"
}

rotation_strategy {
  max_docs_per_index = 20000000
  max_size           = 0
  type               = "org.graylog2.indexer.rotation.strategies.MessageCountRotationStrategyConfig"
}
```

TO BE

```tf
retention_strategy = jsonencode({
  max_number_of_indices = 20
  type                  = "org.graylog2.indexer.retention.strategies.DeletionRetentionStrategyConfig"
})

rotation_strategy = jsonencode({
  max_docs_per_index = 20000000
  max_size           = 0
  type               = "org.graylog2.indexer.rotation.strategies.MessageCountRotationStrategyConfig"
})
```

## graylog_input

* Change attribute value to JSON string
  * attributes

AS IS

```tf
attributes {
  bind_address          = "0.0.0.0"
  port                  = 12201
  recv_buffer_size      = 262144
  decompress_size_limit = 8388608
}
```

TO BE

```tf
attributes = jsonencode({
  bind_address          = "0.0.0.0"
  port                  = 12201
  recv_buffer_size      = 262144
  decompress_size_limit = 8388608
})
```

## graylog_stream_rule

### Change the attribute `.id` to `.rule_id`

AS IS

```tf
graylog_stream_rule.foo.id
```

TO BE

```tf
graylog_stream_rule.foo.rule_id
```
