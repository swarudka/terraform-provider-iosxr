// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIosxrClassMap(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrClassMapConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxr_class_map.test", "match_any", "true"),
					resource.TestCheckResourceAttr("data.iosxr_class_map.test", "description", "description1"),
					resource.TestCheckResourceAttr("data.iosxr_class_map.test", "match_dscp_value", "46"),
					resource.TestCheckResourceAttr("data.iosxr_class_map.test", "match_mpls_experimental_topmost_label", "5"),
				),
			},
		},
	})
}

const testAccDataSourceIosxrClassMapConfig = `

resource "iosxr_class_map" "test" {
	class_map_name = "TEST"
	match_any = true
	description = "description1"
	match_dscp_value = "46"
	match_mpls_experimental_topmost_label = 5
}

data "iosxr_class_map" "test" {
	class_map_name = "TEST"
	depends_on = [iosxr_class_map.test]
}
`
