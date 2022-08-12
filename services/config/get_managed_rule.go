package config

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

// GetManagedRule invokes the config.GetManagedRule API synchronously
func (client *Client) GetManagedRule(request *GetManagedRuleRequest) (response *GetManagedRuleResponse, err error) {
	response = CreateGetManagedRuleResponse()
	err = client.DoAction(request, response)
	return
}

// GetManagedRuleWithChan invokes the config.GetManagedRule API asynchronously
func (client *Client) GetManagedRuleWithChan(request *GetManagedRuleRequest) (<-chan *GetManagedRuleResponse, <-chan error) {
	responseChan := make(chan *GetManagedRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetManagedRule(request)
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

// GetManagedRuleWithCallback invokes the config.GetManagedRule API asynchronously
func (client *Client) GetManagedRuleWithCallback(request *GetManagedRuleRequest, callback func(response *GetManagedRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetManagedRuleResponse
		var err error
		defer close(result)
		response, err = client.GetManagedRule(request)
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

// GetManagedRuleRequest is the request struct for api GetManagedRule
type GetManagedRuleRequest struct {
	*requests.RpcRequest
	Identifier string `position:"Query" name:"Identifier"`
}

// GetManagedRuleResponse is the response struct for api GetManagedRule
type GetManagedRuleResponse struct {
	*responses.BaseResponse
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	ManagedRule ManagedRule `json:"ManagedRule" xml:"ManagedRule"`
}

// CreateGetManagedRuleRequest creates a request to invoke GetManagedRule API
func CreateGetManagedRuleRequest() (request *GetManagedRuleRequest) {
	request = &GetManagedRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Config", "2020-09-07", "GetManagedRule", "", "")
	request.Method = requests.POST
	return
}

// CreateGetManagedRuleResponse creates a response to parse from GetManagedRule response
func CreateGetManagedRuleResponse() (response *GetManagedRuleResponse) {
	response = &GetManagedRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
