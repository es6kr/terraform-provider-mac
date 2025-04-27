terraform {
  required_version = "~> 1.1.4"
  required_providers {
    mac = {
      source  = "es6kr/mac"
      version = "~> 0.0.3"
    }
  }
}

provider "mac" {
}

locals {
  apps = ["git", "nebula"]
}

resource "mac_brew" "this" {
  for_each = toset(local.apps)
  name     = each.key
}
