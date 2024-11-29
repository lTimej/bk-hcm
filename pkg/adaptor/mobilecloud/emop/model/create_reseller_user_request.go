// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import "hcm/pkg/adaptor/mobilecloud/emop/eboprsa/position"

type CreateResellerUserRequestOption CreateResellerUserRequest[CreateResellerUserRequestBody]

type CreateResellerUserRequest[T CreateResellerUserRequestBody] struct {
	CreateResellerUserRequestBody *T `json:"createResellerUserRequestBody,omitempty"`
}

type CreateResellerUserRequestBody struct {
	position.Body
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
