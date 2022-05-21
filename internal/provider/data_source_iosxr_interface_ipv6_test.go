// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIosxrInterfaceIPv6(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrInterfaceIPv6Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxr_interface_ipv6.test", "link_local_address", "fe80::1"),
					resource.TestCheckResourceAttr("data.iosxr_interface_ipv6.test", "link_local_zone", "0"),
					resource.TestCheckResourceAttr("data.iosxr_interface_ipv6.test", "autoconfig", "false"),
					resource.TestCheckResourceAttr("data.iosxr_interface_ipv6.test", "enable", "true"),
				),
			},
		},
	})
}

const testAccDataSourceIosxrInterfaceIPv6Config = `

resource "iosxr_interface_ipv6" "test" {
  interface_name = "GigabitEthernet0/0/0/1"
  link_local_address = "fe80::1"
  link_local_zone = "0"
  autoconfig = false
  enable = true
}

data "iosxr_interface_ipv6" "test" {
  interface_name = "GigabitEthernet0/0/0/1"
  depends_on = [iosxr_interface_ipv6.test]
}
`
