/*
 * TencentBlueKing is pleased to support the open source community by making
 * 蓝鲸智云 - 混合云管理平台 (BlueKing - Hybrid Cloud Management System) available.
 * Copyright (C) 2022 THL A29 Limited,
 * a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on
 * an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 *
 * We undertake not to change the open source license (MIT license) applicable
 *
 * to the current version of the project delivered to anyone in the future.
 */

package account

import (
	hsaccount "hcm/pkg/api/hc-service/account"
	proto "hcm/pkg/api/hc-service/account"
	"hcm/pkg/criteria/errf"
	"hcm/pkg/rest"
)

// MobileCloudCreateResellerUser get account with options
func (a *accountSvc) MobileCloudCreateResellerUser(cts *rest.Contexts) (interface{}, error) {
	req := new(proto.CreateResellerUserRequest)
	if err := cts.DecodeInto(req); err != nil {
		return nil, errf.NewFromErr(errf.DecodeRequestFailed, err)
	}
	if err := req.Validate(); err != nil {
		return nil, errf.NewFromErr(errf.InvalidParameter, err)
	}

	resp, err := a.client.HCService().MobileCloud.Account.MobileCloudCreateResellerUser(
		cts.Kit.Ctx,
		cts.Kit.Header(),
		&hsaccount.CreateResellerUserRequest{},
	)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
