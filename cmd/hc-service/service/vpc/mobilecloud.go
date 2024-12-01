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

// Package vpc defines vpc service.
package vpc

import (
	"hcm/pkg/adaptor/mobilecloud/emop/eboprsa/ecloudsdkcore/config"
	"hcm/pkg/adaptor/types"
	"hcm/pkg/criteria/errf"
	"hcm/pkg/rest"

	"gitlab.ecloud.com/ecloud/ecloudsdkvpc/model"
)

// MobileCloudVpcCreate create mobile cloud vpc.
func (v vpc) MobileCloudVpcCreate(cts *rest.Contexts) (interface{}, error) {
	req := new(model.VpcOrderBody)
	if err := cts.DecodeInto(req); err != nil {
		return nil, errf.NewFromErr(errf.DecodeRequestFailed, err)
	}

	rheader := new(types.MobileCloudCredential)
	if err := cts.DecodeHeader(rheader); err != nil {
		return nil, errf.NewFromErr(errf.DecodeRequestFailed, err)
	}
	if err := rheader.Validate(); err != nil {
		return nil, errf.NewFromErr(errf.InvalidParameter, err)
	}
	runtimeConfig := &config.RuntimeConfig{
		RuntimeHeaderParams: map[string]string{
			"Pool-Id": cts.Request.Request.Header.Get("Pool-Id"),
		},
	}
	client, err := v.ad.Adaptor().MobileCloud(rheader)
	if err != nil {
		return nil, err
	}
	resp, err := client.CreateVpc(cts.Kit, req, runtimeConfig)
	return resp, err

}

// MobileCloudVpcCreate create mobile cloud vpc.
func (v vpc) MobileCloudVpcList(cts *rest.Contexts) (interface{}, error) {
	req := new(model.ListVpcQuery)
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
	runtimeConfig := &config.RuntimeConfig{
		RuntimeHeaderParams: map[string]string{
			"Pool-Id": cts.Request.Request.Header.Get("Pool-Id"),
		},
	}
	client, err := v.ad.Adaptor().MobileCloud(rheader)
	if err != nil {
		return nil, err
	}
	resp, err := client.ListVpc(cts.Kit, req, runtimeConfig)
	return resp, err

}

// MobileCloudListNetworkResps get mobile cloud vpc NetworkResps.
func (v vpc) MobileCloudListNetworkResps(cts *rest.Contexts) (interface{}, error) {
	req := new(model.ListNetworkRespQuery)
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
	runtimeConfig := &config.RuntimeConfig{
		RuntimeHeaderParams: map[string]string{
			"Pool-Id": cts.Request.Request.Header.Get("Pool-Id"),
		},
	}
	client, err := v.ad.Adaptor().MobileCloud(rheader)
	if err != nil {
		return nil, err
	}
	resp, err := client.ListNetworkResps(cts.Kit, req, runtimeConfig)
	return resp, err

}

// MobileCloudListPort get mobile cloud vpc port.
func (v vpc) MobileCloudListPort(cts *rest.Contexts) (interface{}, error) {
	req := new(model.ListPortQuery)
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
	runtimeConfig := &config.RuntimeConfig{
		RuntimeHeaderParams: map[string]string{
			"Pool-Id": cts.Request.Request.Header.Get("Pool-Id"),
		},
	}
	client, err := v.ad.Adaptor().MobileCloud(rheader)
	if err != nil {
		return nil, err
	}
	resp, err := client.ListPort(cts.Kit, req, runtimeConfig)
	return resp, err

}
