package oceanbasepro

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

// DescribeTenant invokes the oceanbasepro.DescribeTenant API synchronously
func (client *Client) DescribeTenant(request *DescribeTenantRequest) (response *DescribeTenantResponse, err error) {
	response = CreateDescribeTenantResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeTenantWithChan invokes the oceanbasepro.DescribeTenant API asynchronously
func (client *Client) DescribeTenantWithChan(request *DescribeTenantRequest) (<-chan *DescribeTenantResponse, <-chan error) {
	responseChan := make(chan *DescribeTenantResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeTenant(request)
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

// DescribeTenantWithCallback invokes the oceanbasepro.DescribeTenant API asynchronously
func (client *Client) DescribeTenantWithCallback(request *DescribeTenantRequest, callback func(response *DescribeTenantResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeTenantResponse
		var err error
		defer close(result)
		response, err = client.DescribeTenant(request)
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

// DescribeTenantRequest is the request struct for api DescribeTenant
type DescribeTenantRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Body" name:"InstanceId"`
	TenantId   string `position:"Body" name:"TenantId"`
}

// DescribeTenantResponse is the response struct for api DescribeTenant
type DescribeTenantResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Tenant    Tenant `json:"Tenant" xml:"Tenant"`
}

// CreateDescribeTenantRequest creates a request to invoke DescribeTenant API
func CreateDescribeTenantRequest() (request *DescribeTenantRequest) {
	request = &DescribeTenantRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OceanBasePro", "2019-09-01", "DescribeTenant", "oceanbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeTenantResponse creates a response to parse from DescribeTenant response
func CreateDescribeTenantResponse() (response *DescribeTenantResponse) {
	response = &DescribeTenantResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
