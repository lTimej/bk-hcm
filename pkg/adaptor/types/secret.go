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

package types

import (
	"hcm/pkg/criteria/errf"
	"hcm/pkg/criteria/validator"
)

// BaseSecret defines the hybrid cloud's base secret info.
type BaseSecret struct {
	// CloudSecretID is the secret id to do credential.
	CloudSecretID string `json:"cloud_secret_id"`
	// CloudSecretKey is the secret key to do credential.
	CloudSecretKey string `json:"cloud_secret_key"`
	// CloudAccountID is the account id to do credential.
	CloudAccountID string `json:"cloud_account_id"`
}

// Validate BaseSecret.
func (b BaseSecret) Validate() error {
	if len(b.CloudSecretID) == 0 {
		return errf.New(errf.InvalidParameter, "secret id is required")
	}

	if len(b.CloudSecretKey) == 0 {
		return errf.New(errf.InvalidParameter, "secret key is required")
	}

	return nil
}

// GcpCredential define gcp credential information.
type GcpCredential struct {
	CloudProjectID string `json:"cloud_project_id" validate:"required"`
	Json           []byte `json:"json,omitempty" validate:"required"`
}

// Validate GcpCredential
func (g *GcpCredential) Validate() error {
	return validator.Validate.Struct(g)
}

// AzureCredential define azure credential information.
type AzureCredential struct {
	CloudTenantID        string `json:"cloud_tenant_id" validate:"required"`
	CloudSubscriptionID  string `json:"cloud_subscription_id" validate:"required"`
	CloudApplicationID   string `json:"cloud_application_id" validate:"required"`
	CloudClientSecretKey string `json:"cloud_client_secret_key" validate:"required"`
}

// Validate AzureCredential
func (a *AzureCredential) Validate() error {
	return validator.Validate.Struct(a)
}

// BaseKey defines the hybrid cloud's base secret info.
type MobileCloudCredential struct {
	AppID      string `json:"app_id" validate:"required"`
	AccessKey  string `json:"access_key" validate:"required"`
	SecretKey  string `json:"secret_key" validate:"required"`
	Privatekey string `json:"private_key" validate:"required"`
	PublicKey  string `json:"public_key" validate:"required"`
	Method     string `json:"method"`
	PoolId     string `json:"poolId"`
}

// Validate BaseKey.
func (mc MobileCloudCredential) Validate() error {
	if len(mc.AppID) == 0 {
		return errf.New(errf.InvalidParameter, "AppID is required")
	}
	if len(mc.AccessKey) == 0 {
		return errf.New(errf.InvalidParameter, "AccessKey is required")
	}

	if len(mc.SecretKey) == 0 {
		return errf.New(errf.InvalidParameter, "SecretKey is required")
	}
	if len(mc.Privatekey) == 0 {
		return errf.New(errf.InvalidParameter, "Privatekey is required")
	}

	if len(mc.PublicKey) == 0 {
		return errf.New(errf.InvalidParameter, "PublicKey is required")
	}

	// if len(mc.Method) == 0 {
	// 	return errf.New(errf.InvalidParameter, "Method is required")
	// }

	return nil
}
