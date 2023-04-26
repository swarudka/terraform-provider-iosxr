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

type RouterOSPFVRF struct {
	Device                                types.String                    `tfsdk:"device"`
	Id                                    types.String                    `tfsdk:"id"`
	ProcessName                           types.String                    `tfsdk:"process_name"`
	VrfName                               types.String                    `tfsdk:"vrf_name"`
	MplsLdpSync                           types.Bool                      `tfsdk:"mpls_ldp_sync"`
	HelloInterval                         types.Int64                     `tfsdk:"hello_interval"`
	DeadInterval                          types.Int64                     `tfsdk:"dead_interval"`
	Priority                              types.Int64                     `tfsdk:"priority"`
	MtuIgnoreEnable                       types.Bool                      `tfsdk:"mtu_ignore_enable"`
	MtuIgnoreDisable                      types.Bool                      `tfsdk:"mtu_ignore_disable"`
	PassiveEnable                         types.Bool                      `tfsdk:"passive_enable"`
	PassiveDisable                        types.Bool                      `tfsdk:"passive_disable"`
	RouterId                              types.String                    `tfsdk:"router_id"`
	RedistributeConnected                 types.Bool                      `tfsdk:"redistribute_connected"`
	RedistributeConnectedTag              types.Int64                     `tfsdk:"redistribute_connected_tag"`
	RedistributeConnectedMetricType       types.String                    `tfsdk:"redistribute_connected_metric_type"`
	RedistributeStatic                    types.Bool                      `tfsdk:"redistribute_static"`
	RedistributeStaticTag                 types.Int64                     `tfsdk:"redistribute_static_tag"`
	RedistributeStaticMetricType          types.String                    `tfsdk:"redistribute_static_metric_type"`
	BfdFastDetect                         types.Bool                      `tfsdk:"bfd_fast_detect"`
	BfdMinimumInterval                    types.Int64                     `tfsdk:"bfd_minimum_interval"`
	BfdMultiplier                         types.Int64                     `tfsdk:"bfd_multiplier"`
	DefaultInformationOriginate           types.Bool                      `tfsdk:"default_information_originate"`
	DefaultInformationOriginateAlways     types.Bool                      `tfsdk:"default_information_originate_always"`
	DefaultInformationOriginateMetricType types.Int64                     `tfsdk:"default_information_originate_metric_type"`
	Areas                                 []RouterOSPFVRFAreas            `tfsdk:"areas"`
	RedistributeBgp                       []RouterOSPFVRFRedistributeBgp  `tfsdk:"redistribute_bgp"`
	RedistributeIsis                      []RouterOSPFVRFRedistributeIsis `tfsdk:"redistribute_isis"`
	RedistributeOspf                      []RouterOSPFVRFRedistributeOspf `tfsdk:"redistribute_ospf"`
}
type RouterOSPFVRFAreas struct {
	AreaId types.String `tfsdk:"area_id"`
}
type RouterOSPFVRFRedistributeBgp struct {
	AsNumber   types.String `tfsdk:"as_number"`
	Tag        types.Int64  `tfsdk:"tag"`
	MetricType types.String `tfsdk:"metric_type"`
}
type RouterOSPFVRFRedistributeIsis struct {
	InstanceName types.String `tfsdk:"instance_name"`
	Level1       types.Bool   `tfsdk:"level_1"`
	Level2       types.Bool   `tfsdk:"level_2"`
	Level12      types.Bool   `tfsdk:"level_1_2"`
	Tag          types.Int64  `tfsdk:"tag"`
	MetricType   types.String `tfsdk:"metric_type"`
}
type RouterOSPFVRFRedistributeOspf struct {
	InstanceName      types.String `tfsdk:"instance_name"`
	MatchInternal     types.Bool   `tfsdk:"match_internal"`
	MatchExternal     types.Bool   `tfsdk:"match_external"`
	MatchNssaExternal types.Bool   `tfsdk:"match_nssa_external"`
	Tag               types.Int64  `tfsdk:"tag"`
	MetricType        types.String `tfsdk:"metric_type"`
}

func (data RouterOSPFVRF) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-router-ospf-cfg:/router/ospf/processes/process[process-name=%s]/vrfs/vrf[vrf-name=%s]", data.ProcessName.ValueString(), data.VrfName.ValueString())
}

