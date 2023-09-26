/*
 * TencentBlueKing is pleased to support the open source community by making
 * 蓝鲸智云 - 混合云管理平台 (BlueKing - Hybrid Cloud Management System) available.
 * Copyright (C) 2022 THL A29 Limited,
 * a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on
 * an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 *
 * We undertake not to change the open source license (MIT license) applicable
 *
 * to the current version of the project delivered to anyone in the future.
 */

package instancetype

import (
	"fmt"
	"strings"

	proto "hcm/pkg/api/cloud-server/instance-type"
	hcproto "hcm/pkg/api/hc-service/instance-type"
	"hcm/pkg/criteria/enumor"
	"hcm/pkg/criteria/errf"
	"hcm/pkg/dal/dao/types"
	"hcm/pkg/iam/meta"
	"hcm/pkg/logs"
	"hcm/pkg/rest"
	"hcm/pkg/tools/converter"
	"hcm/pkg/tools/hooks/handler"
)

// ListInRes ...
func (svc *instanceTypeSvc) ListInRes(cts *rest.Contexts) (interface{}, error) {
	return svc.list(cts, handler.ResValidWithAuth)
}

// ListInBiz ...
func (svc *instanceTypeSvc) ListInBiz(cts *rest.Contexts) (interface{}, error) {
	return svc.list(cts, handler.BizValidWithAuth)
}

// list ...
func (svc *instanceTypeSvc) list(cts *rest.Contexts, validHandler handler.ValidWithAuthHandler) (interface{}, error) {
	req := new(proto.ListReq)
	if err := cts.DecodeInto(req); err != nil {
		return nil, err
	}

	if err := req.Validate(); err != nil {
		return nil, errf.NewFromErr(errf.InvalidParameter, err)
	}

	// validate biz and authorize
	err := validHandler(cts, &handler.ValidWithAuthOption{Authorizer: svc.authorizer, ResType: meta.InstanceType,
		Action: meta.Find, DisableBizIDEqual: true, BasicInfo: &types.CloudResourceBasicInfo{
			AccountID: req.AccountID,
		}})
	if err != nil {
		return nil, err
	}

	switch req.Vendor {
	case enumor.TCloud:
		return svc.ListForTCloud(cts, req)
	case enumor.Aws:
		return svc.ListForAws(cts, req)
	case enumor.HuaWei:
		return svc.ListForHuaWei(cts, req)
	case enumor.Azure:
		return svc.ListForAzure(cts, req)
	case enumor.Gcp:
		return svc.ListForGcp(cts, req)
	}

	return nil, nil
}

// ListForTCloud ...
func (svc *instanceTypeSvc) ListForTCloud(cts *rest.Contexts, req *proto.ListReq) (interface{}, error) {
	listReq := &hcproto.TCloudInstanceTypeListReq{
		AccountID:          req.AccountID,
		Region:             req.Region,
		Zone:               req.Zone,
		InstanceChargeType: req.InstanceChargeType,
	}
	list, err := svc.client.HCService().TCloud.InstanceType.List(cts.Kit, listReq)
	if err != nil {
		logs.Errorf("call hc-service to list instance type failed, err: %v, rid: %s", err, cts.Kit.Rid)
		return nil, err
	}

	familyMap := make(map[string]struct{})
	familyTypeMap := make(map[string]struct{})
	for _, one := range list {
		familyMap[one.InstanceFamily] = struct{}{}
		// e.g: 标准型SA3 -> 标准型
		familyTypeMap[one.TypeName[:len(one.TypeName)-len(one.InstanceFamily)]] = struct{}{}
	}

	result := &proto.ListResult[hcproto.TCloudInstanceTypeResp]{
		InstanceFamilyTypeNames: converter.MapKeyToStringSlice(familyTypeMap),
		InstanceFamilies:        converter.MapKeyToStringSlice(familyMap),
		InstanceTypes:           list,
	}

	return result, nil
}

