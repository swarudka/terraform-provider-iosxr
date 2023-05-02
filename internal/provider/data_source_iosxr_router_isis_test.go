// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIosxrRouterISIS(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrRouterISISConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxr_router_isis.test", "is_type", "level-1"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis.test", "set_overload_bit_levels.0.level_id", "1"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis.test", "set_overload_bit_levels.0.on_startup_advertise_as_overloaded", "true"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis.test", "set_overload_bit_levels.0.on_startup_advertise_as_overloaded_time_to_advertise", "10"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis.test", "set_overload_bit_levels.0.on_startup_wait_for_bgp", "false"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis.test", "set_overload_bit_levels.0.advertise_external", "true"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis.test", "set_overload_bit_levels.0.advertise_interlevel", "true"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis.test", "nsr", "true"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis.test", "nsf_cisco", "true"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis.test", "nsf_ietf", "false"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis.test", "nsf_lifetime", "10"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis.test", "nsf_interface_timer", "5"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis.test", "nsf_interface_expires", "2"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis.test", "log_adjacency_changes", "true"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis.test", "lsp_gen_interval_maximum_wait", "5000"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis.test", "lsp_gen_interval_initial_wait", "50"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis.test", "lsp_gen_interval_secondary_wait", "200"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis.test", "lsp_refresh_interval", "65000"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis.test", "max_lsp_lifetime", "65535"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis.test", "lsp_password_keychain", "ISIS-KEY"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis.test", "distribute_link_state_instance_id", "32"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis.test", "affinity_maps.0.name", "22"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis.test", "affinity_maps.0.bit_position", "4"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis.test", "flex_algos.0.algorithm_number", "128"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis.test", "flex_algos.0.advertise_definition", "true"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis.test", "flex_algos.0.metric_type_delay", "true"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis.test", "nets.0.net_id", "49.0001.2222.2222.2222.00"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis.test", "interfaces.0.interface_name", "GigabitEthernet0/0/0/1"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis.test", "interfaces.0.circuit_type", "level-1"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis.test", "interfaces.0.hello_padding_disable", "true"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis.test", "interfaces.0.hello_padding_sometimes", "false"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis.test", "interfaces.0.priority", "10"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis.test", "interfaces.0.point_to_point", "false"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis.test", "interfaces.0.passive", "false"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis.test", "interfaces.0.suppressed", "false"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis.test", "interfaces.0.shutdown", "false"),
				),
			},
		},
	})
}

const testAccDataSourceIosxrRouterISISConfig = `

resource "iosxr_router_isis" "test" {
	process_id = "P1"
	is_type = "level-1"
	set_overload_bit_levels = [{
		level_id = 1
		on_startup_advertise_as_overloaded = true
		on_startup_advertise_as_overloaded_time_to_advertise = 10
		on_startup_wait_for_bgp = false
		advertise_external = true
		advertise_interlevel = true
	}]
	nsr = true
	nsf_cisco = true
	nsf_ietf = false
	nsf_lifetime = 10
	nsf_interface_timer = 5
	nsf_interface_expires = 2
	log_adjacency_changes = true
	lsp_gen_interval_maximum_wait = 5000
	lsp_gen_interval_initial_wait = 50
	lsp_gen_interval_secondary_wait = 200
	lsp_refresh_interval = 65000
	max_lsp_lifetime = 65535
	lsp_password_keychain = "ISIS-KEY"
	distribute_link_state_instance_id = 32
	affinity_maps = [{
		name = "22"
		bit_position = 4
	}]
	flex_algos = [{
		algorithm_number = 128
		advertise_definition = true
		metric_type_delay = true
	}]
	nets = [{
		net_id = "49.0001.2222.2222.2222.00"
	}]
	interfaces = [{
		interface_name = "GigabitEthernet0/0/0/1"
		circuit_type = "level-1"
		hello_padding_disable = true
		hello_padding_sometimes = false
		priority = 10
		point_to_point = false
		passive = false
		suppressed = false
		shutdown = false
	}]
}

data "iosxr_router_isis" "test" {
	process_id = "P1"
	depends_on = [iosxr_router_isis.test]
}
`
