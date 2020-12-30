resource "graylog_ldap_setting" "system" {
  enabled                   = false
  system_username           = ""
  ldap_uri                  = "ldap://localhost:389"
  use_start_tls             = false
  trust_all_certificates    = false
  active_directory          = false
  search_base               = ""
  search_pattern            = ""
  display_name_attribute    = ""
  default_group             = "Reader"
  group_mapping             = null
  group_search_base         = null
  group_id_attribute        = null
  additional_default_groups = null
  group_search_pattern      = null
}
