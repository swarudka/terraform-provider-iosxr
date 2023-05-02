// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIosxrInterface(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxrInterfaceConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxr_interface.test", "interface_name", "GigabitEthernet0/0/0/1"),
					resource.TestCheckResourceAttr("iosxr_interface.test", "l2transport", "false"),
					resource.TestCheckResourceAttr("iosxr_interface.test", "point_to_point", "false"),
					resource.TestCheckResourceAttr("iosxr_interface.test", "multipoint", "false"),
					resource.TestCheckResourceAttr("iosxr_interface.test", "dampening_decay_half_life_value", "2"),
					resource.TestCheckResourceAttr("iosxr_interface.test", "ipv4_point_to_point", "true"),
					resource.TestCheckResourceAttr("iosxr_interface.test", "service_policy_input.0.name", "CORE-INPUT-POLICY"),
					resource.TestCheckResourceAttr("iosxr_interface.test", "service_policy_output.0.name", "CORE-OUTPUT-POLICY"),
					resource.TestCheckResourceAttr("iosxr_interface.test", "bfd_mode_ietf", "true"),
					resource.TestCheckResourceAttr("iosxr_interface.test", "shutdown", "true"),
					resource.TestCheckResourceAttr("iosxr_interface.test", "mtu", "9000"),
					resource.TestCheckResourceAttr("iosxr_interface.test", "bandwidth", "100000"),
					resource.TestCheckResourceAttr("iosxr_interface.test", "description", "My Interface Description"),
					resource.TestCheckResourceAttr("iosxr_interface.test", "load_interval", "30"),
					resource.TestCheckResourceAttr("iosxr_interface.test", "vrf", "VRF1"),
					resource.TestCheckResourceAttr("iosxr_interface.test", "ipv4_address", "1.1.1.1"),
					resource.TestCheckResourceAttr("iosxr_interface.test", "ipv4_netmask", "255.255.255.0"),
					resource.TestCheckResourceAttr("iosxr_interface.test", "ipv6_link_local_address", "fe80::1"),
					resource.TestCheckResourceAttr("iosxr_interface.test", "ipv6_link_local_zone", "0"),
					resource.TestCheckResourceAttr("iosxr_interface.test", "ipv6_autoconfig", "false"),
					resource.TestCheckResourceAttr("iosxr_interface.test", "ipv6_enable", "true"),
					resource.TestCheckResourceAttr("iosxr_interface.test", "ipv6_addresses.0.address", "2001::1"),
					resource.TestCheckResourceAttr("iosxr_interface.test", "ipv6_addresses.0.prefix_length", "64"),
					resource.TestCheckResourceAttr("iosxr_interface.test", "ipv6_addresses.0.zone", "0"),
				),
			},
			{
				ResourceName:  "iosxr_interface.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XR-um-interface-cfg:/interfaces/interface[interface-name=GigabitEthernet0/0/0/1]",
			},
		},
	})
}

func testAccIosxrInterfaceConfig_minimum() string {
	return `
	resource "iosxr_interface" "test" {
		interface_name = "GigabitEthernet0/0/0/1"
		dampening_decay_half_life_value = 2
	}
	`
}

func testAccIosxrInterfaceConfig_all() string {
	return `
	resource "iosxr_interface" "test" {
		interface_name = "GigabitEthernet0/0/0/1"
		l2transport = false
		point_to_point = false
		multipoint = false
		dampening_decay_half_life_value = 2
		ipv4_point_to_point = true
		service_policy_input = [{
			name = "CORE-INPUT-POLICY"
		}]
		service_policy_output = [{
			name = "CORE-OUTPUT-POLICY"
		}]
		bfd_mode_ietf = true
		shutdown = true
		mtu = 9000
		bandwidth = 100000
		description = "My Interface Description"
		load_interval = 30
		vrf = "VRF1"
		ipv4_address = "1.1.1.1"
		ipv4_netmask = "255.255.255.0"
		ipv6_link_local_address = "fe80::1"
		ipv6_link_local_zone = "0"
		ipv6_autoconfig = false
		ipv6_enable = true
		ipv6_addresses = [{
			address = "2001::1"
			prefix_length = 64
			zone = "0"
		}]
	}
	`
}
