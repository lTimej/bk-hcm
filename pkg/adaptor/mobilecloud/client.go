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

package mobilecloud

import (
	"fmt"
	"hcm/pkg/adaptor/mobilecloud/emop"
	"hcm/pkg/adaptor/mobilecloud/emop/eboprsa/ecloudsdkcore/auth"
	"hcm/pkg/adaptor/mobilecloud/emop/eboprsa/ecloudsdkcore/auth/provider"
	"hcm/pkg/adaptor/mobilecloud/emop/eboprsa/ecloudsdkcore/config"
	"hcm/pkg/adaptor/types"

	vpcconfig "gitlab.ecloud.com/ecloud/ecloudsdkcore/config"
)

// NewCredentialsFunc ...
type NewCredentialsFunc func() *auth.Credential

type clientSet struct {
	ConfigBuilder    *config.ConfigBuilder
	VpcConfigBuilder *vpcconfig.ConfigBuilder
}

func newClientSet(secret *types.MobileCloudCredential) *clientSet {
	var cretentialType auth.CredentialType
	var globalQueryParams map[string]string
	if secret.Method != "" {
		cretentialType = auth.CredentialMop
		globalQueryParams = map[string]string{
			"appId":  secret.AppID,
			"format": emop.Format,
			"method": secret.Method,
			"status": emop.Status,
		}
	} else {
		cretentialType = auth.CredentialAkSk
		globalQueryParams = map[string]string{
			"Version": emop.Version,
		}
		secret.SecretKey, _ = auth.DecryptSK("6d81d8e09ed840128aabf050f4eb5c5f", secret.SecretKey)
	}
	return &clientSet{
		ConfigBuilder: config.NewConfigBuilder().CentralTransportEnabled(false).IgnoreGateway(true).GlobalQueryParams(globalQueryParams).Provider(provider.NewBasicCredentialProvider(&auth.Credential{
			SecretKey:  &secret.SecretKey,
			AccessKey:  &secret.AccessKey,
			PrivateKey: &secret.Privatekey,
			PublicKey:  &secret.PublicKey,
			// EncryptionType: auth.EncryptionTypePointer(auth.EncrytMopRsa),
			CredentialType: auth.CredentialTypePointer(cretentialType),
		}),
		),
	}
}

func (c *clientSet) emopClient() (client *emop.Client, err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("mobile cloud emop error recovered, err: %v", p)
		}
	}()
	client = emop.NewClient(c.ConfigBuilder.Build())

	return client, nil
}

// func (c *clientSet) vpcClient() (cli *ecloudsdkvpc.Client, err error) {
// 	defer func() {
// 		if p := recover(); p != nil {
// 			err = fmt.Errorf("mobile cloud vpc error recovered, err: %v", p)
// 		}
// 	}()

// 	client := ecloudsdkvpc.NewClient(c.VpcConfigBuilder.GlobalHeaderParams(map[string]string{}).Build())

// 	return client, nil
// }
