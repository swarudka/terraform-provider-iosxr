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

var _ resource.Resource = (*InterfaceResource)(nil)

func NewInterfaceResource() resource.Resource {
	return &InterfaceResource{}
}

type InterfaceResource struct {
	client *client.Client
}

func (r *InterfaceResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface"
}

func (r *InterfaceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This resource can manage the Interface configuration.",

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
			"interface_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Interface configuration subcommands").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`[a-zA-Z0-9.:_/-]+`), ""),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"l2transport": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("l2transport sub-interface").String,
				Optional:            true,
			},
			"point_to_point": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("point-to-point sub-interface").String,
				Optional:            true,
			},
			"multipoint": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("multipoint sub-interface").String,
				Optional:            true,
			},
			"dampening_decay_half_life_value": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Decay half life (in minutes)").AddIntegerRangeDescription(1, 45).String,
				Required:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 45),
				},
			},
			"ipv4_point_to_point": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable point-to-point handling for this interface.").String,
				Optional:            true,
			},
			"service_policy_input": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure a policy in the input direction").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Name of the service policy. Set 'input' for 'service-ipsec and 'service-gre' interfaces").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`[a-zA-Z0-9][a-zA-Z0-9._@$%+#:=<>\\-]{0,62}`), ""),
							},
						},
					},
				},
			},
			"service_policy_output": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("direction of service policy application").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Name of the service policy. Set 'output' for 'service-ipsec and 'service-gre' interfaces").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`[a-zA-Z0-9][a-zA-Z0-9._@$%+#:=<>\\-]{0,62}`), ""),
							},
						},
					},
				},
			},
			"bfd_mode_ietf": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Use IETF standard for BoB").String,
				Optional:            true,
			},
			"encapsulation_dot1q_vlan_id": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure first (outer) VLAN ID on the subinterface").AddIntegerRangeDescription(1, 4094).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 4094),
				},
			},
			"l2transport_encapsulation_dot1q_vlan_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Single VLAN id or start of VLAN range").String,
				Optional:            true,
			},
			"l2transport_encapsulation_dot1q_second_dot1q": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("End of VLAN range").String,
				Optional:            true,
			},
			"rewrite_ingress_tag_pop_one": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Remove outer tag only").String,
				Optional:            true,
			},
			"rewrite_ingress_tag_pop_two": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Remove two outermost tags").String,
				Optional:            true,
			},
			"shutdown": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("shutdown the given interface").String,
				Optional:            true,
			},
			"mtu": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set the MTU on an interface").AddIntegerRangeDescription(64, 65535).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(64, 65535),
				},
			},
			"bandwidth": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set the bandwidth of an interface").AddIntegerRangeDescription(0, 9223372036854775807).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 9223372036854775807),
				},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set description for this interface").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 1024),
				},
			},
			"load_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Specify interval for load calculation for an interface").String,
				Optional:            true,
			},
			"vrf": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set VRF in which the interface operates").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 32),
					stringvalidator.RegexMatches(regexp.MustCompile(`[\w\-\.:,_@#%$\+=\|;]+`), ""),
				},
			},
			"ipv4_address": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IP address").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?`), ""),
					stringvalidator.RegexMatches(regexp.MustCompile(`[0-9\.]*`), ""),
				},
			},
			"ipv4_netmask": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IP subnet mask").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?`), ""),
					stringvalidator.RegexMatches(regexp.MustCompile(`[0-9\.]*`), ""),
				},
			},
			"unnumbered": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable IP processing without an explicit address").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`[a-zA-Z0-9.:_/-]+`), ""),
				},
			},
			"ipv6_link_local_address": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IPv6 address").String,
				Optional:            true,
			},
			"ipv6_link_local_zone": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IPv6 address zone").AddDefaultValueDescription("0").String,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					helpers.StringDefaultModifier("0"),
				},
			},
			"ipv6_autoconfig": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable slaac on Mgmt interface").String,
				Optional:            true,
			},
			"ipv6_enable": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable IPv6 on interface").String,
				Optional:            true,
			},
			"ipv6_addresses": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IPv6 address").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IPv6 name or address").String,
							Required:            true,
						},
						"prefix_length": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Prefix length in bits").AddIntegerRangeDescription(0, 128).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 128),
							},
						},
						"zone": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IPv6 address zone").AddDefaultValueDescription("0").String,
							Optional:            true,
							Computed:            true,
							PlanModifiers: []planmodifier.String{
								helpers.StringDefaultModifier("0"),
							},
						},
					},
				},
			},
		},
	}
}

func (r *InterfaceResource) Configure(ctx context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*client.Client)
}

func (r *InterfaceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan Interface

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

func (r *InterfaceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state Interface

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

func (r *InterfaceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state Interface

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

func (r *InterfaceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state Interface

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

func (r *InterfaceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
