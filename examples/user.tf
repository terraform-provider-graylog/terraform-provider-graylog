resource "graylog_user" "test" {
  username  = "test"
  email     = "test@example.com"
  password  = "password"
  full_name = "test test"
  roles = [
    "Reader",
  ]
}
