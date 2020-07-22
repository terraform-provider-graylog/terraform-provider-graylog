# Changes from go-graylog

In this page, we describe what is changed from [go-graylog](https://github.com/suzuki-shunsuke/go-graylog) to this provider.

## Changes

* Increase [Acceptance Tests](https://www.terraform.io/docs/extend/testing/acceptance-tests/index.html)
* [Use JSON string positively](/json-string-attribute)
* Make the development much easier than go-graylog by design change

## Design changes

* Don't define Go's struct which represent Graylog resources
* Convert Graylog API's request/response body and Terraform resource data directly
* Develop API client for only internal, which means we don't expect API client to be used as library

At go-graylog, data is converted to Go's struct once.

```
Graylog API (JSON) <=> Go's struct <=> Terraform ResourceData
```

So we have to do the following tasks.

* Define Go's struct and interface
* Implement MarshalJSON and UnmarshalJSON
* Implement mapping of Go's struct and Terraform ResourceData

We found that intermediation of Go's struct is unneeded, and it is much easier to map Graylog API's request/response body and Terraform ResourceData directly.

```
Graylog API (JSON) <=> Terraform ResourceData
```
