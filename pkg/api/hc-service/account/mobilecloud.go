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

package hsaccount

import "hcm/pkg/criteria/validator"

// MobileCloud CreateResellerUserRequest...
type CreateResellerUserRequest struct {
	// 经销商平台编码,测试账号申请后会提供
	DistributionChannel *string `json:"distributionChannel,omitempty"`
	// 经销商平台名称缩写,缩写由云能信息系统部对接人统一提供 例如：PBS
	DistributionChannelName *string `json:"distributionChannelName,omitempty"`
	// 手机号
	Phone *string `json:"phone,omitempty"`
	// 证件类型
	IcType *string `json:"icType,omitempty"`
	// 证件号码
	IcNo *string `json:"icNo,omitempty"`
	// 所在省
	Province *string `json:"province,omitempty"`
	// 所在市
	City *string `json:"city,omitempty"`
	// 所在区
	County *string `json:"county,omitempty"`
	// 详细地址信息
	Address *string `json:"address,omitempty"`
}

func (r *CreateResellerUserRequest) Validate() error {
	// TODO: 是否还需要添加其他规则校验呢？
	return validator.Validate.Struct(r)
}
