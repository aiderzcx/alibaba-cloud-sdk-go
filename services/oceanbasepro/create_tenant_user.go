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

// CreateTenantUser invokes the oceanbasepro.CreateTenantUser API synchronously
func (client *Client) CreateTenantUser(request *CreateTenantUserRequest) (response *CreateTenantUserResponse, err error) {
	response = CreateCreateTenantUserResponse()
	err = client.DoAction(request, response)
	return
}

// CreateTenantUserWithChan invokes the oceanbasepro.CreateTenantUser API asynchronously
func (client *Client) CreateTenantUserWithChan(request *CreateTenantUserRequest) (<-chan *CreateTenantUserResponse, <-chan error) {
	responseChan := make(chan *CreateTenantUserResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateTenantUser(request)
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

// CreateTenantUserWithCallback invokes the oceanbasepro.CreateTenantUser API asynchronously
func (client *Client) CreateTenantUserWithCallback(request *CreateTenantUserRequest, callback func(response *CreateTenantUserResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateTenantUserResponse
		var err error
		defer close(result)
		response, err = client.CreateTenantUser(request)
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

// CreateTenantUserRequest is the request struct for api CreateTenantUser
type CreateTenantUserRequest struct {
	*requests.RpcRequest
	Roles        string `position:"Body" name:"Roles"`
	UserType     string `position:"Body" name:"UserType"`
	Description  string `position:"Body" name:"Description"`
	UserPassword string `position:"Body" name:"UserPassword"`
	InstanceId   string `position:"Body" name:"InstanceId"`
	TenantId     string `position:"Body" name:"TenantId"`
	UserName     string `position:"Body" name:"UserName"`
}

// CreateTenantUserResponse is the response struct for api CreateTenantUser
type CreateTenantUserResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	TenantUser []Data `json:"TenantUser" xml:"TenantUser"`
}

// CreateCreateTenantUserRequest creates a request to invoke CreateTenantUser API
func CreateCreateTenantUserRequest() (request *CreateTenantUserRequest) {
	request = &CreateTenantUserRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OceanBasePro", "2019-09-01", "CreateTenantUser", "oceanbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateTenantUserResponse creates a response to parse from CreateTenantUser response
func CreateCreateTenantUserResponse() (response *CreateTenantUserResponse) {
	response = &CreateTenantUserResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
