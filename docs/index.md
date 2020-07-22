---
page_title: "Provider: Graylog"
---

The Graylog provider is used to manage [Graylog](https://docs.graylog.org/)'s various resources.

## Example Usage

```tf
# Configure Graylog Provider
provider "graylog" {
  web_endpoint_uri = "http://example.com/api"
  api_version      = "v3"
}

# Create a Input
resource "graylog_input" "gelf_udp" {
  title  = "gelf udp"
  type   = "org.graylog2.inputs.gelf.udp.GELFUDPInput"
  global = true

  attributes = jsonencode({
    bind_address          = "0.0.0.0"
    port                  = 12201
    recv_buffer_size      = 262144
    decompress_size_limit = 8388608
  })
}
```

## Argument Reference

* `web_endpoint_uri` - (Required) Graylog API endpoint, for example https://example.com/api . It can also be sourced from the `GRAYLOG_WEB_ENDPOINT_URI` environment variable.
* `auth_name` - (Required) Username or API token or Session Token. It can also be sourced from the `GRAYLOG_AUTH_NAME` environment variable.
* `auth_password` - (Required) Password or the literal `"token"` or `"session"`. It can also be sourced from the `GRAYLOG_AUTH_PASSWORD` environment variable.
* `x_requested_by` - (Optional) [X-Requested-By Header](https://github.com/Graylog2/graylog2-server/blob/370dd700bc8ada5448bf66459dec9a85fcd22d58/UPGRADING.rst#protecting-against-csrf-http-header-required). The default value is `terraform-provider-graylog`. It can also be sourced from the `GRAYLOG_X_REQUESTED_BY` environment variable.
* `api_version` - (Optional) Graylog API's version. The default value is `v3`. It can also be sourced from the `GRAYLOG_API_VERSION ` environment variable.

## Authentication

About `auth_name` and `auth_password`, please see the [Graylog's Documentation](https://docs.graylog.org/en/latest/pages/configuration/rest_api.html).

You can authenticate with either password or access token or session token.

password

```
auth_name = "<user name>"
auth_password = "<password>"
```

access token

```
auth_name = "<access token>"
auth_password = "token"
```

session token

```
auth_name = "<session token>"
auth_password = "session"
```
