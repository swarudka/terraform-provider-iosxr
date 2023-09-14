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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type ESISet struct {
	Device  types.String `tfsdk:"device"`
	Id      types.String `tfsdk:"id"`
	SetName types.String `tfsdk:"set_name"`
	Rpl     types.String `tfsdk:"rpl"`
}

type ESISetData struct {
	Device  types.String `tfsdk:"device"`
	Id      types.String `tfsdk:"id"`
	SetName types.String `tfsdk:"set_name"`
	Rpl     types.String `tfsdk:"rpl"`
}

func (data ESISet) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-route-policy-cfg:/routing-policy/sets/esi-sets/esi-set[set-name=%s]", data.SetName.ValueString())
}

func (data ESISetData) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-route-policy-cfg:/routing-policy/sets/esi-sets/esi-set[set-name=%s]", data.SetName.ValueString())
}

func (data ESISet) toBody(ctx context.Context) string {
	body := "{}"
	if !data.SetName.IsNull() && !data.SetName.IsUnknown() {
		body, _ = sjson.Set(body, "set-name", data.SetName.ValueString())
	}
	if !data.Rpl.IsNull() && !data.Rpl.IsUnknown() {
		body, _ = sjson.Set(body, "esi-set-as-text", data.Rpl.ValueString())
	}
	return body
}

func (data *ESISet) updateFromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "esi-set-as-text"); value.Exists() && !data.Rpl.IsNull() {
		data.Rpl = types.StringValue(value.String())
	} else {
		data.Rpl = types.StringNull()
	}
}

func (data *ESISetData) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "esi-set-as-text"); value.Exists() {
		data.Rpl = types.StringValue(value.String())
	}
}

func (data *ESISet) getDeletedItems(ctx context.Context, state ESISet) []string {
	deletedItems := make([]string, 0)
	return deletedItems
}

func (data *ESISet) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	return emptyLeafsDelete
}

func (data *ESISet) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.Rpl.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/esi-set-as-text", data.getPath()))
	}
	return deletePaths
}
