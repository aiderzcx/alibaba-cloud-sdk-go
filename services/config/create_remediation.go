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

// CreateRemediation invokes the config.CreateRemediation API synchronously
func (client *Client) CreateRemediation(request *CreateRemediationRequest) (response *CreateRemediationResponse, err error) {
	response = CreateCreateRemediationResponse()
	err = client.DoAction(request, response)
	return
}

// CreateRemediationWithChan invokes the config.CreateRemediation API asynchronously
func (client *Client) CreateRemediationWithChan(request *CreateRemediationRequest) (<-chan *CreateRemediationResponse, <-chan error) {
	responseChan := make(chan *CreateRemediationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateRemediation(request)
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

// CreateRemediationWithCallback invokes the config.CreateRemediation API asynchronously
func (client *Client) CreateRemediationWithCallback(request *CreateRemediationRequest, callback func(response *CreateRemediationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateRemediationResponse
		var err error
		defer close(result)
		response, err = client.CreateRemediation(request)
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

// CreateRemediationRequest is the request struct for api CreateRemediation
type CreateRemediationRequest struct {
	*requests.RpcRequest
	ConfigRuleId          string `position:"Body" name:"ConfigRuleId"`
	RemediationType       string `position:"Body" name:"RemediationType"`
	ClientToken           string `position:"Body" name:"ClientToken"`
	SourceType            string `position:"Body" name:"SourceType"`
	RemediationTemplateId string `position:"Body" name:"RemediationTemplateId"`
	Params                string `position:"Body" name:"Params"`
	InvokeType            string `position:"Body" name:"InvokeType"`
}

// CreateRemediationResponse is the response struct for api CreateRemediation
type CreateRemediationResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	RemediationId string `json:"RemediationId" xml:"RemediationId"`
}

// CreateCreateRemediationRequest creates a request to invoke CreateRemediation API
func CreateCreateRemediationRequest() (request *CreateRemediationRequest) {
	request = &CreateRemediationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Config", "2020-09-07", "CreateRemediation", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateRemediationResponse creates a response to parse from CreateRemediation response
func CreateCreateRemediationResponse() (response *CreateRemediationResponse) {
	response = &CreateRemediationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
