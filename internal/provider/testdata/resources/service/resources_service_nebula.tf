resource "mac_service" "nebula" {
  label = "com.slackhq.nebula"
  plist = "/Library/LaunchDaemons/com.slackhq.nebula.plist"
  start = true

  depends_on = [mac_brew.nebula]
}
