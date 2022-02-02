resource "graylog_sidecar_configuration" "test" {
  name         = "foo"
  color        = "#00796b"
  collector_id = graylog_sidecar_collector.test.id
  template     = <<EOF
fields_under_root: true
EOF
}
