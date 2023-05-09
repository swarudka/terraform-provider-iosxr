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

type SegmentRoutingTrafficEngineering struct {
	Device         types.String                                     `tfsdk:"device"`
	Id             types.String                                     `tfsdk:"id"`
	OnDemandColors []SegmentRoutingTrafficEngineeringOnDemandColors `tfsdk:"on_demand_colors"`
	Policies       []SegmentRoutingTrafficEngineeringPolicies       `tfsdk:"policies"`
}
type SegmentRoutingTrafficEngineeringOnDemandColors struct {
	Color                                     types.Int64  `tfsdk:"color"`
	Srv6Enable                                types.Bool   `tfsdk:"srv6_enable"`
	Srv6LocatorLocatorName                    types.String `tfsdk:"srv6_locator_locator_name"`
	Srv6LocatorBehavior                       types.String `tfsdk:"srv6_locator_behavior"`
	Srv6LocatorBindingSidType                 types.String `tfsdk:"srv6_locator_binding_sid_type"`
	SourceAddressSourceAddress                types.String `tfsdk:"source_address_source_address"`
	SourceAddressIpAddressType                types.String `tfsdk:"source_address_ip_address_type"`
	EffectiveMetricEnable                     types.Bool   `tfsdk:"effective_metric_enable"`
	EffectiveMetricMetricValueTypeMetricValue types.Int64  `tfsdk:"effective_metric_metric_value_type_metric_value"`
	EffectiveMetricMetricValueTypeMetricType  types.String `tfsdk:"effective_metric_metric_value_type_metric_type"`
	ConstraintSegmentsProtectionType          types.String `tfsdk:"constraint_segments_protection_type"`
	ConstraintSegmentsSidAlgorithm            types.Int64  `tfsdk:"constraint_segments_sid_algorithm"`
}
type SegmentRoutingTrafficEngineeringPolicies struct {
	PolicyName                         types.String `tfsdk:"policy_name"`
	Srv6Enable                         types.Bool   `tfsdk:"srv6_enable"`
	Srv6LocatorLocatorName             types.String `tfsdk:"srv6_locator_locator_name"`
	Srv6LocatorBindingSidType          types.String `tfsdk:"srv6_locator_binding_sid_type"`
	Srv6LocatorBehavior                types.String `tfsdk:"srv6_locator_behavior"`
	SourceAddressSourceAddress         types.String `tfsdk:"source_address_source_address"`
	SourceAddressIpAddressType         types.String `tfsdk:"source_address_ip_address_type"`
	PolicyColorEndpointColor           types.Int64  `tfsdk:"policy_color_endpoint_color"`
	PolicyColorEndpointEndPointType    types.String `tfsdk:"policy_color_endpoint_end_point_type"`
	PolicyColorEndpointEndPointAddress types.String `tfsdk:"policy_color_endpoint_end_point_address"`
}

func (data SegmentRoutingTrafficEngineering) getPath() string {
	return "Cisco-IOS-XR-segment-routing-ms-cfg:/sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering"
}

