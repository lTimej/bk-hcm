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

package mobilecloud

import (
	"hcm/pkg/adaptor/mobilecloud/emop/model"
	"hcm/pkg/criteria/errf"
	"hcm/pkg/kit"
)

// PoolInfo ...
func (cli *client) PoolInfo(kt *kit.Kit, opt *model.PoolInfoRequestBody) (*model.PoolInfoResponse[model.PoolInfoResponseBody], error) {
	if err := opt.Validate(); err != nil {
		return nil, errf.NewFromErr(errf.InvalidParameter, err)
	}
	resp, err := cli.cloudCli.PoolInfo(kt, opt)
	return resp, err
}

// QryOffer ...
func (cli *client) QryOffer(kt *kit.Kit, opt *model.QryOfferRequestBody) (*model.QryOfferResponse[model.QryOfferResponseBody], error) {
	if err := opt.Validate(); err != nil {
		return nil, errf.NewFromErr(errf.InvalidParameter, err)
	}
	resp, err := cli.cloudCli.QryOffer(kt, opt)
	return resp, err
}

// ResellerOrderOperateVerify ...
func (cli *client) ResellerOrderOperateVerify(kt *kit.Kit, opt *model.ResellerOrderOperateVerifyRequestBody) (*model.ResellerOrderOperateVerifyResponse[model.ResellerOrderOperateVerifyResponseBody], error) {
	if err := opt.Validate(); err != nil {
		return nil, errf.NewFromErr(errf.InvalidParameter, err)
	}
	resp, err := cli.cloudCli.ResellerOrderOperateVerify(kt, opt)
	return resp, err
}
