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
	"reflect"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type LoggingVRF struct {
	Device            types.String                  `tfsdk:"device"`
	Id                types.String                  `tfsdk:"id"`
	DeleteMode        types.String                  `tfsdk:"delete_mode"`
	VrfName           types.String                  `tfsdk:"vrf_name"`
	HostIpv4Addresses []LoggingVRFHostIpv4Addresses `tfsdk:"host_ipv4_addresses"`
	HostIpv6Addresses []LoggingVRFHostIpv6Addresses `tfsdk:"host_ipv6_addresses"`
}

type LoggingVRFData struct {
	Device            types.String                  `tfsdk:"device"`
	Id                types.String                  `tfsdk:"id"`
	VrfName           types.String                  `tfsdk:"vrf_name"`
	HostIpv4Addresses []LoggingVRFHostIpv4Addresses `tfsdk:"host_ipv4_addresses"`
	HostIpv6Addresses []LoggingVRFHostIpv6Addresses `tfsdk:"host_ipv6_addresses"`
}
type LoggingVRFHostIpv4Addresses struct {
	Ipv4Address types.String `tfsdk:"ipv4_address"`
	Severity    types.String `tfsdk:"severity"`
	Port        types.Int64  `tfsdk:"port"`
	Operator    types.String `tfsdk:"operator"`
}
type LoggingVRFHostIpv6Addresses struct {
	Ipv6Address types.String `tfsdk:"ipv6_address"`
	Severity    types.String `tfsdk:"severity"`
	Port        types.Int64  `tfsdk:"port"`
	Operator    types.String `tfsdk:"operator"`
}

func (data LoggingVRF) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-logging-cfg:/logging/vrfs/vrf[vrf-name=%s]", data.VrfName.ValueString())
}

func (data LoggingVRFData) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-logging-cfg:/logging/vrfs/vrf[vrf-name=%s]", data.VrfName.ValueString())
}

func (data LoggingVRF) toBody(ctx context.Context) string {
	body := "{}"
	if !data.VrfName.IsNull() && !data.VrfName.IsUnknown() {
		body, _ = sjson.Set(body, "vrf-name", data.VrfName.ValueString())
	}
	if len(data.HostIpv4Addresses) > 0 {
		body, _ = sjson.Set(body, "host-ipv4-addresses.host-ipv4-address", []interface{}{})
		for index, item := range data.HostIpv4Addresses {
			if !item.Ipv4Address.IsNull() && !item.Ipv4Address.IsUnknown() {
				body, _ = sjson.Set(body, "host-ipv4-addresses.host-ipv4-address"+"."+strconv.Itoa(index)+"."+"ipv4-address", item.Ipv4Address.ValueString())
			}
			if !item.Severity.IsNull() && !item.Severity.IsUnknown() {
				body, _ = sjson.Set(body, "host-ipv4-addresses.host-ipv4-address"+"."+strconv.Itoa(index)+"."+"severity", item.Severity.ValueString())
			}
			if !item.Port.IsNull() && !item.Port.IsUnknown() {
				body, _ = sjson.Set(body, "host-ipv4-addresses.host-ipv4-address"+"."+strconv.Itoa(index)+"."+"port", strconv.FormatInt(item.Port.ValueInt64(), 10))
			}
			if !item.Operator.IsNull() && !item.Operator.IsUnknown() {
				body, _ = sjson.Set(body, "host-ipv4-addresses.host-ipv4-address"+"."+strconv.Itoa(index)+"."+"operator", item.Operator.ValueString())
			}
		}
	}
	if len(data.HostIpv6Addresses) > 0 {
		body, _ = sjson.Set(body, "host-ipv6-addresses.host-ipv6-address", []interface{}{})
		for index, item := range data.HostIpv6Addresses {
			if !item.Ipv6Address.IsNull() && !item.Ipv6Address.IsUnknown() {
				body, _ = sjson.Set(body, "host-ipv6-addresses.host-ipv6-address"+"."+strconv.Itoa(index)+"."+"ipv6-address", item.Ipv6Address.ValueString())
			}
			if !item.Severity.IsNull() && !item.Severity.IsUnknown() {
				body, _ = sjson.Set(body, "host-ipv6-addresses.host-ipv6-address"+"."+strconv.Itoa(index)+"."+"severity", item.Severity.ValueString())
			}
			if !item.Port.IsNull() && !item.Port.IsUnknown() {
				body, _ = sjson.Set(body, "host-ipv6-addresses.host-ipv6-address"+"."+strconv.Itoa(index)+"."+"port", strconv.FormatInt(item.Port.ValueInt64(), 10))
			}
			if !item.Operator.IsNull() && !item.Operator.IsUnknown() {
				body, _ = sjson.Set(body, "host-ipv6-addresses.host-ipv6-address"+"."+strconv.Itoa(index)+"."+"operator", item.Operator.ValueString())
			}
		}
	}
	return body
}

