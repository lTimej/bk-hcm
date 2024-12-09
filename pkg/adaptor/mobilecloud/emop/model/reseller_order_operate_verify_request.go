// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"hcm/pkg/adaptor/mobilecloud/emop/eboprsa/position"
	"hcm/pkg/criteria/validator"
)

type ResellerOrderOperateVerifyRequestOption ResellerOrderOperateVerifyRequest[ResellerOrderOperateVerifyRequestBody]

type ResellerOrderOperateVerifyRequest[T ResellerOrderOperateVerifyRequestBody] struct {
	ResellerOrderOperateVerifyRequestBody *T `json:"resellerOrderOperateVerifyRequestBody,omitempty"`
}

type ResellerOrderOperateVerifyRequestBody struct {
	position.Body
	// 用户编码
	UserId *string `json:"userId,omitempty" `
	// 操作类型;1-订购；2-退订;3-变更；4-续订
	OrderType *string `json:"orderType,omitempty"`
	// 产品集合
	Products []Products `json:"products,omitempty"`
}

type Products struct {
	// 商品编码
	OfferId *string `json:"offerId,omitempty" `
	// 资源池编码
	PoolId *string `json:"poolId,omitempty"`
	// 规格组标识
	GroupId *string `json:"groupId,omitempty"`
	// 产品类型
	ProductType *string `json:"productType,omitempty" `
	// 计费模式;1：包周期计费2：按需计费
	ChargingMode *int `json:"chargingMode,omitempty"`
	// 周期类型：年、月、日等。包周期计费需传;hour：小month：月year：年once：一次性
	PeriodType *string `json:"periodType,omitempty"`
	// 订购时长;非0：根据计费类型偏移（包周期）0：无到期时间（按量）
	Duration *int `json:"duration,omitempty" `
	// 产品订购数量
	Quantity *int `json:"quantity,omitempty"`
	// 可用区域编码
	ZoneId    *string `json:"zoneId,omitempty"`
	AutoRenew *string `json:"autoRenew,omitempty"`
	// 资源实例编码
	InstanceId *string `json:"useinstanceIdrId,omitempty" `
	// 变更后新资源实例
	NewInstanceId *string `json:"newInstanceId,omitempty"`
	// 变更后新时长
	DurationChange *int `json:"durationChange,omitempty"`
	// 额外eboss开通参数
	EbossParams map[string]interface{} `json:"ebossParams,omitempty"`
	// 产品属性
	ExtParams map[string]interface{} `json:"extParams,omitempty" `
	// 资费列表
	PriceList []PriceList `json:"priceList,omitempty"`
}

type PriceList struct {
	// 价格id
	PriceId *string `json:"priceId,omitempty"`
}

func (r *ResellerOrderOperateVerifyRequestBody) Validate() error {
	// TODO: 是否还需要添加其他规则校验呢？
	return validator.Validate.Struct(r)
}
