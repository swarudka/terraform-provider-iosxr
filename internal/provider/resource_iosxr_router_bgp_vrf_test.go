// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIosxrRouterBGPVRF(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxrRouterBGPVRFConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxr_router_bgp_vrf.test", "vrf_name", "VRF2"),
					resource.TestCheckResourceAttr("iosxr_router_bgp_vrf.test", "rd_auto", "false"),
					resource.TestCheckResourceAttr("iosxr_router_bgp_vrf.test", "rd_ip_address_ipv4_address", "14.14.14.14"),
					resource.TestCheckResourceAttr("iosxr_router_bgp_vrf.test", "rd_ip_address_index", "3"),
					resource.TestCheckResourceAttr("iosxr_router_bgp_vrf.test", "default_information_originate", "true"),
					resource.TestCheckResourceAttr("iosxr_router_bgp_vrf.test", "default_metric", "125"),
					resource.TestCheckResourceAttr("iosxr_router_bgp_vrf.test", "timers_bgp_keepalive_interval", "5"),
					resource.TestCheckResourceAttr("iosxr_router_bgp_vrf.test", "timers_bgp_holdtime", "20"),
					resource.TestCheckResourceAttr("iosxr_router_bgp_vrf.test", "bfd_minimum_interval", "10"),
					resource.TestCheckResourceAttr("iosxr_router_bgp_vrf.test", "bfd_multiplier", "4"),
					resource.TestCheckResourceAttr("iosxr_router_bgp_vrf.test", "neighbors.0.neighbor_address", "10.1.1.2"),
					resource.TestCheckResourceAttr("iosxr_router_bgp_vrf.test", "neighbors.0.remote_as", "65002"),
					resource.TestCheckResourceAttr("iosxr_router_bgp_vrf.test", "neighbors.0.description", "My Neighbor Description"),
					resource.TestCheckResourceAttr("iosxr_router_bgp_vrf.test", "neighbors.0.ignore_connected_check", "true"),
					resource.TestCheckResourceAttr("iosxr_router_bgp_vrf.test", "neighbors.0.ebgp_multihop_maximum_hop_count", "10"),
					resource.TestCheckResourceAttr("iosxr_router_bgp_vrf.test", "neighbors.0.bfd_minimum_interval", "10"),
					resource.TestCheckResourceAttr("iosxr_router_bgp_vrf.test", "neighbors.0.bfd_multiplier", "4"),
					resource.TestCheckResourceAttr("iosxr_router_bgp_vrf.test", "neighbors.0.local_as", "65003"),
					resource.TestCheckResourceAttr("iosxr_router_bgp_vrf.test", "neighbors.0.local_as_no_prepend", "true"),
					resource.TestCheckResourceAttr("iosxr_router_bgp_vrf.test", "neighbors.0.local_as_replace_as", "true"),
					resource.TestCheckResourceAttr("iosxr_router_bgp_vrf.test", "neighbors.0.local_as_dual_as", "true"),
					resource.TestCheckResourceAttr("iosxr_router_bgp_vrf.test", "neighbors.0.password", "12341C2713181F13253920"),
					resource.TestCheckResourceAttr("iosxr_router_bgp_vrf.test", "neighbors.0.shutdown", "false"),
					resource.TestCheckResourceAttr("iosxr_router_bgp_vrf.test", "neighbors.0.timers_keepalive_interval", "5"),
					resource.TestCheckResourceAttr("iosxr_router_bgp_vrf.test", "neighbors.0.timers_holdtime", "20"),
					resource.TestCheckResourceAttr("iosxr_router_bgp_vrf.test", "neighbors.0.update_source", "GigabitEthernet0/0/0/1"),
					resource.TestCheckResourceAttr("iosxr_router_bgp_vrf.test", "neighbors.0.ttl_security", "false"),
				),
			},
			{
				ResourceName:  "iosxr_router_bgp_vrf.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XR-um-router-bgp-cfg:/router/bgp/as[as-number=65001]/vrfs/vrf[vrf-name=VRF2]",
			},
		},
	})
}

func testAccIosxrRouterBGPVRFConfig_minimum() string {
	return `
	resource "iosxr_router_bgp_vrf" "test" {
		as_number = "65001"
		vrf_name = "VRF2"
		timers_bgp_keepalive_interval = 5
		timers_bgp_holdtime = "20"
	}
	`
}

func testAccIosxrRouterBGPVRFConfig_all() string {
	return `
	resource "iosxr_router_bgp_vrf" "test" {
		as_number = "65001"
		vrf_name = "VRF2"
		rd_auto = false
		rd_ip_address_ipv4_address = "14.14.14.14"
		rd_ip_address_index = 3
		default_information_originate = true
		default_metric = 125
		timers_bgp_keepalive_interval = 5
		timers_bgp_holdtime = "20"
		bfd_minimum_interval = 10
		bfd_multiplier = 4
		neighbors = [{
			neighbor_address = "10.1.1.2"
			remote_as = "65002"
			description = "My Neighbor Description"
			ignore_connected_check = true
			ebgp_multihop_maximum_hop_count = 10
			bfd_minimum_interval = 10
			bfd_multiplier = 4
			local_as = "65003"
			local_as_no_prepend = true
			local_as_replace_as = true
			local_as_dual_as = true
			password = "12341C2713181F13253920"
			shutdown = false
			timers_keepalive_interval = 5
			timers_holdtime = "20"
			update_source = "GigabitEthernet0/0/0/1"
			ttl_security = false
		}]
	}
	`
}
