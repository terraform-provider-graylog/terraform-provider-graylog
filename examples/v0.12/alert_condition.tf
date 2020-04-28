resource "graylog_alert_condition" "test" {
  type      = "field_content_value"
  stream_id = graylog_stream.test.id
  in_grace  = false
  title     = "test"

  parameters = <<EOF
{
  "backlog": 2,
  "repeat_notifications": false,
  "field": "message",
  "query": "*",
  "grace": 0,
  "value": "hoge hoge"
}
EOF
}
