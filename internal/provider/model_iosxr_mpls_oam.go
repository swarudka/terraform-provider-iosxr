// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type MPLSOAM struct {
	Device                                        types.String `tfsdk:"device"`
	Id                                            types.String `tfsdk:"id"`
	DeleteMode                                    types.String `tfsdk:"delete_mode"`
	Oam                                           types.Bool   `tfsdk:"oam"`
	OamEchoDisableVendorExtension                 types.Bool   `tfsdk:"oam_echo_disable_vendor_extension"`
	OamEchoReplyModeControlChannelAllowReverseLsp types.Bool   `tfsdk:"oam_echo_reply_mode_control_channel_allow_reverse_lsp"`
	OamDpmPps                                     types.Int64  `tfsdk:"oam_dpm_pps"`
	OamDpmInterval                                types.Int64  `tfsdk:"oam_dpm_interval"`
}

type MPLSOAMData struct {
	Device                                        types.String `tfsdk:"device"`
	Id                                            types.String `tfsdk:"id"`
	Oam                                           types.Bool   `tfsdk:"oam"`
	OamEchoDisableVendorExtension                 types.Bool   `tfsdk:"oam_echo_disable_vendor_extension"`
	OamEchoReplyModeControlChannelAllowReverseLsp types.Bool   `tfsdk:"oam_echo_reply_mode_control_channel_allow_reverse_lsp"`
	OamDpmPps                                     types.Int64  `tfsdk:"oam_dpm_pps"`
	OamDpmInterval                                types.Int64  `tfsdk:"oam_dpm_interval"`
}

func (data MPLSOAM) getPath() string {
	return "Cisco-IOS-XR-um-mpls-oam-cfg:/mpls"
}

func (data MPLSOAMData) getPath() string {
	return "Cisco-IOS-XR-um-mpls-oam-cfg:/mpls"
}

func (data MPLSOAM) toBody(ctx context.Context) string {
	body := "{}"
	if !data.Oam.IsNull() && !data.Oam.IsUnknown() {
		if data.Oam.ValueBool() {
			body, _ = sjson.Set(body, "oam", map[string]string{})
		}
	}
	if !data.OamEchoDisableVendorExtension.IsNull() && !data.OamEchoDisableVendorExtension.IsUnknown() {
		if data.OamEchoDisableVendorExtension.ValueBool() {
			body, _ = sjson.Set(body, "oam.echo.disable-vendor-extension", map[string]string{})
		}
	}
	if !data.OamEchoReplyModeControlChannelAllowReverseLsp.IsNull() && !data.OamEchoReplyModeControlChannelAllowReverseLsp.IsUnknown() {
		if data.OamEchoReplyModeControlChannelAllowReverseLsp.ValueBool() {
			body, _ = sjson.Set(body, "oam.echo.reply-mode.control-channel.allow-reverse-lsp", map[string]string{})
		}
	}
	if !data.OamDpmPps.IsNull() && !data.OamDpmPps.IsUnknown() {
		body, _ = sjson.Set(body, "oam.dpm.pps", strconv.FormatInt(data.OamDpmPps.ValueInt64(), 10))
	}
	if !data.OamDpmInterval.IsNull() && !data.OamDpmInterval.IsUnknown() {
		body, _ = sjson.Set(body, "oam.dpm.interval", strconv.FormatInt(data.OamDpmInterval.ValueInt64(), 10))
	}
	return body
}

func (data *MPLSOAM) updateFromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "oam"); !data.Oam.IsNull() {
		if value.Exists() {
			data.Oam = types.BoolValue(true)
		} else {
			data.Oam = types.BoolValue(false)
		}
	} else {
		data.Oam = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "oam.echo.disable-vendor-extension"); !data.OamEchoDisableVendorExtension.IsNull() {
		if value.Exists() {
			data.OamEchoDisableVendorExtension = types.BoolValue(true)
		} else {
			data.OamEchoDisableVendorExtension = types.BoolValue(false)
		}
	} else {
		data.OamEchoDisableVendorExtension = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "oam.echo.reply-mode.control-channel.allow-reverse-lsp"); !data.OamEchoReplyModeControlChannelAllowReverseLsp.IsNull() {
		if value.Exists() {
			data.OamEchoReplyModeControlChannelAllowReverseLsp = types.BoolValue(true)
		} else {
			data.OamEchoReplyModeControlChannelAllowReverseLsp = types.BoolValue(false)
		}
	} else {
		data.OamEchoReplyModeControlChannelAllowReverseLsp = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "oam.dpm.pps"); value.Exists() && !data.OamDpmPps.IsNull() {
		data.OamDpmPps = types.Int64Value(value.Int())
	} else {
		data.OamDpmPps = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "oam.dpm.interval"); value.Exists() && !data.OamDpmInterval.IsNull() {
		data.OamDpmInterval = types.Int64Value(value.Int())
	} else {
		data.OamDpmInterval = types.Int64Null()
	}
}

func (data *MPLSOAMData) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "oam"); value.Exists() {
		data.Oam = types.BoolValue(true)
	} else {
		data.Oam = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "oam.echo.disable-vendor-extension"); value.Exists() {
		data.OamEchoDisableVendorExtension = types.BoolValue(true)
	} else {
		data.OamEchoDisableVendorExtension = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "oam.echo.reply-mode.control-channel.allow-reverse-lsp"); value.Exists() {
		data.OamEchoReplyModeControlChannelAllowReverseLsp = types.BoolValue(true)
	} else {
		data.OamEchoReplyModeControlChannelAllowReverseLsp = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "oam.dpm.pps"); value.Exists() {
		data.OamDpmPps = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "oam.dpm.interval"); value.Exists() {
		data.OamDpmInterval = types.Int64Value(value.Int())
	}
}

func (data *MPLSOAM) getDeletedListItems(ctx context.Context, state MPLSOAM) []string {
	deletedListItems := make([]string, 0)
	return deletedListItems
}

func (data *MPLSOAM) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	if !data.Oam.IsNull() && !data.Oam.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/oam", data.getPath()))
	}
	if !data.OamEchoDisableVendorExtension.IsNull() && !data.OamEchoDisableVendorExtension.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/oam/echo/disable-vendor-extension", data.getPath()))
	}
	if !data.OamEchoReplyModeControlChannelAllowReverseLsp.IsNull() && !data.OamEchoReplyModeControlChannelAllowReverseLsp.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/oam/echo/reply-mode/control-channel/allow-reverse-lsp", data.getPath()))
	}
	return emptyLeafsDelete
}

func (data *MPLSOAM) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.Oam.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/oam", data.getPath()))
	}
	if !data.OamEchoDisableVendorExtension.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/oam/echo/disable-vendor-extension", data.getPath()))
	}
	if !data.OamEchoReplyModeControlChannelAllowReverseLsp.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/oam/echo/reply-mode/control-channel/allow-reverse-lsp", data.getPath()))
	}
	if !data.OamDpmPps.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/oam/dpm/pps", data.getPath()))
	}
	if !data.OamDpmInterval.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/oam/dpm/interval", data.getPath()))
	}
	return deletePaths
}
