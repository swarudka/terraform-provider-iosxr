// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/terraform-provider-iosxr/internal/provider/client"
)

func New() provider.Provider {
	return &iosxrProvider{}
}

// provider satisfies the tfsdk.Provider interface and usually is included
// with all Resource and DataSource implementations.
type iosxrProvider struct{}

// providerData can be used to store data from the Terraform configuration.
type providerData struct {
	Username types.String         `tfsdk:"username"`
	Password types.String         `tfsdk:"password"`
	Host     types.String         `tfsdk:"host"`
	Devices  []providerDataDevice `tfsdk:"devices"`
}

type providerDataDevice struct {
	Name types.String `tfsdk:"name"`
	Host types.String `tfsdk:"host"`
}

// Metadata returns the provider type name.
func (p *iosxrProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "iosxr"
}

func (p *iosxrProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"username": schema.StringAttribute{
				MarkdownDescription: "Username for the IOS-XR device. This can also be set as the IOSXR_USERNAME environment variable.",
				Optional:            true,
			},
			"password": schema.StringAttribute{
				MarkdownDescription: "Password for the IOS-XR device. This can also be set as the IOSXR_PASSWORD environment variable.",
				Optional:            true,
				Sensitive:           true,
			},
			"host": schema.StringAttribute{
				MarkdownDescription: "IP or name of the Cisco IOS-XR device. Optionally a port can be added with `:12345`. The default port is `57400`. This can also be set as the IOSXR_HOST environment variable. If no `host` is provided, the `host` of the first device from the `devices` list is being used.",
				Optional:            true,
			},
			"devices": schema.ListNestedAttribute{
				MarkdownDescription: "This can be used to manage a list of devices from a single provider. All devices must use the same credentials. Each resource and data source has an optional attribute named `device`, which can then select a device by its name from this list.",
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: "Device name.",
							Required:            true,
						},
						"host": schema.StringAttribute{
							MarkdownDescription: "IP of the Cisco IOS-XR device.",
							Required:            true,
						},
					},
				},
			},
		},
	}
}

func (p *iosxrProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	// Retrieve provider data from configuration
	var config providerData
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// User must provide a username to the provider
	var username string
	if config.Username.IsUnknown() {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as username",
		)
		return
	}

	if config.Username.IsNull() {
		username = os.Getenv("IOSXR_USERNAME")
	} else {
		username = config.Username.ValueString()
	}

	if username == "" {
		// Error vs warning - empty value must stop execution
		resp.Diagnostics.AddError(
			"Unable to find username",
			"Username cannot be an empty string",
		)
		return
	}

	// User must provide a password to the provider
	var password string
	if config.Password.IsUnknown() {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as password",
		)
		return
	}

	if config.Password.IsNull() {
		password = os.Getenv("IOSXR_PASSWORD")
	} else {
		password = config.Password.ValueString()
	}

	if password == "" {
		// Error vs warning - empty value must stop execution
		resp.Diagnostics.AddError(
			"Unable to find password",
			"Password cannot be an empty string",
		)
		return
	}

	// User must provide a username to the provider
	var host string
	if config.Host.IsUnknown() {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as host",
		)
		return
	}

	if config.Host.IsNull() {
		host = os.Getenv("IOSXR_HOST")
		if host == "" && len(config.Devices) > 0 {
			host = config.Devices[0].Host.ValueString()
		}
	} else {
		host = config.Host.ValueString()
	}

	if host == "" {
		// Error vs warning - empty value must stop execution
		resp.Diagnostics.AddError(
			"Unable to find host",
			"Host cannot be an empty string",
		)
		return
	}

	client := client.NewClient()

	diags = client.AddTarget(ctx, "", host, username, password)
	resp.Diagnostics.Append(diags...)

	for _, device := range config.Devices {
		diags = client.AddTarget(ctx, device.Name.ValueString(), device.Host.ValueString(), username, password)
		resp.Diagnostics.Append(diags...)
	}

	resp.DataSourceData = &client
	resp.ResourceData = &client
}

func (p *iosxrProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewGnmiResource,
		NewBGPASFormatResource,
		NewEVPNResource,
		NewEVPNEVIResource,
		NewEVPNGroupResource,
		NewEVPNInterfaceResource,
		NewHostnameResource,
		NewInterfaceResource,
		NewL2VPNResource,
		NewL2VPNXconnectGroupP2PResource,
		NewMPLSLDPResource,
		NewOCSystemConfigResource,
		NewPrefixSetResource,
		NewRoutePolicyResource,
		NewRouterBGPResource,
		NewRouterBGPAddressFamilyResource,
		NewRouterBGPVRFResource,
		NewRouterBGPVRFAddressFamilyResource,
		NewRouterISISResource,
		NewRouterISISInterfaceAddressFamilyResource,
		NewRouterOSPFResource,
		NewRouterOSPFAreaInterfaceResource,
		NewRouterOSPFVRFResource,
		NewRouterOSPFVRFAreaInterfaceResource,
		NewSSHResource,
		NewVRFResource,
	}
}

func (p *iosxrProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewGnmiDataSource,
		NewBGPASFormatDataSource,
		NewEVPNDataSource,
		NewEVPNEVIDataSource,
		NewEVPNGroupDataSource,
		NewEVPNInterfaceDataSource,
		NewHostnameDataSource,
		NewInterfaceDataSource,
		NewL2VPNDataSource,
		NewL2VPNXconnectGroupP2PDataSource,
		NewMPLSLDPDataSource,
		NewOCSystemConfigDataSource,
		NewPrefixSetDataSource,
		NewRoutePolicyDataSource,
		NewRouterBGPDataSource,
		NewRouterBGPAddressFamilyDataSource,
		NewRouterBGPVRFDataSource,
		NewRouterBGPVRFAddressFamilyDataSource,
		NewRouterISISDataSource,
		NewRouterISISInterfaceAddressFamilyDataSource,
		NewRouterOSPFDataSource,
		NewRouterOSPFAreaInterfaceDataSource,
		NewRouterOSPFVRFDataSource,
		NewRouterOSPFVRFAreaInterfaceDataSource,
		NewSSHDataSource,
		NewVRFDataSource,
	}
}
