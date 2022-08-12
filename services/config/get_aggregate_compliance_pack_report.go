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

// GetAggregateCompliancePackReport invokes the config.GetAggregateCompliancePackReport API synchronously
func (client *Client) GetAggregateCompliancePackReport(request *GetAggregateCompliancePackReportRequest) (response *GetAggregateCompliancePackReportResponse, err error) {
	response = CreateGetAggregateCompliancePackReportResponse()
	err = client.DoAction(request, response)
	return
}

// GetAggregateCompliancePackReportWithChan invokes the config.GetAggregateCompliancePackReport API asynchronously
func (client *Client) GetAggregateCompliancePackReportWithChan(request *GetAggregateCompliancePackReportRequest) (<-chan *GetAggregateCompliancePackReportResponse, <-chan error) {
	responseChan := make(chan *GetAggregateCompliancePackReportResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAggregateCompliancePackReport(request)
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

// GetAggregateCompliancePackReportWithCallback invokes the config.GetAggregateCompliancePackReport API asynchronously
func (client *Client) GetAggregateCompliancePackReportWithCallback(request *GetAggregateCompliancePackReportRequest, callback func(response *GetAggregateCompliancePackReportResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAggregateCompliancePackReportResponse
		var err error
		defer close(result)
		response, err = client.GetAggregateCompliancePackReport(request)
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

// GetAggregateCompliancePackReportRequest is the request struct for api GetAggregateCompliancePackReport
type GetAggregateCompliancePackReportRequest struct {
	*requests.RpcRequest
	AggregatorId     string `position:"Query" name:"AggregatorId"`
	CompliancePackId string `position:"Query" name:"CompliancePackId"`
}

// GetAggregateCompliancePackReportResponse is the response struct for api GetAggregateCompliancePackReport
type GetAggregateCompliancePackReportResponse struct {
	*responses.BaseResponse
	RequestId            string               `json:"RequestId" xml:"RequestId"`
	CompliancePackReport CompliancePackReport `json:"CompliancePackReport" xml:"CompliancePackReport"`
}

// CreateGetAggregateCompliancePackReportRequest creates a request to invoke GetAggregateCompliancePackReport API
func CreateGetAggregateCompliancePackReportRequest() (request *GetAggregateCompliancePackReportRequest) {
	request = &GetAggregateCompliancePackReportRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Config", "2020-09-07", "GetAggregateCompliancePackReport", "", "")
	request.Method = requests.GET
	return
}

// CreateGetAggregateCompliancePackReportResponse creates a response to parse from GetAggregateCompliancePackReport response
func CreateGetAggregateCompliancePackReportResponse() (response *GetAggregateCompliancePackReportResponse) {
	response = &GetAggregateCompliancePackReportResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
