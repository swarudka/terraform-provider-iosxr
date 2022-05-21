// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIosxrRouterISISAddressFamily(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrRouterISISAddressFamilyConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "mpls_ldp_auto_config", "false"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "metric_style_narrow", "false"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "metric_style_wide", "true"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "metric_style_transition", "false"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "router_id_ip_address", "1.2.3.4"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "default_information_originate", "true"),
				),
			},
		},
	})
}

const testAccDataSourceIosxrRouterISISAddressFamilyConfig = `

resource "iosxr_router_isis_address_family" "test" {
  process_id = "P1"
  af_name = "ipv4"
  saf_name = "unicast"
  mpls_ldp_auto_config = false
  metric_style_narrow = false
  metric_style_wide = true
  metric_style_transition = false
  router_id_ip_address = "1.2.3.4"
  default_information_originate = true
}

data "iosxr_router_isis_address_family" "test" {
  process_id = "P1"
  af_name = "ipv4"
  saf_name = "unicast"
  depends_on = [iosxr_router_isis_address_family.test]
}
`
