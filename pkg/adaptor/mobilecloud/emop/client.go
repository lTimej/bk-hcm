// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package emop

import (
	"fmt"
	"hcm/pkg/adaptor/mobilecloud/emop/model"
	"sync"

	imsmodel "gitlab.ecloud.com/ecloud/ecloudsdkims/model"
	vpcmodel "gitlab.ecloud.com/ecloud/ecloudsdkvpc/model"

	"hcm/pkg/adaptor/mobilecloud/emop/eboprsa/ecloudsdkcore"

	"hcm/pkg/adaptor/mobilecloud/emop/eboprsa/ecloudsdkcore/config"

	"gitlab.ecloud.com/ecloud/ecloudsdkcore/param"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/request"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

// CIDC-A-8909aa58d8e640ccbec83f5074d89a5d

var MobileCloudSecret sync.Map

type Client struct {
	apiClient   *ecloudsdkcore.APIClient
	config      *config.Config
	httpRequest *request.HttpRequest
	allRegions  map[string]string
}

func NewClient(config *config.Config) *Client {
	httpRequest := request.DefaultHttpRequest()
	httpRequest.Product = product
	httpRequest.Version = version
	httpRequest.SdkVersion = sdkVersion
	ecloudsdkcore.InitConfig(config)
	apiClient := ecloudsdkcore.DefaultApiClient(config, httpRequest)
	client := &Client{
		apiClient:   apiClient,
		config:      config,
		httpRequest: httpRequest,
	}
	client.allRegions = client.initRegions()
	client.setEndpoint(config, httpRequest)
	return client
}

const (
	product    string = "emop"
	version           = "v1"
	sdkVersion        = "1.0.0"
	Version           = "2016-12-05"
	Format            = "json"
	Status            = "0"
)

func (c *Client) initRegions() map[string]string {
	m := map[string]string{
		"mop": "https://36.133.25.49:31015",
	}
	return m
}

func (c *Client) setEndpoint(config *config.Config, httpRequest *request.HttpRequest) {
	if utils.IsUnSet(config.PoolId) {
		httpRequest.Endpoint = c.allRegions["mop"]
		return
	}
	endpoint := c.allRegions[*config.PoolId]
	if utils.IsUnSet(endpoint) {
		httpRequest.Endpoint = utils.DefaultEndpoint
		return
	}
	httpRequest.Endpoint = endpoint
}

// 经销商终端客户注册订购基础口令
func (c *Client) CreateResellerUser(request *model.CreateResellerUserRequest[model.CreateResellerUserRequestBody]) (*model.CreateResellerUserResponse[model.CreateResellerUserResponseBody], error) {
	return c.CreateResellerUserWithConfig(request, nil)
}

func (c *Client) CreateResellerUserWithConfig(request *model.CreateResellerUserRequest[model.CreateResellerUserRequestBody], runtimeConfig *config.RuntimeConfig) (*model.CreateResellerUserResponse[model.CreateResellerUserResponseBody], error) {
	params := param.NewParamsBuilder().
		Uri("/api/query/emop").
		Protocol("https").
		ContentType("application/json").
		Method("POST").
		Request(request).
		Build()
	returnValue := &model.CreateResellerUserResponse[model.CreateResellerUserResponseBody]{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// 用户ak/sk创建
func (c *Client) CreateAkSk(request *model.CreateAkSkRequest[model.CreateAkSkRequestBody], runtimeConfig *config.RuntimeConfig) (*model.CreateAkSkResponse[model.CreateAkSkResponseBody], error) {
	return c.CreateAkSkWithConfig(request, runtimeConfig)
}

func (c *Client) CreateAkSkWithConfig(request *model.CreateAkSkRequest[model.CreateAkSkRequestBody], runtimeConfig *config.RuntimeConfig) (*model.CreateAkSkResponse[model.CreateAkSkResponseBody], error) {

	params := param.NewParamsBuilder().
		Uri("/api/access/admin/subsystem/admin/create/decrypt").
		Protocol("https").
		ContentType("multipart/form-data").
		Method("POST").
		Request(request).
		Build()
	returnValue := &model.CreateAkSkResponse[model.CreateAkSkResponseBody]{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// 用户ak/sk查询
func (c *Client) QueryAkSk(request *model.QueryAkSkRequest[model.QueryAkSkRequestBody]) (*model.QueryAkSkResponse[model.QueryAkSkResponseBody], error) {
	return c.QueryAkSkWithConfig(request, nil)
}

func (c *Client) QueryAkSkWithConfig(request *model.QueryAkSkRequest[model.QueryAkSkRequestBody], runtimeConfig *config.RuntimeConfig) (*model.QueryAkSkResponse[model.QueryAkSkResponseBody], error) {

	params := param.NewParamsBuilder().
		Uri(fmt.Sprintf("/api/accessKey/query/accesskey/decrypt/aes/python/" + request.QueryAkSkRequestBody.ToString(request.QueryAkSkRequestBody.UserId))).
		Protocol("https").
		ContentType("application/json").
		Method("GET").
		Request(request).
		Build()
	returnValue := &model.QueryAkSkResponse[model.QueryAkSkResponseBody]{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// 经销商终端客户操作（暂停、恢复、注销）
func (c *Client) ResellerUserOperate(request *model.ResellerUserOperateRequest[model.ResellerUserOperateRequestBody]) (*model.ResellerUserOperateResponse[model.ResellerUserOperateResponseBody], error) {
	return c.ResellerUserOperateWithConfig(request, nil)
}

func (c *Client) ResellerUserOperateWithConfig(request *model.ResellerUserOperateRequest[model.ResellerUserOperateRequestBody], runtimeConfig *config.RuntimeConfig) (*model.ResellerUserOperateResponse[model.ResellerUserOperateResponseBody], error) {
	params := param.NewParamsBuilder().
		Uri("/api/query/emop").
		Protocol("https").
		ContentType("application/json").
		Method("POST").
		Request(request).
		Build()
	returnValue := &model.ResellerUserOperateResponse[model.ResellerUserOperateResponseBody]{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// OP或第三方通过此接口通过产品类型查询所支持的资源池信息
func (c *Client) PoolInfo(request *model.PoolInfoRequest[model.PoolInfoRequestBody]) (*model.PoolInfoResponse[model.PoolInfoResponseBody], error) {
	return c.PoolInfoWithConfig(request, nil)
}

func (c *Client) PoolInfoWithConfig(request *model.PoolInfoRequest[model.PoolInfoRequestBody], runtimeConfig *config.RuntimeConfig) (*model.PoolInfoResponse[model.PoolInfoResponseBody], error) {
	params := param.NewParamsBuilder().
		Uri("/api/query/emop").
		Protocol("https").
		ContentType("application/json").
		Method("POST").
		Request(request).
		Build()
	returnValue := &model.PoolInfoResponse[model.PoolInfoResponseBody]{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// 资源池智能推荐
func (c *Client) GetDefaultPoolInfo(request *model.PoolInfoRequest[model.PoolInfoRequestBody]) (*model.PoolInfoResponse[model.PoolInfoResponseBody], error) {
	return c.GetDefaultPoolInfoWithConfig(request, nil)
}

func (c *Client) GetDefaultPoolInfoWithConfig(request *model.PoolInfoRequest[model.PoolInfoRequestBody], runtimeConfig *config.RuntimeConfig) (*model.PoolInfoResponse[model.PoolInfoResponseBody], error) {
	params := param.NewParamsBuilder().
		Uri("/api/query/emop").
		Protocol("https").
		ContentType("application/json").
		Method("POST").
		Request(request).
		Build()
	returnValue := &model.PoolInfoResponse[model.PoolInfoResponseBody]{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// 查询产品信息
func (c *Client) QryOffer(request *model.QryOfferRequest[model.QryOfferRequestBody]) (*model.QryOfferResponse[model.QryOfferResponseBody], error) {
	return c.QryOfferWithConfig(request, nil)
}

func (c *Client) QryOfferWithConfig(request *model.QryOfferRequest[model.QryOfferRequestBody], runtimeConfig *config.RuntimeConfig) (*model.QryOfferResponse[model.QryOfferResponseBody], error) {
	params := param.NewParamsBuilder().
		Uri("/api/query/emop").
		Protocol("https").
		ContentType("application/json").
		Method("POST").
		Request(request).
		Build()
	returnValue := &model.QryOfferResponse[model.QryOfferResponseBody]{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// 经销商业务操作校验接口
func (c *Client) ResellerOrderOperateVerify(request *model.ResellerOrderOperateVerifyRequest[model.ResellerOrderOperateVerifyRequestBody]) (*model.ResellerOrderOperateVerifyResponse[model.ResellerOrderOperateVerifyResponseBody], error) {
	return c.ResellerOrderOperateVerifyWithConfig(request, nil)
}

func (c *Client) ResellerOrderOperateVerifyWithConfig(request *model.ResellerOrderOperateVerifyRequest[model.ResellerOrderOperateVerifyRequestBody], runtimeConfig *config.RuntimeConfig) (*model.ResellerOrderOperateVerifyResponse[model.ResellerOrderOperateVerifyResponseBody], error) {
	params := param.NewParamsBuilder().
		Uri("/api/query/emop").
		Protocol("https").
		ContentType("application/json").
		Method("POST").
		Request(request).
		Build()
	returnValue := &model.ResellerOrderOperateVerifyResponse[model.ResellerOrderOperateVerifyResponseBody]{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// 配额校验
func (c *Client) ManageQuotaCheck(request *model.QuotaCheckRequest[model.QuotaCheckRequestBody], runtimeConfig *config.RuntimeConfig) (*model.QuotaCheckResponse, error) {
	return c.ManageQuotaCheckWithConfig(request, runtimeConfig)
}

func (c *Client) ManageQuotaCheckWithConfig(request *model.QuotaCheckRequest[model.QuotaCheckRequestBody], runtimeConfig *config.RuntimeConfig) (*model.QuotaCheckResponse, error) {
	params := param.NewParamsBuilder().
		Uri("/api/quota/manage/v4/quota/check").
		Protocol("https").
		ContentType("application/json").
		Method("PUT").
		Request(request).
		Build()
	returnValue := &model.QuotaCheckResponse{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// 冻结配额
func (c *Client) ManageQuotaFreeze(request *model.QuotaCheckRequest[model.QuotaCheckRequestBody]) (*model.QuotaCheckResponse, error) {
	return c.ManageQuotaCheckWithConfig(request, nil)
}

func (c *Client) ManageQuotaFreezeWithConfig(request *model.QuotaCheckRequest[model.QuotaCheckRequestBody], runtimeConfig *config.RuntimeConfig) (*model.QuotaCheckResponse, error) {
	params := param.NewParamsBuilder().
		Uri("/api/query/manage/v4/quota/check").
		Protocol("https").
		ContentType("application/json").
		Method("PUT").
		Request(request).
		Build()
	returnValue := &model.QuotaCheckResponse{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// 创建vpc
func (c *Client) CreateVpc(request *vpcmodel.VpcOrderRequest, runtimeConfig *config.RuntimeConfig) (*vpcmodel.VpcOrderResponse, error) {
	return c.CreateVpcWithConfig(request, runtimeConfig)
}

func (c *Client) CreateVpcWithConfig(request *vpcmodel.VpcOrderRequest, runtimeConfig *config.RuntimeConfig) (*vpcmodel.VpcOrderResponse, error) {
	params := param.NewParamsBuilder().
		Uri("/api/openapi-vpc/customer/v3/order/create/vpc").
		Protocol("https").
		ContentType("application/json").
		Method("POST").
		Request(request).
		Build()
	returnValue := &vpcmodel.VpcOrderResponse{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// 查看vpc
func (c *Client) ListVpc(request *vpcmodel.ListVpcRequest, runtimeConfig *config.RuntimeConfig) (*vpcmodel.ListVpcResponse, error) {
	return c.ListVpcWithConfig(request, runtimeConfig)
}

func (c *Client) ListVpcWithConfig(request *vpcmodel.ListVpcRequest, runtimeConfig *config.RuntimeConfig) (*vpcmodel.ListVpcResponse, error) {
	params := param.NewParamsBuilder().
		Uri("/api/openapi-vpc/customer/v3/vpc").
		Protocol("https").
		ContentType("application/json").
		Method("GET").
		Request(request).
		Build()
	returnValue := &vpcmodel.ListVpcResponse{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// 查看vpc下网络
func (c *Client) ListNetworkResps(request *vpcmodel.ListNetworkRespRequest, runtimeConfig *config.RuntimeConfig) (*vpcmodel.ListNetworkRespResponse, error) {
	return c.ListNetworkRespsWithConfig(request, runtimeConfig)
}

func (c *Client) ListNetworkRespsWithConfig(request *vpcmodel.ListNetworkRespRequest, runtimeConfig *config.RuntimeConfig) (*vpcmodel.ListNetworkRespResponse, error) {
	params := param.NewParamsBuilder().
		Uri("/api/openapi-vpc/customer/v3/network/NetworkResps").
		Protocol("https").
		ContentType("application/json").
		Method("GET").
		Request(request).
		Build()
	returnValue := &vpcmodel.ListNetworkRespResponse{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// 查看网卡
func (c *Client) ListPort(request *vpcmodel.ListPortRequest, runtimeConfig *config.RuntimeConfig) (*vpcmodel.ListPortResponse, error) {
	return c.ListPortWithConfig(request, runtimeConfig)
}

func (c *Client) ListPortWithConfig(request *vpcmodel.ListPortRequest, runtimeConfig *config.RuntimeConfig) (*vpcmodel.ListPortResponse, error) {
	params := param.NewParamsBuilder().
		Uri("/api/openapi-vpc/customer/v3/port").
		Protocol("https").
		ContentType("application/json").
		Method("GET").
		Request(request).
		Build()
	returnValue := &vpcmodel.ListPortResponse{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// 创建安全组
func (c *Client) CreateSecurityGroup(request *vpcmodel.CreateSecurityGroupRequest, runtimeConfig *config.RuntimeConfig) (*vpcmodel.CreateSecurityGroupResponse, error) {
	return c.CreateSecurityGroupWithConfig(request, runtimeConfig)
}

func (c *Client) CreateSecurityGroupWithConfig(request *vpcmodel.CreateSecurityGroupRequest, runtimeConfig *config.RuntimeConfig) (*vpcmodel.CreateSecurityGroupResponse, error) {
	params := param.NewParamsBuilder().
		Uri("/api/openapi-vpc/customer/v3/SecurityGroup").
		Protocol("https").
		ContentType("application/json").
		Method("POST").
		Request(request).
		Build()
	returnValue := &vpcmodel.CreateSecurityGroupResponse{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// 创建安全组规则
func (c *Client) CreateSecurityGroupRule(request *vpcmodel.CreateSecurityGroupRuleRequest, runtimeConfig *config.RuntimeConfig) (*vpcmodel.CreateSecurityGroupRuleResponse, error) {
	return c.CreateSecurityGroupRuleWithConfig(request, runtimeConfig)
}

func (c *Client) CreateSecurityGroupRuleWithConfig(request *vpcmodel.CreateSecurityGroupRuleRequest, runtimeConfig *config.RuntimeConfig) (*vpcmodel.CreateSecurityGroupRuleResponse, error) {
	params := param.NewParamsBuilder().
		Uri("/api/openapi-vpc/customer/v3/SecurityGroupRule").
		Protocol("https").
		ContentType("application/json").
		Method("POST").
		Request(request).
		Build()
	returnValue := &vpcmodel.CreateSecurityGroupRuleResponse{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// 查看安全组
func (c *Client) ListSecurityGroup(request *vpcmodel.ListSecGroupRequest, runtimeConfig *config.RuntimeConfig) (*vpcmodel.ListSecGroupResponse, error) {
	return c.ListSecurityGroupWithConfig(request, runtimeConfig)
}

func (c *Client) ListSecurityGroupWithConfig(request *vpcmodel.ListSecGroupRequest, runtimeConfig *config.RuntimeConfig) (*vpcmodel.ListSecGroupResponse, error) {
	params := param.NewParamsBuilder().
		Uri("/api/openapi-vpc/customer/v3/SecurityGroup").
		Protocol("https").
		ContentType("application/json").
		Method("GET").
		Request(request).
		Build()
	returnValue := &vpcmodel.ListSecGroupResponse{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// 订购
func (c *Client) CreateOrderUnify(request *model.CreateOrderUnifyRequest[model.CreateOrderUnifyRequestBody]) (*model.CreateOrderUnifyResponse[model.CreateOrderUnifyResponseBody], error) {
	return c.CreateOrderUnifyWithConfig(request, nil)
}

func (c *Client) CreateOrderUnifyWithConfig(request *model.CreateOrderUnifyRequest[model.CreateOrderUnifyRequestBody], runtimeConfig *config.RuntimeConfig) (*model.CreateOrderUnifyResponse[model.CreateOrderUnifyResponseBody], error) {
	params := param.NewParamsBuilder().
		Uri("/api/query/emop").
		Protocol("https").
		ContentType("application/json").
		Method("POST").
		Request(request).
		Build()
	returnValue := &model.CreateOrderUnifyResponse[model.CreateOrderUnifyResponseBody]{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// 获取镜像
func (c *Client) ListPublicImages(request *imsmodel.ListServerPublicImageV2Request, runtimeConfig *config.RuntimeConfig) (*imsmodel.ListServerPublicImageV2Response, error) {
	return c.ListPublicImagesWithConfig(request, runtimeConfig)
}

func (c *Client) ListPublicImagesWithConfig(request *imsmodel.ListServerPublicImageV2Request, runtimeConfig *config.RuntimeConfig) (*imsmodel.ListServerPublicImageV2Response, error) {
	params := param.NewParamsBuilder().
		Uri("/api/web/routes/ims-backend-console/user/v3/image/server/public").
		Protocol("https").
		ContentType("application/json").
		Method("GET").
		Request(request).
		Build()
	returnValue := &imsmodel.ListServerPublicImageV2Response{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}
