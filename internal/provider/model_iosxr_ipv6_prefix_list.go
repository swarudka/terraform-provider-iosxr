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

type IPv6PrefixList struct {
	Device         types.String              `tfsdk:"device"`
	Id             types.String              `tfsdk:"id"`
	PrefixListName types.String              `tfsdk:"prefix_list_name"`
	Sequences      []IPv6PrefixListSequences `tfsdk:"sequences"`
}

type IPv6PrefixListData struct {
	Device         types.String              `tfsdk:"device"`
	Id             types.String              `tfsdk:"id"`
	PrefixListName types.String              `tfsdk:"prefix_list_name"`
	Sequences      []IPv6PrefixListSequences `tfsdk:"sequences"`
}
type IPv6PrefixListSequences struct {
	SequenceNumber      types.Int64  `tfsdk:"sequence_number"`
	Remark              types.String `tfsdk:"remark"`
	Permission          types.String `tfsdk:"permission"`
	Prefix              types.String `tfsdk:"prefix"`
	Mask                types.Int64  `tfsdk:"mask"`
	MatchPrefixLengthEq types.Int64  `tfsdk:"match_prefix_length_eq"`
	MatchPrefixLengthGe types.Int64  `tfsdk:"match_prefix_length_ge"`
	MatchPrefixLengthLe types.Int64  `tfsdk:"match_prefix_length_le"`
}

func (data IPv6PrefixList) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-ipv6-prefix-list-cfg:/ipv6/prefix-lists/prefix-list[prefix-list-name=%s]", data.PrefixListName.ValueString())
}

func (data IPv6PrefixListData) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-ipv6-prefix-list-cfg:/ipv6/prefix-lists/prefix-list[prefix-list-name=%s]", data.PrefixListName.ValueString())
}

func (data IPv6PrefixList) toBody(ctx context.Context) string {
	body := "{}"
	if !data.PrefixListName.IsNull() && !data.PrefixListName.IsUnknown() {
		body, _ = sjson.Set(body, "prefix-list-name", data.PrefixListName.ValueString())
	}
	if len(data.Sequences) > 0 {
		body, _ = sjson.Set(body, "sequences.sequence", []interface{}{})
		for index, item := range data.Sequences {
			if !item.SequenceNumber.IsNull() && !item.SequenceNumber.IsUnknown() {
				body, _ = sjson.Set(body, "sequences.sequence"+"."+strconv.Itoa(index)+"."+"sequence-number", strconv.FormatInt(item.SequenceNumber.ValueInt64(), 10))
			}
			if !item.Remark.IsNull() && !item.Remark.IsUnknown() {
				body, _ = sjson.Set(body, "sequences.sequence"+"."+strconv.Itoa(index)+"."+"remark", item.Remark.ValueString())
			}
			if !item.Permission.IsNull() && !item.Permission.IsUnknown() {
				body, _ = sjson.Set(body, "sequences.sequence"+"."+strconv.Itoa(index)+"."+"permission", item.Permission.ValueString())
			}
			if !item.Prefix.IsNull() && !item.Prefix.IsUnknown() {
				body, _ = sjson.Set(body, "sequences.sequence"+"."+strconv.Itoa(index)+"."+"prefix", item.Prefix.ValueString())
			}
			if !item.Mask.IsNull() && !item.Mask.IsUnknown() {
				body, _ = sjson.Set(body, "sequences.sequence"+"."+strconv.Itoa(index)+"."+"mask", strconv.FormatInt(item.Mask.ValueInt64(), 10))
			}
			if !item.MatchPrefixLengthEq.IsNull() && !item.MatchPrefixLengthEq.IsUnknown() {
				body, _ = sjson.Set(body, "sequences.sequence"+"."+strconv.Itoa(index)+"."+"match-prefix-length.eq", strconv.FormatInt(item.MatchPrefixLengthEq.ValueInt64(), 10))
			}
			if !item.MatchPrefixLengthGe.IsNull() && !item.MatchPrefixLengthGe.IsUnknown() {
				body, _ = sjson.Set(body, "sequences.sequence"+"."+strconv.Itoa(index)+"."+"match-prefix-length.ge", strconv.FormatInt(item.MatchPrefixLengthGe.ValueInt64(), 10))
			}
			if !item.MatchPrefixLengthLe.IsNull() && !item.MatchPrefixLengthLe.IsUnknown() {
				body, _ = sjson.Set(body, "sequences.sequence"+"."+strconv.Itoa(index)+"."+"match-prefix-length.le", strconv.FormatInt(item.MatchPrefixLengthLe.ValueInt64(), 10))
			}
		}
	}
	return body
}

