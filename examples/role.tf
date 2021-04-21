resource "graylog_role" "read-stream-test" {
  name        = "read-stream-test"
  description = "read the stream 'test'"

  permissions = [
    "streams:read:${graylog_stream.test.id}"
  ]
}

resource "graylog_role" "terraform" {
  name        = "terraform"
  description = "terraform"

  permissions = [
    "dashboards:*",
    "indexsets:*",
    "inputs:*",
    "roles:*",
    "streams:*",
    "users:*",
    "pipeline_rule:*",
  ]
}
