// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

type ResellerUserOperateResponseOption ResellerUserOperateResponse[ResellerUserOperateResponseBody]

type ResellerUserOperateResponse[T ResellerUserOperateResponseBody] struct {

	// 每个请求的序列号
	RespCode *string `json:"respCode,omitempty"`
	// 页面国际化错误提示
	RespDesc *string `json:"respDesc,omitempty"`
	// 统一错误码
	Result *T `json:"result,omitempty"`
}

type ResellerUserOperateResponseBody struct {
	// 操作任务新建结果
	Flag *string `json:"flag,omitempty"`
}
