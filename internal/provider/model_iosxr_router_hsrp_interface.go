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
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type RouterHSRPInterface struct {
	Device                 types.String `tfsdk:"device"`
	Id                     types.String `tfsdk:"id"`
	DeleteMode             types.String `tfsdk:"delete_mode"`
	InterfaceName          types.String `tfsdk:"interface_name"`
	HsrpUseBia             types.Bool   `tfsdk:"hsrp_use_bia"`
	HsrpRedirectsDisable   types.Bool   `tfsdk:"hsrp_redirects_disable"`
	HsrpDelayMinimum       types.Int64  `tfsdk:"hsrp_delay_minimum"`
	HsrpDelayReload        types.Int64  `tfsdk:"hsrp_delay_reload"`
	HsrpBfdMinimumInterval types.Int64  `tfsdk:"hsrp_bfd_minimum_interval"`
	HsrpBfdMultiplier      types.Int64  `tfsdk:"hsrp_bfd_multiplier"`
	HsrpMacRefresh         types.Int64  `tfsdk:"hsrp_mac_refresh"`
}

type RouterHSRPInterfaceData struct {
	Device                 types.String `tfsdk:"device"`
	Id                     types.String `tfsdk:"id"`
	InterfaceName          types.String `tfsdk:"interface_name"`
	HsrpUseBia             types.Bool   `tfsdk:"hsrp_use_bia"`
	HsrpRedirectsDisable   types.Bool   `tfsdk:"hsrp_redirects_disable"`
	HsrpDelayMinimum       types.Int64  `tfsdk:"hsrp_delay_minimum"`
	HsrpDelayReload        types.Int64  `tfsdk:"hsrp_delay_reload"`
	HsrpBfdMinimumInterval types.Int64  `tfsdk:"hsrp_bfd_minimum_interval"`
	HsrpBfdMultiplier      types.Int64  `tfsdk:"hsrp_bfd_multiplier"`
	HsrpMacRefresh         types.Int64  `tfsdk:"hsrp_mac_refresh"`
}

func (data RouterHSRPInterface) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-router-hsrp-cfg:/router/hsrp/interfaces/interface[interface-name=%s]", data.InterfaceName.ValueString())
}

func (data RouterHSRPInterfaceData) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-router-hsrp-cfg:/router/hsrp/interfaces/interface[interface-name=%s]", data.InterfaceName.ValueString())
}

func (data RouterHSRPInterface) toBody(ctx context.Context) string {
	body := "{}"
	if !data.InterfaceName.IsNull() && !data.InterfaceName.IsUnknown() {
		body, _ = sjson.Set(body, "interface-name", data.InterfaceName.ValueString())
	}
	if !data.HsrpUseBia.IsNull() && !data.HsrpUseBia.IsUnknown() {
		if data.HsrpUseBia.ValueBool() {
			body, _ = sjson.Set(body, "hsrp.use-bia", map[string]string{})
		}
	}
	if !data.HsrpRedirectsDisable.IsNull() && !data.HsrpRedirectsDisable.IsUnknown() {
		if data.HsrpRedirectsDisable.ValueBool() {
			body, _ = sjson.Set(body, "hsrp.redirects.disable", map[string]string{})
		}
	}
	if !data.HsrpDelayMinimum.IsNull() && !data.HsrpDelayMinimum.IsUnknown() {
		body, _ = sjson.Set(body, "hsrp.delay.minimum", strconv.FormatInt(data.HsrpDelayMinimum.ValueInt64(), 10))
	}
	if !data.HsrpDelayReload.IsNull() && !data.HsrpDelayReload.IsUnknown() {
		body, _ = sjson.Set(body, "hsrp.delay.reload", strconv.FormatInt(data.HsrpDelayReload.ValueInt64(), 10))
	}
	if !data.HsrpBfdMinimumInterval.IsNull() && !data.HsrpBfdMinimumInterval.IsUnknown() {
		body, _ = sjson.Set(body, "hsrp.bfd.minimum-interval", strconv.FormatInt(data.HsrpBfdMinimumInterval.ValueInt64(), 10))
	}
	if !data.HsrpBfdMultiplier.IsNull() && !data.HsrpBfdMultiplier.IsUnknown() {
		body, _ = sjson.Set(body, "hsrp.bfd.multiplier", strconv.FormatInt(data.HsrpBfdMultiplier.ValueInt64(), 10))
	}
	if !data.HsrpMacRefresh.IsNull() && !data.HsrpMacRefresh.IsUnknown() {
		body, _ = sjson.Set(body, "hsrp.mac-refresh", strconv.FormatInt(data.HsrpMacRefresh.ValueInt64(), 10))
	}
	return body
}

