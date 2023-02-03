package mse

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

// ModifyLosslessRule invokes the mse.ModifyLosslessRule API synchronously
func (client *Client) ModifyLosslessRule(request *ModifyLosslessRuleRequest) (response *ModifyLosslessRuleResponse, err error) {
	response = CreateModifyLosslessRuleResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyLosslessRuleWithChan invokes the mse.ModifyLosslessRule API asynchronously
func (client *Client) ModifyLosslessRuleWithChan(request *ModifyLosslessRuleRequest) (<-chan *ModifyLosslessRuleResponse, <-chan error) {
	responseChan := make(chan *ModifyLosslessRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyLosslessRule(request)
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

// ModifyLosslessRuleWithCallback invokes the mse.ModifyLosslessRule API asynchronously
func (client *Client) ModifyLosslessRuleWithCallback(request *ModifyLosslessRuleRequest, callback func(response *ModifyLosslessRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyLosslessRuleResponse
		var err error
		defer close(result)
		response, err = client.ModifyLosslessRule(request)
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

// ModifyLosslessRuleRequest is the request struct for api ModifyLosslessRule
type ModifyLosslessRuleRequest struct {
	*requests.RpcRequest
	MseSessionId        string           `position:"Query" name:"MseSessionId"`
	DelayTime           requests.Integer `position:"Query" name:"DelayTime"`
	Source              string           `position:"Query" name:"Source"`
	WarmupTime          requests.Integer `position:"Query" name:"WarmupTime"`
	AppName             string           `position:"Query" name:"AppName"`
	Related             requests.Boolean `position:"Query" name:"Related"`
	Enable              requests.Boolean `position:"Query" name:"Enable"`
	Aligned             requests.Boolean `position:"Query" name:"Aligned"`
	ShutdownWaitSeconds requests.Integer `position:"Query" name:"ShutdownWaitSeconds"`
	Notice              requests.Boolean `position:"Query" name:"Notice"`
	LossLessDetail      requests.Boolean `position:"Query" name:"LossLessDetail"`
	FuncType            requests.Integer `position:"Query" name:"FuncType"`
	AppId               string           `position:"Query" name:"AppId"`
	Namespace           string           `position:"Query" name:"Namespace"`
	AcceptLanguage      string           `position:"Query" name:"AcceptLanguage"`
}

// ModifyLosslessRuleResponse is the response struct for api ModifyLosslessRule
type ModifyLosslessRuleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	HttpCode  string `json:"HttpCode" xml:"HttpCode"`
	Message   string `json:"Message" xml:"Message"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateModifyLosslessRuleRequest creates a request to invoke ModifyLosslessRule API
func CreateModifyLosslessRuleRequest() (request *ModifyLosslessRuleRequest) {
	request = &ModifyLosslessRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "ModifyLosslessRule", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyLosslessRuleResponse creates a response to parse from ModifyLosslessRule response
func CreateModifyLosslessRuleResponse() (response *ModifyLosslessRuleResponse) {
	response = &ModifyLosslessRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
