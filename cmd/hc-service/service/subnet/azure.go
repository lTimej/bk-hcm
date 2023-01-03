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

// Package subnet defines subnet service.
package subnet

import (
	"fmt"

	"hcm/pkg/adaptor/types"
	adcore "hcm/pkg/adaptor/types/core"
	cloudcore "hcm/pkg/api/core/cloud"
	dataservice "hcm/pkg/api/data-service"
	"hcm/pkg/api/data-service/cloud"
	hcservice "hcm/pkg/api/hc-service"
	"hcm/pkg/criteria/enumor"
	"hcm/pkg/criteria/errf"
	"hcm/pkg/dal/dao/tools"
	"hcm/pkg/logs"
	"hcm/pkg/rest"
	"hcm/pkg/tools/converter"
	"hcm/pkg/tools/uuid"
)

// AzureSubnetUpdate update azure subnet.
func (v subnet) AzureSubnetUpdate(cts *rest.Contexts) (interface{}, error) {
	id := cts.PathParameter("id").String()

	req := new(hcservice.SubnetUpdateReq)
	if err := cts.DecodeInto(req); err != nil {
		return nil, errf.NewFromErr(errf.DecodeRequestFailed, err)
	}
	if err := req.Validate(); err != nil {
		return nil, errf.NewFromErr(errf.InvalidParameter, err)
	}

	getRes, err := v.cs.DataService().Azure.Subnet.Get(cts.Kit.Ctx, cts.Kit.Header(), id)
	if err != nil {
		return nil, err
	}

	cli, err := v.ad.Azure(cts.Kit, getRes.AccountID)
	if err != nil {
		return nil, err
	}

	updateOpt := new(types.AzureSubnetUpdateOption)
	err = cli.UpdateSubnet(cts.Kit, updateOpt)
	if err != nil {
		return nil, err
	}

	updateReq := &cloud.SubnetBatchUpdateReq[cloud.AzureSubnetUpdateExt]{
		Subnets: []cloud.SubnetUpdateReq[cloud.AzureSubnetUpdateExt]{{
			ID: id,
			SubnetUpdateBaseInfo: cloud.SubnetUpdateBaseInfo{
				Memo: req.Memo,
			},
		}},
	}
	err = v.cs.DataService().Azure.Subnet.BatchUpdate(cts.Kit.Ctx, cts.Kit.Header(), updateReq)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// AzureSubnetDelete delete azure subnet.
func (v subnet) AzureSubnetDelete(cts *rest.Contexts) (interface{}, error) {
	id := cts.PathParameter("id").String()

	getRes, err := v.cs.DataService().Azure.Subnet.Get(cts.Kit.Ctx, cts.Kit.Header(), id)
	if err != nil {
		return nil, err
	}

	cli, err := v.ad.Azure(cts.Kit, getRes.AccountID)
	if err != nil {
		return nil, err
	}

	delOpt := &adcore.AzureDeleteOption{
		BaseDeleteOption:  adcore.BaseDeleteOption{ResourceID: getRes.CloudID},
		ResourceGroupName: getRes.Extension.ResourceGroup,
	}
	err = cli.DeleteSubnet(cts.Kit, delOpt)
	if err != nil {
		return nil, err
	}

	deleteReq := &dataservice.BatchDeleteReq{
		Filter: tools.EqualExpression("id", id),
	}
	err = v.cs.DataService().Global.Subnet.BatchDelete(cts.Kit.Ctx, cts.Kit.Header(), deleteReq)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// AzureSubnetSync sync azure cloud subnet.
func (v subnet) AzureSubnetSync(cts *rest.Contexts) (interface{}, error) {
	req := new(hcservice.ResourceSyncReq)
	if err := cts.DecodeInto(req); err != nil {
		return nil, errf.NewFromErr(errf.DecodeRequestFailed, err)
	}
	if err := req.Validate(); err != nil {
		return nil, errf.NewFromErr(errf.InvalidParameter, err)
	}
	if len(req.ResourceGroupName) == 0 {
		return nil, errf.NewFromErr(errf.InvalidParameter, fmt.Errorf("resource_group_name is required"))
	}
	if len(req.VpcName) == 0 {
		return nil, errf.NewFromErr(errf.InvalidParameter, fmt.Errorf("vpc_name is required"))
	}

	var (
		vendorName = enumor.Azure
		rsp        = hcservice.ResourceSyncResult{
			TaskID: uuid.UUID(),
		}
	)

	// batch get subnet list from cloudapi.
	list, err := v.BatchGetAzureSubnetList(cts, req)
	if err != nil || list == nil {
		logs.Errorf("[%s-subnet] request cloudapi response failed. accountID:%s, region:%s, err:%v",
			vendorName, req.AccountID, req.Region, err)
		return nil, err
	}

	// batch get subnet map from db.
	resourceDBMap, err := v.BatchGetSubnetMapFromDB(cts, req, vendorName)
	if err != nil {
		logs.Errorf("[%s-subnet] batch get subnetdblist failed. accountID:%s, region:%s, err:%v",
			vendorName, req.AccountID, req.Region, err)
		return nil, err
	}

	// batch compare vendor subnet list.
	_, err = v.BatchCompareAzureSubnetList(cts, req, list, resourceDBMap)
	if err != nil {
		logs.Errorf("[%s-subnet] compare api and dblist failed. accountID:%s, region:%s, err:%v",
			vendorName, req.AccountID, req.Region, err)
		return nil, err
	}

	return rsp, nil
}

// BatchGetAzureSubnetList batch get subnet list from cloudapi.
func (v subnet) BatchGetAzureSubnetList(cts *rest.Contexts, req *hcservice.ResourceSyncReq) (
	*types.AzureSubnetListResult, error) {
	cli, err := v.ad.Azure(cts.Kit, req.AccountID)
	if err != nil {
		return nil, err
	}

	opt := &types.AzureSubnetListOption{
		VpcName: req.VpcName,
	}
	opt.ResourceGroupName = req.ResourceGroupName
	list, err := cli.ListSubnet(cts.Kit, opt)
	if err != nil || list == nil {
		logs.Errorf("[%s-subnet]batch get cloud api failed. accountID:%s, region:%s, err:%v",
			enumor.Azure, req.AccountID, req.Region, err)
		return nil, err
	}
	return list, nil
}

// BatchCompareAzureSubnetList batch compare vendor subnet list.
func (v subnet) BatchCompareAzureSubnetList(cts *rest.Contexts, req *hcservice.ResourceSyncReq,
	list *types.AzureSubnetListResult, resourceDBMap map[string]cloudcore.BaseSubnet) (interface{}, error) {
	var (
		createResources []cloud.SubnetCreateReq[cloudcore.AzureSubnetExtension]
		updateResources []cloud.SubnetUpdateReq[cloud.AzureSubnetUpdateExt]
		existIDMap      = map[string]bool{}
		deleteIDs       []string
	)

	err := v.filterAzureSubnetList(req, list, resourceDBMap, &createResources, &updateResources, existIDMap)
	if err != nil {
		return nil, err
	}

	// update resource data
	if len(updateResources) > 0 {
		if err = v.cs.DataService().Azure.Subnet.BatchUpdate(cts.Kit.Ctx, cts.Kit.Header(),
			&cloud.SubnetBatchUpdateReq[cloud.AzureSubnetUpdateExt]{
				Subnets: updateResources,
			}); err != nil {
			logs.Errorf("[%s-subnet]batch compare db update failed. accountID:%s, region:%s, err:%v",
				enumor.Azure, req.AccountID, req.Region, err)
			return nil, err
		}
	}

	// add resource data
	if len(createResources) > 0 {
		if _, err = v.cs.DataService().Azure.Subnet.BatchCreate(cts.Kit.Ctx, cts.Kit.Header(),
			&cloud.SubnetBatchCreateReq[cloudcore.AzureSubnetExtension]{
				Subnets: createResources,
			}); err != nil {
			logs.Errorf("[%s-subnet]batch compare db create failed. accountID:%s, region:%s, err:%v",
				enumor.Azure, req.AccountID, req.Region, err)
			return nil, err
		}
	}

	// delete resource data
	for _, resItem := range resourceDBMap {
		if _, ok := existIDMap[resItem.ID]; !ok {
			deleteIDs = append(deleteIDs, resItem.ID)
		}
	}
	if len(deleteIDs) > 0 {
		if err = v.cs.DataService().Global.Subnet.BatchDelete(cts.Kit.Ctx, cts.Kit.Header(),
			&dataservice.BatchDeleteReq{
				Filter: tools.ContainersExpression("id", deleteIDs),
			}); err != nil {
			logs.Errorf("[%s-subnet]batch compare db delete failed. accountID:%s, region:%s, delIDs:%v, err:%v",
				enumor.Azure, req.AccountID, req.Region, deleteIDs, err)
			return nil, err
		}
	}
	return nil, nil
}

func (v subnet) filterAzureSubnetList(req *hcservice.ResourceSyncReq, list *types.AzureSubnetListResult,
	resourceDBMap map[string]cloudcore.BaseSubnet,
	createResources *[]cloud.SubnetCreateReq[cloudcore.AzureSubnetExtension],
	updateResources *[]cloud.SubnetUpdateReq[cloud.AzureSubnetUpdateExt], existIDMap map[string]bool) error {
	if list == nil || len(list.Details) == 0 {
		return fmt.Errorf("cloudapi subnetlist is empty, accountID:%s, region:%s", req.AccountID, req.Region)
	}

	for _, item := range list.Details {
		// need compare and update subnet data
		if resourceInfo, ok := resourceDBMap[item.CloudID]; ok {
			tmpRes := cloud.SubnetUpdateReq[cloud.AzureSubnetUpdateExt]{
				ID: resourceInfo.ID,
				Extension: &cloud.AzureSubnetUpdateExt{
					NatGateway:           converter.ValToPtr(item.Extension.NatGateway),
					NetworkSecurityGroup: converter.ValToPtr(item.Extension.NetworkSecurityGroup),
				},
			}
			tmpRes.Name = converter.ValToPtr(item.Name)
			tmpRes.Ipv4Cidr = item.Ipv4Cidr
			if len(item.Ipv6Cidr) > 0 {
				tmpRes.Ipv6Cidr = item.Ipv6Cidr
			} else {
				tmpRes.Ipv6Cidr = []string{""}
			}
			tmpRes.Memo = item.Memo

			*updateResources = append(*updateResources, tmpRes)
			existIDMap[resourceInfo.ID] = true
		} else {
			// need add subnet data
			tmpRes := cloud.SubnetCreateReq[cloudcore.AzureSubnetExtension]{
				AccountID:  req.AccountID,
				CloudVpcID: item.CloudVpcID,
				CloudID:    item.CloudID,
				Name:       converter.ValToPtr(item.Name),
				Ipv4Cidr:   item.Ipv4Cidr,
				Memo:       item.Memo,
				Extension: &cloudcore.AzureSubnetExtension{
					ResourceGroup:        item.Extension.ResourceGroup,
					NatGateway:           item.Extension.NatGateway,
					NetworkSecurityGroup: item.Extension.NetworkSecurityGroup,
				},
			}
			if len(item.Ipv6Cidr) > 0 {
				tmpRes.Ipv6Cidr = item.Ipv6Cidr
			} else {
				tmpRes.Ipv6Cidr = []string{""}
			}

			*createResources = append(*createResources, tmpRes)
		}
	}
	return nil
}
