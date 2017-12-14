resource "systemd_hostnamed" "alarmpi" {
    user         = "alarm"
    hostname     = "alarmpi.xlan"
    new_hostname = "node0.xlan"
}
