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

type RouterHSRPInterfaceAddressFamilyIPv6V1 struct {
	Device                                     types.String                                            `tfsdk:"device"`
	Id                                         types.String                                            `tfsdk:"id"`
	DeleteMode                                 types.String                                            `tfsdk:"delete_mode"`
	InterfaceName                              types.String                                            `tfsdk:"interface_name"`
	GroupNumberVersion1Id                      types.Int64                                             `tfsdk:"group_number_version_1_id"`
	Name                                       types.String                                            `tfsdk:"name"`
	MacAddress                                 types.String                                            `tfsdk:"mac_address"`
	TimersHoldTime                             types.Int64                                             `tfsdk:"timers_hold_time"`
	TimersHoldTime2                            types.Int64                                             `tfsdk:"timers_hold_time2"`
	TimersMsec                                 types.Int64                                             `tfsdk:"timers_msec"`
	TimersMsec2                                types.Int64                                             `tfsdk:"timers_msec2"`
	PreemptDelay                               types.Int64                                             `tfsdk:"preempt_delay"`
	Priority                                   types.Int64                                             `tfsdk:"priority"`
	BfdFastDetectPeerIpv6                      types.String                                            `tfsdk:"bfd_fast_detect_peer_ipv6"`
	BfdFastDetectPeerInterface                 types.String                                            `tfsdk:"bfd_fast_detect_peer_interface"`
	TrackObjects                               []RouterHSRPInterfaceAddressFamilyIPv6V1TrackObjects    `tfsdk:"track_objects"`
	TrackInterfaces                            []RouterHSRPInterfaceAddressFamilyIPv6V1TrackInterfaces `tfsdk:"track_interfaces"`
	AddressGlobals                             []RouterHSRPInterfaceAddressFamilyIPv6V1AddressGlobals  `tfsdk:"address_globals"`
	AddressLinkLocalAutoconfigLegacyCompatible types.Bool                                              `tfsdk:"address_link_local_autoconfig_legacy_compatible"`
	AddressLinkLocalIpv6Address                types.String                                            `tfsdk:"address_link_local_ipv6_address"`
}

type RouterHSRPInterfaceAddressFamilyIPv6V1Data struct {
	Device                                     types.String                                            `tfsdk:"device"`
	Id                                         types.String                                            `tfsdk:"id"`
	InterfaceName                              types.String                                            `tfsdk:"interface_name"`
	GroupNumberVersion1Id                      types.Int64                                             `tfsdk:"group_number_version_1_id"`
	Name                                       types.String                                            `tfsdk:"name"`
	MacAddress                                 types.String                                            `tfsdk:"mac_address"`
	TimersHoldTime                             types.Int64                                             `tfsdk:"timers_hold_time"`
	TimersHoldTime2                            types.Int64                                             `tfsdk:"timers_hold_time2"`
	TimersMsec                                 types.Int64                                             `tfsdk:"timers_msec"`
	TimersMsec2                                types.Int64                                             `tfsdk:"timers_msec2"`
	PreemptDelay                               types.Int64                                             `tfsdk:"preempt_delay"`
	Priority                                   types.Int64                                             `tfsdk:"priority"`
	BfdFastDetectPeerIpv6                      types.String                                            `tfsdk:"bfd_fast_detect_peer_ipv6"`
	BfdFastDetectPeerInterface                 types.String                                            `tfsdk:"bfd_fast_detect_peer_interface"`
	TrackObjects                               []RouterHSRPInterfaceAddressFamilyIPv6V1TrackObjects    `tfsdk:"track_objects"`
	TrackInterfaces                            []RouterHSRPInterfaceAddressFamilyIPv6V1TrackInterfaces `tfsdk:"track_interfaces"`
	AddressGlobals                             []RouterHSRPInterfaceAddressFamilyIPv6V1AddressGlobals  `tfsdk:"address_globals"`
	AddressLinkLocalAutoconfigLegacyCompatible types.Bool                                              `tfsdk:"address_link_local_autoconfig_legacy_compatible"`
	AddressLinkLocalIpv6Address                types.String                                            `tfsdk:"address_link_local_ipv6_address"`
}
type RouterHSRPInterfaceAddressFamilyIPv6V1TrackObjects struct {
	ObjectName        types.String `tfsdk:"object_name"`
	PriorityDecrement types.Int64  `tfsdk:"priority_decrement"`
}
type RouterHSRPInterfaceAddressFamilyIPv6V1TrackInterfaces struct {
	TrackName         types.String `tfsdk:"track_name"`
	PriorityDecrement types.Int64  `tfsdk:"priority_decrement"`
}
type RouterHSRPInterfaceAddressFamilyIPv6V1AddressGlobals struct {
	Address types.String `tfsdk:"address"`
}

