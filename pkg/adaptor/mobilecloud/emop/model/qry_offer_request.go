// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"hcm/pkg/adaptor/mobilecloud/emop/eboprsa/position"
	"hcm/pkg/criteria/validator"
)

type QryOfferRequestOption QryOfferRequest[QryOfferRequestBody]

type QryOfferRequest[T QryOfferRequestBody] struct {
	QryOfferRequestBody *T `json:"qryOfferRequestBody,omitempty"`
}

type QryOfferRequestBody struct {
	position.Body
	// 产品类型;capebs、ssd
	ProductType *string `json:"productType,omitempty" `
	// 商品编码，一串UUID，联系移动云获取
	OfferId *string `json:"offerId,omitempty"`
	// 产品规格编码
	ChaId *string `json:"chaId,omitempty"`
	// 资源池编码
	PoolId *string `json:"poolId,omitempty"`
	// 计费方式;1 包周期；2 按量
	FeeType *string `json:"feeType,omitempty"`
	// 计费单位;month,year,ci,gb,mb
	FeeUnit *string `feeUnit:"chaId,omitempty"`
	// 付费方式;0 预付（互联网客户）；1 后付（政企客户）
	ChargeType *string `json:"chargeType,omitempty"`
	// 属性集合;{ "cpu" : "1" }
	Attrs *map[string]interface{} `json:"attrs,omitempty"`
	// 是否包含失效规格;枚举：0：包含失效规格，1：不包含，不传默认为1当传0时chaId字段为必传
	ValidStatus *string `json:"validStatus,omitempty"`
	// 登陆账号OP分配的userId;接口非空，经销商必传，进行可用资源池过滤避免下单后无法使用
	UserId *string `json:"userId,omitempty"`
}

func (r *QryOfferRequestBody) Validate() error {
	// TODO: 是否还需要添加其他规则校验呢？
	return validator.Validate.Struct(r)
}