func (data *IPv6PrefixList) updateFromBody(ctx context.Context, res []byte) {
	for i := range data.Sequences {
		keys := [...]string{"sequence-number"}
		keyValues := [...]string{strconv.FormatInt(data.Sequences[i].SequenceNumber.ValueInt64(), 10)}

		var r gjson.Result
		gjson.GetBytes(res, "sequences.sequence").ForEach(
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
		if value := r.Get("sequence-number"); value.Exists() && !data.Sequences[i].SequenceNumber.IsNull() {
			data.Sequences[i].SequenceNumber = types.Int64Value(value.Int())
		} else {
			data.Sequences[i].SequenceNumber = types.Int64Null()
		}
		if value := r.Get("remark"); value.Exists() && !data.Sequences[i].Remark.IsNull() {
			data.Sequences[i].Remark = types.StringValue(value.String())
		} else {
			data.Sequences[i].Remark = types.StringNull()
		}
		if value := r.Get("permission"); value.Exists() && !data.Sequences[i].Permission.IsNull() {
			data.Sequences[i].Permission = types.StringValue(value.String())
		} else {
			data.Sequences[i].Permission = types.StringNull()
		}
		if value := r.Get("prefix"); value.Exists() && !data.Sequences[i].Prefix.IsNull() {
			data.Sequences[i].Prefix = types.StringValue(value.String())
		} else {
			data.Sequences[i].Prefix = types.StringNull()
		}
		if value := r.Get("mask"); value.Exists() && !data.Sequences[i].Mask.IsNull() {
			data.Sequences[i].Mask = types.Int64Value(value.Int())
		} else {
			data.Sequences[i].Mask = types.Int64Null()
		}
		if value := r.Get("match-prefix-length.eq"); value.Exists() && !data.Sequences[i].MatchPrefixLengthEq.IsNull() {
			data.Sequences[i].MatchPrefixLengthEq = types.Int64Value(value.Int())
		} else {
			data.Sequences[i].MatchPrefixLengthEq = types.Int64Null()
		}
		if value := r.Get("match-prefix-length.ge"); value.Exists() && !data.Sequences[i].MatchPrefixLengthGe.IsNull() {
			data.Sequences[i].MatchPrefixLengthGe = types.Int64Value(value.Int())
		} else {
			data.Sequences[i].MatchPrefixLengthGe = types.Int64Null()
		}
		if value := r.Get("match-prefix-length.le"); value.Exists() && !data.Sequences[i].MatchPrefixLengthLe.IsNull() {
			data.Sequences[i].MatchPrefixLengthLe = types.Int64Value(value.Int())
		} else {
			data.Sequences[i].MatchPrefixLengthLe = types.Int64Null()
		}
	}
}

func (data *IPv6PrefixListData) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "sequences.sequence"); value.Exists() {
		data.Sequences = make([]IPv6PrefixListSequences, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := IPv6PrefixListSequences{}
			if cValue := v.Get("sequence-number"); cValue.Exists() {
				item.SequenceNumber = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("remark"); cValue.Exists() {
				item.Remark = types.StringValue(cValue.String())
			}
			if cValue := v.Get("permission"); cValue.Exists() {
				item.Permission = types.StringValue(cValue.String())
			}
			if cValue := v.Get("prefix"); cValue.Exists() {
				item.Prefix = types.StringValue(cValue.String())
			}
			if cValue := v.Get("mask"); cValue.Exists() {
				item.Mask = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("match-prefix-length.eq"); cValue.Exists() {
				item.MatchPrefixLengthEq = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("match-prefix-length.ge"); cValue.Exists() {
				item.MatchPrefixLengthGe = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("match-prefix-length.le"); cValue.Exists() {
				item.MatchPrefixLengthLe = types.Int64Value(cValue.Int())
			}
			data.Sequences = append(data.Sequences, item)
			return true
		})
	}
}

func (data *IPv6PrefixList) getDeletedListItems(ctx context.Context, state IPv6PrefixList) []string {
	deletedListItems := make([]string, 0)
	for i := range state.Sequences {
		keys := [...]string{"sequence-number"}
		stateKeyValues := [...]string{strconv.FormatInt(state.Sequences[i].SequenceNumber.ValueInt64(), 10)}

		emptyKeys := true
		if !reflect.ValueOf(state.Sequences[i].SequenceNumber.ValueInt64()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.Sequences {
			found = true
			if state.Sequences[i].SequenceNumber.ValueInt64() != data.Sequences[j].SequenceNumber.ValueInt64() {
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
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/sequences/sequence%v", state.getPath(), keyString))
		}
	}
	return deletedListItems
}

func (data *IPv6PrefixList) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	for i := range data.Sequences {
		keys := [...]string{"sequence-number"}
		keyValues := [...]string{strconv.FormatInt(data.Sequences[i].SequenceNumber.ValueInt64(), 10)}
		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
	}
	return emptyLeafsDelete
}

func (data *IPv6PrefixList) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	for i := range data.Sequences {
		keys := [...]string{"sequence-number"}
		keyValues := [...]string{strconv.FormatInt(data.Sequences[i].SequenceNumber.ValueInt64(), 10)}

		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
		deletePaths = append(deletePaths, fmt.Sprintf("%v/sequences/sequence%v", data.getPath(), keyString))
	}
	return deletePaths
}
