terraform {
  required_version = "~> 1.1.4"
  required_providers {
    installer = {
      source  = "es6kr/mac"
      version = "~> 0.6.0"
    }
  }
}

provider "mac" {
}

locals {
  apps = ["git", "starship"]
}

resource "installer_brew" "this" {
  for_each = toset(local.apps)
  name     = each.key
}
