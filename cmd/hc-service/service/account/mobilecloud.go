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
	"hcm/pkg/adaptor/mobilecloud/emop/eboprsa/ecloudsdkcore/config"
	"hcm/pkg/adaptor/mobilecloud/emop/model"
	"hcm/pkg/adaptor/types"
	"hcm/pkg/adaptor/types/account"
	proto "hcm/pkg/api/hc-service/account"
	"hcm/pkg/criteria/errf"
	"hcm/pkg/rest"
)

// MobileCloudCreateResellerUser 经销商终端客户注册订购基础口令
func (svc *service) MobileCloudCreateResellerUser(cts *rest.Contexts) (interface{}, error) {
	req := new(proto.CreateResellerUserRequest)
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
	client, err := svc.ad.Adaptor().MobileCloud(rheader)
	if err != nil {
		return nil, err
	}
	mobileCloudCreateUser := &account.MobileCloudCreateResellerUserReq{
		DistributionChannel:     req.DistributionChannel,
		DistributionChannelName: req.DistributionChannelName,
		Phone:                   req.Phone,
		IcType:                  req.IcType,
		IcNo:                    req.IcNo,
		Province:                req.Province,
		City:                    req.City,
		County:                  req.County,
		Address:                 req.Address,
	}
	resp, err := client.CreateResellerUser(cts.Kit, mobileCloudCreateUser)
	return resp, err

}

// QueryAkSk 用户ak/sk创建
func (svc *service) MobileCloudCreateAkSk(cts *rest.Contexts) (interface{}, error) {
	req := new(model.CreateAkSkRequestBody)
	hreq, err := cts.DecodeFormData(req)
	if err != nil {
		return nil, errf.NewFromErr(errf.DecodeRequestFailed, err)
	}
	runtimeConfig := &config.RuntimeConfig{
		HttpRequest: hreq,
	}
	rheader := new(types.MobileCloudCredential)
	if err := cts.DecodeHeader(rheader); err != nil {
		return nil, errf.NewFromErr(errf.DecodeRequestFailed, err)
	}
	if err := rheader.Validate(); err != nil {
		return nil, errf.NewFromErr(errf.InvalidParameter, err)
	}
	client, err := svc.ad.Adaptor().MobileCloud(rheader)
	if err != nil {
		return nil, err
	}
	resp, err := client.CreateAkSk(cts.Kit, req, runtimeConfig)
	return resp, err

}

// QueryAkSk 用户ak/sk查询
func (svc *service) MobileCloudQueryAkSk(cts *rest.Contexts) (interface{}, error) {
	req := new(model.QueryAkSkRequestBody)
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
	client, err := svc.ad.Adaptor().MobileCloud(rheader)
	if err != nil {
		return nil, err
	}
	resp, err := client.QueryAkSk(cts.Kit, req)
	return resp, err

}

// MobileCloudCreateResellerUser 经销商终端客户注册订购基础口令
func (svc *service) MobileCloudResellerUserOperate(cts *rest.Contexts) (interface{}, error) {
	req := new(model.ResellerUserOperateRequestBody)
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
	client, err := svc.ad.Adaptor().MobileCloud(rheader)
	if err != nil {
		return nil, err
	}
	resp, err := client.ResellerUserOperate(cts.Kit, req)
	return resp, err

}
