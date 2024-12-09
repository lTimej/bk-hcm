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
	"hcm/pkg/adaptor/mobilecloud/emop"
	"hcm/pkg/adaptor/mobilecloud/emop/model"
	"hcm/pkg/adaptor/types"
	"hcm/pkg/criteria/errf"
	"hcm/pkg/rest"
)

// SyncPoolInfo ....
func (svc *service) SyncPoolInfo(cts *rest.Contexts) (interface{}, error) {
	req := new(model.PoolInfoRequestBody)
	if err := cts.DecodeInto(req); err != nil {
		return nil, errf.NewFromErr(errf.DecodeRequestFailed, err)
	}

	if err := req.Validate(); err != nil {
		return nil, errf.NewFromErr(errf.InvalidParameter, err)
	}
	rheader := new(types.MobileCloudCredential)
	if err := cts.DecodeHeader(rheader); err != nil {
		return nil, errf.NewFromErr(errf.DecodeRequestFailed, err)
	}
	if err := rheader.Validate(); err != nil {
		return nil, errf.NewFromErr(errf.InvalidParameter, err)
	}
	emop.MobileCloudSecret.Store(rheader.AppID, rheader)
	syncCli, err := svc.syncCli.MobileCloud(cts.Kit, rheader.AppID)
	if err != nil {
		return nil, err
	}
	return syncCli.PoolInfo(cts.Kit, req)
}

// QryOffer ....
func (svc *service) QryOffer(cts *rest.Contexts) (interface{}, error) {
	req := new(model.QryOfferRequestBody)
	if err := cts.DecodeInto(req); err != nil {
		return nil, errf.NewFromErr(errf.DecodeRequestFailed, err)
	}

	if err := req.Validate(); err != nil {
		return nil, errf.NewFromErr(errf.InvalidParameter, err)
	}
	rheader := new(types.MobileCloudCredential)
	if err := cts.DecodeHeader(rheader); err != nil {
		return nil, errf.NewFromErr(errf.DecodeRequestFailed, err)
	}
	if err := rheader.Validate(); err != nil {
		return nil, errf.NewFromErr(errf.InvalidParameter, err)
	}
	emop.MobileCloudSecret.Store(rheader.AppID, rheader)
	syncCli, err := svc.syncCli.MobileCloud(cts.Kit, rheader.AppID)
	if err != nil {
		return nil, err
	}
	return syncCli.QryOffer(cts.Kit, req)
}

// ResellerOrderOperateVerify ....
func (svc *service) ResellerOrderOperateVerify(cts *rest.Contexts) (interface{}, error) {
	req := new(model.ResellerOrderOperateVerifyRequestBody)
	if err := cts.DecodeInto(req); err != nil {
		return nil, errf.NewFromErr(errf.DecodeRequestFailed, err)
	}

	if err := req.Validate(); err != nil {
		return nil, errf.NewFromErr(errf.InvalidParameter, err)
	}
	rheader := new(types.MobileCloudCredential)
	if err := cts.DecodeHeader(rheader); err != nil {
		return nil, errf.NewFromErr(errf.DecodeRequestFailed, err)
	}
	if err := rheader.Validate(); err != nil {
		return nil, errf.NewFromErr(errf.InvalidParameter, err)
	}
	emop.MobileCloudSecret.Store(rheader.AppID, rheader)
	syncCli, err := svc.syncCli.MobileCloud(cts.Kit, rheader.AppID)
	if err != nil {
		return nil, err
	}
	return syncCli.ResellerOrderOperateVerify(cts.Kit, req)
}

// ResellerOrderOperateVerify ....
func (svc *service) CreateOrderUnify(cts *rest.Contexts) (interface{}, error) {
	req := new(model.CreateOrderUnifyRequestBody)
	if err := cts.DecodeInto(req); err != nil {
		return nil, errf.NewFromErr(errf.DecodeRequestFailed, err)
	}

	if err := req.Validate(); err != nil {
		return nil, errf.NewFromErr(errf.InvalidParameter, err)
	}
	rheader := new(types.MobileCloudCredential)
	if err := cts.DecodeHeader(rheader); err != nil {
		return nil, errf.NewFromErr(errf.DecodeRequestFailed, err)
	}
	if err := rheader.Validate(); err != nil {
		return nil, errf.NewFromErr(errf.InvalidParameter, err)
	}
	emop.MobileCloudSecret.Store(rheader.AppID, rheader)
	syncCli, err := svc.syncCli.MobileCloud(cts.Kit, rheader.AppID)
	if err != nil {
		return nil, err
	}
	return syncCli.CreateOrderUnify(cts.Kit, req)
}
