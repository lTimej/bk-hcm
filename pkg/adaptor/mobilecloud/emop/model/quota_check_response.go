// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

type QuotaCheckResponse struct {
	// 请求成功时返回的数据
	Body map[string]interface{} `json:"body,omitempty"`
	// 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
	// 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
	// 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
	// 返回状态码
	State *string `json:"state,omitempty"`
}
