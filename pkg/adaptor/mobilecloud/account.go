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
	"fmt"
	"hcm/pkg/adaptor/mobilecloud/emop/model"
	"hcm/pkg/adaptor/types/account"
	"hcm/pkg/kit"
	"hcm/pkg/logs"
)

// 经销商终端客户注册订购基础口令
func (mc *MobileCloud) CreateResellerUser(kt *kit.Kit, opt *account.MobileCloudCreateResellerUserReq) (*account.CreateResellerUserResponse[account.CreateResellerUserResponseBody], error) {
	client, err := mc.clientSet.emopClient()
	if err != nil {
		logs.Errorf("new iam client failed, err: %v, rid: %s", err, kt.Rid)
		return nil, err
	}
	resp, err := client.CreateResellerUser(&model.CreateResellerUserRequest[model.CreateResellerUserRequestBody]{
		CreateResellerUserRequestBody: &model.CreateResellerUserRequestBody{
			DistributionChannel:     opt.DistributionChannel,
			DistributionChannelName: opt.DistributionChannelName,
			Phone:                   opt.Phone,
			IcType:                  opt.IcType,
			IcNo:                    opt.IcNo,
			Province:                opt.Province,
			City:                    opt.City,
			County:                  opt.County,
			Address:                 opt.Address,
		},
	})

	if err != nil {
		logs.Errorf("ShowPermanentAccessKey failed, err: %v, rid: %s", err, kt.Rid)
		return nil, fmt.Errorf("ShowPermanentAccessKey failed, err: %v", err)

	}
	if resp.Result != nil {
		return &account.CreateResellerUserResponse[account.CreateResellerUserResponseBody]{
			RespCode: resp.RespCode,
			RespDesc: resp.RespDesc,
			Result: &account.CreateResellerUserResponseBody{
				CustomerId: resp.Result.CustomerId,
			},
		}, nil
	} else {
		return &account.CreateResellerUserResponse[account.CreateResellerUserResponseBody]{
			RespCode: resp.RespCode,
			RespDesc: resp.RespDesc,
		}, nil
	}

}

// 用户ak/sk创建
func (mc *MobileCloud) CreateAkSk(kt *kit.Kit, opt *model.QueryAkSkRequestBody) (*model.QueryAkSkResponse[model.QueryAkSkResponseBody], error) {
	client, err := mc.clientSet.emopClient()
	if err != nil {
		logs.Errorf("new iam client failed, err: %v, rid: %s", err, kt.Rid)
		return nil, err
	}
	resp, err := client.CreateAkSk(&model.QueryAkSkRequest[model.QueryAkSkRequestBody]{
		QueryAkSkRequestBody: &model.QueryAkSkRequestBody{
			UserId: opt.UserId,
		},
	})
	if err != nil {
		logs.Errorf("ShowPermanentAccessKey failed, err: %v, rid: %s", err, kt.Rid)
		return nil, fmt.Errorf("ShowPermanentAccessKey failed, err: %v", err)

	}
	return resp, nil

}

// 用户ak/sk查询
func (mc *MobileCloud) QueryAkSk(kt *kit.Kit, opt *model.QueryAkSkRequestBody) (*model.QueryAkSkResponse[model.QueryAkSkResponseBody], error) {
	client, err := mc.clientSet.emopClient()
	if err != nil {
		logs.Errorf("new iam client failed, err: %v, rid: %s", err, kt.Rid)
		return nil, err
	}
	resp, err := client.QueryAkSk(&model.QueryAkSkRequest[model.QueryAkSkRequestBody]{
		QueryAkSkRequestBody: &model.QueryAkSkRequestBody{
			UserId: opt.UserId,
		},
	})
	if err != nil {
		logs.Errorf("ShowPermanentAccessKey failed, err: %v, rid: %s", err, kt.Rid)
		return nil, fmt.Errorf("ShowPermanentAccessKey failed, err: %v", err)

	}
	return resp, nil

}

// 经销商终端客户操作（暂停、恢复、注销）
func (mc *MobileCloud) ResellerUserOperate(kt *kit.Kit, opt *model.ResellerUserOperateRequestBody) (*model.ResellerUserOperateResponse[model.ResellerUserOperateResponseBody], error) {
	client, err := mc.clientSet.emopClient()
	if err != nil {
		logs.Errorf("new iam client failed, err: %v, rid: %s", err, kt.Rid)
		return nil, err
	}
	resp, err := client.ResellerUserOperate(&model.ResellerUserOperateRequest[model.ResellerUserOperateRequestBody]{
		ResellerUserOperateRequestBody: opt,
	})

	if err != nil {
		logs.Errorf("ShowPermanentAccessKey failed, err: %v, rid: %s", err, kt.Rid)
		return nil, fmt.Errorf("ShowPermanentAccessKey failed, err: %v", err)
	}
	return resp, nil
}
