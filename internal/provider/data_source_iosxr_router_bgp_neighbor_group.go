// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/CiscoDevNet/terraform-provider-iosxr/internal/provider/client"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &RouterBGPNeighborGroupDataSource{}
	_ datasource.DataSourceWithConfigure = &RouterBGPNeighborGroupDataSource{}
)

func NewRouterBGPNeighborGroupDataSource() datasource.DataSource {
	return &RouterBGPNeighborGroupDataSource{}
}

type RouterBGPNeighborGroupDataSource struct {
	client *client.Client
}

func (d *RouterBGPNeighborGroupDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_router_bgp_neighbor_group"
}

func (d *RouterBGPNeighborGroupDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Router BGP Neighbor Group configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the retrieved object.",
				Computed:            true,
			},
			"as_number": schema.StringAttribute{
				MarkdownDescription: "bgp as-number",
				Required:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Specify a Neighbor-group",
				Required:            true,
			},
			"remote_as": schema.StringAttribute{
				MarkdownDescription: "bgp as-number",
				Computed:            true,
			},
			"update_source": schema.StringAttribute{
				MarkdownDescription: "Source of routing updates",
				Computed:            true,
			},
			"ao_key_chain_name": schema.StringAttribute{
				MarkdownDescription: "Name of the key chain - maximum 32 characters",
				Computed:            true,
			},
			"ao_include_tcp_options_enable": schema.BoolAttribute{
				MarkdownDescription: "Include other TCP options in the header",
				Computed:            true,
			},
			"bfd_minimum_interval": schema.Int64Attribute{
				MarkdownDescription: "Hello interval",
				Computed:            true,
			},
			"bfd_fast_detect": schema.BoolAttribute{
				MarkdownDescription: "Enable Fast detection",
				Computed:            true,
			},
			"address_families": schema.ListNestedAttribute{
				MarkdownDescription: "Enter Address Family command mode",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"af_name": schema.StringAttribute{
							MarkdownDescription: "Enter Address Family command mode",
							Computed:            true,
						},
						"soft_reconfiguration_inbound_always": schema.BoolAttribute{
							MarkdownDescription: "Always use soft reconfig, even if route refresh is supported",
							Computed:            true,
						},
						"next_hop_self_inheritance_disable": schema.BoolAttribute{
							MarkdownDescription: "Prevent next-hop-self from being inherited from the parent",
							Computed:            true,
						},
						"route_reflector_client_inheritance_disable": schema.BoolAttribute{
							MarkdownDescription: "Prevent route-reflector-client from being inherited from the parent",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *RouterBGPNeighborGroupDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*client.Client)
}

func (d *RouterBGPNeighborGroupDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config RouterBGPNeighborGroup

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.getPath()))

	getResp, diags := d.client.Get(ctx, config.Device.ValueString(), config.getPath())
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	config.fromBody(ctx, getResp.Notification[0].Update[0].Val.GetJsonIetfVal())
	config.Id = types.StringValue(config.getPath())

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.getPath()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}
