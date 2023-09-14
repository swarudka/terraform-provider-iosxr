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

type CommunitySet struct {
	Device  types.String `tfsdk:"device"`
	Id      types.String `tfsdk:"id"`
	SetName types.String `tfsdk:"set_name"`
	Rpl     types.String `tfsdk:"rpl"`
}

type CommunitySetData struct {
	Device  types.String `tfsdk:"device"`
	Id      types.String `tfsdk:"id"`
	SetName types.String `tfsdk:"set_name"`
	Rpl     types.String `tfsdk:"rpl"`
}

func (data CommunitySet) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-route-policy-cfg:/routing-policy/sets/community-sets/community-set[set-name=%s]", data.SetName.ValueString())
}

func (data CommunitySetData) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-route-policy-cfg:/routing-policy/sets/community-sets/community-set[set-name=%s]", data.SetName.ValueString())
}

func (data CommunitySet) toBody(ctx context.Context) string {
	body := "{}"
	if !data.SetName.IsNull() && !data.SetName.IsUnknown() {
		body, _ = sjson.Set(body, "set-name", data.SetName.ValueString())
	}
	if !data.Rpl.IsNull() && !data.Rpl.IsUnknown() {
		body, _ = sjson.Set(body, "rpl-community-set", data.Rpl.ValueString())
	}
	return body
}

func (data *CommunitySet) updateFromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "rpl-community-set"); value.Exists() && !data.Rpl.IsNull() {
		data.Rpl = types.StringValue(value.String())
	} else {
		data.Rpl = types.StringNull()
	}
}

func (data *CommunitySetData) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "rpl-community-set"); value.Exists() {
		data.Rpl = types.StringValue(value.String())
	}
}

func (data *CommunitySet) getDeletedItems(ctx context.Context, state CommunitySet) []string {
	deletedItems := make([]string, 0)
	return deletedItems
}

func (data *CommunitySet) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	return emptyLeafsDelete
}

func (data *CommunitySet) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.Rpl.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/rpl-community-set", data.getPath()))
	}
	return deletePaths
}