func (data *LoggingVRF) updateFromBody(ctx context.Context, res []byte) {
	for i := range data.HostIpv4Addresses {
		keys := [...]string{"ipv4-address"}
		keyValues := [...]string{data.HostIpv4Addresses[i].Ipv4Address.ValueString()}

		var r gjson.Result
		gjson.GetBytes(res, "host-ipv4-addresses.host-ipv4-address").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() == keyValues[ik] {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if value := r.Get("ipv4-address"); value.Exists() && !data.HostIpv4Addresses[i].Ipv4Address.IsNull() {
			data.HostIpv4Addresses[i].Ipv4Address = types.StringValue(value.String())
		} else {
			data.HostIpv4Addresses[i].Ipv4Address = types.StringNull()
		}
		if value := r.Get("severity"); value.Exists() && !data.HostIpv4Addresses[i].Severity.IsNull() {
			data.HostIpv4Addresses[i].Severity = types.StringValue(value.String())
		} else {
			data.HostIpv4Addresses[i].Severity = types.StringNull()
		}
		if value := r.Get("port"); value.Exists() && !data.HostIpv4Addresses[i].Port.IsNull() {
			data.HostIpv4Addresses[i].Port = types.Int64Value(value.Int())
		} else {
			data.HostIpv4Addresses[i].Port = types.Int64Null()
		}
		if value := r.Get("operator"); value.Exists() && !data.HostIpv4Addresses[i].Operator.IsNull() {
			data.HostIpv4Addresses[i].Operator = types.StringValue(value.String())
		} else {
			data.HostIpv4Addresses[i].Operator = types.StringNull()
		}
	}
	for i := range data.HostIpv6Addresses {
		keys := [...]string{"ipv6-address"}
		keyValues := [...]string{data.HostIpv6Addresses[i].Ipv6Address.ValueString()}

		var r gjson.Result
		gjson.GetBytes(res, "host-ipv6-addresses.host-ipv6-address").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() == keyValues[ik] {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if value := r.Get("ipv6-address"); value.Exists() && !data.HostIpv6Addresses[i].Ipv6Address.IsNull() {
			data.HostIpv6Addresses[i].Ipv6Address = types.StringValue(value.String())
		} else {
			data.HostIpv6Addresses[i].Ipv6Address = types.StringNull()
		}
		if value := r.Get("severity"); value.Exists() && !data.HostIpv6Addresses[i].Severity.IsNull() {
			data.HostIpv6Addresses[i].Severity = types.StringValue(value.String())
		} else {
			data.HostIpv6Addresses[i].Severity = types.StringNull()
		}
		if value := r.Get("port"); value.Exists() && !data.HostIpv6Addresses[i].Port.IsNull() {
			data.HostIpv6Addresses[i].Port = types.Int64Value(value.Int())
		} else {
			data.HostIpv6Addresses[i].Port = types.Int64Null()
		}
		if value := r.Get("operator"); value.Exists() && !data.HostIpv6Addresses[i].Operator.IsNull() {
			data.HostIpv6Addresses[i].Operator = types.StringValue(value.String())
		} else {
			data.HostIpv6Addresses[i].Operator = types.StringNull()
		}
	}
}

func (data *LoggingVRFData) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "host-ipv4-addresses.host-ipv4-address"); value.Exists() {
		data.HostIpv4Addresses = make([]LoggingVRFHostIpv4Addresses, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := LoggingVRFHostIpv4Addresses{}
			if cValue := v.Get("ipv4-address"); cValue.Exists() {
				item.Ipv4Address = types.StringValue(cValue.String())
			}
			if cValue := v.Get("severity"); cValue.Exists() {
				item.Severity = types.StringValue(cValue.String())
			}
			if cValue := v.Get("port"); cValue.Exists() {
				item.Port = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("operator"); cValue.Exists() {
				item.Operator = types.StringValue(cValue.String())
			}
			data.HostIpv4Addresses = append(data.HostIpv4Addresses, item)
			return true
		})
	}
	if value := gjson.GetBytes(res, "host-ipv6-addresses.host-ipv6-address"); value.Exists() {
		data.HostIpv6Addresses = make([]LoggingVRFHostIpv6Addresses, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := LoggingVRFHostIpv6Addresses{}
			if cValue := v.Get("ipv6-address"); cValue.Exists() {
				item.Ipv6Address = types.StringValue(cValue.String())
			}
			if cValue := v.Get("severity"); cValue.Exists() {
				item.Severity = types.StringValue(cValue.String())
			}
			if cValue := v.Get("port"); cValue.Exists() {
				item.Port = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("operator"); cValue.Exists() {
				item.Operator = types.StringValue(cValue.String())
			}
			data.HostIpv6Addresses = append(data.HostIpv6Addresses, item)
			return true
		})
	}
}

