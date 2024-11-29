// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

type CreateV2ResponseStateEnum string

// List of State
const (
	CreateV2ResponseStateEnumOk        CreateV2ResponseStateEnum = "OK"
	CreateV2ResponseStateEnumError     CreateV2ResponseStateEnum = "ERROR"
	CreateV2ResponseStateEnumException CreateV2ResponseStateEnum = "EXCEPTION"
	CreateV2ResponseStateEnumAlarm     CreateV2ResponseStateEnum = "ALARM"
	CreateV2ResponseStateEnumForbidden CreateV2ResponseStateEnum = "FORBIDDEN"
)

type CreateResellerUserResponseOption CreateResellerUserResponse[CreateResellerUserResponseBody]

type CreateResellerUserResponse[T CreateResellerUserResponseBody] struct {

	// 每个请求的序列号
	RespCode *string `json:"respCode,omitempty"`
	// 页面国际化错误提示
	RespDesc *string `json:"respDesc,omitempty"`
	// 统一错误码
	Result *T `json:"result,omitempty"`
}

type CreateResellerUserResponseBody struct {
	CustomerId *string `json:"customerId,omitempty"`
}
