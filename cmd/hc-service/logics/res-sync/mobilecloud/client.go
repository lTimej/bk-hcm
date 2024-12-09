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
	"hcm/pkg/adaptor/mobilecloud"
	"hcm/pkg/adaptor/mobilecloud/emop/eboprsa/ecloudsdkcore/config"
	"hcm/pkg/adaptor/mobilecloud/emop/model"
	dataservice "hcm/pkg/client/data-service"
	"hcm/pkg/criteria/validator"
	"hcm/pkg/kit"

	imsmodel "gitlab.ecloud.com/ecloud/ecloudsdkims/model"
)

type SyncCvmOption struct {
}

// Validate ...
func (opt SyncCvmOption) Validate() error {
	return validator.Validate.Struct(opt)
}

// Interface support resource sync.
type Interface interface {
	CloudCli() *mobilecloud.MobileCloud
	Image(kt *kit.Kit, opt *imsmodel.ListServerPublicImageV2Query, runtimeConfig *config.RuntimeConfig) (*imsmodel.ListServerPublicImageV2Response, error)
	PoolInfo(kt *kit.Kit, opt *model.PoolInfoRequestBody) (*model.PoolInfoResponse[model.PoolInfoResponseBody], error)
	QryOffer(kt *kit.Kit, opt *model.QryOfferRequestBody) (*model.QryOfferResponse[model.QryOfferResponseBody], error)
	ResellerOrderOperateVerify(kt *kit.Kit, opt *model.ResellerOrderOperateVerifyRequestBody) (*model.ResellerOrderOperateVerifyResponse[model.ResellerOrderOperateVerifyResponseBody], error)
	CreateOrderUnify(kt *kit.Kit, opt *model.CreateOrderUnifyRequestBody) (*model.CreateOrderUnifyResponse[model.CreateOrderUnifyResponseBody], error)
}

var _ Interface = new(client)

// NewClient new client.
func NewClient(dbCli *dataservice.Client, cloudCli *mobilecloud.MobileCloud) Interface {
	return &client{
		dbCli:    dbCli,
		cloudCli: cloudCli,
	}
}

type client struct {
	accountID string
	cloudCli  *mobilecloud.MobileCloud
	dbCli     *dataservice.Client
}

// CloudCli ...
func (cli *client) CloudCli() *mobilecloud.MobileCloud {
	return cli.cloudCli
}
