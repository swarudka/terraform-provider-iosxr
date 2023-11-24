// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

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
	_ datasource.DataSource              = &RouterBGPNeighborAddressFamilyDataSource{}
	_ datasource.DataSourceWithConfigure = &RouterBGPNeighborAddressFamilyDataSource{}
)

func NewRouterBGPNeighborAddressFamilyDataSource() datasource.DataSource {
	return &RouterBGPNeighborAddressFamilyDataSource{}
}

type RouterBGPNeighborAddressFamilyDataSource struct {
	client *client.Client
}

func (d *RouterBGPNeighborAddressFamilyDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_router_bgp_neighbor_address_family"
}

func (d *RouterBGPNeighborAddressFamilyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Router BGP Neighbor Address Family configuration.",

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
			"neighbor_address": schema.StringAttribute{
				MarkdownDescription: "Neighbor address",
				Required:            true,
			},
			"af_name": schema.StringAttribute{
				MarkdownDescription: "Enter Address Family command mode",
				Required:            true,
			},
			"import_stitching_rt_re_originate_stitching_rt": schema.BoolAttribute{
				MarkdownDescription: "Reoriginate imported routes by attaching stitching RTs",
				Computed:            true,
			},
			"route_reflector_client": schema.BoolAttribute{
				MarkdownDescription: "Configure a neighbor as Route Reflector client",
				Computed:            true,
			},
			"route_reflector_client_inheritance_disable": schema.BoolAttribute{
				MarkdownDescription: "Prevent route-reflector-client from being inherited from the parent",
				Computed:            true,
			},
			"advertise_vpnv4_unicast_enable_re_originated_stitching_rt": schema.BoolAttribute{
				MarkdownDescription: "Advertise re-originated and local routes with stitching Route-Targets",
				Computed:            true,
			},
			"next_hop_self_inheritance_disable": schema.BoolAttribute{
				MarkdownDescription: "Prevent next-hop-self from being inherited from the parent",
				Computed:            true,
			},
			"encapsulation_type_srv6": schema.BoolAttribute{
				MarkdownDescription: "SRv6 encapsulation",
				Computed:            true,
			},
			"route_policy_in": schema.StringAttribute{
				MarkdownDescription: "Apply route policy to inbound routes",
				Computed:            true,
			},
			"route_policy_out": schema.StringAttribute{
				MarkdownDescription: "Apply route policy to outbound routes",
				Computed:            true,
			},
		},
	}
}

func (d *RouterBGPNeighborAddressFamilyDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*client.Client)
}

func (d *RouterBGPNeighborAddressFamilyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config RouterBGPNeighborAddressFamilyData

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
