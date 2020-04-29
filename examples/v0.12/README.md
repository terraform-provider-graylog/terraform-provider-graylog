# Example

This is sample which we can execute.

## Requirements

* Docker Engine
* Docker Compose
* Terraform
* terraform-provider-graylog
* [graylog-plugin-slack](https://github.com/graylog-labs/graylog-plugin-slack)
  * Download a file and install it to the plugin dir

```
$ curl -L -o plugin/graylog-plugin-slack-3.1.0.jar https://github.com/graylog-labs/graylog-plugin-slack/releases/download/3.1.0/graylog-plugin-slack-3.1.0.jar
```

```
$ docker-compose up -d
```

Access http://127.0.0.1:9000 and login as admin

* username: admin
* password: admin

To create a HTTP Alarm Callback, add `https://example.com` to the URL Whitelist Configuration.

https://docs.graylog.org/en/3.2/pages/secure/sec_url_whitelist.html

```
$ terraform init
```

```
$ terraform plan
```
