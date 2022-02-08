package iot

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

// GetLoraNodesTask invokes the iot.GetLoraNodesTask API synchronously
func (client *Client) GetLoraNodesTask(request *GetLoraNodesTaskRequest) (response *GetLoraNodesTaskResponse, err error) {
	response = CreateGetLoraNodesTaskResponse()
	err = client.DoAction(request, response)
	return
}

// GetLoraNodesTaskWithChan invokes the iot.GetLoraNodesTask API asynchronously
func (client *Client) GetLoraNodesTaskWithChan(request *GetLoraNodesTaskRequest) (<-chan *GetLoraNodesTaskResponse, <-chan error) {
	responseChan := make(chan *GetLoraNodesTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetLoraNodesTask(request)
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

// GetLoraNodesTaskWithCallback invokes the iot.GetLoraNodesTask API asynchronously
func (client *Client) GetLoraNodesTaskWithCallback(request *GetLoraNodesTaskRequest, callback func(response *GetLoraNodesTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetLoraNodesTaskResponse
		var err error
		defer close(result)
		response, err = client.GetLoraNodesTask(request)
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

// GetLoraNodesTaskRequest is the request struct for api GetLoraNodesTask
type GetLoraNodesTaskRequest struct {
	*requests.RpcRequest
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
	TaskId        string `position:"Query" name:"TaskId"`
	ApiProduct    string `position:"Body" name:"ApiProduct"`
	ApiRevision   string `position:"Body" name:"ApiRevision"`
}

// GetLoraNodesTaskResponse is the response struct for api GetLoraNodesTask
type GetLoraNodesTaskResponse struct {
	*responses.BaseResponse
	RequestId      string                           `json:"RequestId" xml:"RequestId"`
	Success        bool                             `json:"Success" xml:"Success"`
	Code           string                           `json:"Code" xml:"Code"`
	ErrorMessage   string                           `json:"ErrorMessage" xml:"ErrorMessage"`
	TaskId         string                           `json:"TaskId" xml:"TaskId"`
	TaskState      string                           `json:"TaskState" xml:"TaskState"`
	TotalCount     int64                            `json:"TotalCount" xml:"TotalCount"`
	SuccessCount   int64                            `json:"SuccessCount" xml:"SuccessCount"`
	SuccessDevEuis SuccessDevEuisInGetLoraNodesTask `json:"SuccessDevEuis" xml:"SuccessDevEuis"`
}

// CreateGetLoraNodesTaskRequest creates a request to invoke GetLoraNodesTask API
func CreateGetLoraNodesTaskRequest() (request *GetLoraNodesTaskRequest) {
	request = &GetLoraNodesTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "GetLoraNodesTask", "", "")
	request.Method = requests.POST
	return
}

// CreateGetLoraNodesTaskResponse creates a response to parse from GetLoraNodesTask response
func CreateGetLoraNodesTaskResponse() (response *GetLoraNodesTaskResponse) {
	response = &GetLoraNodesTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