func (data RouterHSRPInterfaceAddressFamilyIPv6V1) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-router-hsrp-cfg:/router/hsrp/interfaces/interface[interface-name=%s]/address-family/ipv6/hsrp/group-number-version-1s/group-number-version-1[group-number-version-1-id=%v]", data.InterfaceName.ValueString(), data.GroupNumberVersion1Id.ValueInt64())
}

func (data RouterHSRPInterfaceAddressFamilyIPv6V1Data) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-router-hsrp-cfg:/router/hsrp/interfaces/interface[interface-name=%s]/address-family/ipv6/hsrp/group-number-version-1s/group-number-version-1[group-number-version-1-id=%v]", data.InterfaceName.ValueString(), data.GroupNumberVersion1Id.ValueInt64())
}

func (data RouterHSRPInterfaceAddressFamilyIPv6V1) toBody(ctx context.Context) string {
	body := "{}"
	if !data.GroupNumberVersion1Id.IsNull() && !data.GroupNumberVersion1Id.IsUnknown() {
		body, _ = sjson.Set(body, "group-number-version-1-id", strconv.FormatInt(data.GroupNumberVersion1Id.ValueInt64(), 10))
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.MacAddress.IsNull() && !data.MacAddress.IsUnknown() {
		body, _ = sjson.Set(body, "mac-address", data.MacAddress.ValueString())
	}
	if !data.TimersHoldTime.IsNull() && !data.TimersHoldTime.IsUnknown() {
		body, _ = sjson.Set(body, "timers.hold-time", strconv.FormatInt(data.TimersHoldTime.ValueInt64(), 10))
	}
	if !data.TimersHoldTime2.IsNull() && !data.TimersHoldTime2.IsUnknown() {
		body, _ = sjson.Set(body, "timers.hold-time2", strconv.FormatInt(data.TimersHoldTime2.ValueInt64(), 10))
	}
	if !data.TimersMsec.IsNull() && !data.TimersMsec.IsUnknown() {
		body, _ = sjson.Set(body, "timers.msec", strconv.FormatInt(data.TimersMsec.ValueInt64(), 10))
	}
	if !data.TimersMsec2.IsNull() && !data.TimersMsec2.IsUnknown() {
		body, _ = sjson.Set(body, "timers.msec2", strconv.FormatInt(data.TimersMsec2.ValueInt64(), 10))
	}
	if !data.PreemptDelay.IsNull() && !data.PreemptDelay.IsUnknown() {
		body, _ = sjson.Set(body, "preempt.delay", strconv.FormatInt(data.PreemptDelay.ValueInt64(), 10))
	}
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		body, _ = sjson.Set(body, "priority", strconv.FormatInt(data.Priority.ValueInt64(), 10))
	}
	if !data.BfdFastDetectPeerIpv6.IsNull() && !data.BfdFastDetectPeerIpv6.IsUnknown() {
		body, _ = sjson.Set(body, "bfd.fast-detect.peer.ipv6", data.BfdFastDetectPeerIpv6.ValueString())
	}
	if !data.BfdFastDetectPeerInterface.IsNull() && !data.BfdFastDetectPeerInterface.IsUnknown() {
		body, _ = sjson.Set(body, "bfd.fast-detect.peer.interface", data.BfdFastDetectPeerInterface.ValueString())
	}
	if !data.AddressLinkLocalAutoconfigLegacyCompatible.IsNull() && !data.AddressLinkLocalAutoconfigLegacyCompatible.IsUnknown() {
		if data.AddressLinkLocalAutoconfigLegacyCompatible.ValueBool() {
			body, _ = sjson.Set(body, "address.link-local.autoconfig.legacy-compatible", map[string]string{})
		}
	}
	if !data.AddressLinkLocalIpv6Address.IsNull() && !data.AddressLinkLocalIpv6Address.IsUnknown() {
		body, _ = sjson.Set(body, "address.link-local.ipv6-address", data.AddressLinkLocalIpv6Address.ValueString())
	}
	if len(data.TrackObjects) > 0 {
		body, _ = sjson.Set(body, "track-objects.track-object", []interface{}{})
		for index, item := range data.TrackObjects {
			if !item.ObjectName.IsNull() && !item.ObjectName.IsUnknown() {
				body, _ = sjson.Set(body, "track-objects.track-object"+"."+strconv.Itoa(index)+"."+"object-name", item.ObjectName.ValueString())
			}
			if !item.PriorityDecrement.IsNull() && !item.PriorityDecrement.IsUnknown() {
				body, _ = sjson.Set(body, "track-objects.track-object"+"."+strconv.Itoa(index)+"."+"priority-decrement", strconv.FormatInt(item.PriorityDecrement.ValueInt64(), 10))
			}
		}
	}
	if len(data.TrackInterfaces) > 0 {
		body, _ = sjson.Set(body, "track-interfaces.track-interface", []interface{}{})
		for index, item := range data.TrackInterfaces {
			if !item.TrackName.IsNull() && !item.TrackName.IsUnknown() {
				body, _ = sjson.Set(body, "track-interfaces.track-interface"+"."+strconv.Itoa(index)+"."+"track-name", item.TrackName.ValueString())
			}
			if !item.PriorityDecrement.IsNull() && !item.PriorityDecrement.IsUnknown() {
				body, _ = sjson.Set(body, "track-interfaces.track-interface"+"."+strconv.Itoa(index)+"."+"priority-decrement", strconv.FormatInt(item.PriorityDecrement.ValueInt64(), 10))
			}
		}
	}
	if len(data.AddressGlobals) > 0 {
		body, _ = sjson.Set(body, "address.globals.global", []interface{}{})
		for index, item := range data.AddressGlobals {
			if !item.Address.IsNull() && !item.Address.IsUnknown() {
				body, _ = sjson.Set(body, "address.globals.global"+"."+strconv.Itoa(index)+"."+"address", item.Address.ValueString())
			}
		}
	}
	return body
}