func (data RouterOSPFVRF) toBody(ctx context.Context) string {
	body := "{}"
	if !data.VrfName.IsNull() && !data.VrfName.IsUnknown() {
		body, _ = sjson.Set(body, "vrf-name", data.VrfName.ValueString())
	}
	if !data.MplsLdpSync.IsNull() && !data.MplsLdpSync.IsUnknown() {
		if data.MplsLdpSync.ValueBool() {
			body, _ = sjson.Set(body, "mpls.ldp.sync", map[string]string{})
		}
	}
	if !data.HelloInterval.IsNull() && !data.HelloInterval.IsUnknown() {
		body, _ = sjson.Set(body, "hello-interval", strconv.FormatInt(data.HelloInterval.ValueInt64(), 10))
	}
	if !data.DeadInterval.IsNull() && !data.DeadInterval.IsUnknown() {
		body, _ = sjson.Set(body, "dead-interval", strconv.FormatInt(data.DeadInterval.ValueInt64(), 10))
	}
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		body, _ = sjson.Set(body, "priority", strconv.FormatInt(data.Priority.ValueInt64(), 10))
	}
	if !data.MtuIgnoreEnable.IsNull() && !data.MtuIgnoreEnable.IsUnknown() {
		if data.MtuIgnoreEnable.ValueBool() {
			body, _ = sjson.Set(body, "mtu-ignore.enable", map[string]string{})
		}
	}
	if !data.MtuIgnoreDisable.IsNull() && !data.MtuIgnoreDisable.IsUnknown() {
		if data.MtuIgnoreDisable.ValueBool() {
			body, _ = sjson.Set(body, "mtu-ignore.disable", map[string]string{})
		}
	}
	if !data.PassiveEnable.IsNull() && !data.PassiveEnable.IsUnknown() {
		if data.PassiveEnable.ValueBool() {
			body, _ = sjson.Set(body, "passive.enable", map[string]string{})
		}
	}
	if !data.PassiveDisable.IsNull() && !data.PassiveDisable.IsUnknown() {
		if data.PassiveDisable.ValueBool() {
			body, _ = sjson.Set(body, "passive.disable", map[string]string{})
		}
	}
	if !data.RouterId.IsNull() && !data.RouterId.IsUnknown() {
		body, _ = sjson.Set(body, "router-id", data.RouterId.ValueString())
	}
	if !data.RedistributeConnected.IsNull() && !data.RedistributeConnected.IsUnknown() {
		if data.RedistributeConnected.ValueBool() {
			body, _ = sjson.Set(body, "redistribute.connected", map[string]string{})
		}
	}
	if !data.RedistributeConnectedTag.IsNull() && !data.RedistributeConnectedTag.IsUnknown() {
		body, _ = sjson.Set(body, "redistribute.connected.tag", strconv.FormatInt(data.RedistributeConnectedTag.ValueInt64(), 10))
	}
	if !data.RedistributeConnectedMetricType.IsNull() && !data.RedistributeConnectedMetricType.IsUnknown() {
		body, _ = sjson.Set(body, "redistribute.connected.metric-type", data.RedistributeConnectedMetricType.ValueString())
	}
	if !data.RedistributeStatic.IsNull() && !data.RedistributeStatic.IsUnknown() {
		if data.RedistributeStatic.ValueBool() {
			body, _ = sjson.Set(body, "redistribute.static", map[string]string{})
		}
	}
	if !data.RedistributeStaticTag.IsNull() && !data.RedistributeStaticTag.IsUnknown() {
		body, _ = sjson.Set(body, "redistribute.static.tag", strconv.FormatInt(data.RedistributeStaticTag.ValueInt64(), 10))
	}
	if !data.RedistributeStaticMetricType.IsNull() && !data.RedistributeStaticMetricType.IsUnknown() {
		body, _ = sjson.Set(body, "redistribute.static.metric-type", data.RedistributeStaticMetricType.ValueString())
	}
	if !data.BfdFastDetect.IsNull() && !data.BfdFastDetect.IsUnknown() {
		if data.BfdFastDetect.ValueBool() {
			body, _ = sjson.Set(body, "bfd.fast-detect", map[string]string{})
		}
	}
	if !data.BfdMinimumInterval.IsNull() && !data.BfdMinimumInterval.IsUnknown() {
		body, _ = sjson.Set(body, "bfd.minimum-interval", strconv.FormatInt(data.BfdMinimumInterval.ValueInt64(), 10))
	}
	if !data.BfdMultiplier.IsNull() && !data.BfdMultiplier.IsUnknown() {
		body, _ = sjson.Set(body, "bfd.multiplier", strconv.FormatInt(data.BfdMultiplier.ValueInt64(), 10))
	}
	if !data.DefaultInformationOriginate.IsNull() && !data.DefaultInformationOriginate.IsUnknown() {
		if data.DefaultInformationOriginate.ValueBool() {
			body, _ = sjson.Set(body, "default-information.originate", map[string]string{})
		}
	}
	if !data.DefaultInformationOriginateAlways.IsNull() && !data.DefaultInformationOriginateAlways.IsUnknown() {
		if data.DefaultInformationOriginateAlways.ValueBool() {
			body, _ = sjson.Set(body, "default-information.originate.always", map[string]string{})
		}
	}
	if !data.DefaultInformationOriginateMetricType.IsNull() && !data.DefaultInformationOriginateMetricType.IsUnknown() {
		body, _ = sjson.Set(body, "default-information.originate.metric-type", strconv.FormatInt(data.DefaultInformationOriginateMetricType.ValueInt64(), 10))
	}
	if len(data.Areas) > 0 {
		body, _ = sjson.Set(body, "areas.area", []interface{}{})
		for index, item := range data.Areas {
			if !item.AreaId.IsNull() && !item.AreaId.IsUnknown() {
				body, _ = sjson.Set(body, "areas.area"+"."+strconv.Itoa(index)+"."+"area-id", item.AreaId.ValueString())
			}
		}
	}
	if len(data.RedistributeBgp) > 0 {
		body, _ = sjson.Set(body, "redistribute.bgp.as", []interface{}{})
		for index, item := range data.RedistributeBgp {
			if !item.AsNumber.IsNull() && !item.AsNumber.IsUnknown() {
				body, _ = sjson.Set(body, "redistribute.bgp.as"+"."+strconv.Itoa(index)+"."+"as-number", item.AsNumber.ValueString())
			}
			if !item.Tag.IsNull() && !item.Tag.IsUnknown() {
				body, _ = sjson.Set(body, "redistribute.bgp.as"+"."+strconv.Itoa(index)+"."+"tag", strconv.FormatInt(item.Tag.ValueInt64(), 10))
			}
			if !item.MetricType.IsNull() && !item.MetricType.IsUnknown() {
				body, _ = sjson.Set(body, "redistribute.bgp.as"+"."+strconv.Itoa(index)+"."+"metric-type", item.MetricType.ValueString())
			}
		}
	}
	if len(data.RedistributeIsis) > 0 {
		body, _ = sjson.Set(body, "redistribute.isis", []interface{}{})
		for index, item := range data.RedistributeIsis {
			if !item.InstanceName.IsNull() && !item.InstanceName.IsUnknown() {
				body, _ = sjson.Set(body, "redistribute.isis"+"."+strconv.Itoa(index)+"."+"instance-name", item.InstanceName.ValueString())
			}
			if !item.Level1.IsNull() && !item.Level1.IsUnknown() {
				if item.Level1.ValueBool() {
					body, _ = sjson.Set(body, "redistribute.isis"+"."+strconv.Itoa(index)+"."+"level-1", map[string]string{})
				}
			}
			if !item.Level2.IsNull() && !item.Level2.IsUnknown() {
				if item.Level2.ValueBool() {
					body, _ = sjson.Set(body, "redistribute.isis"+"."+strconv.Itoa(index)+"."+"level-2", map[string]string{})
				}
			}
			if !item.Level12.IsNull() && !item.Level12.IsUnknown() {
				if item.Level12.ValueBool() {
					body, _ = sjson.Set(body, "redistribute.isis"+"."+strconv.Itoa(index)+"."+"level-1-2", map[string]string{})
				}
			}
			if !item.Tag.IsNull() && !item.Tag.IsUnknown() {
				body, _ = sjson.Set(body, "redistribute.isis"+"."+strconv.Itoa(index)+"."+"tag", strconv.FormatInt(item.Tag.ValueInt64(), 10))
			}
			if !item.MetricType.IsNull() && !item.MetricType.IsUnknown() {
				body, _ = sjson.Set(body, "redistribute.isis"+"."+strconv.Itoa(index)+"."+"metric-type", item.MetricType.ValueString())
			}
		}
	}
	if len(data.RedistributeOspf) > 0 {
		body, _ = sjson.Set(body, "redistribute.ospf", []interface{}{})
		for index, item := range data.RedistributeOspf {
			if !item.InstanceName.IsNull() && !item.InstanceName.IsUnknown() {
				body, _ = sjson.Set(body, "redistribute.ospf"+"."+strconv.Itoa(index)+"."+"instance-name", item.InstanceName.ValueString())
			}
			if !item.MatchInternal.IsNull() && !item.MatchInternal.IsUnknown() {
				if item.MatchInternal.ValueBool() {
					body, _ = sjson.Set(body, "redistribute.ospf"+"."+strconv.Itoa(index)+"."+"match.internal", map[string]string{})
				}
			}
			if !item.MatchExternal.IsNull() && !item.MatchExternal.IsUnknown() {
				if item.MatchExternal.ValueBool() {
					body, _ = sjson.Set(body, "redistribute.ospf"+"."+strconv.Itoa(index)+"."+"match.external", map[string]string{})
				}
			}
			if !item.MatchNssaExternal.IsNull() && !item.MatchNssaExternal.IsUnknown() {
				if item.MatchNssaExternal.ValueBool() {
					body, _ = sjson.Set(body, "redistribute.ospf"+"."+strconv.Itoa(index)+"."+"match.nssa-external", map[string]string{})
				}
			}
			if !item.Tag.IsNull() && !item.Tag.IsUnknown() {
				body, _ = sjson.Set(body, "redistribute.ospf"+"."+strconv.Itoa(index)+"."+"tag", strconv.FormatInt(item.Tag.ValueInt64(), 10))
			}
			if !item.MetricType.IsNull() && !item.MetricType.IsUnknown() {
				body, _ = sjson.Set(body, "redistribute.ospf"+"."+strconv.Itoa(index)+"."+"metric-type", item.MetricType.ValueString())
			}
		}
	}
	return body
}

