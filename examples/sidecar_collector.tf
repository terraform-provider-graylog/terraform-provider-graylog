resource "graylog_sidecar_collector" "test" {
  name                  = "foo"
  service_type          = "exec"
  node_operating_system = "linux"
  executable_path       = "/usr/share/filebeat/bin/filebeat"
}
