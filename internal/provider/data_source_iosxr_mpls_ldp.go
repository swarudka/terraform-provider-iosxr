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
	_ datasource.DataSource              = &MPLSLDPDataSource{}
	_ datasource.DataSourceWithConfigure = &MPLSLDPDataSource{}
)

func NewMPLSLDPDataSource() datasource.DataSource {
	return &MPLSLDPDataSource{}
}

type MPLSLDPDataSource struct {
	client *client.Client
}

func (d *MPLSLDPDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mpls_ldp"
}

func (d *MPLSLDPDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the MPLS LDP configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the retrieved object.",
				Computed:            true,
			},
			"router_id": schema.StringAttribute{
				MarkdownDescription: "Configure router Id",
				Computed:            true,
			},
			"address_families": schema.ListNestedAttribute{
				MarkdownDescription: "Configure Address Family and its parameters",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"af_name": schema.StringAttribute{
							MarkdownDescription: "Configure Address Family and its parameters",
							Computed:            true,
						},
					},
				},
			},
			"interfaces": schema.ListNestedAttribute{
				MarkdownDescription: "Enable LDP on an interface and enter interface submode",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"interface_name": schema.StringAttribute{
							MarkdownDescription: "Enable LDP on an interface and enter interface submode",
							Computed:            true,
						},
					},
				},
			},
			"capabilities_sac_ipv4_disable": schema.BoolAttribute{
				MarkdownDescription: "Disable exchanging IPv4 prefix label bindings",
				Computed:            true,
			},
			"capabilities_sac_ipv6_disable": schema.BoolAttribute{
				MarkdownDescription: "Disable exchanging IPv6 prefix label bindings",
				Computed:            true,
			},
			"capabilities_sac_fec128_disable": schema.BoolAttribute{
				MarkdownDescription: "Disable exchanging PW FEC128 label bindings",
				Computed:            true,
			},
			"capabilities_sac_fec129_disable": schema.BoolAttribute{
				MarkdownDescription: "Disable exchanging PW FEC129 label bindings",
				Computed:            true,
			},
			"mldp_logging_notifications": schema.BoolAttribute{
				MarkdownDescription: "MLDP logging notifications",
				Computed:            true,
			},
			"mldp_address_families": schema.ListNestedAttribute{
				MarkdownDescription: "Configure Address Family and its parameters",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: "Configure Address Family and its parameters",
							Computed:            true,
						},
						"make_before_break_delay": schema.Int64Attribute{
							MarkdownDescription: "MBB delay",
							Computed:            true,
						},
						"forwarding_recursive": schema.BoolAttribute{
							MarkdownDescription: "Enable recursive forwarding",
							Computed:            true,
						},
						"forwarding_recursive_route_policy": schema.StringAttribute{
							MarkdownDescription: "Route policy",
							Computed:            true,
						},
						"recursive_fec": schema.BoolAttribute{
							MarkdownDescription: "MLDP Recursive FEC enable",
							Computed:            true,
						},
					},
				},
			},
			"session_protection": schema.BoolAttribute{
				MarkdownDescription: "Configure session protection parameters",
				Computed:            true,
			},
		},
	}
}

func (d *MPLSLDPDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*client.Client)
}

func (d *MPLSLDPDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config MPLSLDP

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
