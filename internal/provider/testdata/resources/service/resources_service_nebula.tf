resource "mac_service" "nebula" {
  label = "com.slackhq.nebula"
  plist = "/Library/LaunchDaemons/com.slackhq.nebula.plist"

  depends_on = [mac_brew.nebula]
}
