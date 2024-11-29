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

	"hcm/pkg/adaptor/mobilecloud/emop/eboprsa/ecloudsdkcore/config"
	"hcm/pkg/kit"
	"hcm/pkg/logs"

	"gitlab.ecloud.com/ecloud/ecloudsdkvpc/model"
)

// CreateVpc create vpc.
func (mc *MobileCloud) CreateVpc(kt *kit.Kit, opt *model.VpcOrderBody, runtimeConfig *config.RuntimeConfig) (*model.VpcOrderResponse, error) {
	client, err := mc.clientSet.emopClient()
	if err != nil {
		logs.Errorf("new iam client failed, err: %v, rid: %s", err, kt.Rid)
		return nil, err
	}
	resp, err := client.CreateVpc(&model.VpcOrderRequest{
		VpcOrderBody: opt,
	}, runtimeConfig)
	if err != nil {
		logs.Errorf("ShowPermanentAccessKey failed, err: %v, rid: %s", err, kt.Rid)
		return nil, fmt.Errorf("ShowPermanentAccessKey failed, err: %v", err)

	}
	return resp, nil
}

// ListVpc list vpc.
func (mc *MobileCloud) ListVpc(kt *kit.Kit, opt *model.ListVpcQuery, runtimeConfig *config.RuntimeConfig) (*model.ListVpcResponse, error) {
	client, err := mc.clientSet.emopClient()
	if err != nil {
		logs.Errorf("new iam client failed, err: %v, rid: %s", err, kt.Rid)
		return nil, err
	}
	resp, err := client.ListVpc(&model.ListVpcRequest{
		ListVpcQuery: opt,
	}, runtimeConfig)
	if err != nil {
		logs.Errorf("ShowPermanentAccessKey failed, err: %v, rid: %s", err, kt.Rid)
		return nil, fmt.Errorf("ShowPermanentAccessKey failed, err: %v", err)

	}
	return resp, nil
}

// ListNetworkResps list vpc network.
func (mc *MobileCloud) ListNetworkResps(kt *kit.Kit, opt *model.ListNetworkRespQuery, runtimeConfig *config.RuntimeConfig) (*model.ListNetworkRespResponse, error) {
	client, err := mc.clientSet.emopClient()
	if err != nil {
		logs.Errorf("new iam client failed, err: %v, rid: %s", err, kt.Rid)
		return nil, err
	}
	resp, err := client.ListNetworkResps(&model.ListNetworkRespRequest{
		ListNetworkRespQuery: opt,
	}, runtimeConfig)
	if err != nil {
		logs.Errorf("ShowPermanentAccessKey failed, err: %v, rid: %s", err, kt.Rid)
		return nil, fmt.Errorf("ShowPermanentAccessKey failed, err: %v", err)

	}
	return resp, nil
}

// ListPort list vpc port.
func (mc *MobileCloud) ListPort(kt *kit.Kit, opt *model.ListPortQuery, runtimeConfig *config.RuntimeConfig) (*model.ListPortResponse, error) {
	client, err := mc.clientSet.emopClient()
	if err != nil {
		logs.Errorf("new iam client failed, err: %v, rid: %s", err, kt.Rid)
		return nil, err
	}
	resp, err := client.ListPort(&model.ListPortRequest{
		ListPortQuery: opt,
	}, runtimeConfig)
	if err != nil {
		logs.Errorf("ShowPermanentAccessKey failed, err: %v, rid: %s", err, kt.Rid)
		return nil, fmt.Errorf("ShowPermanentAccessKey failed, err: %v", err)

	}
	return resp, nil
}
