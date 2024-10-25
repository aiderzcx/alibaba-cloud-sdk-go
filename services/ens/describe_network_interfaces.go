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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeNetworkInterfaces invokes the ens.DescribeNetworkInterfaces API synchronously
func (client *Client) DescribeNetworkInterfaces(request *DescribeNetworkInterfacesRequest) (response *DescribeNetworkInterfacesResponse, err error) {
	response = CreateDescribeNetworkInterfacesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeNetworkInterfacesWithChan invokes the ens.DescribeNetworkInterfaces API asynchronously
func (client *Client) DescribeNetworkInterfacesWithChan(request *DescribeNetworkInterfacesRequest) (<-chan *DescribeNetworkInterfacesResponse, <-chan error) {
	responseChan := make(chan *DescribeNetworkInterfacesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeNetworkInterfaces(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeNetworkInterfacesWithCallback invokes the ens.DescribeNetworkInterfaces API asynchronously
func (client *Client) DescribeNetworkInterfacesWithCallback(request *DescribeNetworkInterfacesRequest, callback func(response *DescribeNetworkInterfacesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeNetworkInterfacesResponse
		var err error
		defer close(result)
		response, err = client.DescribeNetworkInterfaces(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeNetworkInterfacesRequest is the request struct for api DescribeNetworkInterfaces
type DescribeNetworkInterfacesRequest struct {
	*requests.RpcRequest
	Type                 string           `position:"Query" name:"Type"`
	EnsRegionId          string           `position:"Query" name:"EnsRegionId"`
	NetworkInterfaceName string           `position:"Query" name:"NetworkInterfaceName"`
	VSwitchId            string           `position:"Query" name:"VSwitchId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	NetworkId            string           `position:"Query" name:"NetworkId"`
	Ipv6Address          *[]string        `position:"Query" name:"Ipv6Address"  type:"Repeated"`
	Status               string           `position:"Query" name:"Status"`
	SecurityGroupId      string           `position:"Query" name:"SecurityGroupId"`
	PageNumber           string           `position:"Query" name:"PageNumber"`
	ShowDetail           requests.Boolean `position:"Query" name:"ShowDetail"`
	PageSize             string           `position:"Query" name:"PageSize"`
	PrimaryIpAddress     string           `position:"Query" name:"PrimaryIpAddress"`
	NetworkInterfaceId   string           `position:"Query" name:"NetworkInterfaceId"`
}

// DescribeNetworkInterfacesResponse is the response struct for api DescribeNetworkInterfaces
type DescribeNetworkInterfacesResponse struct {
	*responses.BaseResponse
	PageNumber           int                  `json:"PageNumber" xml:"PageNumber"`
	PageSize             int                  `json:"PageSize" xml:"PageSize"`
	TotalCount           int                  `json:"TotalCount" xml:"TotalCount"`
	RequestId            string               `json:"RequestId" xml:"RequestId"`
	NetworkInterfaceSets NetworkInterfaceSets `json:"NetworkInterfaceSets" xml:"NetworkInterfaceSets"`
}

// CreateDescribeNetworkInterfacesRequest creates a request to invoke DescribeNetworkInterfaces API
func CreateDescribeNetworkInterfacesRequest() (request *DescribeNetworkInterfacesRequest) {
	request = &DescribeNetworkInterfacesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "DescribeNetworkInterfaces", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeNetworkInterfacesResponse creates a response to parse from DescribeNetworkInterfaces response
func CreateDescribeNetworkInterfacesResponse() (response *DescribeNetworkInterfacesResponse) {
	response = &DescribeNetworkInterfacesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
