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

type L2VPN struct {
	Device         types.String          `tfsdk:"device"`
	Id             types.String          `tfsdk:"id"`
	DeleteMode     types.String          `tfsdk:"delete_mode"`
	Description    types.String          `tfsdk:"description"`
	RouterId       types.String          `tfsdk:"router_id"`
	XconnectGroups []L2VPNXconnectGroups `tfsdk:"xconnect_groups"`
}

type L2VPNData struct {
	Device         types.String          `tfsdk:"device"`
	Id             types.String          `tfsdk:"id"`
	Description    types.String          `tfsdk:"description"`
	RouterId       types.String          `tfsdk:"router_id"`
	XconnectGroups []L2VPNXconnectGroups `tfsdk:"xconnect_groups"`
}
type L2VPNXconnectGroups struct {
	GroupName types.String `tfsdk:"group_name"`
}

func (data L2VPN) getPath() string {
	return "Cisco-IOS-XR-um-l2vpn-cfg:/l2vpn"
}

func (data L2VPNData) getPath() string {
	return "Cisco-IOS-XR-um-l2vpn-cfg:/l2vpn"
}

func (data L2VPN) toBody(ctx context.Context) string {
	body := "{}"
	if !data.Description.IsNull() && !data.Description.IsUnknown() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if !data.RouterId.IsNull() && !data.RouterId.IsUnknown() {
		body, _ = sjson.Set(body, "router-id", data.RouterId.ValueString())
	}
	if len(data.XconnectGroups) > 0 {
		body, _ = sjson.Set(body, "xconnect.groups.group", []interface{}{})
		for index, item := range data.XconnectGroups {
			if !item.GroupName.IsNull() && !item.GroupName.IsUnknown() {
				body, _ = sjson.Set(body, "xconnect.groups.group"+"."+strconv.Itoa(index)+"."+"group-name", item.GroupName.ValueString())
			}
		}
	}
	return body
}

func (data *L2VPN) updateFromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "description"); value.Exists() && !data.Description.IsNull() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := gjson.GetBytes(res, "router-id"); value.Exists() && !data.RouterId.IsNull() {
		data.RouterId = types.StringValue(value.String())
	} else {
		data.RouterId = types.StringNull()
	}
	for i := range data.XconnectGroups {
		keys := [...]string{"group-name"}
		keyValues := [...]string{data.XconnectGroups[i].GroupName.ValueString()}

		var r gjson.Result
		gjson.GetBytes(res, "xconnect.groups.group").ForEach(
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
		if value := r.Get("group-name"); value.Exists() && !data.XconnectGroups[i].GroupName.IsNull() {
			data.XconnectGroups[i].GroupName = types.StringValue(value.String())
		} else {
			data.XconnectGroups[i].GroupName = types.StringNull()
		}
	}
}

func (data *L2VPNData) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "router-id"); value.Exists() {
		data.RouterId = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "xconnect.groups.group"); value.Exists() {
		data.XconnectGroups = make([]L2VPNXconnectGroups, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := L2VPNXconnectGroups{}
			if cValue := v.Get("group-name"); cValue.Exists() {
				item.GroupName = types.StringValue(cValue.String())
			}
			data.XconnectGroups = append(data.XconnectGroups, item)
			return true
		})
	}
}

func (data *L2VPN) getDeletedListItems(ctx context.Context, state L2VPN) []string {
	deletedListItems := make([]string, 0)
	for i := range state.XconnectGroups {
		keys := [...]string{"group-name"}
		stateKeyValues := [...]string{state.XconnectGroups[i].GroupName.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.XconnectGroups[i].GroupName.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.XconnectGroups {
			found = true
			if state.XconnectGroups[i].GroupName.ValueString() != data.XconnectGroups[j].GroupName.ValueString() {
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
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/xconnect/groups/group%v", state.getPath(), keyString))
		}
	}
	return deletedListItems
}

func (data *L2VPN) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	for i := range data.XconnectGroups {
		keys := [...]string{"group-name"}
		keyValues := [...]string{data.XconnectGroups[i].GroupName.ValueString()}
		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
	}
	return emptyLeafsDelete
}

func (data *L2VPN) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.Description.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/description", data.getPath()))
	}
	if !data.RouterId.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/router-id", data.getPath()))
	}
	for i := range data.XconnectGroups {
		keys := [...]string{"group-name"}
		keyValues := [...]string{data.XconnectGroups[i].GroupName.ValueString()}

		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
		deletePaths = append(deletePaths, fmt.Sprintf("%v/xconnect/groups/group%v", data.getPath(), keyString))
	}
	return deletePaths
}
