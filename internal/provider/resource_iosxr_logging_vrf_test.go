// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIosxrLoggingVRF(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxrLoggingVRFConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxr_logging_vrf.test", "vrf_name", "VRF1"),
					resource.TestCheckResourceAttr("iosxr_logging_vrf.test", "host_ipv4_addresses.0.ipv4_address", "1.1.1.1"),
					resource.TestCheckResourceAttr("iosxr_logging_vrf.test", "host_ipv4_addresses.0.severity", "info"),
					resource.TestCheckResourceAttr("iosxr_logging_vrf.test", "host_ipv6_addresses.0.ipv6_address", "2001:0db8:85a3:0000:0000:8a2e:0370:7334"),
					resource.TestCheckResourceAttr("iosxr_logging_vrf.test", "host_ipv6_addresses.0.severity", "info"),
				),
			},
			{
				ResourceName:  "iosxr_logging_vrf.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XR-um-logging-cfg:logging/vrfs/vrf[vrf-name=VRF1]",
			},
		},
	})
}

func testAccIosxrLoggingVRFConfig_minimum() string {
	return `
	resource "iosxr_logging_vrf" "test" {
		vrf_name = "VRF1"
	}
	`
}

func testAccIosxrLoggingVRFConfig_all() string {
	return `
	resource "iosxr_logging_vrf" "test" {
		vrf_name = "VRF1"
		host_ipv4_addresses = [{
			ipv4_address = "1.1.1.1"
			severity = "info"
		}]
		host_ipv6_addresses = [{
			ipv6_address = "2001:0db8:85a3:0000:0000:8a2e:0370:7334"
			severity = "info"
		}]
	}
	`
}
