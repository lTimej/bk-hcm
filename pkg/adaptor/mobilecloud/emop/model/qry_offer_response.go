// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

type QryOfferResponseOption QryOfferResponse[QryOfferResponseBody]

type QryOfferResponse[T QryOfferResponseBody] struct {

	// 每个请求的序列号
	RespCode *string `json:"respCode,omitempty"`
	// 页面国际化错误提示
	RespDesc *string `json:"respDesc,omitempty"`
	// 统一错误码
	Result *T `json:"result,omitempty"`
}

type Body struct {
	// 商品名称
	OfferName *string `json:"offerName,omitempty"`
	// 产码;全网产品编码，MQ消息的tag
	ProductId *string `json:"productId,omitempty"`
	//
	BusiTestUserOrderFlag *string `json:"busiTestUserOrderFlag,omitempty"`
	//
	IsNewData *string `json:"isNewData,omitempty"`
	// 允许重复订购;false 不允许；true 允许
	Repeat *bool `json:"repeat,omitempty"`
	// 商品编码
	OfferId *string `json:"offerId,omitempty"`
	// 商品规格集合
	ProductChas []ProductChas `json:"productChas,omitempty"`
	// 当前资源池下所有可用区信息
	OfferChas map[string]interface{} `json:"offerChas,omitempty"`
	// 产品类型;Vm, emrVm
	ProductType *string `json:"productType,omitempty"`
	// 当前资源池下所有可用区信息
	Status *string `json:"status,omitempty"`
}

type ProductChas struct {
	// 规格ID
	ChaId *string `json:"chaId,omitempty"`
	// 规格排序
	ChaSeq *int `json:"chaSeq,omitempty"`
	// 资源池信息
	Pools []Pools `json:"pools,omitempty"`
	// 规格名称
	ChaName *string `json:"chaName,omitempty"`
	//
	IsTrail *string `json:"isTrail,omitempty"`
	// 价格
	Prices []Prices `json:"prices,omitempty"`
	// 规格属性集合
	Attrs []Attrs `json:"attrs,omitempty"`
}

type Pools struct {
	// 可用区列表
	ZoneList []ZoneList `json:"zoneList,omitempty"`
	// 资源池编码
	PoolId *string `json:"poolId,omitempty"`
	// 资源池状态
	ChaStatus *string `json:"chaStatus,omitempty"`
	// 资源池名称
	PoolName *string `json:"poolName,omitempty"`
}

type ZoneList struct {
	// 容量数
	CapacityCount *string `json:"capacityCount,omitempty"`
	// 可用区id
	ZoneId *string `json:"zoneId,omitempty"`
	// 可用区类型
	ZoneType *string `json:"zoneType,omitempty"`
	// 可用区详情
	ZoneDesc *string `json:"zoneDesc,omitempty"`
	// 可用区名
	ZoneName *string `json:"zoneName,omitempty"`
	// 可用区状态
	ZoneStatus *string `json:"zoneStatus,omitempty"`
	// 售罄标记;1售罄，0未售罄
	SoldOut *string `json:"soldOut,omitempty"`
	// 剩余量
	RemainCount *string `json:"remainCount,omitempty"`
	// 可用区编码
	ZoneCode *string `json:"zoneCode,omitempty"`
}

type Prices struct {
	// 计费单位
	FeeUnit *string `json:"feeUnit,omitempty"`
	// 计费方式
	PriceType *string `json:"priceType,omitempty"`
	// 付费方式
	ChargeType *string `json:"chargeType,omitempty"`
	// 资费场景
	Scenes []Scenes `json:"scenes,omitempty"`
	// 资费编码
	PriceId *string `json:"priceId,omitempty"`
	// 资费名称
	PriceName *string `json:"priceName,omitempty"`
	// 可用区编码
	FeeType *string `json:"feeType,omitempty"`
	// 资费排序
	PriceSeq *int `json:"priceSeq,omitempty"`
	// 自费状态
	Status *string `json:"status,omitempty"`
}

type Scenes struct {
	// 场景资费单位;元、元/M
	Unit *string `json:"unit,omitempty"`
	// 是否阶梯场景；true表示是，false表示否
	Ladder *bool `json:"ladder,omitempty"`
	// 话单类型;话单场景下的chargeMode
	ChargeMode *string `json:"chargeMode,omitempty"`
	// 场景名称
	SceneName *string `json:"sceneName,omitempty"`
	// 单价费用
	Fee *string `json:"fee,omitempty"`
	// 场景ID
	SceneId *string `json:"sceneId,omitempty"`
	// 是否话单场景
	Bill *bool `json:"bill,omitempty"`
	// 折扣
	Discount *string `json:"discount,omitempty"`
}

type Attrs struct {
	// 属性ID
	AttrId *string `json:"attrId,omitempty"`
	// 属性编码
	AttrCode *string `json:"attrCode,omitempty"`
	// // 可用区编码
	// AttrSeqMap map[string]interface{} `json:"attrSeqMap,omitempty"`
	// 属性值
	AttrVal *string `json:"attrVal,omitempty"`
	// 属性排序
	AttrSeq *int `json:"attrSeq,omitempty"`
	// 属性名称
	AttrName *string `json:"attrName,omitempty"`
	// 属性类型;2 展示属性（仅用作展示），10 订购属性（需要订购时传递回来）
	AttrType *string `json:"attrType,omitempty"`
}

type QryOfferResponseBody struct {
	// 返回结果集
	Body []Body `json:"body,omitempty"`
}
