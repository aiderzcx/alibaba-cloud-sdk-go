package ccc

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

// SubmitBatchJobs invokes the ccc.SubmitBatchJobs API synchronously
func (client *Client) SubmitBatchJobs(request *SubmitBatchJobsRequest) (response *SubmitBatchJobsResponse, err error) {
	response = CreateSubmitBatchJobsResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitBatchJobsWithChan invokes the ccc.SubmitBatchJobs API asynchronously
func (client *Client) SubmitBatchJobsWithChan(request *SubmitBatchJobsRequest) (<-chan *SubmitBatchJobsResponse, <-chan error) {
	responseChan := make(chan *SubmitBatchJobsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitBatchJobs(request)
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

// SubmitBatchJobsWithCallback invokes the ccc.SubmitBatchJobs API asynchronously
func (client *Client) SubmitBatchJobsWithCallback(request *SubmitBatchJobsRequest, callback func(response *SubmitBatchJobsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitBatchJobsResponse
		var err error
		defer close(result)
		response, err = client.SubmitBatchJobs(request)
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

// SubmitBatchJobsRequest is the request struct for api SubmitBatchJobs
type SubmitBatchJobsRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	JobGroupId string `position:"Query" name:"JobGroupId"`
}

// SubmitBatchJobsResponse is the response struct for api SubmitBatchJobs
type SubmitBatchJobsResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateSubmitBatchJobsRequest creates a request to invoke SubmitBatchJobs API
func CreateSubmitBatchJobsRequest() (request *SubmitBatchJobsRequest) {
	request = &SubmitBatchJobsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "SubmitBatchJobs", "CCC", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSubmitBatchJobsResponse creates a response to parse from SubmitBatchJobs response
func CreateSubmitBatchJobsResponse() (response *SubmitBatchJobsResponse) {
	response = &SubmitBatchJobsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
