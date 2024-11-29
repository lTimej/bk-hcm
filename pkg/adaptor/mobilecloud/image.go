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

	imsmodel "gitlab.ecloud.com/ecloud/ecloudsdkims/model"
)

// ListImage 查询公共镜像列表
func (mc *MobileCloud) ListPublicImages(kt *kit.Kit, opt *imsmodel.ListServerPublicImageV2Query, runtimeConfig *config.RuntimeConfig) (*imsmodel.ListServerPublicImageV2Response, error) {

	client, err := mc.clientSet.emopClient()
	if err != nil {
		logs.Errorf("new iam client failed, err: %v, rid: %s", err, kt.Rid)
		return nil, err
	}
	resp, err := client.ListPublicImages(&imsmodel.ListServerPublicImageV2Request{
		ListServerPublicImageV2Query: opt,
	}, runtimeConfig)
	if err != nil {
		logs.Errorf("ShowPermanentAccessKey failed, err: %v, rid: %s", err, kt.Rid)
		return nil, fmt.Errorf("ShowPermanentAccessKey failed, err: %v", err)

	}
	return resp, nil
}
