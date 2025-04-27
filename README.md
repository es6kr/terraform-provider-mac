# terraform-provider-mac

[![Checks](https://github.com/es6kr/terraform-provider-mac/actions/workflows/checks.yml/badge.svg)](https://github.com/es6kr/terraform-provider-mac/actions/workflows/checks.yml)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/es6kr/terraform-provider-mac)
[![GitHub license](https://img.shields.io/github/license/es6kr/terraform-provider-mac)](https://github.com/es6kr/terraform-provider-mac/blob/trunk/LICENSE)
[![pre-commit](https://img.shields.io/badge/pre--commit-enabled-brightgreen?logo=pre-commit&logoColor=white)](https://github.com/pre-commit/pre-commit)

**terraform-provider-mac** is a [Terraform](https://www.terraform.io/) provider for installing softwares via various package management tools. Currently, **terraform-provider-mac** supports

- [Homebrew](https://brew.sh/)
- Shell script
- [asdf](https://asdf-vm.com/)

The following shows how to install **git** and **starship** through Homebrew using **terraform-provider-mac** provider. See <https://registry.terraform.io/providers/es6kr/mac/latest/docs> for complete documentation.

```tf
terraform {
  required_version = "~> 1.1.4"
  required_providers {
    mac = {
      source  = "es6kr/mac"
    }
  }
}

locals {
  apps = ["git", "starship"]
}

resource "mac_brew" "this" {
  for_each = toset(local.apps)
  name     = each.key
}
```

## Development

### Code quality

We use [pre-commit](https://pre-commit.com/) to maintain the code quality of this project. Refer to [.pre-commit-config.yaml](./.pre-commit-config.yaml) for the list of linters that we are using. Refer to [this page](https://pre-commit.com/#install) to install pre-commit and the git hook script.

```sh
pre-commit install
```

### Running automated tests

Run unit tests (no resources will be created/destroyed) with

```sh
make test
```

Run acceptance tests with

```sh
make TESTARGS="-tags=asdf" testacc
```

We must provide the value for `-tags` because some tests only run on a specific platform. Currently, the valid values for `-tags` are:

- `brew` for the environment that uses [Homebrew](https://brew.sh/).

We can also use `go test` in a specific directory. For example, the following runs acceptance tests inside the `brew` directory.

```sh
cd ./internal/brew/
TF_ACC=1 go test -tags=brew
```

### Testing with development version

You added a new feature or fixed a bug in **terraform-provider-mac**. You want to test it directly with your Terraform configurations on your local machine. Here is what you can do.

1. Run `make install`. This command installs the provider in `/tmp/tfproviders`. We've set up Terraform in `.terraformrc` to use provider from `/tmp/tfproviders`. Specify the OS and architecture using `OS_ARCH`, e.g., on Apple Silicon macOS, use

   ```sh
   OS_ARCH=darwin_arm64 make install
   ```

2. Use the `.terraformrc` file in this repository to override the published provider.

   ```sh
   export TF_CLI_CONFIG_FILE=$(pwd)/.terraformrc
   terraform init
   terraform plan
   ```

#### Tips

1. Use `export TF_LOG_PROVIDER=DEBUG` for debugging. See <https://www.terraform.io/internals/debugging>.
2. Use [Doc Preview Tool](https://registry.terraform.io/tools/doc-preview) to preview generated documentations.

## References

- [Terraform Provider Installer](https://github.com/shihanng/terraform-provider-installer)
- [Custom Providers](https://learn.hashicorp.com/collections/terraform/providers)
- [Terraform Provider Scaffolding](https://github.com/hashicorp/terraform-provider-scaffolding)
- [Terraform Provider Hashicups](https://github.com/hashicorp/terraform-provider-hashicups)
- [Terraform Provider GitHub](https://github.com/integrations/terraform-provider-github)