func (data *LoggingVRF) getDeletedItems(ctx context.Context, state LoggingVRF) []string {
	deletedItems := make([]string, 0)
	for i := range state.HostIpv4Addresses {
		keys := [...]string{"ipv4-address"}
		stateKeyValues := [...]string{state.HostIpv4Addresses[i].Ipv4Address.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.HostIpv4Addresses[i].Ipv4Address.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.HostIpv4Addresses {
			found = true
			if state.HostIpv4Addresses[i].Ipv4Address.ValueString() != data.HostIpv4Addresses[j].Ipv4Address.ValueString() {
				found = false
			}
			if found {
				break
			}
		}
		if !found {
			keyString := ""
			for ki := range keys {
				keyString += "[" + keys[ki] + "=" + stateKeyValues[ki] + "]"
			}
			deletedItems = append(deletedItems, fmt.Sprintf("%v/host-ipv4-addresses/host-ipv4-address%v", state.getPath(), keyString))
		}
	}
	for i := range state.HostIpv6Addresses {
		keys := [...]string{"ipv6-address"}
		stateKeyValues := [...]string{state.HostIpv6Addresses[i].Ipv6Address.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.HostIpv6Addresses[i].Ipv6Address.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.HostIpv6Addresses {
			found = true
			if state.HostIpv6Addresses[i].Ipv6Address.ValueString() != data.HostIpv6Addresses[j].Ipv6Address.ValueString() {
				found = false
			}
			if found {
				break
			}
		}
		if !found {
			keyString := ""
			for ki := range keys {
				keyString += "[" + keys[ki] + "=" + stateKeyValues[ki] + "]"
			}
			deletedItems = append(deletedItems, fmt.Sprintf("%v/host-ipv6-addresses/host-ipv6-address%v", state.getPath(), keyString))
		}
	}
	return deletedItems
}

func (data *LoggingVRF) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	for i := range data.HostIpv4Addresses {
		keys := [...]string{"ipv4-address"}
		keyValues := [...]string{data.HostIpv4Addresses[i].Ipv4Address.ValueString()}
		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
	}
	for i := range data.HostIpv6Addresses {
		keys := [...]string{"ipv6-address"}
		keyValues := [...]string{data.HostIpv6Addresses[i].Ipv6Address.ValueString()}
		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
	}
	return emptyLeafsDelete
}

func (data *LoggingVRF) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	for i := range data.HostIpv4Addresses {
		keys := [...]string{"ipv4-address"}
		keyValues := [...]string{data.HostIpv4Addresses[i].Ipv4Address.ValueString()}

		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
		deletePaths = append(deletePaths, fmt.Sprintf("%v/host-ipv4-addresses/host-ipv4-address%v", data.getPath(), keyString))
	}
	for i := range data.HostIpv6Addresses {
		keys := [...]string{"ipv6-address"}
		keyValues := [...]string{data.HostIpv6Addresses[i].Ipv6Address.ValueString()}

		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
		deletePaths = append(deletePaths, fmt.Sprintf("%v/host-ipv6-addresses/host-ipv6-address%v", data.getPath(), keyString))
	}
	return deletePaths
}
