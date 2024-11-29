// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"hcm/pkg/adaptor/mobilecloud/emop/eboprsa/position"
	"hcm/pkg/criteria/validator"
)

type PoolInfoRequestOption PoolInfoRequest[PoolInfoRequestBody]

type PoolInfoRequest[T PoolInfoRequestBody] struct {
	PoolInfoRequestBody *T `json:"poolInfoRequestBody,omitempty"`
}

type PoolInfoRequestBody struct {
	position.Body
	// 产品类型;capebs、ssd
	ProductType *string `json:"productType,omitempty" `
	// 登陆账号平台用户编码; CIDC-U-cd8e01a325ea4eb0bcf0b83f973d5d9f
	UserId *string `json:"userId,omitempty"`
	// true，false，默认为false
	IgnoreInvalid *bool `json:"ignoreInvalid,omitempty"`
}

func (r *PoolInfoRequestBody) Validate() error {
	// TODO: 是否还需要添加其他规则校验呢？
	return validator.Validate.Struct(r)
}