func (data *RouterHSRPInterface) updateFromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "hsrp.use-bia"); !data.HsrpUseBia.IsNull() {
		if value.Exists() {
			data.HsrpUseBia = types.BoolValue(true)
		} else {
			data.HsrpUseBia = types.BoolValue(false)
		}
	} else {
		data.HsrpUseBia = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "hsrp.redirects.disable"); !data.HsrpRedirectsDisable.IsNull() {
		if value.Exists() {
			data.HsrpRedirectsDisable = types.BoolValue(true)
		} else {
			data.HsrpRedirectsDisable = types.BoolValue(false)
		}
	} else {
		data.HsrpRedirectsDisable = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "hsrp.delay.minimum"); value.Exists() && !data.HsrpDelayMinimum.IsNull() {
		data.HsrpDelayMinimum = types.Int64Value(value.Int())
	} else {
		data.HsrpDelayMinimum = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "hsrp.delay.reload"); value.Exists() && !data.HsrpDelayReload.IsNull() {
		data.HsrpDelayReload = types.Int64Value(value.Int())
	} else {
		data.HsrpDelayReload = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "hsrp.bfd.minimum-interval"); value.Exists() && !data.HsrpBfdMinimumInterval.IsNull() {
		data.HsrpBfdMinimumInterval = types.Int64Value(value.Int())
	} else {
		data.HsrpBfdMinimumInterval = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "hsrp.bfd.multiplier"); value.Exists() && !data.HsrpBfdMultiplier.IsNull() {
		data.HsrpBfdMultiplier = types.Int64Value(value.Int())
	} else {
		data.HsrpBfdMultiplier = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "hsrp.mac-refresh"); value.Exists() && !data.HsrpMacRefresh.IsNull() {
		data.HsrpMacRefresh = types.Int64Value(value.Int())
	} else {
		data.HsrpMacRefresh = types.Int64Null()
	}
}

func (data *RouterHSRPInterfaceData) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "hsrp.use-bia"); value.Exists() {
		data.HsrpUseBia = types.BoolValue(true)
	} else {
		data.HsrpUseBia = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "hsrp.redirects.disable"); value.Exists() {
		data.HsrpRedirectsDisable = types.BoolValue(true)
	} else {
		data.HsrpRedirectsDisable = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "hsrp.delay.minimum"); value.Exists() {
		data.HsrpDelayMinimum = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "hsrp.delay.reload"); value.Exists() {
		data.HsrpDelayReload = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "hsrp.bfd.minimum-interval"); value.Exists() {
		data.HsrpBfdMinimumInterval = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "hsrp.bfd.multiplier"); value.Exists() {
		data.HsrpBfdMultiplier = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "hsrp.mac-refresh"); value.Exists() {
		data.HsrpMacRefresh = types.Int64Value(value.Int())
	}
}

func (data *RouterHSRPInterface) getDeletedItems(ctx context.Context, state RouterHSRPInterface) []string {
	deletedItems := make([]string, 0)
	return deletedItems
}

func (data *RouterHSRPInterface) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	if !data.HsrpUseBia.IsNull() && !data.HsrpUseBia.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/hsrp/use-bia", data.getPath()))
	}
	if !data.HsrpRedirectsDisable.IsNull() && !data.HsrpRedirectsDisable.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/hsrp/redirects/disable", data.getPath()))
	}
	return emptyLeafsDelete
}

func (data *RouterHSRPInterface) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.HsrpUseBia.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/hsrp/use-bia", data.getPath()))
	}
	if !data.HsrpRedirectsDisable.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/hsrp/redirects/disable", data.getPath()))
	}
	if !data.HsrpDelayMinimum.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/hsrp/delay/minimum", data.getPath()))
	}
	if !data.HsrpDelayReload.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/hsrp/delay/reload", data.getPath()))
	}
	if !data.HsrpBfdMinimumInterval.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/hsrp/bfd/minimum-interval", data.getPath()))
	}
	if !data.HsrpBfdMultiplier.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/hsrp/bfd/multiplier", data.getPath()))
	}
	if !data.HsrpMacRefresh.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/hsrp/mac-refresh", data.getPath()))
	}
	return deletePaths
}