func (data *RouterOSPFVRF) updateFromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "mpls.ldp.sync"); !data.MplsLdpSync.IsNull() {
		if value.Exists() {
			data.MplsLdpSync = types.BoolValue(true)
		} else {
			data.MplsLdpSync = types.BoolValue(false)
		}
	} else {
		data.MplsLdpSync = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "hello-interval"); value.Exists() && !data.HelloInterval.IsNull() {
		data.HelloInterval = types.Int64Value(value.Int())
	} else {
		data.HelloInterval = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "dead-interval"); value.Exists() && !data.DeadInterval.IsNull() {
		data.DeadInterval = types.Int64Value(value.Int())
	} else {
		data.DeadInterval = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "priority"); value.Exists() && !data.Priority.IsNull() {
		data.Priority = types.Int64Value(value.Int())
	} else {
		data.Priority = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "mtu-ignore.enable"); !data.MtuIgnoreEnable.IsNull() {
		if value.Exists() {
			data.MtuIgnoreEnable = types.BoolValue(true)
		} else {
			data.MtuIgnoreEnable = types.BoolValue(false)
		}
	} else {
		data.MtuIgnoreEnable = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "mtu-ignore.disable"); !data.MtuIgnoreDisable.IsNull() {
		if value.Exists() {
			data.MtuIgnoreDisable = types.BoolValue(true)
		} else {
			data.MtuIgnoreDisable = types.BoolValue(false)
		}
	} else {
		data.MtuIgnoreDisable = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "passive.enable"); !data.PassiveEnable.IsNull() {
		if value.Exists() {
			data.PassiveEnable = types.BoolValue(true)
		} else {
			data.PassiveEnable = types.BoolValue(false)
		}
	} else {
		data.PassiveEnable = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "passive.disable"); !data.PassiveDisable.IsNull() {
		if value.Exists() {
			data.PassiveDisable = types.BoolValue(true)
		} else {
			data.PassiveDisable = types.BoolValue(false)
		}
	} else {
		data.PassiveDisable = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "router-id"); value.Exists() && !data.RouterId.IsNull() {
		data.RouterId = types.StringValue(value.String())
	} else {
		data.RouterId = types.StringNull()
	}
	if value := gjson.GetBytes(res, "redistribute.connected"); !data.RedistributeConnected.IsNull() {
		if value.Exists() {
			data.RedistributeConnected = types.BoolValue(true)
		} else {
			data.RedistributeConnected = types.BoolValue(false)
		}
	} else {
		data.RedistributeConnected = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "redistribute.connected.tag"); value.Exists() && !data.RedistributeConnectedTag.IsNull() {
		data.RedistributeConnectedTag = types.Int64Value(value.Int())
	} else {
		data.RedistributeConnectedTag = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "redistribute.connected.metric-type"); value.Exists() && !data.RedistributeConnectedMetricType.IsNull() {
		data.RedistributeConnectedMetricType = types.StringValue(value.String())
	} else {
		data.RedistributeConnectedMetricType = types.StringNull()
	}
	if value := gjson.GetBytes(res, "redistribute.static"); !data.RedistributeStatic.IsNull() {
		if value.Exists() {
			data.RedistributeStatic = types.BoolValue(true)
		} else {
			data.RedistributeStatic = types.BoolValue(false)
		}
	} else {
		data.RedistributeStatic = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "redistribute.static.tag"); value.Exists() && !data.RedistributeStaticTag.IsNull() {
		data.RedistributeStaticTag = types.Int64Value(value.Int())
	} else {
		data.RedistributeStaticTag = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "redistribute.static.metric-type"); value.Exists() && !data.RedistributeStaticMetricType.IsNull() {
		data.RedistributeStaticMetricType = types.StringValue(value.String())
	} else {
		data.RedistributeStaticMetricType = types.StringNull()
	}
	if value := gjson.GetBytes(res, "bfd.fast-detect"); !data.BfdFastDetect.IsNull() {
		if value.Exists() {
			data.BfdFastDetect = types.BoolValue(true)
		} else {
			data.BfdFastDetect = types.BoolValue(false)
		}
	} else {
		data.BfdFastDetect = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "bfd.minimum-interval"); value.Exists() && !data.BfdMinimumInterval.IsNull() {
		data.BfdMinimumInterval = types.Int64Value(value.Int())
	} else {
		data.BfdMinimumInterval = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "bfd.multiplier"); value.Exists() && !data.BfdMultiplier.IsNull() {
		data.BfdMultiplier = types.Int64Value(value.Int())
	} else {
		data.BfdMultiplier = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "default-information.originate"); !data.DefaultInformationOriginate.IsNull() {
		if value.Exists() {
			data.DefaultInformationOriginate = types.BoolValue(true)
		} else {
			data.DefaultInformationOriginate = types.BoolValue(false)
		}
	} else {
		data.DefaultInformationOriginate = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "default-information.originate.always"); !data.DefaultInformationOriginateAlways.IsNull() {
		if value.Exists() {
			data.DefaultInformationOriginateAlways = types.BoolValue(true)
		} else {
			data.DefaultInformationOriginateAlways = types.BoolValue(false)
		}
	} else {
		data.DefaultInformationOriginateAlways = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "default-information.originate.metric-type"); value.Exists() && !data.DefaultInformationOriginateMetricType.IsNull() {
		data.DefaultInformationOriginateMetricType = types.Int64Value(value.Int())
	} else {
		data.DefaultInformationOriginateMetricType = types.Int64Null()
	}
	for i := range data.Areas {
		keys := [...]string{"area-id"}
		keyValues := [...]string{data.Areas[i].AreaId.ValueString()}

		var r gjson.Result
		gjson.GetBytes(res, "areas.area").ForEach(
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
		if value := r.Get("area-id"); value.Exists() && !data.Areas[i].AreaId.IsNull() {
			data.Areas[i].AreaId = types.StringValue(value.String())
		} else {
			data.Areas[i].AreaId = types.StringNull()
		}
	}
	for i := range data.RedistributeBgp {
		keys := [...]string{"as-number"}
		keyValues := [...]string{data.RedistributeBgp[i].AsNumber.ValueString()}

		var r gjson.Result
		gjson.GetBytes(res, "redistribute.bgp.as").ForEach(
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
		if value := r.Get("as-number"); value.Exists() && !data.RedistributeBgp[i].AsNumber.IsNull() {
			data.RedistributeBgp[i].AsNumber = types.StringValue(value.String())
		} else {
			data.RedistributeBgp[i].AsNumber = types.StringNull()
		}
		if value := r.Get("tag"); value.Exists() && !data.RedistributeBgp[i].Tag.IsNull() {
			data.RedistributeBgp[i].Tag = types.Int64Value(value.Int())
		} else {
			data.RedistributeBgp[i].Tag = types.Int64Null()
		}
		if value := r.Get("metric-type"); value.Exists() && !data.RedistributeBgp[i].MetricType.IsNull() {
			data.RedistributeBgp[i].MetricType = types.StringValue(value.String())
		} else {
			data.RedistributeBgp[i].MetricType = types.StringNull()
		}
	}
	for i := range data.RedistributeIsis {
		keys := [...]string{"instance-name"}
		keyValues := [...]string{data.RedistributeIsis[i].InstanceName.ValueString()}

		var r gjson.Result
		gjson.GetBytes(res, "redistribute.isis").ForEach(
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
		if value := r.Get("instance-name"); value.Exists() && !data.RedistributeIsis[i].InstanceName.IsNull() {
			data.RedistributeIsis[i].InstanceName = types.StringValue(value.String())
		} else {
			data.RedistributeIsis[i].InstanceName = types.StringNull()
		}
		if value := r.Get("level-1"); !data.RedistributeIsis[i].Level1.IsNull() {
			if value.Exists() {
				data.RedistributeIsis[i].Level1 = types.BoolValue(true)
			} else {
				data.RedistributeIsis[i].Level1 = types.BoolValue(false)
			}
		} else {
			data.RedistributeIsis[i].Level1 = types.BoolNull()
		}
		if value := r.Get("level-2"); !data.RedistributeIsis[i].Level2.IsNull() {
			if value.Exists() {
				data.RedistributeIsis[i].Level2 = types.BoolValue(true)
			} else {
				data.RedistributeIsis[i].Level2 = types.BoolValue(false)
			}
		} else {
			data.RedistributeIsis[i].Level2 = types.BoolNull()
		}
		if value := r.Get("level-1-2"); !data.RedistributeIsis[i].Level12.IsNull() {
			if value.Exists() {
				data.RedistributeIsis[i].Level12 = types.BoolValue(true)
			} else {
				data.RedistributeIsis[i].Level12 = types.BoolValue(false)
			}
		} else {
			data.RedistributeIsis[i].Level12 = types.BoolNull()
		}
		if value := r.Get("tag"); value.Exists() && !data.RedistributeIsis[i].Tag.IsNull() {
			data.RedistributeIsis[i].Tag = types.Int64Value(value.Int())
		} else {
			data.RedistributeIsis[i].Tag = types.Int64Null()
		}
		if value := r.Get("metric-type"); value.Exists() && !data.RedistributeIsis[i].MetricType.IsNull() {
			data.RedistributeIsis[i].MetricType = types.StringValue(value.String())
		} else {
			data.RedistributeIsis[i].MetricType = types.StringNull()
		}
	}
	for i := range data.RedistributeOspf {
		keys := [...]string{"instance-name"}
		keyValues := [...]string{data.RedistributeOspf[i].InstanceName.ValueString()}

		var r gjson.Result
		gjson.GetBytes(res, "redistribute.ospf").ForEach(
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
		if value := r.Get("instance-name"); value.Exists() && !data.RedistributeOspf[i].InstanceName.IsNull() {
			data.RedistributeOspf[i].InstanceName = types.StringValue(value.String())
		} else {
			data.RedistributeOspf[i].InstanceName = types.StringNull()
		}
		if value := r.Get("match.internal"); !data.RedistributeOspf[i].MatchInternal.IsNull() {
			if value.Exists() {
				data.RedistributeOspf[i].MatchInternal = types.BoolValue(true)
			} else {
				data.RedistributeOspf[i].MatchInternal = types.BoolValue(false)
			}
		} else {
			data.RedistributeOspf[i].MatchInternal = types.BoolNull()
		}
		if value := r.Get("match.external"); !data.RedistributeOspf[i].MatchExternal.IsNull() {
			if value.Exists() {
				data.RedistributeOspf[i].MatchExternal = types.BoolValue(true)
			} else {
				data.RedistributeOspf[i].MatchExternal = types.BoolValue(false)
			}
		} else {
			data.RedistributeOspf[i].MatchExternal = types.BoolNull()
		}
		if value := r.Get("match.nssa-external"); !data.RedistributeOspf[i].MatchNssaExternal.IsNull() {
			if value.Exists() {
				data.RedistributeOspf[i].MatchNssaExternal = types.BoolValue(true)
			} else {
				data.RedistributeOspf[i].MatchNssaExternal = types.BoolValue(false)
			}
		} else {
			data.RedistributeOspf[i].MatchNssaExternal = types.BoolNull()
		}
		if value := r.Get("tag"); value.Exists() && !data.RedistributeOspf[i].Tag.IsNull() {
			data.RedistributeOspf[i].Tag = types.Int64Value(value.Int())
		} else {
			data.RedistributeOspf[i].Tag = types.Int64Null()
		}
		if value := r.Get("metric-type"); value.Exists() && !data.RedistributeOspf[i].MetricType.IsNull() {
			data.RedistributeOspf[i].MetricType = types.StringValue(value.String())
		} else {
			data.RedistributeOspf[i].MetricType = types.StringNull()
		}
	}
}

func (data *RouterOSPFVRF) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "mpls.ldp.sync"); value.Exists() {
		data.MplsLdpSync = types.BoolValue(true)
	} else {
		data.MplsLdpSync = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "hello-interval"); value.Exists() {
		data.HelloInterval = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "dead-interval"); value.Exists() {
		data.DeadInterval = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "priority"); value.Exists() {
		data.Priority = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "mtu-ignore.enable"); value.Exists() {
		data.MtuIgnoreEnable = types.BoolValue(true)
	} else {
		data.MtuIgnoreEnable = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "mtu-ignore.disable"); value.Exists() {
		data.MtuIgnoreDisable = types.BoolValue(true)
	} else {
		data.MtuIgnoreDisable = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "passive.enable"); value.Exists() {
		data.PassiveEnable = types.BoolValue(true)
	} else {
		data.PassiveEnable = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "passive.disable"); value.Exists() {
		data.PassiveDisable = types.BoolValue(true)
	} else {
		data.PassiveDisable = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "router-id"); value.Exists() {
		data.RouterId = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "redistribute.connected"); value.Exists() {
		data.RedistributeConnected = types.BoolValue(true)
	} else {
		data.RedistributeConnected = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "redistribute.connected.tag"); value.Exists() {
		data.RedistributeConnectedTag = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "redistribute.connected.metric-type"); value.Exists() {
		data.RedistributeConnectedMetricType = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "redistribute.static"); value.Exists() {
		data.RedistributeStatic = types.BoolValue(true)
	} else {
		data.RedistributeStatic = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "redistribute.static.tag"); value.Exists() {
		data.RedistributeStaticTag = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "redistribute.static.metric-type"); value.Exists() {
		data.RedistributeStaticMetricType = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "bfd.fast-detect"); value.Exists() {
		data.BfdFastDetect = types.BoolValue(true)
	} else {
		data.BfdFastDetect = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "bfd.minimum-interval"); value.Exists() {
		data.BfdMinimumInterval = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "bfd.multiplier"); value.Exists() {
		data.BfdMultiplier = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "default-information.originate"); value.Exists() {
		data.DefaultInformationOriginate = types.BoolValue(true)
	} else {
		data.DefaultInformationOriginate = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "default-information.originate.always"); value.Exists() {
		data.DefaultInformationOriginateAlways = types.BoolValue(true)
	} else {
		data.DefaultInformationOriginateAlways = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "default-information.originate.metric-type"); value.Exists() {
		data.DefaultInformationOriginateMetricType = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "areas.area"); value.Exists() {
		data.Areas = make([]RouterOSPFVRFAreas, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := RouterOSPFVRFAreas{}
			if cValue := v.Get("area-id"); cValue.Exists() {
				item.AreaId = types.StringValue(cValue.String())
			}
			data.Areas = append(data.Areas, item)
			return true
		})
	}
	if value := gjson.GetBytes(res, "redistribute.bgp.as"); value.Exists() {
		data.RedistributeBgp = make([]RouterOSPFVRFRedistributeBgp, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := RouterOSPFVRFRedistributeBgp{}
			if cValue := v.Get("as-number"); cValue.Exists() {
				item.AsNumber = types.StringValue(cValue.String())
			}
			if cValue := v.Get("tag"); cValue.Exists() {
				item.Tag = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("metric-type"); cValue.Exists() {
				item.MetricType = types.StringValue(cValue.String())
			}
			data.RedistributeBgp = append(data.RedistributeBgp, item)
			return true
		})
	}
	if value := gjson.GetBytes(res, "redistribute.isis"); value.Exists() {
		data.RedistributeIsis = make([]RouterOSPFVRFRedistributeIsis, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := RouterOSPFVRFRedistributeIsis{}
			if cValue := v.Get("instance-name"); cValue.Exists() {
				item.InstanceName = types.StringValue(cValue.String())
			}
			if cValue := v.Get("level-1"); cValue.Exists() {
				item.Level1 = types.BoolValue(true)
			} else {
				item.Level1 = types.BoolValue(false)
			}
			if cValue := v.Get("level-2"); cValue.Exists() {
				item.Level2 = types.BoolValue(true)
			} else {
				item.Level2 = types.BoolValue(false)
			}
			if cValue := v.Get("level-1-2"); cValue.Exists() {
				item.Level12 = types.BoolValue(true)
			} else {
				item.Level12 = types.BoolValue(false)
			}
			if cValue := v.Get("tag"); cValue.Exists() {
				item.Tag = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("metric-type"); cValue.Exists() {
				item.MetricType = types.StringValue(cValue.String())
			}
			data.RedistributeIsis = append(data.RedistributeIsis, item)
			return true
		})
	}
	if value := gjson.GetBytes(res, "redistribute.ospf"); value.Exists() {
		data.RedistributeOspf = make([]RouterOSPFVRFRedistributeOspf, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := RouterOSPFVRFRedistributeOspf{}
			if cValue := v.Get("instance-name"); cValue.Exists() {
				item.InstanceName = types.StringValue(cValue.String())
			}
			if cValue := v.Get("match.internal"); cValue.Exists() {
				item.MatchInternal = types.BoolValue(true)
			} else {
				item.MatchInternal = types.BoolValue(false)
			}
			if cValue := v.Get("match.external"); cValue.Exists() {
				item.MatchExternal = types.BoolValue(true)
			} else {
				item.MatchExternal = types.BoolValue(false)
			}
			if cValue := v.Get("match.nssa-external"); cValue.Exists() {
				item.MatchNssaExternal = types.BoolValue(true)
			} else {
				item.MatchNssaExternal = types.BoolValue(false)
			}
			if cValue := v.Get("tag"); cValue.Exists() {
				item.Tag = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("metric-type"); cValue.Exists() {
				item.MetricType = types.StringValue(cValue.String())
			}
			data.RedistributeOspf = append(data.RedistributeOspf, item)
			return true
		})
	}
}

