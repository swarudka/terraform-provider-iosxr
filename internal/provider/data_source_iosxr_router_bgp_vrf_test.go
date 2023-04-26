// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIosxrRouterBGPVRF(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrRouterBGPVRFConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "rd_auto", "true"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "rd_two_byte_as_as_number", "65004"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "rd_two_byte_as_index", "1"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "rd_four_byte_as_as_number", "65005"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "rd_four_byte_as_index", "2"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "rd_ip_address_ipv4_address", "14.14.14.14"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "rd_ip_address_index", "3"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "default_information_originate", "true"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "default_metric", "125"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "timers_bgp_keepalive_interval", "5"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "timers_bgp_holdtime", "20"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "bfd_minimum_interval", "10"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "bfd_multiplier", "4"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "neighbors.0.neighbor_address", "10.1.1.2"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "neighbors.0.remote_as", "65002"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "neighbors.0.description", "My Neighbor Description"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "neighbors.0.ignore_connected_check", "true"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "neighbors.0.ebgp_multihop_maximum_hop_count", "10"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "neighbors.0.bfd_minimum_interval", "10"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "neighbors.0.bfd_multiplier", "4"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "neighbors.0.local_as", "65003"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "neighbors.0.local_as_no_prepend", "true"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "neighbors.0.local_as_replace_as", "true"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "neighbors.0.local_as_dual_as", "true"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "neighbors.0.password", "12341C2713181F13253920"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "neighbors.0.shutdown", "false"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "neighbors.0.timers_keepalive_interval", "5"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "neighbors.0.timers_holdtime", "20"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "neighbors.0.update_source", "GigabitEthernet0/0/0/1"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "neighbors.0.ttl_security", "false"),
				),
			},
		},
	})
}

const testAccDataSourceIosxrRouterBGPVRFConfig = `

resource "iosxr_router_bgp_vrf" "test" {
	as_number = "65001"
	vrf_name = "VRF1"
	rd_auto = true
	rd_two_byte_as_as_number = "65004"
	rd_two_byte_as_index = 1
	rd_four_byte_as_as_number = "65005"
	rd_four_byte_as_index = 2
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

data "iosxr_router_bgp_vrf" "test" {
	as_number = "65001"
	vrf_name = "VRF1"
	depends_on = [iosxr_router_bgp_vrf.test]
}
`