func (data SegmentRoutingTrafficEngineering) toBody(ctx context.Context) string {
	body := "{}"
	if len(data.OnDemandColors) > 0 {
		body, _ = sjson.Set(body, "on-demand-colors.on-demand-color", []interface{}{})
		for index, item := range data.OnDemandColors {
			if !item.Color.IsNull() && !item.Color.IsUnknown() {
				body, _ = sjson.Set(body, "on-demand-colors.on-demand-color"+"."+strconv.Itoa(index)+"."+"color", strconv.FormatInt(item.Color.ValueInt64(), 10))
			}
			if !item.Srv6Enable.IsNull() && !item.Srv6Enable.IsUnknown() {
				if item.Srv6Enable.ValueBool() {
					body, _ = sjson.Set(body, "on-demand-colors.on-demand-color"+"."+strconv.Itoa(index)+"."+"srv6.enable", map[string]string{})
				}
			}
			if !item.Srv6LocatorLocatorName.IsNull() && !item.Srv6LocatorLocatorName.IsUnknown() {
				body, _ = sjson.Set(body, "on-demand-colors.on-demand-color"+"."+strconv.Itoa(index)+"."+"srv6.locator.locator-name", item.Srv6LocatorLocatorName.ValueString())
			}
			if !item.Srv6LocatorBehavior.IsNull() && !item.Srv6LocatorBehavior.IsUnknown() {
				body, _ = sjson.Set(body, "on-demand-colors.on-demand-color"+"."+strconv.Itoa(index)+"."+"srv6.locator.behavior", item.Srv6LocatorBehavior.ValueString())
			}
			if !item.Srv6LocatorBindingSidType.IsNull() && !item.Srv6LocatorBindingSidType.IsUnknown() {
				body, _ = sjson.Set(body, "on-demand-colors.on-demand-color"+"."+strconv.Itoa(index)+"."+"srv6.locator.binding-sid-type", item.Srv6LocatorBindingSidType.ValueString())
			}
			if !item.SourceAddressSourceAddress.IsNull() && !item.SourceAddressSourceAddress.IsUnknown() {
				body, _ = sjson.Set(body, "on-demand-colors.on-demand-color"+"."+strconv.Itoa(index)+"."+"source-address.source-address", item.SourceAddressSourceAddress.ValueString())
			}
			if !item.SourceAddressIpAddressType.IsNull() && !item.SourceAddressIpAddressType.IsUnknown() {
				body, _ = sjson.Set(body, "on-demand-colors.on-demand-color"+"."+strconv.Itoa(index)+"."+"source-address.ip-address-type", item.SourceAddressIpAddressType.ValueString())
			}
			if !item.EffectiveMetricEnable.IsNull() && !item.EffectiveMetricEnable.IsUnknown() {
				if item.EffectiveMetricEnable.ValueBool() {
					body, _ = sjson.Set(body, "on-demand-colors.on-demand-color"+"."+strconv.Itoa(index)+"."+"effective-metric.enable", map[string]string{})
				}
			}
			if !item.EffectiveMetricMetricValueTypeMetricValue.IsNull() && !item.EffectiveMetricMetricValueTypeMetricValue.IsUnknown() {
				body, _ = sjson.Set(body, "on-demand-colors.on-demand-color"+"."+strconv.Itoa(index)+"."+"effective-metric.metric-value-type.metric-value", strconv.FormatInt(item.EffectiveMetricMetricValueTypeMetricValue.ValueInt64(), 10))
			}
			if !item.EffectiveMetricMetricValueTypeMetricType.IsNull() && !item.EffectiveMetricMetricValueTypeMetricType.IsUnknown() {
				body, _ = sjson.Set(body, "on-demand-colors.on-demand-color"+"."+strconv.Itoa(index)+"."+"effective-metric.metric-value-type.metric-type", item.EffectiveMetricMetricValueTypeMetricType.ValueString())
			}
			if !item.ConstraintSegmentsProtectionType.IsNull() && !item.ConstraintSegmentsProtectionType.IsUnknown() {
				body, _ = sjson.Set(body, "on-demand-colors.on-demand-color"+"."+strconv.Itoa(index)+"."+"constraint.segments.protection-type", item.ConstraintSegmentsProtectionType.ValueString())
			}
			if !item.ConstraintSegmentsSidAlgorithm.IsNull() && !item.ConstraintSegmentsSidAlgorithm.IsUnknown() {
				body, _ = sjson.Set(body, "on-demand-colors.on-demand-color"+"."+strconv.Itoa(index)+"."+"constraint.segments.sid-algorithm", strconv.FormatInt(item.ConstraintSegmentsSidAlgorithm.ValueInt64(), 10))
			}
		}
	}
	if len(data.Policies) > 0 {
		body, _ = sjson.Set(body, "policies.policy", []interface{}{})
		for index, item := range data.Policies {
			if !item.PolicyName.IsNull() && !item.PolicyName.IsUnknown() {
				body, _ = sjson.Set(body, "policies.policy"+"."+strconv.Itoa(index)+"."+"policy-name", item.PolicyName.ValueString())
			}
			if !item.Srv6Enable.IsNull() && !item.Srv6Enable.IsUnknown() {
				if item.Srv6Enable.ValueBool() {
					body, _ = sjson.Set(body, "policies.policy"+"."+strconv.Itoa(index)+"."+"srv6.enable", map[string]string{})
				}
			}
			if !item.Srv6LocatorLocatorName.IsNull() && !item.Srv6LocatorLocatorName.IsUnknown() {
				body, _ = sjson.Set(body, "policies.policy"+"."+strconv.Itoa(index)+"."+"srv6.locator.locator-name", item.Srv6LocatorLocatorName.ValueString())
			}
			if !item.Srv6LocatorBindingSidType.IsNull() && !item.Srv6LocatorBindingSidType.IsUnknown() {
				body, _ = sjson.Set(body, "policies.policy"+"."+strconv.Itoa(index)+"."+"srv6.locator.binding-sid-type", item.Srv6LocatorBindingSidType.ValueString())
			}
			if !item.Srv6LocatorBehavior.IsNull() && !item.Srv6LocatorBehavior.IsUnknown() {
				body, _ = sjson.Set(body, "policies.policy"+"."+strconv.Itoa(index)+"."+"srv6.locator.behavior", item.Srv6LocatorBehavior.ValueString())
			}
			if !item.SourceAddressSourceAddress.IsNull() && !item.SourceAddressSourceAddress.IsUnknown() {
				body, _ = sjson.Set(body, "policies.policy"+"."+strconv.Itoa(index)+"."+"source-address.source-address", item.SourceAddressSourceAddress.ValueString())
			}
			if !item.SourceAddressIpAddressType.IsNull() && !item.SourceAddressIpAddressType.IsUnknown() {
				body, _ = sjson.Set(body, "policies.policy"+"."+strconv.Itoa(index)+"."+"source-address.ip-address-type", item.SourceAddressIpAddressType.ValueString())
			}
			if !item.PolicyColorEndpointColor.IsNull() && !item.PolicyColorEndpointColor.IsUnknown() {
				body, _ = sjson.Set(body, "policies.policy"+"."+strconv.Itoa(index)+"."+"policy-color-endpoint.color", strconv.FormatInt(item.PolicyColorEndpointColor.ValueInt64(), 10))
			}
			if !item.PolicyColorEndpointEndPointType.IsNull() && !item.PolicyColorEndpointEndPointType.IsUnknown() {
				body, _ = sjson.Set(body, "policies.policy"+"."+strconv.Itoa(index)+"."+"policy-color-endpoint.end-point-type", item.PolicyColorEndpointEndPointType.ValueString())
			}
			if !item.PolicyColorEndpointEndPointAddress.IsNull() && !item.PolicyColorEndpointEndPointAddress.IsUnknown() {
				body, _ = sjson.Set(body, "policies.policy"+"."+strconv.Itoa(index)+"."+"policy-color-endpoint.end-point-address", item.PolicyColorEndpointEndPointAddress.ValueString())
			}
		}
	}
	return body
}

