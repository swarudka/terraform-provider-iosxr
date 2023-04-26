// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIosxrRouterISISInterfaceAddressFamily(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxrRouterISISInterfaceAddressFamilyConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxr_router_isis_interface_address_family.test", "af_name", "ipv4"),
					resource.TestCheckResourceAttr("iosxr_router_isis_interface_address_family.test", "saf_name", "unicast"),
					resource.TestCheckResourceAttr("iosxr_router_isis_interface_address_family.test", "fast_reroute_per_prefix_levels.0.level_id", "1"),
					resource.TestCheckResourceAttr("iosxr_router_isis_interface_address_family.test", "fast_reroute_per_prefix_levels.0.ti_lfa", "true"),
					resource.TestCheckResourceAttr("iosxr_router_isis_interface_address_family.test", "tag", "100"),
				),
			},
			{
				ResourceName:  "iosxr_router_isis_interface_address_family.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XR-um-router-isis-cfg:/router/isis/processes/process[process-id=P1]/interfaces/interface[interface-name=GigabitEthernet0/0/0/1]/address-families/address-family[af-name=ipv4][saf-name=unicast]",
			},
		},
	})
}

func testAccIosxrRouterISISInterfaceAddressFamilyConfig_minimum() string {
	return `
	resource "iosxr_router_isis_interface_address_family" "test" {
		process_id = "P1"
		interface_name = "GigabitEthernet0/0/0/1"
		af_name = "ipv4"
		saf_name = "unicast"
	}
	`
}

func testAccIosxrRouterISISInterfaceAddressFamilyConfig_all() string {
	return `
	resource "iosxr_router_isis_interface_address_family" "test" {
		process_id = "P1"
		interface_name = "GigabitEthernet0/0/0/1"
		af_name = "ipv4"
		saf_name = "unicast"
		fast_reroute_per_prefix_levels = [{
			level_id = 1
			ti_lfa = true
		}]
		tag = 100
	}
	`
}
