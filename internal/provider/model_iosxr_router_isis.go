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

type RouterISIS struct {
	Device                        types.String                     `tfsdk:"device"`
	Id                            types.String                     `tfsdk:"id"`
	ProcessId                     types.String                     `tfsdk:"process_id"`
	IsType                        types.String                     `tfsdk:"is_type"`
	SetOverloadBitLevels          []RouterISISSetOverloadBitLevels `tfsdk:"set_overload_bit_levels"`
	Nsr                           types.Bool                       `tfsdk:"nsr"`
	NsfCisco                      types.Bool                       `tfsdk:"nsf_cisco"`
	NsfIetf                       types.Bool                       `tfsdk:"nsf_ietf"`
	NsfLifetime                   types.Int64                      `tfsdk:"nsf_lifetime"`
	NsfInterfaceTimer             types.Int64                      `tfsdk:"nsf_interface_timer"`
	NsfInterfaceExpires           types.Int64                      `tfsdk:"nsf_interface_expires"`
	LogAdjacencyChanges           types.Bool                       `tfsdk:"log_adjacency_changes"`
	LspGenIntervalMaximumWait     types.Int64                      `tfsdk:"lsp_gen_interval_maximum_wait"`
	LspGenIntervalInitialWait     types.Int64                      `tfsdk:"lsp_gen_interval_initial_wait"`
	LspGenIntervalSecondaryWait   types.Int64                      `tfsdk:"lsp_gen_interval_secondary_wait"`
	LspRefreshInterval            types.Int64                      `tfsdk:"lsp_refresh_interval"`
	MaxLspLifetime                types.Int64                      `tfsdk:"max_lsp_lifetime"`
	LspPasswordKeychain           types.String                     `tfsdk:"lsp_password_keychain"`
	DistributeLinkStateInstanceId types.Int64                      `tfsdk:"distribute_link_state_instance_id"`
	DistributeLinkStateThrottle   types.Int64                      `tfsdk:"distribute_link_state_throttle"`
	DistributeLinkStateLevel      types.Int64                      `tfsdk:"distribute_link_state_level"`
	AffinityMaps                  []RouterISISAffinityMaps         `tfsdk:"affinity_maps"`
	FlexAlgos                     []RouterISISFlexAlgos            `tfsdk:"flex_algos"`
	Nets                          []RouterISISNets                 `tfsdk:"nets"`
	Interfaces                    []RouterISISInterfaces           `tfsdk:"interfaces"`
}
type RouterISISSetOverloadBitLevels struct {
	LevelId                                       types.Int64 `tfsdk:"level_id"`
	OnStartupAdvertiseAsOverloaded                types.Bool  `tfsdk:"on_startup_advertise_as_overloaded"`
	OnStartupAdvertiseAsOverloadedTimeToAdvertise types.Int64 `tfsdk:"on_startup_advertise_as_overloaded_time_to_advertise"`
	OnStartupWaitForBgp                           types.Bool  `tfsdk:"on_startup_wait_for_bgp"`
	AdvertiseExternal                             types.Bool  `tfsdk:"advertise_external"`
	AdvertiseInterlevel                           types.Bool  `tfsdk:"advertise_interlevel"`
}
type RouterISISAffinityMaps struct {
	Name        types.String `tfsdk:"name"`
	BitPosition types.Int64  `tfsdk:"bit_position"`
}
type RouterISISFlexAlgos struct {
	AlgorithmNumber     types.Int64 `tfsdk:"algorithm_number"`
	AdvertiseDefinition types.Bool  `tfsdk:"advertise_definition"`
	MetricTypeDelay     types.Bool  `tfsdk:"metric_type_delay"`
}
type RouterISISNets struct {
	NetId types.String `tfsdk:"net_id"`
}
type RouterISISInterfaces struct {
	InterfaceName         types.String `tfsdk:"interface_name"`
	CircuitType           types.String `tfsdk:"circuit_type"`
	HelloPaddingDisable   types.Bool   `tfsdk:"hello_padding_disable"`
	HelloPaddingSometimes types.Bool   `tfsdk:"hello_padding_sometimes"`
	Priority              types.Int64  `tfsdk:"priority"`
	PointToPoint          types.Bool   `tfsdk:"point_to_point"`
	Passive               types.Bool   `tfsdk:"passive"`
	Suppressed            types.Bool   `tfsdk:"suppressed"`
	Shutdown              types.Bool   `tfsdk:"shutdown"`
}

func (data RouterISIS) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-router-isis-cfg:/router/isis/processes/process[process-id=%s]", data.ProcessId.ValueString())
}

