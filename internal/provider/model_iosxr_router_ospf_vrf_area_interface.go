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

type RouterOSPFVRFAreaInterface struct {
	Device                   types.String `tfsdk:"device"`
	Id                       types.String `tfsdk:"id"`
	ProcessName              types.String `tfsdk:"process_name"`
	VrfName                  types.String `tfsdk:"vrf_name"`
	AreaId                   types.String `tfsdk:"area_id"`
	InterfaceName            types.String `tfsdk:"interface_name"`
	NetworkBroadcast         types.Bool   `tfsdk:"network_broadcast"`
	NetworkNonBroadcast      types.Bool   `tfsdk:"network_non_broadcast"`
	NetworkPointToPoint      types.Bool   `tfsdk:"network_point_to_point"`
	NetworkPointToMultipoint types.Bool   `tfsdk:"network_point_to_multipoint"`
	Cost                     types.Int64  `tfsdk:"cost"`
	Priority                 types.Int64  `tfsdk:"priority"`
	PassiveEnable            types.Bool   `tfsdk:"passive_enable"`
	PassiveDisable           types.Bool   `tfsdk:"passive_disable"`
}

func (data RouterOSPFVRFAreaInterface) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-router-ospf-cfg:/router/ospf/processes/process[process-name=%s]/vrfs/vrf[vrf-name=%s]/areas/area[area-id=%s]/interfaces/interface[interface-name=%s]", data.ProcessName.ValueString(), data.VrfName.ValueString(), data.AreaId.ValueString(), data.InterfaceName.ValueString())
}

func (data RouterOSPFVRFAreaInterface) toBody(ctx context.Context) string {
	body := "{}"
	if !data.InterfaceName.IsNull() && !data.InterfaceName.IsUnknown() {
		body, _ = sjson.Set(body, "interface-name", data.InterfaceName.ValueString())
	}
	if !data.NetworkBroadcast.IsNull() && !data.NetworkBroadcast.IsUnknown() {
		if data.NetworkBroadcast.ValueBool() {
			body, _ = sjson.Set(body, "network.broadcast", map[string]string{})
		}
	}
	if !data.NetworkNonBroadcast.IsNull() && !data.NetworkNonBroadcast.IsUnknown() {
		if data.NetworkNonBroadcast.ValueBool() {
			body, _ = sjson.Set(body, "network.non-broadcast", map[string]string{})
		}
	}
	if !data.NetworkPointToPoint.IsNull() && !data.NetworkPointToPoint.IsUnknown() {
		if data.NetworkPointToPoint.ValueBool() {
			body, _ = sjson.Set(body, "network.point-to-point", map[string]string{})
		}
	}
	if !data.NetworkPointToMultipoint.IsNull() && !data.NetworkPointToMultipoint.IsUnknown() {
		if data.NetworkPointToMultipoint.ValueBool() {
			body, _ = sjson.Set(body, "network.point-to-multipoint", map[string]string{})
		}
	}
	if !data.Cost.IsNull() && !data.Cost.IsUnknown() {
		body, _ = sjson.Set(body, "cost", strconv.FormatInt(data.Cost.ValueInt64(), 10))
	}
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		body, _ = sjson.Set(body, "priority", strconv.FormatInt(data.Priority.ValueInt64(), 10))
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
	return body
}

func (data *RouterOSPFVRFAreaInterface) updateFromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "network.broadcast"); !data.NetworkBroadcast.IsNull() {
		if value.Exists() {
			data.NetworkBroadcast = types.BoolValue(true)
		} else {
			data.NetworkBroadcast = types.BoolValue(false)
		}
	} else {
		data.NetworkBroadcast = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "network.non-broadcast"); !data.NetworkNonBroadcast.IsNull() {
		if value.Exists() {
			data.NetworkNonBroadcast = types.BoolValue(true)
		} else {
			data.NetworkNonBroadcast = types.BoolValue(false)
		}
	} else {
		data.NetworkNonBroadcast = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "network.point-to-point"); !data.NetworkPointToPoint.IsNull() {
		if value.Exists() {
			data.NetworkPointToPoint = types.BoolValue(true)
		} else {
			data.NetworkPointToPoint = types.BoolValue(false)
		}
	} else {
		data.NetworkPointToPoint = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "network.point-to-multipoint"); !data.NetworkPointToMultipoint.IsNull() {
		if value.Exists() {
			data.NetworkPointToMultipoint = types.BoolValue(true)
		} else {
			data.NetworkPointToMultipoint = types.BoolValue(false)
		}
	} else {
		data.NetworkPointToMultipoint = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "cost"); value.Exists() && !data.Cost.IsNull() {
		data.Cost = types.Int64Value(value.Int())
	} else {
		data.Cost = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "priority"); value.Exists() && !data.Priority.IsNull() {
		data.Priority = types.Int64Value(value.Int())
	} else {
		data.Priority = types.Int64Null()
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
}

func (data *RouterOSPFVRFAreaInterface) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "network.broadcast"); value.Exists() {
		data.NetworkBroadcast = types.BoolValue(true)
	} else {
		data.NetworkBroadcast = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "network.non-broadcast"); value.Exists() {
		data.NetworkNonBroadcast = types.BoolValue(true)
	} else {
		data.NetworkNonBroadcast = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "network.point-to-point"); value.Exists() {
		data.NetworkPointToPoint = types.BoolValue(true)
	} else {
		data.NetworkPointToPoint = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "network.point-to-multipoint"); value.Exists() {
		data.NetworkPointToMultipoint = types.BoolValue(true)
	} else {
		data.NetworkPointToMultipoint = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "cost"); value.Exists() {
		data.Cost = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "priority"); value.Exists() {
		data.Priority = types.Int64Value(value.Int())
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
}

func (data *RouterOSPFVRFAreaInterface) fromPlan(ctx context.Context, plan RouterOSPFVRFAreaInterface) {
	data.Device = plan.Device
	data.ProcessName = types.StringValue(plan.ProcessName.ValueString())
	data.VrfName = types.StringValue(plan.VrfName.ValueString())
	data.AreaId = types.StringValue(plan.AreaId.ValueString())
	data.InterfaceName = types.StringValue(plan.InterfaceName.ValueString())
}

func (data *RouterOSPFVRFAreaInterface) getDeletedListItems(ctx context.Context, state RouterOSPFVRFAreaInterface) []string {
	deletedListItems := make([]string, 0)
	return deletedListItems
}

func (data *RouterOSPFVRFAreaInterface) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	return emptyLeafsDelete
}
