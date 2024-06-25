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

// EnableWorkflow invokes the schedulerx2.EnableWorkflow API synchronously
func (client *Client) EnableWorkflow(request *EnableWorkflowRequest) (response *EnableWorkflowResponse, err error) {
	response = CreateEnableWorkflowResponse()
	err = client.DoAction(request, response)
	return
}

// EnableWorkflowWithChan invokes the schedulerx2.EnableWorkflow API asynchronously
func (client *Client) EnableWorkflowWithChan(request *EnableWorkflowRequest) (<-chan *EnableWorkflowResponse, <-chan error) {
	responseChan := make(chan *EnableWorkflowResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EnableWorkflow(request)
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

// EnableWorkflowWithCallback invokes the schedulerx2.EnableWorkflow API asynchronously
func (client *Client) EnableWorkflowWithCallback(request *EnableWorkflowRequest, callback func(response *EnableWorkflowResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EnableWorkflowResponse
		var err error
		defer close(result)
		response, err = client.EnableWorkflow(request)
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

// EnableWorkflowRequest is the request struct for api EnableWorkflow
type EnableWorkflowRequest struct {
	*requests.RpcRequest
	NamespaceSource string           `position:"Query" name:"NamespaceSource"`
	GroupId         string           `position:"Query" name:"GroupId"`
	Namespace       string           `position:"Query" name:"Namespace"`
	WorkflowId      requests.Integer `position:"Query" name:"WorkflowId"`
}

// EnableWorkflowResponse is the response struct for api EnableWorkflow
type EnableWorkflowResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateEnableWorkflowRequest creates a request to invoke EnableWorkflow API
func CreateEnableWorkflowRequest() (request *EnableWorkflowRequest) {
	request = &EnableWorkflowRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("schedulerx2", "2019-04-30", "EnableWorkflow", "schedulerx2", "openAPI")
	request.Method = requests.GET
	return
}

// CreateEnableWorkflowResponse creates a response to parse from EnableWorkflow response
func CreateEnableWorkflowResponse() (response *EnableWorkflowResponse) {
	response = &EnableWorkflowResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
