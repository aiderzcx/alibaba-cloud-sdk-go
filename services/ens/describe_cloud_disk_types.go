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

// DescribeCloudDiskTypes invokes the ens.DescribeCloudDiskTypes API synchronously
func (client *Client) DescribeCloudDiskTypes(request *DescribeCloudDiskTypesRequest) (response *DescribeCloudDiskTypesResponse, err error) {
	response = CreateDescribeCloudDiskTypesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCloudDiskTypesWithChan invokes the ens.DescribeCloudDiskTypes API asynchronously
func (client *Client) DescribeCloudDiskTypesWithChan(request *DescribeCloudDiskTypesRequest) (<-chan *DescribeCloudDiskTypesResponse, <-chan error) {
	responseChan := make(chan *DescribeCloudDiskTypesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCloudDiskTypes(request)
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

// DescribeCloudDiskTypesWithCallback invokes the ens.DescribeCloudDiskTypes API asynchronously
func (client *Client) DescribeCloudDiskTypesWithCallback(request *DescribeCloudDiskTypesRequest, callback func(response *DescribeCloudDiskTypesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCloudDiskTypesResponse
		var err error
		defer close(result)
		response, err = client.DescribeCloudDiskTypes(request)
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

// DescribeCloudDiskTypesRequest is the request struct for api DescribeCloudDiskTypes
type DescribeCloudDiskTypesRequest struct {
	*requests.RpcRequest
	EnsRegionId  string    `position:"Query" name:"EnsRegionId"`
	EnsRegionIds *[]string `position:"Query" name:"EnsRegionIds"  type:"Json"`
}

// DescribeCloudDiskTypesResponse is the response struct for api DescribeCloudDiskTypes
type DescribeCloudDiskTypesResponse struct {
	*responses.BaseResponse
	RequestId        string                                   `json:"RequestId" xml:"RequestId"`
	SupportResources SupportResourcesInDescribeCloudDiskTypes `json:"SupportResources" xml:"SupportResources"`
}

// CreateDescribeCloudDiskTypesRequest creates a request to invoke DescribeCloudDiskTypes API
func CreateDescribeCloudDiskTypesRequest() (request *DescribeCloudDiskTypesRequest) {
	request = &DescribeCloudDiskTypesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "DescribeCloudDiskTypes", "ens", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeCloudDiskTypesResponse creates a response to parse from DescribeCloudDiskTypes response
func CreateDescribeCloudDiskTypesResponse() (response *DescribeCloudDiskTypesResponse) {
	response = &DescribeCloudDiskTypesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
