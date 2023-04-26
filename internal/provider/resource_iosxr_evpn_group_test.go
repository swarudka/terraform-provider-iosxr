// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIosxrEVPNGroup(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxrEVPNGroupConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxr_evpn_group.test", "group_id", "1"),
					resource.TestCheckResourceAttr("iosxr_evpn_group.test", "core_interfaces.0.interface_name", "Bundle-Ether111"),
				),
			},
			{
				ResourceName:  "iosxr_evpn_group.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XR-um-l2vpn-cfg:/evpn/groups/group[group-name=1]",
			},
		},
	})
}

func testAccIosxrEVPNGroupConfig_minimum() string {
	return `
	resource "iosxr_evpn_group" "test" {
		group_id = 1
	}
	`
}

func testAccIosxrEVPNGroupConfig_all() string {
	return `
	resource "iosxr_evpn_group" "test" {
		group_id = 1
		core_interfaces = [{
			interface_name = "Bundle-Ether111"
		}]
	}
	`
}
