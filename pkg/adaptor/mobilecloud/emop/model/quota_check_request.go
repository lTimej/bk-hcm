// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"hcm/pkg/adaptor/mobilecloud/emop/eboprsa/position"
	"hcm/pkg/criteria/validator"
)

type QuotaCheckRequestOption QuotaCheckRequest[QuotaCheckRequestBody]

type QuotaCheckRequest[T QuotaCheckRequestBody] struct {
	QuotaCheckRequestBody *T `json:"QuotaCheckRequestBody,omitempty"`
}

type QuotaCheckRequestBody struct {
	position.Body
	// 校验是否超出配额请求体
	QuotaCheckListReq []QuotaCheckListReq `json:"quotaCheckListReq,omitempty"`
	// 用户ID
	UserId *string `json:"userId,omitempty"`
}

type QuotaCheckListReq struct {
	// 资源池编号
	PoolId *string `json:"poolId,omitempty"`
	// 要增加的配额数量
	Quantity *int `json:"quantity,omitempty"`
	// 配额名称
	QuotaName *string `json:"quotaName,omitempty"`
	// 可用区编号
	RegionId *string `json:"regionId,omitempty"`
	// 资源配
	ResourceId *string `json:"resourceId,omitempty"`
}

func (r *QuotaCheckRequestBody) Validate() error {
	// TODO: 是否还需要添加其他规则校验呢？
	return validator.Validate.Struct(r)
}
