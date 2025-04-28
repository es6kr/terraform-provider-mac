---
layout: ""
page_title: "Provider: macOS"
description: |-
  Install/uninstall apps in the local machine using the macOS provider.
---

# macOS Provider

The macOS provider provides resources to manage apps (install/uninstall)
in your local machine in a declarative manner. It currently supports systems that use

- [Homebrew](https://brew.sh/)

It also supports shell script via the `mac_script` resource.

## Example Usage

There is no configuration on the provider level.
The following shows how to ensure the system has git and starship installed via Homebrew.

```terraform
terraform {
  required_version = "~> 1.1.4"
  required_providers {
    mac = {
      source  = "es6kr/mac"
      version = "~> 0.0.4"
    }
  }
}

provider "mac" {
}

locals {
  apps = ["git", "starship"]
}

resource "mac_brew" "this" {
  for_each = toset(local.apps)
  name     = each.key
}
```
