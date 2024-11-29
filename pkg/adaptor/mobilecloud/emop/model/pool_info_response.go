// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

type PoolInfoResponseOption PoolInfoResponse[PoolInfoResponseBody]

type PoolInfoResponse[T PoolInfoResponseBody] struct {

	// 每个请求的序列号
	RespCode *string `json:"respCode,omitempty"`
	// 页面国际化错误提示
	RespDesc *string `json:"respDesc,omitempty"`
	// 统一错误码
	Result *T `json:"result,omitempty"`
}

type PoolList struct {
	// 资源池id;如：CIDC-RP-25
	PoolId *string `json:"poolId,omitempty"`
	// 资源池区域名称
	PoolArea *string `json:"poolArea,omitempty"`
	// 产品类型;capebs、ssd
	ProductType *string `json:"productType,omitempty"`
	// 资源池名称;如：华东节点4
	PoolName *string `json:"poolName,omitempty"`
	// 资源池状态;1-生效；0-失效
	Status *string `json:"status,omitempty"`
	// 是否推荐资源池;1是0否,默认0
	IsRecommend *string `json:"isRecommend,omitempty"`
	// 当前资源池下所有可用区信息
	ZoneInfo []ZoneInfo `json:"zoneInfo,omitempty"`
}

type ZoneInfo struct {
	// 可用区ID
	ZoneId *string `json:"zoneId,omitempty"`
	// 可用区名称
	ZoneName *string `json:"zoneName,omitempty"`
	// 可用区编码
	ZoneCode *string `json:"zoneCode,omitempty"`
}

type PoolInfoResponseBody struct {
	// 返回结果集
	PoolList []PoolList `json:"poolList,omitempty"`
}
