// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIosxrL2VPNXconnectGroupP2PInterface(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrL2VPNXconnectGroupP2PInterfaceConfig,
				Check:  resource.ComposeTestCheckFunc(),
			},
		},
	})
}

const testAccDataSourceIosxrL2VPNXconnectGroupP2PInterfaceConfig = `

resource "iosxr_l2vpn_xconnect_group_p2p_interface" "test" {
  group_name = "P2P"
  p2p_xconnect_name = "XC"
  interface_name = "GigabitEthernet0/0/0/2"
}

data "iosxr_l2vpn_xconnect_group_p2p_interface" "test" {
  group_name = "P2P"
  p2p_xconnect_name = "XC"
  interface_name = "GigabitEthernet0/0/0/2"
  depends_on = [iosxr_l2vpn_xconnect_group_p2p_interface.test]
}
`
