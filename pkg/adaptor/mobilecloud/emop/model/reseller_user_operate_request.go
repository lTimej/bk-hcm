// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"hcm/pkg/adaptor/mobilecloud/emop/eboprsa/position"
	"hcm/pkg/criteria/validator"
)

type ResellerUserOperateRequestOption ResellerUserOperateRequest[ResellerUserOperateRequestBody]

type ResellerUserOperateRequest[T ResellerUserOperateRequestBody] struct {
	ResellerUserOperateRequestBody *T `json:"resellerUserOperateRequestBody,omitempty"`
}

type ResellerUserOperateRequestBody struct {
	position.Body
	// 经销商终端客户的客户编码
	CustomerId *string `json:"customerId,omitempty"`
	// 操作类型；2:注销经销商终端客户基础口令 3:暂停经销商终端客户基础口令 4:恢复经销商终端客户基础口令
	OperateType *string `json:"operateType,omitempty"`
}

func (r *ResellerUserOperateRequestBody) Validate() error {
	// TODO: 是否还需要添加其他规则校验呢？
	return validator.Validate.Struct(r)
}
