data "iosxr_router_ospf_area_interface" "example" {
  process_name   = "OSPF11"
  area_id        = "0"
  interface_name = "GigabitEthernet0/0/0/1"
}
