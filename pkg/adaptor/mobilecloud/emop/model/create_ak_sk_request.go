// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"hcm/pkg/adaptor/mobilecloud/emop/eboprsa/position"
	"hcm/pkg/criteria/validator"
)

type CreateAkSkRequestOption CreateAkSkRequest[CreateAkSkRequestBody]

type CreateAkSkRequest[T CreateAkSkRequestBody] struct {
	CreateAkSkRequestBody *T `json:"CreateAkSkRequestBody,omitempty"`
}

type CreateAkSkRequestBody struct {
	position.Body
	// 登陆账号平台用户编码; CIDC-U-cd8e01a325ea4eb0bcf0b83f973d5d9f
	UserId *string `json:"userId,omitempty"`
}

func (r *CreateAkSkRequestBody) Validate() error {
	// TODO: 是否还需要添加其他规则校验呢？
	return validator.Validate.Struct(r)
}

func (r *CreateAkSkRequestBody) ToString(userId *string) string {
	return *userId
}
