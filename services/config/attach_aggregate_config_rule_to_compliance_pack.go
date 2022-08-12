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

// AttachAggregateConfigRuleToCompliancePack invokes the config.AttachAggregateConfigRuleToCompliancePack API synchronously
func (client *Client) AttachAggregateConfigRuleToCompliancePack(request *AttachAggregateConfigRuleToCompliancePackRequest) (response *AttachAggregateConfigRuleToCompliancePackResponse, err error) {
	response = CreateAttachAggregateConfigRuleToCompliancePackResponse()
	err = client.DoAction(request, response)
	return
}

// AttachAggregateConfigRuleToCompliancePackWithChan invokes the config.AttachAggregateConfigRuleToCompliancePack API asynchronously
func (client *Client) AttachAggregateConfigRuleToCompliancePackWithChan(request *AttachAggregateConfigRuleToCompliancePackRequest) (<-chan *AttachAggregateConfigRuleToCompliancePackResponse, <-chan error) {
	responseChan := make(chan *AttachAggregateConfigRuleToCompliancePackResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AttachAggregateConfigRuleToCompliancePack(request)
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

// AttachAggregateConfigRuleToCompliancePackWithCallback invokes the config.AttachAggregateConfigRuleToCompliancePack API asynchronously
func (client *Client) AttachAggregateConfigRuleToCompliancePackWithCallback(request *AttachAggregateConfigRuleToCompliancePackRequest, callback func(response *AttachAggregateConfigRuleToCompliancePackResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AttachAggregateConfigRuleToCompliancePackResponse
		var err error
		defer close(result)
		response, err = client.AttachAggregateConfigRuleToCompliancePack(request)
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

// AttachAggregateConfigRuleToCompliancePackRequest is the request struct for api AttachAggregateConfigRuleToCompliancePack
type AttachAggregateConfigRuleToCompliancePackRequest struct {
	*requests.RpcRequest
	ConfigRuleIds    string `position:"Query" name:"ConfigRuleIds"`
	AggregatorId     string `position:"Query" name:"AggregatorId"`
	CompliancePackId string `position:"Query" name:"CompliancePackId"`
}

// AttachAggregateConfigRuleToCompliancePackResponse is the response struct for api AttachAggregateConfigRuleToCompliancePack
type AttachAggregateConfigRuleToCompliancePackResponse struct {
	*responses.BaseResponse
	RequestId         string            `json:"RequestId" xml:"RequestId"`
	OperateRuleResult OperateRuleResult `json:"OperateRuleResult" xml:"OperateRuleResult"`
}

// CreateAttachAggregateConfigRuleToCompliancePackRequest creates a request to invoke AttachAggregateConfigRuleToCompliancePack API
func CreateAttachAggregateConfigRuleToCompliancePackRequest() (request *AttachAggregateConfigRuleToCompliancePackRequest) {
	request = &AttachAggregateConfigRuleToCompliancePackRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Config", "2020-09-07", "AttachAggregateConfigRuleToCompliancePack", "", "")
	request.Method = requests.POST
	return
}

// CreateAttachAggregateConfigRuleToCompliancePackResponse creates a response to parse from AttachAggregateConfigRuleToCompliancePack response
func CreateAttachAggregateConfigRuleToCompliancePackResponse() (response *AttachAggregateConfigRuleToCompliancePackResponse) {
	response = &AttachAggregateConfigRuleToCompliancePackResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
