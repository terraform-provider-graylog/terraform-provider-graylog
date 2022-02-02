terraform {
  required_providers {
    graylog = {
      source = "terraform-provider-graylog/graylog"
    }
    random = {
      source = "hashicorp/random"
    }
  }
  required_version = ">= 0.13"
}
