resource "mac_brew" "nebula" {
  name = "nebula"
}

resource "mac_service" "nebula" {
  label = "homebrew.mxcl.nebula"
  plist = "/opt/homebrew/opt/nebula/homebrew.mxcl.nebula.plist"
  start = true

  depends_on = [mac_brew.nebula]
}
