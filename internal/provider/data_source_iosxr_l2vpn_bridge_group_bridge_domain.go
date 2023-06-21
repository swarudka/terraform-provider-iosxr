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
	_ datasource.DataSource              = &L2VPNBridgeGroupBridgeDomainDataSource{}
	_ datasource.DataSourceWithConfigure = &L2VPNBridgeGroupBridgeDomainDataSource{}
)

func NewL2VPNBridgeGroupBridgeDomainDataSource() datasource.DataSource {
	return &L2VPNBridgeGroupBridgeDomainDataSource{}
}

type L2VPNBridgeGroupBridgeDomainDataSource struct {
	client *client.Client
}

func (d *L2VPNBridgeGroupBridgeDomainDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_l2vpn_bridge_group_bridge_domain"
}

func (d *L2VPNBridgeGroupBridgeDomainDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the L2VPN Bridge Group Bridge Domain configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the retrieved object.",
				Computed:            true,
			},
			"bridge_group_name": schema.StringAttribute{
				MarkdownDescription: "Specify the group the bridge belongs to",
				Required:            true,
			},
			"bridge_domain_name": schema.StringAttribute{
				MarkdownDescription: "Configure bridge domain",
				Required:            true,
			},
			"evis": schema.ListNestedAttribute{
				MarkdownDescription: "Ethernet VPN identifier",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"vpn_id": schema.Int64Attribute{
							MarkdownDescription: "Ethernet VPN identifier",
							Computed:            true,
						},
					},
				},
			},
			"vnis": schema.ListNestedAttribute{
				MarkdownDescription: "VxLAN VPN identifier",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"vni_id": schema.Int64Attribute{
							MarkdownDescription: "VxLAN VPN identifier",
							Computed:            true,
						},
					},
				},
			},
			"mtu": schema.Int64Attribute{
				MarkdownDescription: "Maximum transmission unit (payload) for this Bridge Domain",
				Computed:            true,
			},
			"storm_control_broadcast_pps": schema.Int64Attribute{
				MarkdownDescription: "Set the storm control pps",
				Computed:            true,
			},
			"storm_control_broadcast_kbps": schema.Int64Attribute{
				MarkdownDescription: "Set the storm control kbps",
				Computed:            true,
			},
			"storm_control_multicast_pps": schema.Int64Attribute{
				MarkdownDescription: "Set the storm control pps",
				Computed:            true,
			},
			"storm_control_multicast_kbps": schema.Int64Attribute{
				MarkdownDescription: "Set the storm control kbps",
				Computed:            true,
			},
			"storm_control_unknown_unicast_pps": schema.Int64Attribute{
				MarkdownDescription: "Set the storm control pps",
				Computed:            true,
			},
			"storm_control_unknown_unicast_kbps": schema.Int64Attribute{
				MarkdownDescription: "Set the storm control kbps",
				Computed:            true,
			},
			"interfaces": schema.ListNestedAttribute{
				MarkdownDescription: "Specify interface name",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"interface_name": schema.StringAttribute{
							MarkdownDescription: "Specify interface name",
							Computed:            true,
						},
						"split_horizon_group": schema.BoolAttribute{
							MarkdownDescription: "Configure split-horizon group",
							Computed:            true,
						},
					},
				},
			},
			"segment_routing_srv6_evis": schema.ListNestedAttribute{
				MarkdownDescription: "Ethernet VPN identifier for srv6",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"vpn_id": schema.Int64Attribute{
							MarkdownDescription: "Ethernet VPN identifier for srv6",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *L2VPNBridgeGroupBridgeDomainDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*client.Client)
}

func (d *L2VPNBridgeGroupBridgeDomainDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config L2VPNBridgeGroupBridgeDomain

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