func (data *SegmentRoutingTrafficEngineering) updateFromBody(ctx context.Context, res []byte) {
	for i := range data.OnDemandColors {
		keys := [...]string{"color"}
		keyValues := [...]string{strconv.FormatInt(data.OnDemandColors[i].Color.ValueInt64(), 10)}

		var r gjson.Result
		gjson.GetBytes(res, "on-demand-colors.on-demand-color").ForEach(
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
		if value := r.Get("color"); value.Exists() && !data.OnDemandColors[i].Color.IsNull() {
			data.OnDemandColors[i].Color = types.Int64Value(value.Int())
		} else {
			data.OnDemandColors[i].Color = types.Int64Null()
		}
		if value := r.Get("srv6.enable"); !data.OnDemandColors[i].Srv6Enable.IsNull() {
			if value.Exists() {
				data.OnDemandColors[i].Srv6Enable = types.BoolValue(true)
			} else {
				data.OnDemandColors[i].Srv6Enable = types.BoolValue(false)
			}
		} else {
			data.OnDemandColors[i].Srv6Enable = types.BoolNull()
		}
		if value := r.Get("srv6.locator.locator-name"); value.Exists() && !data.OnDemandColors[i].Srv6LocatorLocatorName.IsNull() {
			data.OnDemandColors[i].Srv6LocatorLocatorName = types.StringValue(value.String())
		} else {
			data.OnDemandColors[i].Srv6LocatorLocatorName = types.StringNull()
		}
		if value := r.Get("srv6.locator.behavior"); value.Exists() && !data.OnDemandColors[i].Srv6LocatorBehavior.IsNull() {
			data.OnDemandColors[i].Srv6LocatorBehavior = types.StringValue(value.String())
		} else {
			data.OnDemandColors[i].Srv6LocatorBehavior = types.StringNull()
		}
		if value := r.Get("srv6.locator.binding-sid-type"); value.Exists() && !data.OnDemandColors[i].Srv6LocatorBindingSidType.IsNull() {
			data.OnDemandColors[i].Srv6LocatorBindingSidType = types.StringValue(value.String())
		} else {
			data.OnDemandColors[i].Srv6LocatorBindingSidType = types.StringNull()
		}
		if value := r.Get("source-address.source-address"); value.Exists() && !data.OnDemandColors[i].SourceAddressSourceAddress.IsNull() {
			data.OnDemandColors[i].SourceAddressSourceAddress = types.StringValue(value.String())
		} else {
			data.OnDemandColors[i].SourceAddressSourceAddress = types.StringNull()
		}
		if value := r.Get("source-address.ip-address-type"); value.Exists() && !data.OnDemandColors[i].SourceAddressIpAddressType.IsNull() {
			data.OnDemandColors[i].SourceAddressIpAddressType = types.StringValue(value.String())
		} else {
			data.OnDemandColors[i].SourceAddressIpAddressType = types.StringNull()
		}
		if value := r.Get("effective-metric.enable"); !data.OnDemandColors[i].EffectiveMetricEnable.IsNull() {
			if value.Exists() {
				data.OnDemandColors[i].EffectiveMetricEnable = types.BoolValue(true)
			} else {
				data.OnDemandColors[i].EffectiveMetricEnable = types.BoolValue(false)
			}
		} else {
			data.OnDemandColors[i].EffectiveMetricEnable = types.BoolNull()
		}
		if value := r.Get("effective-metric.metric-value-type.metric-value"); value.Exists() && !data.OnDemandColors[i].EffectiveMetricMetricValueTypeMetricValue.IsNull() {
			data.OnDemandColors[i].EffectiveMetricMetricValueTypeMetricValue = types.Int64Value(value.Int())
		} else {
			data.OnDemandColors[i].EffectiveMetricMetricValueTypeMetricValue = types.Int64Null()
		}
		if value := r.Get("effective-metric.metric-value-type.metric-type"); value.Exists() && !data.OnDemandColors[i].EffectiveMetricMetricValueTypeMetricType.IsNull() {
			data.OnDemandColors[i].EffectiveMetricMetricValueTypeMetricType = types.StringValue(value.String())
		} else {
			data.OnDemandColors[i].EffectiveMetricMetricValueTypeMetricType = types.StringNull()
		}
		if value := r.Get("constraint.segments.protection-type"); value.Exists() && !data.OnDemandColors[i].ConstraintSegmentsProtectionType.IsNull() {
			data.OnDemandColors[i].ConstraintSegmentsProtectionType = types.StringValue(value.String())
		} else {
			data.OnDemandColors[i].ConstraintSegmentsProtectionType = types.StringNull()
		}
		if value := r.Get("constraint.segments.sid-algorithm"); value.Exists() && !data.OnDemandColors[i].ConstraintSegmentsSidAlgorithm.IsNull() {
			data.OnDemandColors[i].ConstraintSegmentsSidAlgorithm = types.Int64Value(value.Int())
		} else {
			data.OnDemandColors[i].ConstraintSegmentsSidAlgorithm = types.Int64Null()
		}
	}
	for i := range data.Policies {
		keys := [...]string{"policy-name"}
		keyValues := [...]string{data.Policies[i].PolicyName.ValueString()}

		var r gjson.Result
		gjson.GetBytes(res, "policies.policy").ForEach(
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
		if value := r.Get("policy-name"); value.Exists() && !data.Policies[i].PolicyName.IsNull() {
			data.Policies[i].PolicyName = types.StringValue(value.String())
		} else {
			data.Policies[i].PolicyName = types.StringNull()
		}
		if value := r.Get("srv6.enable"); !data.Policies[i].Srv6Enable.IsNull() {
			if value.Exists() {
				data.Policies[i].Srv6Enable = types.BoolValue(true)
			} else {
				data.Policies[i].Srv6Enable = types.BoolValue(false)
			}
		} else {
			data.Policies[i].Srv6Enable = types.BoolNull()
		}
		if value := r.Get("srv6.locator.locator-name"); value.Exists() && !data.Policies[i].Srv6LocatorLocatorName.IsNull() {
			data.Policies[i].Srv6LocatorLocatorName = types.StringValue(value.String())
		} else {
			data.Policies[i].Srv6LocatorLocatorName = types.StringNull()
		}
		if value := r.Get("srv6.locator.binding-sid-type"); value.Exists() && !data.Policies[i].Srv6LocatorBindingSidType.IsNull() {
			data.Policies[i].Srv6LocatorBindingSidType = types.StringValue(value.String())
		} else {
			data.Policies[i].Srv6LocatorBindingSidType = types.StringNull()
		}
		if value := r.Get("srv6.locator.behavior"); value.Exists() && !data.Policies[i].Srv6LocatorBehavior.IsNull() {
			data.Policies[i].Srv6LocatorBehavior = types.StringValue(value.String())
		} else {
			data.Policies[i].Srv6LocatorBehavior = types.StringNull()
		}
		if value := r.Get("source-address.source-address"); value.Exists() && !data.Policies[i].SourceAddressSourceAddress.IsNull() {
			data.Policies[i].SourceAddressSourceAddress = types.StringValue(value.String())
		} else {
			data.Policies[i].SourceAddressSourceAddress = types.StringNull()
		}
		if value := r.Get("source-address.ip-address-type"); value.Exists() && !data.Policies[i].SourceAddressIpAddressType.IsNull() {
			data.Policies[i].SourceAddressIpAddressType = types.StringValue(value.String())
		} else {
			data.Policies[i].SourceAddressIpAddressType = types.StringNull()
		}
		if value := r.Get("policy-color-endpoint.color"); value.Exists() && !data.Policies[i].PolicyColorEndpointColor.IsNull() {
			data.Policies[i].PolicyColorEndpointColor = types.Int64Value(value.Int())
		} else {
			data.Policies[i].PolicyColorEndpointColor = types.Int64Null()
		}
		if value := r.Get("policy-color-endpoint.end-point-type"); value.Exists() && !data.Policies[i].PolicyColorEndpointEndPointType.IsNull() {
			data.Policies[i].PolicyColorEndpointEndPointType = types.StringValue(value.String())
		} else {
			data.Policies[i].PolicyColorEndpointEndPointType = types.StringNull()
		}
		if value := r.Get("policy-color-endpoint.end-point-address"); value.Exists() && !data.Policies[i].PolicyColorEndpointEndPointAddress.IsNull() {
			data.Policies[i].PolicyColorEndpointEndPointAddress = types.StringValue(value.String())
		} else {
			data.Policies[i].PolicyColorEndpointEndPointAddress = types.StringNull()
		}
	}
}

func (data *SegmentRoutingTrafficEngineering) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "on-demand-colors.on-demand-color"); value.Exists() {
		data.OnDemandColors = make([]SegmentRoutingTrafficEngineeringOnDemandColors, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := SegmentRoutingTrafficEngineeringOnDemandColors{}
			if cValue := v.Get("color"); cValue.Exists() {
				item.Color = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("srv6.enable"); cValue.Exists() {
				item.Srv6Enable = types.BoolValue(true)
			} else {
				item.Srv6Enable = types.BoolValue(false)
			}
			if cValue := v.Get("srv6.locator.locator-name"); cValue.Exists() {
				item.Srv6LocatorLocatorName = types.StringValue(cValue.String())
			}
			if cValue := v.Get("srv6.locator.behavior"); cValue.Exists() {
				item.Srv6LocatorBehavior = types.StringValue(cValue.String())
			}
			if cValue := v.Get("srv6.locator.binding-sid-type"); cValue.Exists() {
				item.Srv6LocatorBindingSidType = types.StringValue(cValue.String())
			}
			if cValue := v.Get("source-address.source-address"); cValue.Exists() {
				item.SourceAddressSourceAddress = types.StringValue(cValue.String())
			}
			if cValue := v.Get("source-address.ip-address-type"); cValue.Exists() {
				item.SourceAddressIpAddressType = types.StringValue(cValue.String())
			}
			if cValue := v.Get("effective-metric.enable"); cValue.Exists() {
				item.EffectiveMetricEnable = types.BoolValue(true)
			} else {
				item.EffectiveMetricEnable = types.BoolValue(false)
			}
			if cValue := v.Get("effective-metric.metric-value-type.metric-value"); cValue.Exists() {
				item.EffectiveMetricMetricValueTypeMetricValue = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("effective-metric.metric-value-type.metric-type"); cValue.Exists() {
				item.EffectiveMetricMetricValueTypeMetricType = types.StringValue(cValue.String())
			}
			if cValue := v.Get("constraint.segments.protection-type"); cValue.Exists() {
				item.ConstraintSegmentsProtectionType = types.StringValue(cValue.String())
			}
			if cValue := v.Get("constraint.segments.sid-algorithm"); cValue.Exists() {
				item.ConstraintSegmentsSidAlgorithm = types.Int64Value(cValue.Int())
			}
			data.OnDemandColors = append(data.OnDemandColors, item)
			return true
		})
	}
	if value := gjson.GetBytes(res, "policies.policy"); value.Exists() {
		data.Policies = make([]SegmentRoutingTrafficEngineeringPolicies, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := SegmentRoutingTrafficEngineeringPolicies{}
			if cValue := v.Get("policy-name"); cValue.Exists() {
				item.PolicyName = types.StringValue(cValue.String())
			}
			if cValue := v.Get("srv6.enable"); cValue.Exists() {
				item.Srv6Enable = types.BoolValue(true)
			} else {
				item.Srv6Enable = types.BoolValue(false)
			}
			if cValue := v.Get("srv6.locator.locator-name"); cValue.Exists() {
				item.Srv6LocatorLocatorName = types.StringValue(cValue.String())
			}
			if cValue := v.Get("srv6.locator.binding-sid-type"); cValue.Exists() {
				item.Srv6LocatorBindingSidType = types.StringValue(cValue.String())
			}
			if cValue := v.Get("srv6.locator.behavior"); cValue.Exists() {
				item.Srv6LocatorBehavior = types.StringValue(cValue.String())
			}
			if cValue := v.Get("source-address.source-address"); cValue.Exists() {
				item.SourceAddressSourceAddress = types.StringValue(cValue.String())
			}
			if cValue := v.Get("source-address.ip-address-type"); cValue.Exists() {
				item.SourceAddressIpAddressType = types.StringValue(cValue.String())
			}
			if cValue := v.Get("policy-color-endpoint.color"); cValue.Exists() {
				item.PolicyColorEndpointColor = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("policy-color-endpoint.end-point-type"); cValue.Exists() {
				item.PolicyColorEndpointEndPointType = types.StringValue(cValue.String())
			}
			if cValue := v.Get("policy-color-endpoint.end-point-address"); cValue.Exists() {
				item.PolicyColorEndpointEndPointAddress = types.StringValue(cValue.String())
			}
			data.Policies = append(data.Policies, item)
			return true
		})
	}
}

func (data *SegmentRoutingTrafficEngineering) fromPlan(ctx context.Context, plan SegmentRoutingTrafficEngineering) {
	data.Device = plan.Device
}

func (data *SegmentRoutingTrafficEngineering) getDeletedListItems(ctx context.Context, state SegmentRoutingTrafficEngineering) []string {
	deletedListItems := make([]string, 0)
	for i := range state.OnDemandColors {
		keys := [...]string{"color"}
		stateKeyValues := [...]string{strconv.FormatInt(state.OnDemandColors[i].Color.ValueInt64(), 10)}

		emptyKeys := true
		if !reflect.ValueOf(state.OnDemandColors[i].Color.ValueInt64()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.OnDemandColors {
			found = true
			if state.OnDemandColors[i].Color.ValueInt64() != data.OnDemandColors[j].Color.ValueInt64() {
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
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/on-demand-colors/on-demand-color%v", state.getPath(), keyString))
		}
	}
	for i := range state.Policies {
		keys := [...]string{"policy-name"}
		stateKeyValues := [...]string{state.Policies[i].PolicyName.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.Policies[i].PolicyName.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.Policies {
			found = true
			if state.Policies[i].PolicyName.ValueString() != data.Policies[j].PolicyName.ValueString() {
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
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/policies/policy%v", state.getPath(), keyString))
		}
	}
	return deletedListItems
}

func (data *SegmentRoutingTrafficEngineering) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)

	for i := range data.OnDemandColors {
		keys := [...]string{"color"}
		keyValues := [...]string{strconv.FormatInt(data.OnDemandColors[i].Color.ValueInt64(), 10)}
		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
		if !data.OnDemandColors[i].Srv6Enable.IsNull() && !data.OnDemandColors[i].Srv6Enable.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/on-demand-colors/on-demand-color%v/srv6/enable", data.getPath(), keyString))
		}
		if !data.OnDemandColors[i].EffectiveMetricEnable.IsNull() && !data.OnDemandColors[i].EffectiveMetricEnable.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/on-demand-colors/on-demand-color%v/effective-metric/enable", data.getPath(), keyString))
		}
	}

	for i := range data.Policies {
		keys := [...]string{"policy-name"}
		keyValues := [...]string{data.Policies[i].PolicyName.ValueString()}
		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
		if !data.Policies[i].Srv6Enable.IsNull() && !data.Policies[i].Srv6Enable.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/policies/policy%v/srv6/enable", data.getPath(), keyString))
		}
	}
	return emptyLeafsDelete
}