// ListForAws ...
func (svc *instanceTypeSvc) ListForAws(cts *rest.Contexts, req *proto.ListReq) (interface{}, error) {

	listReq := &hcproto.AwsInstanceTypeListReq{
		AccountID: req.AccountID,
		Region:    req.Region,
	}
	list, err := svc.client.HCService().Aws.InstanceType.List(cts.Kit, listReq)
	if err != nil {
		logs.Errorf("call hc-service to list instance type failed, err: %v, rid: %s", err, cts.Kit.Rid)
		return nil, err
	}

	familyMap := make(map[string]struct{})
	for _, one := range list {
		familyMap[one.InstanceFamily] = struct{}{}
	}

	result := &proto.ListResult[hcproto.AwsInstanceTypeResp]{
		InstanceFamilies: converter.MapKeyToStringSlice(familyMap),
		InstanceTypes:    list,
	}

	return result, nil
}

// ListForHuaWei ...
func (svc *instanceTypeSvc) ListForHuaWei(cts *rest.Contexts, req *proto.ListReq) (interface{}, error) {

	listReq := &hcproto.HuaWeiInstanceTypeListReq{
		AccountID: req.AccountID,
		Region:    req.Region,
		Zone:      req.Zone,
	}
	list, err := svc.client.HCService().HuaWei.InstanceType.List(cts.Kit, listReq)
	if err != nil {
		logs.Errorf("call hc-service to list instance type failed, err: %v, rid: %s", err, cts.Kit.Rid)
		return nil, err
	}

	familyMap := make(map[string]struct{})
	for _, one := range list {
		familyMap[one.InstanceFamily] = struct{}{}
	}

	result := &proto.ListResult[hcproto.HuaWeiInstanceTypeResp]{
		InstanceFamilies: converter.MapKeyToStringSlice(familyMap),
		InstanceTypes:    list,
	}

	return result, nil
}

// ListForAzure ...
func (svc *instanceTypeSvc) ListForAzure(cts *rest.Contexts, req *proto.ListReq) (interface{}, error) {

	listReq := &hcproto.AzureInstanceTypeListReq{
		AccountID: req.AccountID,
		Region:    req.Region,
	}
	list, err := svc.client.HCService().Azure.InstanceType.List(cts.Kit, listReq)
	if err != nil {
		if strings.Contains(err.Error(), "No registered resource provider found for location") {
			return nil, fmt.Errorf("no instance type found for %s region", req.Region)
		}

		logs.Errorf("call hc-service to list instance type failed, err: %v, rid: %s", err, cts.Kit.Rid)
		return nil, err
	}

	familyMap := make(map[string]struct{})
	for _, one := range list {
		fmt.Println("[ ---- debug ---- ]: ", one)
		familyMap[one.InstanceFamily] = struct{}{}
	}

	result := &proto.ListResult[hcproto.AzureInstanceTypeResp]{
		InstanceFamilies: converter.MapKeyToStringSlice(familyMap),
		InstanceTypes:    list,
	}

	return result, nil
}

// ListForGcp ...
func (svc *instanceTypeSvc) ListForGcp(cts *rest.Contexts, req *proto.ListReq) (interface{}, error) {

	listReq := &hcproto.GcpInstanceTypeListReq{
		AccountID: req.AccountID,
		Zone:      req.Zone,
	}
	list, err := svc.client.HCService().Gcp.InstanceType.List(cts.Kit, listReq)
	if err != nil {
		logs.Errorf("call hc-service to list instance type failed, err: %v, rid: %s", err, cts.Kit.Rid)
		return nil, err
	}

	familyMap := make(map[string]struct{})
	for _, one := range list {
		familyMap[one.InstanceFamily] = struct{}{}
	}

	result := &proto.ListResult[hcproto.GcpInstanceTypeResp]{
		InstanceFamilies: converter.MapKeyToStringSlice(familyMap),
		InstanceTypes:    list,
	}

	return result, nil
}
