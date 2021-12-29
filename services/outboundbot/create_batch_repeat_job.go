package outboundbot

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

// CreateBatchRepeatJob invokes the outboundbot.CreateBatchRepeatJob API synchronously
func (client *Client) CreateBatchRepeatJob(request *CreateBatchRepeatJobRequest) (response *CreateBatchRepeatJobResponse, err error) {
	response = CreateCreateBatchRepeatJobResponse()
	err = client.DoAction(request, response)
	return
}

// CreateBatchRepeatJobWithChan invokes the outboundbot.CreateBatchRepeatJob API asynchronously
func (client *Client) CreateBatchRepeatJobWithChan(request *CreateBatchRepeatJobRequest) (<-chan *CreateBatchRepeatJobResponse, <-chan error) {
	responseChan := make(chan *CreateBatchRepeatJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateBatchRepeatJob(request)
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

// CreateBatchRepeatJobWithCallback invokes the outboundbot.CreateBatchRepeatJob API asynchronously
func (client *Client) CreateBatchRepeatJobWithCallback(request *CreateBatchRepeatJobRequest, callback func(response *CreateBatchRepeatJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateBatchRepeatJobResponse
		var err error
		defer close(result)
		response, err = client.CreateBatchRepeatJob(request)
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

// CreateBatchRepeatJobRequest is the request struct for api CreateBatchRepeatJob
type CreateBatchRepeatJobRequest struct {
	*requests.RpcRequest
	RecallStrategyJson string    `position:"Query" name:"RecallStrategyJson"`
	Description        string    `position:"Query" name:"Description"`
	ScriptId           string    `position:"Query" name:"ScriptId"`
	CallingNumber      *[]string `position:"Query" name:"CallingNumber"  type:"Repeated"`
	InstanceId         string    `position:"Query" name:"InstanceId"`
	FilterStatus       string    `position:"Query" name:"FilterStatus"`
	StrategyJson       string    `position:"Query" name:"StrategyJson"`
	Name               string    `position:"Query" name:"Name"`
	SourceGroupId      string    `position:"Query" name:"SourceGroupId"`
}

// CreateBatchRepeatJobResponse is the response struct for api CreateBatchRepeatJob
type CreateBatchRepeatJobResponse struct {
	*responses.BaseResponse
	HttpStatusCode int      `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code           string   `json:"Code" xml:"Code"`
	Message        string   `json:"Message" xml:"Message"`
	RequestId      string   `json:"RequestId" xml:"RequestId"`
	Success        bool     `json:"Success" xml:"Success"`
	JobGroup       JobGroup `json:"JobGroup" xml:"JobGroup"`
}

// CreateCreateBatchRepeatJobRequest creates a request to invoke CreateBatchRepeatJob API
func CreateCreateBatchRepeatJobRequest() (request *CreateBatchRepeatJobRequest) {
	request = &CreateBatchRepeatJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "CreateBatchRepeatJob", "outboundbot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateBatchRepeatJobResponse creates a response to parse from CreateBatchRepeatJob response
func CreateCreateBatchRepeatJobResponse() (response *CreateBatchRepeatJobResponse) {
	response = &CreateBatchRepeatJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
