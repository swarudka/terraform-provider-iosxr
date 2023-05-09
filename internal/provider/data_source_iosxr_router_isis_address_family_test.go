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
					resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "metric_style_narrow", "false"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "metric_style_wide", "true"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "metric_style_transition", "false"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "metric_style_levels.0.level_id", "1"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "metric_style_levels.0.narrow", "false"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "metric_style_levels.0.wide", "true"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "metric_style_levels.0.transition", "false"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "default_information_originate", "true"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "fast_reroute_delay_interval", "300"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "fast_reroute_per_link_priority_limit_critical", "true"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "fast_reroute_per_link_priority_limit_high", "false"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "fast_reroute_per_link_priority_limit_medium", "false"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "fast_reroute_per_prefix_priority_limit_critical", "true"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "fast_reroute_per_prefix_priority_limit_high", "false"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "fast_reroute_per_prefix_priority_limit_medium", "false"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "microloop_avoidance_protected", "false"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "microloop_avoidance_segment_routing", "true"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "advertise_passive_only", "true"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "advertise_link_attributes", "true"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "mpls_ldp_auto_config", "false"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "mpls_traffic_eng_level_1_2", "false"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "mpls_traffic_eng_level_1", "false"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "mpls_traffic_eng_level_2_only", "true"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "spf_interval_maximum_wait", "5000"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "spf_interval_initial_wait", "50"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "spf_interval_secondary_wait", "200"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "segment_routing_mpls_sr_prefer", "true"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "maximum_redistributed_prefixes", "100"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "maximum_redistributed_prefixes_levels.0.level_id", "1"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "maximum_redistributed_prefixes_levels.0.maximum_prefixes", "1000"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "redistribute_id.0.instance_id", "CORE"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "redistribute_id.0.route_policy", "ROUTE_POLICY_1"),
				),
			},
		},
	})
}

const testAccDataSourceIosxrRouterISISAddressFamilyConfig = `

resource "iosxr_router_isis_address_family" "test" {
	process_id = "P1"
	af_name = "ipv6"
	saf_name = "unicast"
	metric_style_narrow = false
	metric_style_wide = true
	metric_style_transition = false
	metric_style_levels = [{
		level_id = 1
		narrow = false
		wide = true
		transition = false
	}]
	default_information_originate = true
	fast_reroute_delay_interval = 300
	fast_reroute_per_link_priority_limit_critical = true
	fast_reroute_per_link_priority_limit_high = false
	fast_reroute_per_link_priority_limit_medium = false
	fast_reroute_per_prefix_priority_limit_critical = true
	fast_reroute_per_prefix_priority_limit_high = false
	fast_reroute_per_prefix_priority_limit_medium = false
	microloop_avoidance_protected = false
	microloop_avoidance_segment_routing = true
	advertise_passive_only = true
	advertise_link_attributes = true
	mpls_ldp_auto_config = false
	mpls_traffic_eng_level_1_2 = false
	mpls_traffic_eng_level_1 = false
	mpls_traffic_eng_level_2_only = true
	spf_interval_maximum_wait = 5000
	spf_interval_initial_wait = 50
	spf_interval_secondary_wait = 200
	segment_routing_mpls_sr_prefer = true
	maximum_redistributed_prefixes = 100
	maximum_redistributed_prefixes_levels = [{
		level_id = 1
		maximum_prefixes = 1000
	}]
	redistribute_id = [{
		instance_id = "CORE"
		route_policy = "ROUTE_POLICY_1"
	}]
}

data "iosxr_router_isis_address_family" "test" {
	process_id = "P1"
	af_name = "ipv6"
	saf_name = "unicast"
	depends_on = [iosxr_router_isis_address_family.test]
}
`
