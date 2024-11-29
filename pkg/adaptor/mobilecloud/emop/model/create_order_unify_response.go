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

type UnfiyBody struct {
	// 校验结果码
	BizCode *string `json:"bizCode,omitempty"`
	// 校验结果描述
	BizDesc *string `json:"bizDesc,omitempty"`
}

type CreateOrderUnifyResponseBody struct {
	// 返回结果集
	Body UnfiyBody `json:"body,omitempty"`
}