func (data *RouterHSRPInterfaceAddressFamilyIPv6V1) updateFromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := gjson.GetBytes(res, "mac-address"); value.Exists() && !data.MacAddress.IsNull() {
		data.MacAddress = types.StringValue(value.String())
	} else {
		data.MacAddress = types.StringNull()
	}
	if value := gjson.GetBytes(res, "timers.hold-time"); value.Exists() && !data.TimersHoldTime.IsNull() {
		data.TimersHoldTime = types.Int64Value(value.Int())
	} else {
		data.TimersHoldTime = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "timers.hold-time2"); value.Exists() && !data.TimersHoldTime2.IsNull() {
		data.TimersHoldTime2 = types.Int64Value(value.Int())
	} else {
		data.TimersHoldTime2 = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "timers.msec"); value.Exists() && !data.TimersMsec.IsNull() {
		data.TimersMsec = types.Int64Value(value.Int())
	} else {
		data.TimersMsec = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "timers.msec2"); value.Exists() && !data.TimersMsec2.IsNull() {
		data.TimersMsec2 = types.Int64Value(value.Int())
	} else {
		data.TimersMsec2 = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "preempt.delay"); value.Exists() && !data.PreemptDelay.IsNull() {
		data.PreemptDelay = types.Int64Value(value.Int())
	} else {
		data.PreemptDelay = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "priority"); value.Exists() && !data.Priority.IsNull() {
		data.Priority = types.Int64Value(value.Int())
	} else {
		data.Priority = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "bfd.fast-detect.peer.ipv6"); value.Exists() && !data.BfdFastDetectPeerIpv6.IsNull() {
		data.BfdFastDetectPeerIpv6 = types.StringValue(value.String())
	} else {
		data.BfdFastDetectPeerIpv6 = types.StringNull()
	}
	if value := gjson.GetBytes(res, "bfd.fast-detect.peer.interface"); value.Exists() && !data.BfdFastDetectPeerInterface.IsNull() {
		data.BfdFastDetectPeerInterface = types.StringValue(value.String())
	} else {
		data.BfdFastDetectPeerInterface = types.StringNull()
	}
	for i := range data.TrackObjects {
		keys := [...]string{"object-name"}
		keyValues := [...]string{data.TrackObjects[i].ObjectName.ValueString()}

		var r gjson.Result
		gjson.GetBytes(res, "track-objects.track-object").ForEach(
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
		if value := r.Get("object-name"); value.Exists() && !data.TrackObjects[i].ObjectName.IsNull() {
			data.TrackObjects[i].ObjectName = types.StringValue(value.String())
		} else {
			data.TrackObjects[i].ObjectName = types.StringNull()
		}
		if value := r.Get("priority-decrement"); value.Exists() && !data.TrackObjects[i].PriorityDecrement.IsNull() {
			data.TrackObjects[i].PriorityDecrement = types.Int64Value(value.Int())
		} else {
			data.TrackObjects[i].PriorityDecrement = types.Int64Null()
		}
	}
	for i := range data.TrackInterfaces {
		keys := [...]string{"track-name"}
		keyValues := [...]string{data.TrackInterfaces[i].TrackName.ValueString()}

		var r gjson.Result
		gjson.GetBytes(res, "track-interfaces.track-interface").ForEach(
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
		if value := r.Get("track-name"); value.Exists() && !data.TrackInterfaces[i].TrackName.IsNull() {
			data.TrackInterfaces[i].TrackName = types.StringValue(value.String())
		} else {
			data.TrackInterfaces[i].TrackName = types.StringNull()
		}
		if value := r.Get("priority-decrement"); value.Exists() && !data.TrackInterfaces[i].PriorityDecrement.IsNull() {
			data.TrackInterfaces[i].PriorityDecrement = types.Int64Value(value.Int())
		} else {
			data.TrackInterfaces[i].PriorityDecrement = types.Int64Null()
		}
	}
	for i := range data.AddressGlobals {
		keys := [...]string{"address"}
		keyValues := [...]string{data.AddressGlobals[i].Address.ValueString()}

		var r gjson.Result
		gjson.GetBytes(res, "address.globals.global").ForEach(
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
		if value := r.Get("address"); value.Exists() && !data.AddressGlobals[i].Address.IsNull() {
			data.AddressGlobals[i].Address = types.StringValue(value.String())
		} else {
			data.AddressGlobals[i].Address = types.StringNull()
		}
	}
	if value := gjson.GetBytes(res, "address.link-local.autoconfig.legacy-compatible"); !data.AddressLinkLocalAutoconfigLegacyCompatible.IsNull() {
		if value.Exists() {
			data.AddressLinkLocalAutoconfigLegacyCompatible = types.BoolValue(true)
		} else {
			data.AddressLinkLocalAutoconfigLegacyCompatible = types.BoolValue(false)
		}
	} else {
		data.AddressLinkLocalAutoconfigLegacyCompatible = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "address.link-local.ipv6-address"); value.Exists() && !data.AddressLinkLocalIpv6Address.IsNull() {
		data.AddressLinkLocalIpv6Address = types.StringValue(value.String())
	} else {
		data.AddressLinkLocalIpv6Address = types.StringNull()
	}
}

func (data *RouterHSRPInterfaceAddressFamilyIPv6V1Data) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "mac-address"); value.Exists() {
		data.MacAddress = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "timers.hold-time"); value.Exists() {
		data.TimersHoldTime = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "timers.hold-time2"); value.Exists() {
		data.TimersHoldTime2 = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "timers.msec"); value.Exists() {
		data.TimersMsec = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "timers.msec2"); value.Exists() {
		data.TimersMsec2 = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "preempt.delay"); value.Exists() {
		data.PreemptDelay = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "priority"); value.Exists() {
		data.Priority = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "bfd.fast-detect.peer.ipv6"); value.Exists() {
		data.BfdFastDetectPeerIpv6 = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "bfd.fast-detect.peer.interface"); value.Exists() {
		data.BfdFastDetectPeerInterface = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "track-objects.track-object"); value.Exists() {
		data.TrackObjects = make([]RouterHSRPInterfaceAddressFamilyIPv6V1TrackObjects, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := RouterHSRPInterfaceAddressFamilyIPv6V1TrackObjects{}
			if cValue := v.Get("object-name"); cValue.Exists() {
				item.ObjectName = types.StringValue(cValue.String())
			}
			if cValue := v.Get("priority-decrement"); cValue.Exists() {
				item.PriorityDecrement = types.Int64Value(cValue.Int())
			}
			data.TrackObjects = append(data.TrackObjects, item)
			return true
		})
	}
	if value := gjson.GetBytes(res, "track-interfaces.track-interface"); value.Exists() {
		data.TrackInterfaces = make([]RouterHSRPInterfaceAddressFamilyIPv6V1TrackInterfaces, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := RouterHSRPInterfaceAddressFamilyIPv6V1TrackInterfaces{}
			if cValue := v.Get("track-name"); cValue.Exists() {
				item.TrackName = types.StringValue(cValue.String())
			}
			if cValue := v.Get("priority-decrement"); cValue.Exists() {
				item.PriorityDecrement = types.Int64Value(cValue.Int())
			}
			data.TrackInterfaces = append(data.TrackInterfaces, item)
			return true
		})
	}
	if value := gjson.GetBytes(res, "address.globals.global"); value.Exists() {
		data.AddressGlobals = make([]RouterHSRPInterfaceAddressFamilyIPv6V1AddressGlobals, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := RouterHSRPInterfaceAddressFamilyIPv6V1AddressGlobals{}
			if cValue := v.Get("address"); cValue.Exists() {
				item.Address = types.StringValue(cValue.String())
			}
			data.AddressGlobals = append(data.AddressGlobals, item)
			return true
		})
	}
	if value := gjson.GetBytes(res, "address.link-local.autoconfig.legacy-compatible"); value.Exists() {
		data.AddressLinkLocalAutoconfigLegacyCompatible = types.BoolValue(true)
	} else {
		data.AddressLinkLocalAutoconfigLegacyCompatible = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "address.link-local.ipv6-address"); value.Exists() {
		data.AddressLinkLocalIpv6Address = types.StringValue(value.String())
	}
}

