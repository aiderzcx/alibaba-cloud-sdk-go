package schedulerx2

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

// SetWfInstanceSuccess invokes the schedulerx2.SetWfInstanceSuccess API synchronously
func (client *Client) SetWfInstanceSuccess(request *SetWfInstanceSuccessRequest) (response *SetWfInstanceSuccessResponse, err error) {
	response = CreateSetWfInstanceSuccessResponse()
	err = client.DoAction(request, response)
	return
}

// SetWfInstanceSuccessWithChan invokes the schedulerx2.SetWfInstanceSuccess API asynchronously
func (client *Client) SetWfInstanceSuccessWithChan(request *SetWfInstanceSuccessRequest) (<-chan *SetWfInstanceSuccessResponse, <-chan error) {
	responseChan := make(chan *SetWfInstanceSuccessResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetWfInstanceSuccess(request)
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

// SetWfInstanceSuccessWithCallback invokes the schedulerx2.SetWfInstanceSuccess API asynchronously
func (client *Client) SetWfInstanceSuccessWithCallback(request *SetWfInstanceSuccessRequest, callback func(response *SetWfInstanceSuccessResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetWfInstanceSuccessResponse
		var err error
		defer close(result)
		response, err = client.SetWfInstanceSuccess(request)
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

// SetWfInstanceSuccessRequest is the request struct for api SetWfInstanceSuccess
type SetWfInstanceSuccessRequest struct {
	*requests.RpcRequest
	NamespaceSource string           `position:"Query" name:"NamespaceSource"`
	GroupId         string           `position:"Query" name:"GroupId"`
	WfInstanceId    requests.Integer `position:"Query" name:"WfInstanceId"`
	Namespace       string           `position:"Query" name:"Namespace"`
	WorkflowId      requests.Integer `position:"Query" name:"WorkflowId"`
}

// SetWfInstanceSuccessResponse is the response struct for api SetWfInstanceSuccess
type SetWfInstanceSuccessResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      int    `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateSetWfInstanceSuccessRequest creates a request to invoke SetWfInstanceSuccess API
func CreateSetWfInstanceSuccessRequest() (request *SetWfInstanceSuccessRequest) {
	request = &SetWfInstanceSuccessRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("schedulerx2", "2019-04-30", "SetWfInstanceSuccess", "", "")
	request.Method = requests.POST
	return
}

// CreateSetWfInstanceSuccessResponse creates a response to parse from SetWfInstanceSuccess response
func CreateSetWfInstanceSuccessResponse() (response *SetWfInstanceSuccessResponse) {
	response = &SetWfInstanceSuccessResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
