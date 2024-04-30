package rds

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

// ModifyTaskInfo invokes the rds.ModifyTaskInfo API synchronously
func (client *Client) ModifyTaskInfo(request *ModifyTaskInfoRequest) (response *ModifyTaskInfoResponse, err error) {
	response = CreateModifyTaskInfoResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyTaskInfoWithChan invokes the rds.ModifyTaskInfo API asynchronously
func (client *Client) ModifyTaskInfoWithChan(request *ModifyTaskInfoRequest) (<-chan *ModifyTaskInfoResponse, <-chan error) {
	responseChan := make(chan *ModifyTaskInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyTaskInfo(request)
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

// ModifyTaskInfoWithCallback invokes the rds.ModifyTaskInfo API asynchronously
func (client *Client) ModifyTaskInfoWithCallback(request *ModifyTaskInfoRequest, callback func(response *ModifyTaskInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyTaskInfoResponse
		var err error
		defer close(result)
		response, err = client.ModifyTaskInfo(request)
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

// ModifyTaskInfoRequest is the request struct for api ModifyTaskInfo
type ModifyTaskInfoRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	TaskId               string           `position:"Query" name:"TaskId"`
	ResourceOwnerAccount requests.Integer `position:"Query" name:"ResourceOwnerAccount"`
	StepName             string           `position:"Query" name:"StepName"`
	ActionParams         string           `position:"Query" name:"ActionParams"`
	TaskAction           string           `position:"Query" name:"TaskAction"`
}

// ModifyTaskInfoResponse is the response struct for api ModifyTaskInfo
type ModifyTaskInfoResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ErrorTaskId  string `json:"ErrorTaskId" xml:"ErrorTaskId"`
	SuccessCount string `json:"SuccessCount" xml:"SuccessCount"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
}

// CreateModifyTaskInfoRequest creates a request to invoke ModifyTaskInfo API
func CreateModifyTaskInfoRequest() (request *ModifyTaskInfoRequest) {
	request = &ModifyTaskInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "ModifyTaskInfo", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyTaskInfoResponse creates a response to parse from ModifyTaskInfo response
func CreateModifyTaskInfoResponse() (response *ModifyTaskInfoResponse) {
	response = &ModifyTaskInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
