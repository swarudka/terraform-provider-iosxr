// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"regexp"

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

var _ resource.Resource = (*RouterBGPVRFNeighborAddressFamilyResource)(nil)

func NewRouterBGPVRFNeighborAddressFamilyResource() resource.Resource {
	return &RouterBGPVRFNeighborAddressFamilyResource{}
}

type RouterBGPVRFNeighborAddressFamilyResource struct {
	client *client.Client
}

func (r *RouterBGPVRFNeighborAddressFamilyResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_router_bgp_vrf_neighbor_address_family"
}

func (r *RouterBGPVRFNeighborAddressFamilyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This resource can manage the Router BGP VRF Neighbor Address Family configuration.",

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
			"as_number": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("bgp as-number").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"vrf_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Specify a vrf name").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 32),
					stringvalidator.RegexMatches(regexp.MustCompile(`[\w\-\.:,_@#%$\+=\|;]+`), ""),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"neighbor_address": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Neighbor address").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"af_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enter Address Family command mode").AddStringEnumDescription("all-address-family", "ipv4-flowspec", "ipv4-labeled-unicast", "ipv4-mdt", "ipv4-multicast", "ipv4-mvpn", "ipv4-rt-filter", "ipv4-sr-policy", "ipv4-tunnel", "ipv4-unicast", "ipv6-flowspec", "ipv6-labeled-unicast", "ipv6-multicast", "ipv6-mvpn", "ipv6-sr-policy", "ipv6-unicast", "l2vpn-evpn", "l2vpn-mspw", "l2vpn-vpls-vpws", "link-state-link-state", "vpnv4-flowspec", "vpnv4-multicast", "vpnv4-unicast", "vpnv6-flowspec", "vpnv6-multicast", "vpnv6-unicast").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("all-address-family", "ipv4-flowspec", "ipv4-labeled-unicast", "ipv4-mdt", "ipv4-multicast", "ipv4-mvpn", "ipv4-rt-filter", "ipv4-sr-policy", "ipv4-tunnel", "ipv4-unicast", "ipv6-flowspec", "ipv6-labeled-unicast", "ipv6-multicast", "ipv6-mvpn", "ipv6-sr-policy", "ipv6-unicast", "l2vpn-evpn", "l2vpn-mspw", "l2vpn-vpls-vpws", "link-state-link-state", "vpnv4-flowspec", "vpnv4-multicast", "vpnv4-unicast", "vpnv6-flowspec", "vpnv6-multicast", "vpnv6-unicast"),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"route_policy_retention_route_policy_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Apply retention policy to inbound routes").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 255),
				},
			},
			"route_policy_in": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Apply route policy to inbound routes").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 255),
				},
			},
			"route_policy_out": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Apply route policy to outbound routes").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 255),
				},
			},
			"default_originate_route_policy": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Route policy to specify criteria to originate default").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 255),
				},
			},
			"default_originate_inheritance_disable": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Prevent default-originate being inherited from a parent group").String,
				Optional:            true,
			},
			"next_hop_self_inheritance_disable": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Prevent next-hop-self from being inherited from the parent").String,
				Optional:            true,
			},
			"soft_reconfiguration_inbound_always": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Always use soft reconfig, even if route refresh is supported").String,
				Optional:            true,
			},
			"send_community_ebgp_inheritance_disable": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Prevent send-community-ebgp from being inherited from the parent").String,
				Optional:            true,
			},
			"remove_private_as_inheritance_disable": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Prevent remove-private-AS from being inherited from the parent").String,
				Optional:            true,
			},
		},
	}
}

func (r *RouterBGPVRFNeighborAddressFamilyResource) Configure(ctx context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*client.Client)
}

func (r *RouterBGPVRFNeighborAddressFamilyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan RouterBGPVRFNeighborAddressFamily

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

func (r *RouterBGPVRFNeighborAddressFamilyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state RouterBGPVRFNeighborAddressFamily

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

func (r *RouterBGPVRFNeighborAddressFamilyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state RouterBGPVRFNeighborAddressFamily

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

func (r *RouterBGPVRFNeighborAddressFamilyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state RouterBGPVRFNeighborAddressFamily

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

func (r *RouterBGPVRFNeighborAddressFamilyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
