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

// ListResourceEvaluationResults invokes the config.ListResourceEvaluationResults API synchronously
func (client *Client) ListResourceEvaluationResults(request *ListResourceEvaluationResultsRequest) (response *ListResourceEvaluationResultsResponse, err error) {
	response = CreateListResourceEvaluationResultsResponse()
	err = client.DoAction(request, response)
	return
}

// ListResourceEvaluationResultsWithChan invokes the config.ListResourceEvaluationResults API asynchronously
func (client *Client) ListResourceEvaluationResultsWithChan(request *ListResourceEvaluationResultsRequest) (<-chan *ListResourceEvaluationResultsResponse, <-chan error) {
	responseChan := make(chan *ListResourceEvaluationResultsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListResourceEvaluationResults(request)
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

// ListResourceEvaluationResultsWithCallback invokes the config.ListResourceEvaluationResults API asynchronously
func (client *Client) ListResourceEvaluationResultsWithCallback(request *ListResourceEvaluationResultsRequest, callback func(response *ListResourceEvaluationResultsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListResourceEvaluationResultsResponse
		var err error
		defer close(result)
		response, err = client.ListResourceEvaluationResults(request)
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

// ListResourceEvaluationResultsRequest is the request struct for api ListResourceEvaluationResults
type ListResourceEvaluationResultsRequest struct {
	*requests.RpcRequest
	ResourceId     string           `position:"Query" name:"ResourceId"`
	ResourceType   string           `position:"Query" name:"ResourceType"`
	NextToken      string           `position:"Query" name:"NextToken"`
	MaxResults     requests.Integer `position:"Query" name:"MaxResults"`
	Region         string           `position:"Query" name:"Region"`
	ComplianceType string           `position:"Query" name:"ComplianceType"`
}

// ListResourceEvaluationResultsResponse is the response struct for api ListResourceEvaluationResults
type ListResourceEvaluationResultsResponse struct {
	*responses.BaseResponse
	RequestId         string            `json:"RequestId" xml:"RequestId"`
	EvaluationResults EvaluationResults `json:"EvaluationResults" xml:"EvaluationResults"`
}

// CreateListResourceEvaluationResultsRequest creates a request to invoke ListResourceEvaluationResults API
func CreateListResourceEvaluationResultsRequest() (request *ListResourceEvaluationResultsRequest) {
	request = &ListResourceEvaluationResultsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Config", "2020-09-07", "ListResourceEvaluationResults", "", "")
	request.Method = requests.POST
	return
}

// CreateListResourceEvaluationResultsResponse creates a response to parse from ListResourceEvaluationResults response
func CreateListResourceEvaluationResultsResponse() (response *ListResourceEvaluationResultsResponse) {
	response = &ListResourceEvaluationResultsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
