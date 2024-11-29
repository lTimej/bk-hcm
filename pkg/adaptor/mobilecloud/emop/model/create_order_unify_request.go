// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"hcm/pkg/adaptor/mobilecloud/emop/eboprsa/position"
	"hcm/pkg/criteria/validator"
)

type CreateOrderUnifyRequestOption CreateOrderUnifyRequest[CreateOrderUnifyRequestBody]

type CreateOrderUnifyRequest[T CreateOrderUnifyRequestBody] struct {
	CreateOrderUnifyRequestBody *T `json:"CreateOrderUnifyRequestBody,omitempty"`
}

type CreateOrderUnifyRequestBody struct {
	position.Body
	// 订单来源
	OrderSource *string `json:"orderSource,omitempty" `
	// 用户编码
	UserId *string `json:"userId,omitempty"`
	// 商品订购列表
	OfferList []OfferList `json:"offerList,omitempty"`
	// 用户编码
	ChannelInfoList []ChannelInfoList `json:"channelInfoList,omitempty" `
	// 扩展参数
	ExtInfo map[string]interface{} `json:"extInfo,omitempty"`
	// 操作方式
	OperateType *string `json:"operateType,omitempty"`
}

type OfferList struct {
	// 组内产品列表
	GroupOfferList []GroupOfferList `json:"groupOfferList,omitempty" `
	// 资源池编码
	CombOfferId *string `json:"combOfferId,omitempty"`
	// 规格组标识
	GroupQuantity *string `json:"groupQuantity,omitempty"`
}

type GroupOfferList struct {
	// 商品编码
	OfferId *string `json:"offerId,omitempty" `
	// 规格组标识
	ChaGroupId *string `json:"chaGroupId,omitempty"`
	// 资费相关信息
	PriceList []PriceList `json:"priceList,omitempty"`
	// 资源池编码
	PoolId *string `json:"poolId,omitempty" `
	// 可用区域编码
	ZoneId *string `json:"zoneId,omitempty"`
	// 是否开启自动续订
	AutoRenew *bool `json:"autoRenew,omitempty" `
	// 是否开启自动续订
	RenewDuration *string `json:"renewDuration,omitempty"`
	// 自动释放时间
	AutoEndTime *string `json:"autoEndTime,omitempty"`
	// 订购时长;非0：根据计费类型偏移为0：无到期时间
	Duration *string `json:"duration,omitempty"`
	// 产品订购数量
	Quantity *string `json:"quantity,omitempty"`
	// 产品属性
	ExtParams map[string]interface{} `json:"extParams,omitempty"`
}

type ChannelInfoList struct {
	//
	ChannelType *string `json:"channelType,omitempty"`
	//
	ChannelNumber *string `json:"channelNumber,omitempty" `
}

func (r *CreateOrderUnifyRequestBody) Validate() error {
	// TODO: 是否还需要添加其他规则校验呢？
	return validator.Validate.Struct(r)
}
