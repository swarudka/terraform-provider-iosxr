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

func NewRouterBGPResource() resource.Resource {
	return &RouterBGPResource{}
}

type RouterBGPResource struct {
	client *client.Client
}

func (r *RouterBGPResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_router_bgp"
}

func (r *RouterBGPResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This resource can manage the Router BGP configuration.",

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
			"as_number": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("bgp as-number").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"default_information_originate": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Distribute a default route").String,
				Optional:            true,
			},
			"default_metric": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("default redistributed metric").AddIntegerRangeDescription(1, 4294967295).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 4294967295),
				},
			},
			"nsr_disable": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Disable non-stop-routing support for all neighbors").String,
				Optional:            true,
			},
			"bgp_redistribute_internal": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Allow redistribution of iBGP into IGPs (dangerous)").String,
				Optional:            true,
			},
			"segment_routing_srv6_locator": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure locator name").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 1024),
					stringvalidator.RegexMatches(regexp.MustCompile(`[\w\-\.:,_@#%$\+=\|;]+`), ""),
				},
			},
			"timers_bgp_keepalive_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("BGP timers").AddIntegerRangeDescription(0, 65535).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 65535),
				},
			},
			"timers_bgp_holdtime": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Holdtime. Set 0 to disable keepalives/hold time.").String,
				Optional:            true,
			},
			"timers_bgp_minimum_acceptable_holdtime": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Minimum acceptable holdtime from neighbor. Set 0 to disable keepalives/hold time.").String,
				Optional:            true,
			},
			"bgp_router_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure Router-id").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?`), ""),
					stringvalidator.RegexMatches(regexp.MustCompile(`[0-9\.]*`), ""),
				},
			},
			"bgp_graceful_restart_graceful_reset": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Reset gracefully if configuration change forces a peer reset").String,
				Optional:            true,
			},
			"ibgp_policy_out_enforce_modifications": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Allow policy to modify all attributes").String,
				Optional:            true,
			},
			"bgp_log_neighbor_changes_detail": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Include extra detail in change messages").String,
				Optional:            true,
			},
			"bfd_minimum_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Hello interval").AddIntegerRangeDescription(3, 30000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(3, 30000),
				},
			},
			"bfd_multiplier": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Detect multiplier").AddIntegerRangeDescription(2, 16).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(2, 16),
				},
			},
			"nexthop_validation_color_extcomm_sr_policy": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable BGP next-hop reachability validation by SR Policy for color-extcomm paths").String,
				Optional:            true,
			},
			"nexthop_validation_color_extcomm_disable": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Disable next-hop reachability validation for color-extcomm path").String,
				Optional:            true,
			},
			"bgp_bestpath_as_path_ignore": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Ignore as-path length").String,
				Optional:            true,
			},
			"bgp_bestpath_as_path_multipath_relax": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Relax as-path check for multipath selection").String,
				Optional:            true,
			},
			"bgp_bestpath_cost_community_ignore": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Ignore cost-community comparison").String,
				Optional:            true,
			},
			"bgp_bestpath_compare_routerid": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Compare router-id for identical EBGP paths").String,
				Optional:            true,
			},
			"bgp_bestpath_aigp_ignore": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Ignore AIGP attribute").String,
				Optional:            true,
			},
			"bgp_bestpath_igp_metric_ignore": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Ignore IGP metric during path comparison").String,
				Optional:            true,
			},
			"bgp_bestpath_igp_metric_sr_policy": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Use next-hop admin/metric from SR policy at Next Hop metric comparison stage").String,
				Optional:            true,
			},
			"bgp_bestpath_med_always": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Allow comparing MED from different neighbors").String,
				Optional:            true,
			},
			"bgp_bestpath_med_confed": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Compare MED among confederation paths").String,
				Optional:            true,
			},
			"bgp_bestpath_med_missing_as_worst": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Treat missing MED as the least preferred one").String,
				Optional:            true,
			},
			"bgp_bestpath_origin_as_use_validity": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("BGP bestpath selection will use origin-AS validity").String,
				Optional:            true,
			},
			"bgp_bestpath_origin_as_allow_invalid": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("BGP bestpath selection will allow 'invalid' origin-AS").String,
				Optional:            true,
			},
			"bgp_bestpath_sr_policy_prefer": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Consider only paths over SR Policy for bestpath selection, eBGP no-color eligible").String,
				Optional:            true,
			},
			"bgp_bestpath_sr_policy_force": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Consider only paths over SR Policy for bestpath selection, eBGP no-color ineligible").String,
				Optional:            true,
			},
			"neighbors": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Neighbor address").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"neighbor_address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Neighbor address").String,
							Required:            true,
						},
						"remote_as": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("bgp as-number").String,
							Optional:            true,
						},
						"description": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Neighbor specific description").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 1024),
							},
						},
						"use_neighbor_group": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Inherit configuration from a neighbor-group").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 1024),
							},
						},
						"advertisement_interval_seconds": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Minimum interval between sending BGP routing updates").AddIntegerRangeDescription(0, 600).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 600),
							},
						},
						"advertisement_interval_milliseconds": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("time in milliseconds").AddIntegerRangeDescription(0, 999).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 999),
							},
						},
						"ignore_connected_check": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Bypass the directly connected nexthop check for single-hop eBGP peering").String,
							Optional:            true,
						},
						"ebgp_multihop_maximum_hop_count": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("maximum hop count").AddIntegerRangeDescription(1, 255).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 255),
							},
						},
						"bfd_minimum_interval": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Hello interval").AddIntegerRangeDescription(3, 30000).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(3, 30000),
							},
						},
						"bfd_multiplier": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Detect multiplier").AddIntegerRangeDescription(2, 16).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(2, 16),
							},
						},
						"bfd_fast_detect": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable Fast detection").String,
							Optional:            true,
						},
						"bfd_fast_detect_strict_mode": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Hold down neighbor session until BFD session is up").String,
							Optional:            true,
						},
						"bfd_fast_detect_inheritance_disable": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Prevent bfd settings from being inherited from the parent").String,
							Optional:            true,
						},
						"local_as": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("bgp as-number").String,
							Optional:            true,
						},
						"local_as_no_prepend": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Do not prepend local AS to announcements from this neighbor").String,
							Optional:            true,
						},
						"local_as_replace_as": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Prepend only local AS to announcements to this neighbor").String,
							Optional:            true,
						},
						"local_as_dual_as": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Dual-AS mode").String,
							Optional:            true,
						},
						"password": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Specifies an ENCRYPTED password will follow").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`(!.+)|([^!].+)`), ""),
							},
						},
						"shutdown": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Administratively shut down this neighbor").String,
							Optional:            true,
						},
						"timers_keepalive_interval": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("BGP timers").AddIntegerRangeDescription(0, 65535).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 65535),
							},
						},
						"timers_holdtime": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Holdtime. Set 0 to disable keepalives/hold time.").String,
							Optional:            true,
						},
						"timers_minimum_acceptable_holdtime": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Minimum acceptable holdtime from neighbor. Set 0 to disable keepalives/hold time.").String,
							Optional:            true,
						},
						"update_source": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Source of routing updates").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`[a-zA-Z0-9.:_/-]+`), ""),
							},
						},
						"ttl_security": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable EBGP TTL security").String,
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *RouterBGPResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*client.Client)
}

func (r *RouterBGPResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan RouterBGP

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

func (r *RouterBGPResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state RouterBGP

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

func (r *RouterBGPResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state RouterBGP

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

func (r *RouterBGPResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state RouterBGP

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

func (r *RouterBGPResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
