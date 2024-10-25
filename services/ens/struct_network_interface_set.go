package ens

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// NetworkInterfaceSet is a nested struct in ens response
type NetworkInterfaceSet struct {
	CreationTime         string                                      `json:"CreationTime" xml:"CreationTime"`
	EnsRegionId          string                                      `json:"EnsRegionId" xml:"EnsRegionId"`
	InstanceId           string                                      `json:"InstanceId" xml:"InstanceId"`
	InstanceName         string                                      `json:"InstanceName" xml:"InstanceName"`
	MacAddress           string                                      `json:"MacAddress" xml:"MacAddress"`
	NetworkInterfaceId   string                                      `json:"NetworkInterfaceId" xml:"NetworkInterfaceId"`
	PrimaryIp            string                                      `json:"PrimaryIp" xml:"PrimaryIp"`
	PrimaryIpType        string                                      `json:"PrimaryIpType" xml:"PrimaryIpType"`
	Status               string                                      `json:"Status" xml:"Status"`
	NetworkId            string                                      `json:"NetworkId" xml:"NetworkId"`
	VSwitchId            string                                      `json:"VSwitchId" xml:"VSwitchId"`
	NetworkInterfaceName string                                      `json:"NetworkInterfaceName" xml:"NetworkInterfaceName"`
	Description          string                                      `json:"Description" xml:"Description"`
	Type                 string                                      `json:"Type" xml:"Type"`
	SecurityGroupIds     SecurityGroupIdsInDescribeNetworkInterfaces `json:"SecurityGroupIds" xml:"SecurityGroupIds"`
	PrivateIpSets        PrivateIpSetsInDescribeNetworkInterfaces    `json:"PrivateIpSets" xml:"PrivateIpSets"`
	Ipv6Sets             Ipv6SetsInDescribeNetworkInterfaces         `json:"Ipv6Sets" xml:"Ipv6Sets"`
}
