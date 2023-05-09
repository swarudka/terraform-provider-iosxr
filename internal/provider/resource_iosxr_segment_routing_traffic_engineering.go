// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"regexp"

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
	"github.com/netascode/terraform-provider-iosxr/internal/provider/client"
	"github.com/netascode/terraform-provider-iosxr/internal/provider/helpers"
)

var _ resource.Resource = (*SegmentRoutingTrafficEngineeringResource)(nil)

func NewSegmentRoutingTrafficEngineeringResource() resource.Resource {
	return &SegmentRoutingTrafficEngineeringResource{}
}

type SegmentRoutingTrafficEngineeringResource struct {
	client *client.Client
}

func (r *SegmentRoutingTrafficEngineeringResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_segment_routing_traffic_engineering"
}

func (r *SegmentRoutingTrafficEngineeringResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This resource can manage the Segment Routing Traffic Engineering configuration.",

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
			"on_demand_colors": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("On-demand color configuration").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"color": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Color").AddIntegerRangeDescription(1, 4294967295).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 4294967295),
							},
						},
						"srv6_enable": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("True only").String,
							Optional:            true,
						},
						"srv6_locator_locator_name": schema.StringAttribute{
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
						"source_address_source_address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Source address").String,
							Required:            true,
						},
						"source_address_ip_address_type": schema.StringAttribute{
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
						"effective_metric_metric_value_type_metric_value": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Integer value of metric").AddIntegerRangeDescription(0, 4294967295).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 4294967295),
							},
						},
						"effective_metric_metric_value_type_metric_type": schema.StringAttribute{
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
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 59),
								stringvalidator.RegexMatches(regexp.MustCompile(`[\w\-\.:,_@#%$\+=\|;]+`), ""),
							},
						},
						"srv6_enable": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("True only").String,
							Optional:            true,
						},
						"srv6_locator_locator_name": schema.StringAttribute{
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
						"source_address_source_address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Source address").String,
							Required:            true,
						},
						"source_address_ip_address_type": schema.StringAttribute{
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
						"policy_color_endpoint_end_point_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("End point type").AddStringEnumDescription("end-point-type-ipv4", "end-point-type-ipv6").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("end-point-type-ipv4", "end-point-type-ipv6"),
							},
						},
						"policy_color_endpoint_end_point_address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("End point address").String,
							Required:            true,
						},
					},
				},
			},
		},
	}
}

func (r *SegmentRoutingTrafficEngineeringResource) Configure(ctx context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*client.Client)
}

func (r *SegmentRoutingTrafficEngineeringResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SegmentRoutingTrafficEngineering

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.getPath()))

	// Create object
	body := plan.toBody(ctx)

	_, diags = r.client.Set(ctx, plan.Device.ValueString(), plan.getPath(), body, client.Update)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	emptyLeafsDelete := plan.getEmptyLeafsDelete(ctx)
	tflog.Debug(ctx, fmt.Sprintf("List of empty leafs to delete: %+v", emptyLeafsDelete))

	for _, i := range emptyLeafsDelete {
		_, diags = r.client.Set(ctx, plan.Device.ValueString(), i, "", client.Delete)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	plan.Id = types.StringValue(plan.getPath())

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.getPath()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *SegmentRoutingTrafficEngineeringResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SegmentRoutingTrafficEngineering

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

func (r *SegmentRoutingTrafficEngineeringResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state SegmentRoutingTrafficEngineering

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

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Id.ValueString()))

	// Update object
	body := plan.toBody(ctx)

	_, diags = r.client.Set(ctx, plan.Device.ValueString(), plan.getPath(), body, client.Update)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	deletedListItems := plan.getDeletedListItems(ctx, state)
	tflog.Debug(ctx, fmt.Sprintf("List items to delete: %+v", deletedListItems))

	for _, i := range deletedListItems {
		_, diags = r.client.Set(ctx, plan.Device.ValueString(), i, "", client.Delete)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	emptyLeafsDelete := plan.getEmptyLeafsDelete(ctx)
	tflog.Debug(ctx, fmt.Sprintf("List of empty leafs to delete: %+v", emptyLeafsDelete))

	for _, i := range emptyLeafsDelete {
		_, diags = r.client.Set(ctx, plan.Device.ValueString(), i, "", client.Delete)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *SegmentRoutingTrafficEngineeringResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state SegmentRoutingTrafficEngineering

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))

	_, diags = r.client.Set(ctx, state.Device.ValueString(), state.getPath(), "", client.Delete)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

func (r *SegmentRoutingTrafficEngineeringResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