func (data RouterISIS) toBody(ctx context.Context) string {
	body := "{}"
	if !data.ProcessId.IsNull() && !data.ProcessId.IsUnknown() {
		body, _ = sjson.Set(body, "process-id", data.ProcessId.ValueString())
	}
	if !data.IsType.IsNull() && !data.IsType.IsUnknown() {
		body, _ = sjson.Set(body, "is-type", data.IsType.ValueString())
	}
	if !data.Nsr.IsNull() && !data.Nsr.IsUnknown() {
		if data.Nsr.ValueBool() {
			body, _ = sjson.Set(body, "nsr", map[string]string{})
		}
	}
	if !data.NsfCisco.IsNull() && !data.NsfCisco.IsUnknown() {
		if data.NsfCisco.ValueBool() {
			body, _ = sjson.Set(body, "nsf.cisco", map[string]string{})
		}
	}
	if !data.NsfIetf.IsNull() && !data.NsfIetf.IsUnknown() {
		if data.NsfIetf.ValueBool() {
			body, _ = sjson.Set(body, "nsf.ietf", map[string]string{})
		}
	}
	if !data.NsfLifetime.IsNull() && !data.NsfLifetime.IsUnknown() {
		body, _ = sjson.Set(body, "nsf.lifetime", strconv.FormatInt(data.NsfLifetime.ValueInt64(), 10))
	}
	if !data.NsfInterfaceTimer.IsNull() && !data.NsfInterfaceTimer.IsUnknown() {
		body, _ = sjson.Set(body, "nsf.interface-timer", strconv.FormatInt(data.NsfInterfaceTimer.ValueInt64(), 10))
	}
	if !data.NsfInterfaceExpires.IsNull() && !data.NsfInterfaceExpires.IsUnknown() {
		body, _ = sjson.Set(body, "nsf.interface-expires", strconv.FormatInt(data.NsfInterfaceExpires.ValueInt64(), 10))
	}
	if !data.LogAdjacencyChanges.IsNull() && !data.LogAdjacencyChanges.IsUnknown() {
		if data.LogAdjacencyChanges.ValueBool() {
			body, _ = sjson.Set(body, "log.adjacency.changes", map[string]string{})
		}
	}
	if !data.LspGenIntervalMaximumWait.IsNull() && !data.LspGenIntervalMaximumWait.IsUnknown() {
		body, _ = sjson.Set(body, "lsp-gen-interval.maximum-wait", strconv.FormatInt(data.LspGenIntervalMaximumWait.ValueInt64(), 10))
	}
	if !data.LspGenIntervalInitialWait.IsNull() && !data.LspGenIntervalInitialWait.IsUnknown() {
		body, _ = sjson.Set(body, "lsp-gen-interval.initial-wait", strconv.FormatInt(data.LspGenIntervalInitialWait.ValueInt64(), 10))
	}
	if !data.LspGenIntervalSecondaryWait.IsNull() && !data.LspGenIntervalSecondaryWait.IsUnknown() {
		body, _ = sjson.Set(body, "lsp-gen-interval.secondary-wait", strconv.FormatInt(data.LspGenIntervalSecondaryWait.ValueInt64(), 10))
	}
	if !data.LspRefreshInterval.IsNull() && !data.LspRefreshInterval.IsUnknown() {
		body, _ = sjson.Set(body, "lsp-refresh-interval.lsp-refresh-interval-time", strconv.FormatInt(data.LspRefreshInterval.ValueInt64(), 10))
	}
	if !data.MaxLspLifetime.IsNull() && !data.MaxLspLifetime.IsUnknown() {
		body, _ = sjson.Set(body, "max-lsp-lifetime.max-lsp-lifetime", strconv.FormatInt(data.MaxLspLifetime.ValueInt64(), 10))
	}
	if !data.LspPasswordKeychain.IsNull() && !data.LspPasswordKeychain.IsUnknown() {
		body, _ = sjson.Set(body, "lsp-password.keychain.keychain-name", data.LspPasswordKeychain.ValueString())
	}
	if !data.DistributeLinkStateInstanceId.IsNull() && !data.DistributeLinkStateInstanceId.IsUnknown() {
		body, _ = sjson.Set(body, "distribute.link-state.instance-id", strconv.FormatInt(data.DistributeLinkStateInstanceId.ValueInt64(), 10))
	}
	if !data.DistributeLinkStateThrottle.IsNull() && !data.DistributeLinkStateThrottle.IsUnknown() {
		body, _ = sjson.Set(body, "distribute.link-state.throttle", strconv.FormatInt(data.DistributeLinkStateThrottle.ValueInt64(), 10))
	}
	if !data.DistributeLinkStateLevel.IsNull() && !data.DistributeLinkStateLevel.IsUnknown() {
		body, _ = sjson.Set(body, "distribute.link-state.level", strconv.FormatInt(data.DistributeLinkStateLevel.ValueInt64(), 10))
	}
	if len(data.SetOverloadBitLevels) > 0 {
		body, _ = sjson.Set(body, "set-overload-bit-levels.level", []interface{}{})
		for index, item := range data.SetOverloadBitLevels {
			if !item.LevelId.IsNull() && !item.LevelId.IsUnknown() {
				body, _ = sjson.Set(body, "set-overload-bit-levels.level"+"."+strconv.Itoa(index)+"."+"level-id", strconv.FormatInt(item.LevelId.ValueInt64(), 10))
			}
			if !item.OnStartupAdvertiseAsOverloaded.IsNull() && !item.OnStartupAdvertiseAsOverloaded.IsUnknown() {
				if item.OnStartupAdvertiseAsOverloaded.ValueBool() {
					body, _ = sjson.Set(body, "set-overload-bit-levels.level"+"."+strconv.Itoa(index)+"."+"on-startup.advertise-as-overloaded", map[string]string{})
				}
			}
			if !item.OnStartupAdvertiseAsOverloadedTimeToAdvertise.IsNull() && !item.OnStartupAdvertiseAsOverloadedTimeToAdvertise.IsUnknown() {
				body, _ = sjson.Set(body, "set-overload-bit-levels.level"+"."+strconv.Itoa(index)+"."+"on-startup.advertise-as-overloaded.time-to-advertise", strconv.FormatInt(item.OnStartupAdvertiseAsOverloadedTimeToAdvertise.ValueInt64(), 10))
			}
			if !item.OnStartupWaitForBgp.IsNull() && !item.OnStartupWaitForBgp.IsUnknown() {
				if item.OnStartupWaitForBgp.ValueBool() {
					body, _ = sjson.Set(body, "set-overload-bit-levels.level"+"."+strconv.Itoa(index)+"."+"on-startup.wait-for-bgp", map[string]string{})
				}
			}
			if !item.AdvertiseExternal.IsNull() && !item.AdvertiseExternal.IsUnknown() {
				if item.AdvertiseExternal.ValueBool() {
					body, _ = sjson.Set(body, "set-overload-bit-levels.level"+"."+strconv.Itoa(index)+"."+"advertise.external", map[string]string{})
				}
			}
			if !item.AdvertiseInterlevel.IsNull() && !item.AdvertiseInterlevel.IsUnknown() {
				if item.AdvertiseInterlevel.ValueBool() {
					body, _ = sjson.Set(body, "set-overload-bit-levels.level"+"."+strconv.Itoa(index)+"."+"advertise.interlevel", map[string]string{})
				}
			}
		}
	}
	if len(data.AffinityMaps) > 0 {
		body, _ = sjson.Set(body, "affinity-maps.affinity-map", []interface{}{})
		for index, item := range data.AffinityMaps {
			if !item.Name.IsNull() && !item.Name.IsUnknown() {
				body, _ = sjson.Set(body, "affinity-maps.affinity-map"+"."+strconv.Itoa(index)+"."+"affinity-map-name", item.Name.ValueString())
			}
			if !item.BitPosition.IsNull() && !item.BitPosition.IsUnknown() {
				body, _ = sjson.Set(body, "affinity-maps.affinity-map"+"."+strconv.Itoa(index)+"."+"bit-position", strconv.FormatInt(item.BitPosition.ValueInt64(), 10))
			}
		}
	}
	if len(data.FlexAlgos) > 0 {
		body, _ = sjson.Set(body, "flex-algos.flex-algo", []interface{}{})
		for index, item := range data.FlexAlgos {
			if !item.AlgorithmNumber.IsNull() && !item.AlgorithmNumber.IsUnknown() {
				body, _ = sjson.Set(body, "flex-algos.flex-algo"+"."+strconv.Itoa(index)+"."+"algorithm-number", strconv.FormatInt(item.AlgorithmNumber.ValueInt64(), 10))
			}
			if !item.AdvertiseDefinition.IsNull() && !item.AdvertiseDefinition.IsUnknown() {
				if item.AdvertiseDefinition.ValueBool() {
					body, _ = sjson.Set(body, "flex-algos.flex-algo"+"."+strconv.Itoa(index)+"."+"advertise-definition", map[string]string{})
				}
			}
			if !item.MetricTypeDelay.IsNull() && !item.MetricTypeDelay.IsUnknown() {
				if item.MetricTypeDelay.ValueBool() {
					body, _ = sjson.Set(body, "flex-algos.flex-algo"+"."+strconv.Itoa(index)+"."+"metric-type.delay", map[string]string{})
				}
			}
		}
	}
	if len(data.Nets) > 0 {
		body, _ = sjson.Set(body, "nets.net", []interface{}{})
		for index, item := range data.Nets {
			if !item.NetId.IsNull() && !item.NetId.IsUnknown() {
				body, _ = sjson.Set(body, "nets.net"+"."+strconv.Itoa(index)+"."+"net-id", item.NetId.ValueString())
			}
		}
	}
	if len(data.Interfaces) > 0 {
		body, _ = sjson.Set(body, "interfaces.interface", []interface{}{})
		for index, item := range data.Interfaces {
			if !item.InterfaceName.IsNull() && !item.InterfaceName.IsUnknown() {
				body, _ = sjson.Set(body, "interfaces.interface"+"."+strconv.Itoa(index)+"."+"interface-name", item.InterfaceName.ValueString())
			}
			if !item.CircuitType.IsNull() && !item.CircuitType.IsUnknown() {
				body, _ = sjson.Set(body, "interfaces.interface"+"."+strconv.Itoa(index)+"."+"circuit-type", item.CircuitType.ValueString())
			}
			if !item.HelloPaddingDisable.IsNull() && !item.HelloPaddingDisable.IsUnknown() {
				if item.HelloPaddingDisable.ValueBool() {
					body, _ = sjson.Set(body, "interfaces.interface"+"."+strconv.Itoa(index)+"."+"hello-padding.disable", map[string]string{})
				}
			}
			if !item.HelloPaddingSometimes.IsNull() && !item.HelloPaddingSometimes.IsUnknown() {
				if item.HelloPaddingSometimes.ValueBool() {
					body, _ = sjson.Set(body, "interfaces.interface"+"."+strconv.Itoa(index)+"."+"hello-padding.sometimes", map[string]string{})
				}
			}
			if !item.Priority.IsNull() && !item.Priority.IsUnknown() {
				body, _ = sjson.Set(body, "interfaces.interface"+"."+strconv.Itoa(index)+"."+"priority.priority-value", strconv.FormatInt(item.Priority.ValueInt64(), 10))
			}
			if !item.PointToPoint.IsNull() && !item.PointToPoint.IsUnknown() {
				if item.PointToPoint.ValueBool() {
					body, _ = sjson.Set(body, "interfaces.interface"+"."+strconv.Itoa(index)+"."+"point-to-point", map[string]string{})
				}
			}
			if !item.Passive.IsNull() && !item.Passive.IsUnknown() {
				if item.Passive.ValueBool() {
					body, _ = sjson.Set(body, "interfaces.interface"+"."+strconv.Itoa(index)+"."+"passive", map[string]string{})
				}
			}
			if !item.Suppressed.IsNull() && !item.Suppressed.IsUnknown() {
				if item.Suppressed.ValueBool() {
					body, _ = sjson.Set(body, "interfaces.interface"+"."+strconv.Itoa(index)+"."+"suppressed", map[string]string{})
				}
			}
			if !item.Shutdown.IsNull() && !item.Shutdown.IsUnknown() {
				if item.Shutdown.ValueBool() {
					body, _ = sjson.Set(body, "interfaces.interface"+"."+strconv.Itoa(index)+"."+"shutdown", map[string]string{})
				}
			}
		}
	}
	return body
}

