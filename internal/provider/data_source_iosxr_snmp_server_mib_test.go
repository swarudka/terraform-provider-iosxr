// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxrSNMPServerMIB(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrSNMPServerMIBConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxr_snmp_server_mib.test", "ifmib_ifalias_long", "true"),
					resource.TestCheckResourceAttr("data.iosxr_snmp_server_mib.test", "ifindex_persist", "true"),
				),
			},
		},
	})
}

const testAccDataSourceIosxrSNMPServerMIBConfig = `

resource "iosxr_snmp_server_mib" "test" {
	delete_mode = "attributes"
	ifmib_ifalias_long = true
	ifindex_persist = true
}

data "iosxr_snmp_server_mib" "test" {
	depends_on = [iosxr_snmp_server_mib.test]
}
`