func (data *RouterOSPFVRF) fromPlan(ctx context.Context, plan RouterOSPFVRF) {
	data.Device = plan.Device
	data.ProcessName = types.StringValue(plan.ProcessName.ValueString())
	data.VrfName = types.StringValue(plan.VrfName.ValueString())
}

func (data *RouterOSPFVRF) getDeletedListItems(ctx context.Context, state RouterOSPFVRF) []string {
	deletedListItems := make([]string, 0)
	for i := range state.Areas {
		keys := [...]string{"area-id"}
		stateKeyValues := [...]string{state.Areas[i].AreaId.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.Areas[i].AreaId.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.Areas {
			found = true
			if state.Areas[i].AreaId.ValueString() != data.Areas[j].AreaId.ValueString() {
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
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/areas/area%v", state.getPath(), keyString))
		}
	}
	for i := range state.RedistributeBgp {
		keys := [...]string{"as-number"}
		stateKeyValues := [...]string{state.RedistributeBgp[i].AsNumber.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.RedistributeBgp[i].AsNumber.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.RedistributeBgp {
			found = true
			if state.RedistributeBgp[i].AsNumber.ValueString() != data.RedistributeBgp[j].AsNumber.ValueString() {
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
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/redistribute/bgp/as%v", state.getPath(), keyString))
		}
	}
	for i := range state.RedistributeIsis {
		keys := [...]string{"instance-name"}
		stateKeyValues := [...]string{state.RedistributeIsis[i].InstanceName.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.RedistributeIsis[i].InstanceName.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.RedistributeIsis {
			found = true
			if state.RedistributeIsis[i].InstanceName.ValueString() != data.RedistributeIsis[j].InstanceName.ValueString() {
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
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/redistribute/isis%v", state.getPath(), keyString))
		}
	}
	for i := range state.RedistributeOspf {
		keys := [...]string{"instance-name"}
		stateKeyValues := [...]string{state.RedistributeOspf[i].InstanceName.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.RedistributeOspf[i].InstanceName.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.RedistributeOspf {
			found = true
			if state.RedistributeOspf[i].InstanceName.ValueString() != data.RedistributeOspf[j].InstanceName.ValueString() {
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
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/redistribute/ospf%v", state.getPath(), keyString))
		}
	}
	return deletedListItems
}

func (data *RouterOSPFVRF) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)

	return emptyLeafsDelete
}
