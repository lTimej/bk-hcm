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
	"hcm/pkg/adaptor/types"
	"hcm/pkg/criteria/errf"
)

const (
	// emop interface
	Mop        = "mop"
	AppId      = "50381602"
	AccessKey  = "49c0ea9b0f2045d9b61cd95e459ecd60"
	SecretKey  = "6d81d8e09ed840128aabf050f4eb5c5f"
	PrivateKey = "MIIBhQIBADANBgkqhkiG9w0BAQEFAASCAW8wggFrAgEAAkwAi0bTY0NsHsXNDERRLXJByQmi+dU/eblf5qw4zZN+/cc13mRtA+dYSatw+E6QiUCdX5lYWYz5/lcGYawf5gNIu4DLMkxMkwDDgPlrAgMBAAECSzTZwntnaU7gFmgyQG+zbL1B9+NABZ9GNdsNvVxdPRJGFu32Q9vvVU82DoVDxZ2R5fBZXvmOph4EpdLvxiOVMtndYaSL/8W4PVYz6QImDWjaZKhJSh7+dcoKili0TdyNXrlyBwaM71SLGyLrRFGeifsM5rcCJgpi6A/TzD0kzr9iOu3rvN1d+jiw1TGRwKRNI2n/A9NioMtzOa7tAiYFGEWB4P6XjtcXYcBHeBRpUNbVmpfcW3zIodKIaOgCeRBHVH8+WQImAMQe3dv/epsWbONv+VCkE6f05u2ULB3WGchezlizDYp+1cLgBFkCJgY3ORoLB5TMckLlKGPl3+zw1GgiFEJ6ii4Yns6BvZwm3EflcZ7/"
	PublicKey  = "MGcwDQYJKoZIhvcNAQEBBQADVgAwUwJMAItG02NDbB7FzQxEUS1yQckJovnVP3m5X+asOM2Tfv3HNd5kbQPnWEmrcPhOkIlAnV+ZWFmM+f5XBmGsH+YDSLuAyzJMTJMAw4D5awIDAQAB"
)

// NewHuaWei new huawei.
func NewMobileCloud(s *types.MobileCloudCredential) (*MobileCloud, error) {
	if err := validateSecret(s); err != nil {
		return nil, err
	}
	return &MobileCloud{clientSet: newClientSet(s)}, nil
}

// HuaWei is huawei operator.
type MobileCloud struct {
	clientSet *clientSet
}

func validateSecret(s *types.MobileCloudCredential) error {
	if s == nil {
		return errf.New(errf.InvalidParameter, "secret is required")
	}

	if err := s.Validate(); err != nil {
		return err
	}

	return nil
}

// sliceToPtr convert slice to pointer.
func sliceToPtr[T any](slice []T) *[]T {
	ptrArr := make([]T, len(slice))
	for idx, val := range slice {
		ptrArr[idx] = val
	}
	return &ptrArr
}