func (data *RouterISIS) updateFromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "is-type"); value.Exists() && !data.IsType.IsNull() {
		data.IsType = types.StringValue(value.String())
	} else {
		data.IsType = types.StringNull()
	}
	for i := range data.SetOverloadBitLevels {
		keys := [...]string{"level-id"}
		keyValues := [...]string{strconv.FormatInt(data.SetOverloadBitLevels[i].LevelId.ValueInt64(), 10)}

		var r gjson.Result
		gjson.GetBytes(res, "set-overload-bit-levels.level").ForEach(
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
		if value := r.Get("level-id"); value.Exists() && !data.SetOverloadBitLevels[i].LevelId.IsNull() {
			data.SetOverloadBitLevels[i].LevelId = types.Int64Value(value.Int())
		} else {
			data.SetOverloadBitLevels[i].LevelId = types.Int64Null()
		}
		if value := r.Get("on-startup.advertise-as-overloaded"); !data.SetOverloadBitLevels[i].OnStartupAdvertiseAsOverloaded.IsNull() {
			if value.Exists() {
				data.SetOverloadBitLevels[i].OnStartupAdvertiseAsOverloaded = types.BoolValue(true)
			} else {
				data.SetOverloadBitLevels[i].OnStartupAdvertiseAsOverloaded = types.BoolValue(false)
			}
		} else {
			data.SetOverloadBitLevels[i].OnStartupAdvertiseAsOverloaded = types.BoolNull()
		}
		if value := r.Get("on-startup.advertise-as-overloaded.time-to-advertise"); value.Exists() && !data.SetOverloadBitLevels[i].OnStartupAdvertiseAsOverloadedTimeToAdvertise.IsNull() {
			data.SetOverloadBitLevels[i].OnStartupAdvertiseAsOverloadedTimeToAdvertise = types.Int64Value(value.Int())
		} else {
			data.SetOverloadBitLevels[i].OnStartupAdvertiseAsOverloadedTimeToAdvertise = types.Int64Null()
		}
		if value := r.Get("on-startup.wait-for-bgp"); !data.SetOverloadBitLevels[i].OnStartupWaitForBgp.IsNull() {
			if value.Exists() {
				data.SetOverloadBitLevels[i].OnStartupWaitForBgp = types.BoolValue(true)
			} else {
				data.SetOverloadBitLevels[i].OnStartupWaitForBgp = types.BoolValue(false)
			}
		} else {
			data.SetOverloadBitLevels[i].OnStartupWaitForBgp = types.BoolNull()
		}
		if value := r.Get("advertise.external"); !data.SetOverloadBitLevels[i].AdvertiseExternal.IsNull() {
			if value.Exists() {
				data.SetOverloadBitLevels[i].AdvertiseExternal = types.BoolValue(true)
			} else {
				data.SetOverloadBitLevels[i].AdvertiseExternal = types.BoolValue(false)
			}
		} else {
			data.SetOverloadBitLevels[i].AdvertiseExternal = types.BoolNull()
		}
		if value := r.Get("advertise.interlevel"); !data.SetOverloadBitLevels[i].AdvertiseInterlevel.IsNull() {
			if value.Exists() {
				data.SetOverloadBitLevels[i].AdvertiseInterlevel = types.BoolValue(true)
			} else {
				data.SetOverloadBitLevels[i].AdvertiseInterlevel = types.BoolValue(false)
			}
		} else {
			data.SetOverloadBitLevels[i].AdvertiseInterlevel = types.BoolNull()
		}
	}
	if value := gjson.GetBytes(res, "nsr"); !data.Nsr.IsNull() {
		if value.Exists() {
			data.Nsr = types.BoolValue(true)
		} else {
			data.Nsr = types.BoolValue(false)
		}
	} else {
		data.Nsr = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "nsf.cisco"); !data.NsfCisco.IsNull() {
		if value.Exists() {
			data.NsfCisco = types.BoolValue(true)
		} else {
			data.NsfCisco = types.BoolValue(false)
		}
	} else {
		data.NsfCisco = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "nsf.ietf"); !data.NsfIetf.IsNull() {
		if value.Exists() {
			data.NsfIetf = types.BoolValue(true)
		} else {
			data.NsfIetf = types.BoolValue(false)
		}
	} else {
		data.NsfIetf = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "nsf.lifetime"); value.Exists() && !data.NsfLifetime.IsNull() {
		data.NsfLifetime = types.Int64Value(value.Int())
	} else {
		data.NsfLifetime = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "nsf.interface-timer"); value.Exists() && !data.NsfInterfaceTimer.IsNull() {
		data.NsfInterfaceTimer = types.Int64Value(value.Int())
	} else {
		data.NsfInterfaceTimer = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "nsf.interface-expires"); value.Exists() && !data.NsfInterfaceExpires.IsNull() {
		data.NsfInterfaceExpires = types.Int64Value(value.Int())
	} else {
		data.NsfInterfaceExpires = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "log.adjacency.changes"); !data.LogAdjacencyChanges.IsNull() {
		if value.Exists() {
			data.LogAdjacencyChanges = types.BoolValue(true)
		} else {
			data.LogAdjacencyChanges = types.BoolValue(false)
		}
	} else {
		data.LogAdjacencyChanges = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "lsp-gen-interval.maximum-wait"); value.Exists() && !data.LspGenIntervalMaximumWait.IsNull() {
		data.LspGenIntervalMaximumWait = types.Int64Value(value.Int())
	} else {
		data.LspGenIntervalMaximumWait = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "lsp-gen-interval.initial-wait"); value.Exists() && !data.LspGenIntervalInitialWait.IsNull() {
		data.LspGenIntervalInitialWait = types.Int64Value(value.Int())
	} else {
		data.LspGenIntervalInitialWait = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "lsp-gen-interval.secondary-wait"); value.Exists() && !data.LspGenIntervalSecondaryWait.IsNull() {
		data.LspGenIntervalSecondaryWait = types.Int64Value(value.Int())
	} else {
		data.LspGenIntervalSecondaryWait = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "lsp-refresh-interval.lsp-refresh-interval-time"); value.Exists() && !data.LspRefreshInterval.IsNull() {
		data.LspRefreshInterval = types.Int64Value(value.Int())
	} else {
		data.LspRefreshInterval = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "max-lsp-lifetime.max-lsp-lifetime"); value.Exists() && !data.MaxLspLifetime.IsNull() {
		data.MaxLspLifetime = types.Int64Value(value.Int())
	} else {
		data.MaxLspLifetime = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "lsp-password.keychain.keychain-name"); value.Exists() && !data.LspPasswordKeychain.IsNull() {
		data.LspPasswordKeychain = types.StringValue(value.String())
	} else {
		data.LspPasswordKeychain = types.StringNull()
	}
	if value := gjson.GetBytes(res, "distribute.link-state.instance-id"); value.Exists() && !data.DistributeLinkStateInstanceId.IsNull() {
		data.DistributeLinkStateInstanceId = types.Int64Value(value.Int())
	} else {
		data.DistributeLinkStateInstanceId = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "distribute.link-state.throttle"); value.Exists() && !data.DistributeLinkStateThrottle.IsNull() {
		data.DistributeLinkStateThrottle = types.Int64Value(value.Int())
	} else {
		data.DistributeLinkStateThrottle = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "distribute.link-state.level"); value.Exists() && !data.DistributeLinkStateLevel.IsNull() {
		data.DistributeLinkStateLevel = types.Int64Value(value.Int())
	} else {
		data.DistributeLinkStateLevel = types.Int64Null()
	}
	for i := range data.AffinityMaps {
		keys := [...]string{"affinity-map-name"}
		keyValues := [...]string{data.AffinityMaps[i].Name.ValueString()}

		var r gjson.Result
		gjson.GetBytes(res, "affinity-maps.affinity-map").ForEach(
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
		if value := r.Get("affinity-map-name"); value.Exists() && !data.AffinityMaps[i].Name.IsNull() {
			data.AffinityMaps[i].Name = types.StringValue(value.String())
		} else {
			data.AffinityMaps[i].Name = types.StringNull()
		}
		if value := r.Get("bit-position"); value.Exists() && !data.AffinityMaps[i].BitPosition.IsNull() {
			data.AffinityMaps[i].BitPosition = types.Int64Value(value.Int())
		} else {
			data.AffinityMaps[i].BitPosition = types.Int64Null()
		}
	}
	for i := range data.FlexAlgos {
		keys := [...]string{"algorithm-number"}
		keyValues := [...]string{strconv.FormatInt(data.FlexAlgos[i].AlgorithmNumber.ValueInt64(), 10)}

		var r gjson.Result
		gjson.GetBytes(res, "flex-algos.flex-algo").ForEach(
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
		if value := r.Get("algorithm-number"); value.Exists() && !data.FlexAlgos[i].AlgorithmNumber.IsNull() {
			data.FlexAlgos[i].AlgorithmNumber = types.Int64Value(value.Int())
		} else {
			data.FlexAlgos[i].AlgorithmNumber = types.Int64Null()
		}
		if value := r.Get("advertise-definition"); !data.FlexAlgos[i].AdvertiseDefinition.IsNull() {
			if value.Exists() {
				data.FlexAlgos[i].AdvertiseDefinition = types.BoolValue(true)
			} else {
				data.FlexAlgos[i].AdvertiseDefinition = types.BoolValue(false)
			}
		} else {
			data.FlexAlgos[i].AdvertiseDefinition = types.BoolNull()
		}
		if value := r.Get("metric-type.delay"); !data.FlexAlgos[i].MetricTypeDelay.IsNull() {
			if value.Exists() {
				data.FlexAlgos[i].MetricTypeDelay = types.BoolValue(true)
			} else {
				data.FlexAlgos[i].MetricTypeDelay = types.BoolValue(false)
			}
		} else {
			data.FlexAlgos[i].MetricTypeDelay = types.BoolNull()
		}
	}
	for i := range data.Nets {
		keys := [...]string{"net-id"}
		keyValues := [...]string{data.Nets[i].NetId.ValueString()}

		var r gjson.Result
		gjson.GetBytes(res, "nets.net").ForEach(
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
		if value := r.Get("net-id"); value.Exists() && !data.Nets[i].NetId.IsNull() {
			data.Nets[i].NetId = types.StringValue(value.String())
		} else {
			data.Nets[i].NetId = types.StringNull()
		}
	}
	for i := range data.Interfaces {
		keys := [...]string{"interface-name"}
		keyValues := [...]string{data.Interfaces[i].InterfaceName.ValueString()}

		var r gjson.Result
		gjson.GetBytes(res, "interfaces.interface").ForEach(
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
		if value := r.Get("interface-name"); value.Exists() && !data.Interfaces[i].InterfaceName.IsNull() {
			data.Interfaces[i].InterfaceName = types.StringValue(value.String())
		} else {
			data.Interfaces[i].InterfaceName = types.StringNull()
		}
		if value := r.Get("circuit-type"); value.Exists() && !data.Interfaces[i].CircuitType.IsNull() {
			data.Interfaces[i].CircuitType = types.StringValue(value.String())
		} else {
			data.Interfaces[i].CircuitType = types.StringNull()
		}
		if value := r.Get("hello-padding.disable"); !data.Interfaces[i].HelloPaddingDisable.IsNull() {
			if value.Exists() {
				data.Interfaces[i].HelloPaddingDisable = types.BoolValue(true)
			} else {
				data.Interfaces[i].HelloPaddingDisable = types.BoolValue(false)
			}
		} else {
			data.Interfaces[i].HelloPaddingDisable = types.BoolNull()
		}
		if value := r.Get("hello-padding.sometimes"); !data.Interfaces[i].HelloPaddingSometimes.IsNull() {
			if value.Exists() {
				data.Interfaces[i].HelloPaddingSometimes = types.BoolValue(true)
			} else {
				data.Interfaces[i].HelloPaddingSometimes = types.BoolValue(false)
			}
		} else {
			data.Interfaces[i].HelloPaddingSometimes = types.BoolNull()
		}
		if value := r.Get("priority.priority-value"); value.Exists() && !data.Interfaces[i].Priority.IsNull() {
			data.Interfaces[i].Priority = types.Int64Value(value.Int())
		} else {
			data.Interfaces[i].Priority = types.Int64Null()
		}
		if value := r.Get("point-to-point"); !data.Interfaces[i].PointToPoint.IsNull() {
			if value.Exists() {
				data.Interfaces[i].PointToPoint = types.BoolValue(true)
			} else {
				data.Interfaces[i].PointToPoint = types.BoolValue(false)
			}
		} else {
			data.Interfaces[i].PointToPoint = types.BoolNull()
		}
		if value := r.Get("passive"); !data.Interfaces[i].Passive.IsNull() {
			if value.Exists() {
				data.Interfaces[i].Passive = types.BoolValue(true)
			} else {
				data.Interfaces[i].Passive = types.BoolValue(false)
			}
		} else {
			data.Interfaces[i].Passive = types.BoolNull()
		}
		if value := r.Get("suppressed"); !data.Interfaces[i].Suppressed.IsNull() {
			if value.Exists() {
				data.Interfaces[i].Suppressed = types.BoolValue(true)
			} else {
				data.Interfaces[i].Suppressed = types.BoolValue(false)
			}
		} else {
			data.Interfaces[i].Suppressed = types.BoolNull()
		}
		if value := r.Get("shutdown"); !data.Interfaces[i].Shutdown.IsNull() {
			if value.Exists() {
				data.Interfaces[i].Shutdown = types.BoolValue(true)
			} else {
				data.Interfaces[i].Shutdown = types.BoolValue(false)
			}
		} else {
			data.Interfaces[i].Shutdown = types.BoolNull()
		}
	}
}

func (data *RouterISIS) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "is-type"); value.Exists() {
		data.IsType = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "set-overload-bit-levels.level"); value.Exists() {
		data.SetOverloadBitLevels = make([]RouterISISSetOverloadBitLevels, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := RouterISISSetOverloadBitLevels{}
			if cValue := v.Get("level-id"); cValue.Exists() {
				item.LevelId = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("on-startup.advertise-as-overloaded"); cValue.Exists() {
				item.OnStartupAdvertiseAsOverloaded = types.BoolValue(true)
			} else {
				item.OnStartupAdvertiseAsOverloaded = types.BoolValue(false)
			}
			if cValue := v.Get("on-startup.advertise-as-overloaded.time-to-advertise"); cValue.Exists() {
				item.OnStartupAdvertiseAsOverloadedTimeToAdvertise = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("on-startup.wait-for-bgp"); cValue.Exists() {
				item.OnStartupWaitForBgp = types.BoolValue(true)
			} else {
				item.OnStartupWaitForBgp = types.BoolValue(false)
			}
			if cValue := v.Get("advertise.external"); cValue.Exists() {
				item.AdvertiseExternal = types.BoolValue(true)
			} else {
				item.AdvertiseExternal = types.BoolValue(false)
			}
			if cValue := v.Get("advertise.interlevel"); cValue.Exists() {
				item.AdvertiseInterlevel = types.BoolValue(true)
			} else {
				item.AdvertiseInterlevel = types.BoolValue(false)
			}
			data.SetOverloadBitLevels = append(data.SetOverloadBitLevels, item)
			return true
		})
	}
	if value := gjson.GetBytes(res, "nsr"); value.Exists() {
		data.Nsr = types.BoolValue(true)
	} else {
		data.Nsr = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "nsf.cisco"); value.Exists() {
		data.NsfCisco = types.BoolValue(true)
	} else {
		data.NsfCisco = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "nsf.ietf"); value.Exists() {
		data.NsfIetf = types.BoolValue(true)
	} else {
		data.NsfIetf = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "nsf.lifetime"); value.Exists() {
		data.NsfLifetime = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "nsf.interface-timer"); value.Exists() {
		data.NsfInterfaceTimer = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "nsf.interface-expires"); value.Exists() {
		data.NsfInterfaceExpires = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "log.adjacency.changes"); value.Exists() {
		data.LogAdjacencyChanges = types.BoolValue(true)
	} else {
		data.LogAdjacencyChanges = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "lsp-gen-interval.maximum-wait"); value.Exists() {
		data.LspGenIntervalMaximumWait = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "lsp-gen-interval.initial-wait"); value.Exists() {
		data.LspGenIntervalInitialWait = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "lsp-gen-interval.secondary-wait"); value.Exists() {
		data.LspGenIntervalSecondaryWait = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "lsp-refresh-interval.lsp-refresh-interval-time"); value.Exists() {
		data.LspRefreshInterval = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "max-lsp-lifetime.max-lsp-lifetime"); value.Exists() {
		data.MaxLspLifetime = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "lsp-password.keychain.keychain-name"); value.Exists() {
		data.LspPasswordKeychain = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "distribute.link-state.instance-id"); value.Exists() {
		data.DistributeLinkStateInstanceId = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "distribute.link-state.throttle"); value.Exists() {
		data.DistributeLinkStateThrottle = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "distribute.link-state.level"); value.Exists() {
		data.DistributeLinkStateLevel = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "affinity-maps.affinity-map"); value.Exists() {
		data.AffinityMaps = make([]RouterISISAffinityMaps, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := RouterISISAffinityMaps{}
			if cValue := v.Get("affinity-map-name"); cValue.Exists() {
				item.Name = types.StringValue(cValue.String())
			}
			if cValue := v.Get("bit-position"); cValue.Exists() {
				item.BitPosition = types.Int64Value(cValue.Int())
			}
			data.AffinityMaps = append(data.AffinityMaps, item)
			return true
		})
	}
	if value := gjson.GetBytes(res, "flex-algos.flex-algo"); value.Exists() {
		data.FlexAlgos = make([]RouterISISFlexAlgos, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := RouterISISFlexAlgos{}
			if cValue := v.Get("algorithm-number"); cValue.Exists() {
				item.AlgorithmNumber = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("advertise-definition"); cValue.Exists() {
				item.AdvertiseDefinition = types.BoolValue(true)
			} else {
				item.AdvertiseDefinition = types.BoolValue(false)
			}
			if cValue := v.Get("metric-type.delay"); cValue.Exists() {
				item.MetricTypeDelay = types.BoolValue(true)
			} else {
				item.MetricTypeDelay = types.BoolValue(false)
			}
			data.FlexAlgos = append(data.FlexAlgos, item)
			return true
		})
	}
	if value := gjson.GetBytes(res, "nets.net"); value.Exists() {
		data.Nets = make([]RouterISISNets, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := RouterISISNets{}
			if cValue := v.Get("net-id"); cValue.Exists() {
				item.NetId = types.StringValue(cValue.String())
			}
			data.Nets = append(data.Nets, item)
			return true
		})
	}
	if value := gjson.GetBytes(res, "interfaces.interface"); value.Exists() {
		data.Interfaces = make([]RouterISISInterfaces, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := RouterISISInterfaces{}
			if cValue := v.Get("interface-name"); cValue.Exists() {
				item.InterfaceName = types.StringValue(cValue.String())
			}
			if cValue := v.Get("circuit-type"); cValue.Exists() {
				item.CircuitType = types.StringValue(cValue.String())
			}
			if cValue := v.Get("hello-padding.disable"); cValue.Exists() {
				item.HelloPaddingDisable = types.BoolValue(true)
			} else {
				item.HelloPaddingDisable = types.BoolValue(false)
			}
			if cValue := v.Get("hello-padding.sometimes"); cValue.Exists() {
				item.HelloPaddingSometimes = types.BoolValue(true)
			} else {
				item.HelloPaddingSometimes = types.BoolValue(false)
			}
			if cValue := v.Get("priority.priority-value"); cValue.Exists() {
				item.Priority = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("point-to-point"); cValue.Exists() {
				item.PointToPoint = types.BoolValue(true)
			} else {
				item.PointToPoint = types.BoolValue(false)
			}
			if cValue := v.Get("passive"); cValue.Exists() {
				item.Passive = types.BoolValue(true)
			} else {
				item.Passive = types.BoolValue(false)
			}
			if cValue := v.Get("suppressed"); cValue.Exists() {
				item.Suppressed = types.BoolValue(true)
			} else {
				item.Suppressed = types.BoolValue(false)
			}
			if cValue := v.Get("shutdown"); cValue.Exists() {
				item.Shutdown = types.BoolValue(true)
			} else {
				item.Shutdown = types.BoolValue(false)
			}
			data.Interfaces = append(data.Interfaces, item)
			return true
		})
	}
}

func (data *RouterISIS) fromPlan(ctx context.Context, plan RouterISIS) {
	data.Device = plan.Device
	data.ProcessId = types.StringValue(plan.ProcessId.ValueString())
}

func (data *RouterISIS) getDeletedListItems(ctx context.Context, state RouterISIS) []string {
	deletedListItems := make([]string, 0)
	for i := range state.SetOverloadBitLevels {
		keys := [...]string{"level-id"}
		stateKeyValues := [...]string{strconv.FormatInt(state.SetOverloadBitLevels[i].LevelId.ValueInt64(), 10)}

		emptyKeys := true
		if !reflect.ValueOf(state.SetOverloadBitLevels[i].LevelId.ValueInt64()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.SetOverloadBitLevels {
			found = true
			if state.SetOverloadBitLevels[i].LevelId.ValueInt64() != data.SetOverloadBitLevels[j].LevelId.ValueInt64() {
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
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/set-overload-bit-levels/level%v", state.getPath(), keyString))
		}
	}
	for i := range state.AffinityMaps {
		keys := [...]string{"affinity-map-name"}
		stateKeyValues := [...]string{state.AffinityMaps[i].Name.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.AffinityMaps[i].Name.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.AffinityMaps {
			found = true
			if state.AffinityMaps[i].Name.ValueString() != data.AffinityMaps[j].Name.ValueString() {
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
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/affinity-maps/affinity-map%v", state.getPath(), keyString))
		}
	}
	for i := range state.FlexAlgos {
		keys := [...]string{"algorithm-number"}
		stateKeyValues := [...]string{strconv.FormatInt(state.FlexAlgos[i].AlgorithmNumber.ValueInt64(), 10)}

		emptyKeys := true
		if !reflect.ValueOf(state.FlexAlgos[i].AlgorithmNumber.ValueInt64()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.FlexAlgos {
			found = true
			if state.FlexAlgos[i].AlgorithmNumber.ValueInt64() != data.FlexAlgos[j].AlgorithmNumber.ValueInt64() {
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
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/flex-algos/flex-algo%v", state.getPath(), keyString))
		}
	}
	for i := range state.Nets {
		keys := [...]string{"net-id"}
		stateKeyValues := [...]string{state.Nets[i].NetId.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.Nets[i].NetId.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.Nets {
			found = true
			if state.Nets[i].NetId.ValueString() != data.Nets[j].NetId.ValueString() {
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
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/nets/net%v", state.getPath(), keyString))
		}
	}
	for i := range state.Interfaces {
		keys := [...]string{"interface-name"}
		stateKeyValues := [...]string{state.Interfaces[i].InterfaceName.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.Interfaces[i].InterfaceName.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.Interfaces {
			found = true
			if state.Interfaces[i].InterfaceName.ValueString() != data.Interfaces[j].InterfaceName.ValueString() {
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
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/interfaces/interface%v", state.getPath(), keyString))
		}
	}
	return deletedListItems
}

func (data *RouterISIS) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)

	return emptyLeafsDelete
}
