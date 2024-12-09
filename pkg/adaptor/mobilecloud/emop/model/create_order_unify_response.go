// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

type CreateOrderUnifyResponseOption CreateOrderUnifyResponse[CreateOrderUnifyResponseBody]

type CreateOrderUnifyResponse[T CreateOrderUnifyResponseBody] struct {

	// 每个请求的序列号
	RespCode *string `json:"respCode,omitempty"`
	// 页面国际化错误提示
	RespDesc *string `json:"respDesc,omitempty"`
	// 统一错误码
	Result *T `json:"result,omitempty"`
}

type ProductList struct {
	// 校验结果码
	ProductId *string `json:"productId,omitempty"`
	// 产品到期时间
	EndDate *string `json:"endDate,omitempty"`
	// 产品费用（元)
	ProductFee *string `json:"productFee,omitempty"`
	// 订单项编码
	TradeId *string `json:"tradeId,omitempty"`
	// 产品生效时间
	StartDate *string `json:"startDate,omitempty"`
}
type UnfiyBody struct {
	// 订单ID
	OrderId *string `json:"orderId,omitempty"`
	// 订单费用（元）
	OrderFee *string `json:"orderFee,omitempty"`
	// 产品列表
	ProductList []ProductList `json:"productList,omitempty"`
}

type CreateOrderUnifyResponseBody struct {
	// 返回结果集
	Body UnfiyBody `json:"body,omitempty"`
}
