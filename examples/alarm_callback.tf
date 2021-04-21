# You have to add https://example.com to the URL Whitelist Configuration
# https://docs.graylog.org/en/3.2/pages/secure/sec_url_whitelist.html
resource "graylog_alarm_callback" "http" {
  type      = "org.graylog2.alarmcallbacks.HTTPAlarmCallback"
  stream_id = graylog_stream.test.id
  title     = "test"

  configuration = jsonencode({
    url = "https://example.com"
  })
}

resource "graylog_alarm_callback" "email" {
  type      = "org.graylog2.alarmcallbacks.EmailAlarmCallback"
  stream_id = graylog_stream.test.id
  title     = "test"

  configuration = jsonencode({
    user_receivers = [
      "username"
    ]
    body    = "##########\\nAlert Description: $${check_result.resultDescription}\\nDate: $${check_result.triggeredAt}\\nStream ID: $${stream.id}\\nStream title: $${stream.title}\\nStream description: $${stream.description}\\nAlert Condition Title: $${alertCondition.title}\\n$${if stream_url}Stream URL: $${stream_url}$${end}\\n\\nTriggered condition: $${check_result.triggeredCondition}\\n##########\\n\\n$${if backlog}Last messages accounting for this alert:\\n$${foreach backlog message}$${message}\\n\\n$${end}$${else}<No backlog>\\n$${end}\\n"
    sender  = "graylog@example.org"
    subject = "Graylog alert for stream: $${stream.title}: $${check_result.resultDescription}"
    email_receivers = [
      "graylog@example.com"
    ]
  })
}

# Slack plugin is required
# https://github.com/graylog-labs/graylog-plugin-slack
resource "graylog_alarm_callback" "slack" {
  type      = "org.graylog2.plugins.slack.callback.SlackAlarmCallback"
  stream_id = graylog_stream.test.id
  title     = "test"

  configuration = jsonencode({
    graylog2_url   = "https://graylog.example.com"
    custom_message = "$${alert_condition.title}\\n\\n$${foreach backlog message}\\n<https://graylog.example.com/streams/$${stream.id}/search?rangetype=absolute&from=$${message.timestamp}&to=$${message.timestamp} | link> $${message.message}\\n$${end}"
    link_names     = true
    notify_channel = false
    color          = "#FF0000"
    webhook_url    = "https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX"
    user_name      = "Graylog"
    backlog_items  = 5
    channel        = "#general"
  })
}
