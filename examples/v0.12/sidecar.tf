# data "graylog_sidecar" "test" {
#   node_name = "test"
# }

# resource "graylog_sidecars" "all" {
#   sidecars {
#     node_id = data.graylog_sidecar.test.id
#     assignments {
#       collector_id     = graylog_sidecar_collector.test.id
#       configuration_id = graylog_sidecar_configuration.test.id
#     }
#   }
# }
