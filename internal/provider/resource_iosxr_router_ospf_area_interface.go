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
	"regexp"

	"github.com/CiscoDevNet/terraform-provider-iosxr/internal/provider/client"
	"github.com/CiscoDevNet/terraform-provider-iosxr/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func NewRouterOSPFAreaInterfaceResource() resource.Resource {
	return &RouterOSPFAreaInterfaceResource{}
}

type RouterOSPFAreaInterfaceResource struct {
	client *client.Client
}

func (r *RouterOSPFAreaInterfaceResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_router_ospf_area_interface"
}

func (r *RouterOSPFAreaInterfaceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This resource can manage the Router OSPF Area Interface configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the object.",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"delete_mode": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure behavior when deleting/destroying the resource. Either delete the entire object (YANG container) being managed, or only delete the individual resource attributes configured explicitly and leave everything else as-is. Default value is `all`.").AddStringEnumDescription("all", "attributes").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("all", "attributes"),
				},
			},
			"process_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name for this OSPF process").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 32),
					stringvalidator.RegexMatches(regexp.MustCompile(`[\w\-\.:,_@#%$\+=\|;]+`), ""),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"area_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enter the OSPF area configuration submode").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"interface_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable routing on an interface ").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`[a-zA-Z0-9.:_/-]+`), ""),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"network_broadcast": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Specify OSPF broadcast multi-access network").String,
				Optional:            true,
			},
			"network_non_broadcast": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Specify OSPF NBMA network").String,
				Optional:            true,
			},
			"network_point_to_point": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Specify OSPF point-to-point network").String,
				Optional:            true,
			},
			"network_point_to_multipoint": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Specify OSPF point-to-multipoint network").String,
				Optional:            true,
			},
			"cost": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Interface cost").AddIntegerRangeDescription(1, 65535).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
				},
			},
			"priority": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Router priority").AddIntegerRangeDescription(0, 255).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 255),
				},
			},
			"passive_enable": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable passive").String,
				Optional:            true,
			},
			"passive_disable": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Disable passive").String,
				Optional:            true,
			},
			"fast_reroute_per_prefix_ti_lfa": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable TI LFA computation").String,
				Optional:            true,
			},
			"fast_reroute_per_prefix_tiebreaker_srlg_disjoint": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set preference order among tiebreakers").AddIntegerRangeDescription(1, 255).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 255),
				},
			},
			"fast_reroute_per_prefix_tiebreaker_node_protecting": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set preference order among tiebreakers").AddIntegerRangeDescription(1, 255).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 255),
				},
			},
			"prefix_sid_strict_spf_index": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("SID Index").AddIntegerRangeDescription(0, 1048575).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 1048575),
				},
			},
			"prefix_sid_algorithms": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Algorithm Specific Prefix SID Configuration").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"algorithm_number": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Algorithm Specific Prefix SID Configuration").AddIntegerRangeDescription(128, 255).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(128, 255),
							},
						},
						"index_sid_index": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("SID Index").AddIntegerRangeDescription(0, 1048575).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 1048575),
							},
						},
						"index_explicit_null": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Force penultimate hop to send explicit-null label").String,
							Optional:            true,
						},
						"index_n_flag_clear": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Not a node SID (e.g. for anycast SID use)").String,
							Optional:            true,
						},
						"absolute_sid_label": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("SID value").AddIntegerRangeDescription(16000, 1048575).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(16000, 1048575),
							},
						},
						"absolute_explicit_null": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Force penultimate hop to send explicit-null label").String,
							Optional:            true,
						},
						"absolute_n_flag_clear": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Not a node SID (e.g. for anycast SID use)").String,
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *RouterOSPFAreaInterfaceResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*client.Client)
}

func (r *RouterOSPFAreaInterfaceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan RouterOSPFAreaInterface

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var ops []client.SetOperation

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.getPath()))

	// Create object
	body := plan.toBody(ctx)
	ops = append(ops, client.SetOperation{Path: plan.getPath(), Body: body, Operation: client.Update})

	emptyLeafsDelete := plan.getEmptyLeafsDelete(ctx)
	tflog.Debug(ctx, fmt.Sprintf("List of empty leafs to delete: %+v", emptyLeafsDelete))

	for _, i := range emptyLeafsDelete {
		ops = append(ops, client.SetOperation{Path: i, Body: "", Operation: client.Delete})
	}

	_, diags = r.client.Set(ctx, plan.Device.ValueString(), ops...)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan.Id = types.StringValue(plan.getPath())

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.getPath()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *RouterOSPFAreaInterfaceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state RouterOSPFAreaInterface

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.ValueString()))

	getResp, diags := r.client.Get(ctx, state.Device.ValueString(), state.Id.ValueString())
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	state.updateFromBody(ctx, getResp.Notification[0].Update[0].Val.GetJsonIetfVal())

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *RouterOSPFAreaInterfaceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state RouterOSPFAreaInterface

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Read state
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var ops []client.SetOperation

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Id.ValueString()))

	// Update object
	body := plan.toBody(ctx)
	ops = append(ops, client.SetOperation{Path: plan.getPath(), Body: body, Operation: client.Update})

	deletedListItems := plan.getDeletedItems(ctx, state)
	tflog.Debug(ctx, fmt.Sprintf("Removed items to delete: %+v", deletedListItems))

	for _, i := range deletedListItems {
		ops = append(ops, client.SetOperation{Path: i, Body: "", Operation: client.Delete})
	}

	emptyLeafsDelete := plan.getEmptyLeafsDelete(ctx)
	tflog.Debug(ctx, fmt.Sprintf("List of empty leafs to delete: %+v", emptyLeafsDelete))

	for _, i := range emptyLeafsDelete {
		ops = append(ops, client.SetOperation{Path: i, Body: "", Operation: client.Delete})
	}

	_, diags = r.client.Set(ctx, plan.Device.ValueString(), ops...)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *RouterOSPFAreaInterfaceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state RouterOSPFAreaInterface

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	var ops []client.SetOperation
	deleteMode := "all"
	if state.DeleteMode.ValueString() == "all" {
		deleteMode = "all"
	} else if state.DeleteMode.ValueString() == "attributes" {
		deleteMode = "attributes"
	}

	if deleteMode == "all" {
		ops = append(ops, client.SetOperation{Path: state.Id.ValueString(), Body: "", Operation: client.Delete})
	} else {
		deletePaths := state.getDeletePaths(ctx)
		tflog.Debug(ctx, fmt.Sprintf("Paths to delete: %+v", deletePaths))

		for _, i := range deletePaths {
			ops = append(ops, client.SetOperation{Path: i, Body: "", Operation: client.Delete})
		}
	}

	_, diags = r.client.Set(ctx, state.Device.ValueString(), ops...)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

func (r *RouterOSPFAreaInterfaceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
