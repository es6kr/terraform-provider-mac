resource "mac_asdf_plugin" "this" {
  name    = "terraform-ls"
  git_url = "https://github.com/shihanng/asdf-terraform-ls"
}

resource "mac_asdf" "this" {
  name    = "terraform-ls"
  version = "0.25.2"

  depends_on = [
    mac_asdf_plugin.this,
  ]
}
