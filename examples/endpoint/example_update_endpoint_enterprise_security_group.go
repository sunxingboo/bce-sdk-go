/*
 * Copyright 2024 Baidu, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
 * except in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the
 * License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions
 * and limitations under the License.
 */

package endpointexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/endpoint"
)

func UpdateEndpointEnterpriseSecurityGroup() {
	ak, sk, Endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, _ := endpoint.NewClient(ak, sk, Endpoint)         // 初始化client

	endpointId := "endpoint id"
	args := &endpoint.UpdateEndpointESGArgs{
		EnterpriseSecurityGroupIds: []string{"esg-nhwrebdqi4q2"}, // 替换为需要绑定的 enterprise security group 列表
	}

	err := client.UpdateEndpointEnterpriseSecurityGroup(endpointId, args) // 更新endpoint绑定的企业安全组
	if err != nil {
		fmt.Println(err)
	}
}
