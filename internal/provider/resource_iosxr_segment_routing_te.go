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

func NewSegmentRoutingTEResource() resource.Resource {
	return &SegmentRoutingTEResource{}
}

type SegmentRoutingTEResource struct {
	client *client.Client
}

func (r *SegmentRoutingTEResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_segment_routing_te"
}

func (r *SegmentRoutingTEResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This resource can manage the Segment Routing TE configuration.",

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
			"logging_pcep_peer_status": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable logging for pcep peer status").String,
				Optional:            true,
			},
			"logging_policy_status": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable logging for policy status").String,
				Optional:            true,
			},
			"pcc_report_all": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Report all local SR policies to connected PCEP peers").String,
				Optional:            true,
			},
			"pcc_source_address": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Local source IP address to use on PCEP sessions").String,
				Optional:            true,
			},
			"pcc_delegation_timeout": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Maximum time delegated SR-TE policies can remain up without an active connection to a PCE").AddIntegerRangeDescription(0, 1576800000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 1576800000),
				},
			},
			"pcc_dead_timer": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Amount of time after which the peer can declare this session down, if no PCEP message has been received").AddIntegerRangeDescription(1, 255).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 255),
				},
			},
			"pcc_initiated_state": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Amount of time that PCE initiated policy can exist as an orphan before it is cleaned up").AddIntegerRangeDescription(0, 86400).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 86400),
				},
			},
			"pcc_initiated_orphan": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Amount of time that PCE initiated policy remains delegated to a peer that has gone down").AddIntegerRangeDescription(0, 180).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 180),
				},
			},
			"pce_peers": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("PCE peer").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"pce_address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Remote PCE address").String,
							Required:            true,
						},
						"precedence": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Precedence value of this PCE").AddIntegerRangeDescription(0, 255).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 255),
							},
						},
					},
				},
			},
			"on_demand_colors": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("On-demand color configuration").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"dynamic_anycast_sid_inclusion": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Anycast Prefix SID Inclusion. Applicable for SR-MPLS and SRv6 policies").String,
							Optional:            true,
						},
						"dynamic_metric_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Metric Type").AddStringEnumDescription("hopcount", "igp", "latency", "te").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("hopcount", "igp", "latency", "te"),
							},
						},
						"color": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Color").AddIntegerRangeDescription(1, 4294967295).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 4294967295),
							},
						},
						"srv6_enable": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("True only").String,
							Optional:            true,
						},
						"srv6_locator_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("SRv6 locator name").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 64),
								stringvalidator.RegexMatches(regexp.MustCompile(`[\w\-\.:,_@#%$\+=\|;]+`), ""),
							},
						},
						"srv6_locator_behavior": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("SRv6 USID Behavior").AddStringEnumDescription("ub6-insert-reduced").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("ub6-insert-reduced"),
							},
						},
						"srv6_locator_binding_sid_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Binding Segment ID type").AddStringEnumDescription("srv6-dynamic").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("srv6-dynamic"),
							},
						},
						"source_address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Source address").String,
							Required:            true,
						},
						"source_address_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IP address type").AddStringEnumDescription("end-point-type-ipv4", "end-point-type-ipv6").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("end-point-type-ipv4", "end-point-type-ipv6"),
							},
						},
						"effective_metric_enable": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("True only").String,
							Optional:            true,
						},
						"effective_metric_value": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Integer value of metric").AddIntegerRangeDescription(0, 4294967295).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 4294967295),
							},
						},
						"effective_metric_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Metric type, advertised to other protocols").AddStringEnumDescription("default", "hopcount", "igp", "latency", "te").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("default", "hopcount", "igp", "latency", "te"),
							},
						},
						"constraint_segments_protection_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Protection Type").AddStringEnumDescription("protected-only", "protected-preferred", "unprotected-only", "unprotected-preferred").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("protected-only", "protected-preferred", "unprotected-only", "unprotected-preferred"),
							},
						},
						"constraint_segments_sid_algorithm": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("'0' for regular SIDs, '1' for strict-spf SIDs, '128' - '255' for algorithm SIDs").AddIntegerRangeDescription(0, 255).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 255),
							},
						},
					},
				},
			},
			"policies": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Policy configuration").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"policy_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Policy name").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 59),
								stringvalidator.RegexMatches(regexp.MustCompile(`[\w\-\.:,_@#%$\+=\|;]+`), ""),
							},
						},
						"srv6_enable": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("True only").String,
							Optional:            true,
						},
						"srv6_locator_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("SRv6 locator name").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 64),
								stringvalidator.RegexMatches(regexp.MustCompile(`[\w\-\.:,_@#%$\+=\|;]+`), ""),
							},
						},
						"srv6_locator_binding_sid_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Binding Segment ID type").AddStringEnumDescription("srv6-dynamic").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("srv6-dynamic"),
							},
						},
						"srv6_locator_behavior": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("SRv6 USID Behavior").AddStringEnumDescription("ub6-insert-reduced").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("ub6-insert-reduced"),
							},
						},
						"source_address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Source address").String,
							Required:            true,
						},
						"source_address_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IP address type").AddStringEnumDescription("end-point-type-ipv4", "end-point-type-ipv6").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("end-point-type-ipv4", "end-point-type-ipv6"),
							},
						},
						"policy_color_endpoint_color": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Color").AddIntegerRangeDescription(1, 4294967295).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 4294967295),
							},
						},
						"policy_color_endpoint_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("End point type").AddStringEnumDescription("end-point-type-ipv4", "end-point-type-ipv6").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("end-point-type-ipv4", "end-point-type-ipv6"),
							},
						},
						"policy_color_endpoint_address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("End point address").String,
							Required:            true,
						},
					},
				},
			},
		},
	}
}

func (r *SegmentRoutingTEResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*client.Client)
}

func (r *SegmentRoutingTEResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SegmentRoutingTE

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

func (r *SegmentRoutingTEResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SegmentRoutingTE

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

func (r *SegmentRoutingTEResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state SegmentRoutingTE

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

	deletedListItems := plan.getDeletedListItems(ctx, state)
	tflog.Debug(ctx, fmt.Sprintf("List items to delete: %+v", deletedListItems))

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

func (r *SegmentRoutingTEResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state SegmentRoutingTE

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

func (r *SegmentRoutingTEResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
