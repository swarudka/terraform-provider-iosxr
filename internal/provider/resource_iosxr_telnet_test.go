// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxrTelnet(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_telnet.test", "ipv4_client_source_interface", "GigabitEthernet0/0/0/1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_telnet.test", "vrfs.0.vrf_name", "ROI"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_telnet.test", "vrfs.0.ipv4_server_max_servers", "32"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_telnet.test", "vrfs.0.ipv4_server_access_list", "ACCESS1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_telnet.test", "vrfs.0.ipv6_server_max_servers", "34"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_telnet.test", "vrfs.0.ipv6_server_access_list", "ACCESS11"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_telnet.test", "vrfs_dscp.0.vrf_name", "TOI"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_telnet.test", "vrfs_dscp.0.ipv4_dscp", "55"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxrTelnetConfig_minimum(),
			},
			{
				Config: testAccIosxrTelnetConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				ResourceName:  "iosxr_telnet.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XR-um-telnet-cfg:/telnet",
			},
		},
	})
}

func testAccIosxrTelnetConfig_minimum() string {
	config := `resource "iosxr_telnet" "test" {` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxrTelnetConfig_all() string {
	config := `resource "iosxr_telnet" "test" {` + "\n"
	config += `	ipv4_client_source_interface = "GigabitEthernet0/0/0/1"` + "\n"
	config += `	vrfs = [{` + "\n"
	config += `		vrf_name = "ROI"` + "\n"
	config += `		ipv4_server_max_servers = 32` + "\n"
	config += `		ipv4_server_access_list = "ACCESS1"` + "\n"
	config += `		ipv6_server_max_servers = 34` + "\n"
	config += `		ipv6_server_access_list = "ACCESS11"` + "\n"
	config += `	}]` + "\n"
	config += `	vrfs_dscp = [{` + "\n"
	config += `		vrf_name = "TOI"` + "\n"
	config += `		ipv4_dscp = 55` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}
