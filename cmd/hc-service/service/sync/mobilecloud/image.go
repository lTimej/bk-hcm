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
	"hcm/pkg/adaptor/mobilecloud/emop/eboprsa/ecloudsdkcore/config"
	"hcm/pkg/adaptor/types"
	"hcm/pkg/criteria/errf"
	"hcm/pkg/rest"

	imsmodel "gitlab.ecloud.com/ecloud/ecloudsdkims/model"
)

// Sync ...
func (svc *service) SyncListImage(cts *rest.Contexts) (interface{}, error) {
	req := new(imsmodel.ListServerPublicImageV2Query)
	if err := cts.DecodeQuery(req); err != nil {
		return nil, errf.NewFromErr(errf.DecodeRequestFailed, err)
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
	runtimeConfig := &config.RuntimeConfig{
		RuntimeHeaderParams: map[string]string{
			"Pool-Id":        cts.Request.Request.Header.Get("Pool-Id"),
			"request_id":     cts.Request.Request.Header.Get("request_id"),
			"manage_user_id": cts.Request.Request.Header.Get("manage_user_id"),
		},
	}
	return syncCli.Image(cts.Kit, req, runtimeConfig)
}
