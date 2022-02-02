resource "graylog_alert_condition" "test" {
  type      = "field_content_value"
  stream_id = graylog_stream.test.id
  in_grace  = false
  title     = "test"

  parameters = jsonencode({
    field                = "message"
    value                = "hoge hoge"
    backlog              = 2
    repeat_notifications = false
    query                = "*"
    grace                = 0
  })
}
