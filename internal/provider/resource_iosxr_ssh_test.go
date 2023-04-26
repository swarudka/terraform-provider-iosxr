// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIosxrSSH(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxrSSHConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxr_ssh.test", "server_dscp", "48"),
					resource.TestCheckResourceAttr("iosxr_ssh.test", "server_logging", "true"),
					resource.TestCheckResourceAttr("iosxr_ssh.test", "server_rate_limit", "60"),
					resource.TestCheckResourceAttr("iosxr_ssh.test", "server_session_limit", "10"),
					resource.TestCheckResourceAttr("iosxr_ssh.test", "server_v2", "true"),
					resource.TestCheckResourceAttr("iosxr_ssh.test", "server_vrfs.0.vrf_name", "VRF1"),
				),
			},
			{
				ResourceName:  "iosxr_ssh.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XR-um-ssh-cfg:/ssh",
			},
		},
	})
}

func testAccIosxrSSHConfig_minimum() string {
	return `
	resource "iosxr_ssh" "test" {
	}
	`
}

func testAccIosxrSSHConfig_all() string {
	return `
	resource "iosxr_ssh" "test" {
		server_dscp = 48
		server_logging = true
		server_rate_limit = 60
		server_session_limit = 10
		server_v2 = true
		server_vrfs = [{
			vrf_name = "VRF1"
		}]
	}
	`
}
