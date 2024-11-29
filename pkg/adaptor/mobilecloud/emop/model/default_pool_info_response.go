// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

type DefaultPoolInfoResponseOption DefaultPoolInfoResponse[DefaultPoolInfoResponseBody]

type DefaultPoolInfoResponse[T DefaultPoolInfoResponseBody] struct {

	// 每个请求的序列号
	RespCode *string `json:"respCode,omitempty"`
	// 页面国际化错误提示
	RespDesc *string `json:"respDesc,omitempty"`
	// 统一错误码
	Result *T `json:"result,omitempty"`
}

type DefaultPoolList struct {
	PoolList
	// 是否推荐资源池;1是0否,默认0
	IsRecommend *string `json:"isRecommend,omitempty"`
}

type DefaultPoolInfoResponseBody struct {
	// 返回结果集
	PoolList []DefaultPoolList `json:"poolList,omitempty"`
}