func (data *RouterHSRPInterfaceAddressFamilyIPv6V1) getDeletedListItems(ctx context.Context, state RouterHSRPInterfaceAddressFamilyIPv6V1) []string {
	deletedListItems := make([]string, 0)
	for i := range state.TrackObjects {
		keys := [...]string{"object-name"}
		stateKeyValues := [...]string{state.TrackObjects[i].ObjectName.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.TrackObjects[i].ObjectName.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.TrackObjects {
			found = true
			if state.TrackObjects[i].ObjectName.ValueString() != data.TrackObjects[j].ObjectName.ValueString() {
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
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/track-objects/track-object%v", state.getPath(), keyString))
		}
	}
	for i := range state.TrackInterfaces {
		keys := [...]string{"track-name"}
		stateKeyValues := [...]string{state.TrackInterfaces[i].TrackName.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.TrackInterfaces[i].TrackName.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.TrackInterfaces {
			found = true
			if state.TrackInterfaces[i].TrackName.ValueString() != data.TrackInterfaces[j].TrackName.ValueString() {
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
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/track-interfaces/track-interface%v", state.getPath(), keyString))
		}
	}
	for i := range state.AddressGlobals {
		keys := [...]string{"address"}
		stateKeyValues := [...]string{state.AddressGlobals[i].Address.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.AddressGlobals[i].Address.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.AddressGlobals {
			found = true
			if state.AddressGlobals[i].Address.ValueString() != data.AddressGlobals[j].Address.ValueString() {
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
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/address/globals/global%v", state.getPath(), keyString))
		}
	}
	return deletedListItems
}

func (data *RouterHSRPInterfaceAddressFamilyIPv6V1) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	for i := range data.TrackObjects {
		keys := [...]string{"object-name"}
		keyValues := [...]string{data.TrackObjects[i].ObjectName.ValueString()}
		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
	}
	for i := range data.TrackInterfaces {
		keys := [...]string{"track-name"}
		keyValues := [...]string{data.TrackInterfaces[i].TrackName.ValueString()}
		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
	}
	for i := range data.AddressGlobals {
		keys := [...]string{"address"}
		keyValues := [...]string{data.AddressGlobals[i].Address.ValueString()}
		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
	}
	if !data.AddressLinkLocalAutoconfigLegacyCompatible.IsNull() && !data.AddressLinkLocalAutoconfigLegacyCompatible.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/address/link-local/autoconfig/legacy-compatible", data.getPath()))
	}
	return emptyLeafsDelete
}

func (data *RouterHSRPInterfaceAddressFamilyIPv6V1) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.Name.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/name", data.getPath()))
	}
	if !data.MacAddress.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/mac-address", data.getPath()))
	}
	if !data.TimersHoldTime.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/timers/hold-time", data.getPath()))
	}
	if !data.TimersHoldTime2.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/timers/hold-time2", data.getPath()))
	}
	if !data.TimersMsec.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/timers/msec", data.getPath()))
	}
	if !data.TimersMsec2.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/timers/msec2", data.getPath()))
	}
	if !data.PreemptDelay.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/preempt/delay", data.getPath()))
	}
	if !data.Priority.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/priority", data.getPath()))
	}
	if !data.BfdFastDetectPeerIpv6.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/bfd/fast-detect/peer/ipv6", data.getPath()))
	}
	if !data.BfdFastDetectPeerInterface.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/bfd/fast-detect/peer/interface", data.getPath()))
	}
	for i := range data.TrackObjects {
		keys := [...]string{"object-name"}
		keyValues := [...]string{data.TrackObjects[i].ObjectName.ValueString()}

		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
		deletePaths = append(deletePaths, fmt.Sprintf("%v/track-objects/track-object%v", data.getPath(), keyString))
	}
	for i := range data.TrackInterfaces {
		keys := [...]string{"track-name"}
		keyValues := [...]string{data.TrackInterfaces[i].TrackName.ValueString()}

		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
		deletePaths = append(deletePaths, fmt.Sprintf("%v/track-interfaces/track-interface%v", data.getPath(), keyString))
	}
	for i := range data.AddressGlobals {
		keys := [...]string{"address"}
		keyValues := [...]string{data.AddressGlobals[i].Address.ValueString()}

		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
		deletePaths = append(deletePaths, fmt.Sprintf("%v/address/globals/global%v", data.getPath(), keyString))
	}
	if !data.AddressLinkLocalAutoconfigLegacyCompatible.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/address/link-local/autoconfig/legacy-compatible", data.getPath()))
	}
	if !data.AddressLinkLocalIpv6Address.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/address/link-local/ipv6-address", data.getPath()))
	}
	return deletePaths
}
