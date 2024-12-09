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

	"hcm/pkg/kit"
	"hcm/pkg/logs"

	"hcm/pkg/adaptor/mobilecloud/emop/model"
)

// OP或第三方通过此接口通过产品类型查询所支持的资源池信息
func (mc *MobileCloud) PoolInfo(kt *kit.Kit, opt *model.PoolInfoRequestBody) (*model.PoolInfoResponse[model.PoolInfoResponseBody], error) {
	client, err := mc.clientSet.emopClient()
	if err != nil {
		logs.Errorf("new iam client failed, err: %v, rid: %s", err, kt.Rid)
		return nil, err
	}
	resp, err := client.PoolInfo(&model.PoolInfoRequest[model.PoolInfoRequestBody]{
		PoolInfoRequestBody: opt,
	})

	if err != nil {
		logs.Errorf("ShowPermanentAccessKey failed, err: %v, rid: %s", err, kt.Rid)
		return nil, fmt.Errorf("ShowPermanentAccessKey failed, err: %v", err)
	}
	return resp, nil
}

// 查询产品信息
func (mc *MobileCloud) QryOffer(kt *kit.Kit, opt *model.QryOfferRequestBody) (*model.QryOfferResponse[model.QryOfferResponseBody], error) {
	client, err := mc.clientSet.emopClient()
	if err != nil {
		logs.Errorf("new iam client failed, err: %v, rid: %s", err, kt.Rid)
		return nil, err
	}
	resp, err := client.QryOffer(&model.QryOfferRequest[model.QryOfferRequestBody]{
		QryOfferRequestBody: opt,
	})

	if err != nil {
		logs.Errorf("ShowPermanentAccessKey failed, err: %v, rid: %s", err, kt.Rid)
		return nil, fmt.Errorf("ShowPermanentAccessKey failed, err: %v", err)
	}
	return resp, nil
}

// 经销商业务操作校验接口
func (mc *MobileCloud) ResellerOrderOperateVerify(kt *kit.Kit, opt *model.ResellerOrderOperateVerifyRequestBody) (*model.ResellerOrderOperateVerifyResponse[model.ResellerOrderOperateVerifyResponseBody], error) {
	client, err := mc.clientSet.emopClient()
	if err != nil {
		logs.Errorf("new iam client failed, err: %v, rid: %s", err, kt.Rid)
		return nil, err
	}
	resp, err := client.ResellerOrderOperateVerify(&model.ResellerOrderOperateVerifyRequest[model.ResellerOrderOperateVerifyRequestBody]{
		ResellerOrderOperateVerifyRequestBody: opt,
	})

	if err != nil {
		logs.Errorf("ShowPermanentAccessKey failed, err: %v, rid: %s", err, kt.Rid)
		return nil, fmt.Errorf("ShowPermanentAccessKey failed, err: %v", err)
	}
	return resp, nil
}

// 经销商业务操作校验接口
func (mc *MobileCloud) CreateOrderUnify(kt *kit.Kit, opt *model.CreateOrderUnifyRequestBody) (*model.CreateOrderUnifyResponse[model.CreateOrderUnifyResponseBody], error) {
	client, err := mc.clientSet.emopClient()
	if err != nil {
		logs.Errorf("new iam client failed, err: %v, rid: %s", err, kt.Rid)
		return nil, err
	}
	resp, err := client.CreateOrderUnify(&model.CreateOrderUnifyRequest[model.CreateOrderUnifyRequestBody]{
		CreateOrderUnifyRequestBody: opt,
	})

	if err != nil {
		logs.Errorf("ShowPermanentAccessKey failed, err: %v, rid: %s", err, kt.Rid)
		return nil, fmt.Errorf("ShowPermanentAccessKey failed, err: %v", err)
	}
	return resp, nil
}
