// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

type CreateAkSkResponseOption CreateAkSkResponse[CreateAkSkResponseBody]

type CreateAkSkResponse[T CreateAkSkResponseBody] struct {

	// 状态，OK-成功 ERROR-失败，对应有错误码和错误描述
	State *string `json:"state,omitempty"`
	// 每次请求的唯一标识ID
	RequestId *string `json:"requestId,omitempty"`
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty"`
	// 错误描述
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// 响应体具体内容
	Body *T `json:"body,omitempty"`
}

type CreateAkSkResponseBody struct {
	// 公钥
	Accesskey *string `json:"accesskey,omitempty"`
	// 私钥
	Secretkey *string `json:"secretKey,omitempty"`
	// 用户id
	UserId *string `json:"userId,omitempty"`
}
