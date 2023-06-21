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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ resource.Resource = (*RouterStaticResource)(nil)

func NewRouterStaticResource() resource.Resource {
	return &RouterStaticResource{}
}

type RouterStaticResource struct {
	client *client.Client
}

func (r *RouterStaticResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_router_static"
}

func (r *RouterStaticResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This resource can manage the Router Static configuration.",

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
			"prefix_address": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Destination prefix").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?`), ""),
					stringvalidator.RegexMatches(regexp.MustCompile(`[0-9\.]*`), ""),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"prefix_length": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Destination prefix length").AddIntegerRangeDescription(0, 128).String,
				Required:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 128),
				},
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
			},
			"nexthop_interfaces": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Forwarding interface").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"interface_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Forwarding interface").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`[a-zA-Z0-9.:_/-]+`), ""),
							},
						},
						"description": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("description of the static route").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 30),
							},
						},
						"tag": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set tag for this route").AddIntegerRangeDescription(1, 4294967295).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 4294967295),
							},
						},
						"distance_metric": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Distance metric for this route").AddIntegerRangeDescription(1, 254).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 254),
							},
						},
					},
				},
			},
			"nexthop_interface_addresses": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Forwarding interface").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"interface_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Forwarding interface").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`[a-zA-Z0-9.:_/-]+`), ""),
							},
						},
						"address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Forwarding router's address").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?`), ""),
								stringvalidator.RegexMatches(regexp.MustCompile(`[0-9\.]*`), ""),
							},
						},
						"description": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("description of the static route").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 30),
							},
						},
						"tag": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set tag for this route").AddIntegerRangeDescription(1, 4294967295).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 4294967295),
							},
						},
						"distance_metric": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Distance metric for this route").AddIntegerRangeDescription(1, 254).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 254),
							},
						},
					},
				},
			},
			"nexthop_addresses": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Forwarding router's address").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Forwarding router's address").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?`), ""),
								stringvalidator.RegexMatches(regexp.MustCompile(`[0-9\.]*`), ""),
							},
						},
						"description": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("description of the static route").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 30),
							},
						},
						"tag": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set tag for this route").AddIntegerRangeDescription(1, 4294967295).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 4294967295),
							},
						},
						"distance_metric": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Distance metric for this route").AddIntegerRangeDescription(1, 254).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 254),
							},
						},
					},
				},
			},
		},
	}
}

func (r *RouterStaticResource) Configure(ctx context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*client.Client)
}

func (r *RouterStaticResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan RouterStatic

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

func (r *RouterStaticResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state RouterStatic

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

func (r *RouterStaticResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state RouterStatic

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

func (r *RouterStaticResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state RouterStatic

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))

	_, diags = r.client.Set(ctx, state.Device.ValueString(), client.SetOperation{Path: state.getPath(), Body: "", Operation: client.Delete})
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

func (r *RouterStaticResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
