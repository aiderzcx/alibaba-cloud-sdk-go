package csas

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

// GetApproval invokes the csas.GetApproval API synchronously
func (client *Client) GetApproval(request *GetApprovalRequest) (response *GetApprovalResponse, err error) {
	response = CreateGetApprovalResponse()
	err = client.DoAction(request, response)
	return
}

// GetApprovalWithChan invokes the csas.GetApproval API asynchronously
func (client *Client) GetApprovalWithChan(request *GetApprovalRequest) (<-chan *GetApprovalResponse, <-chan error) {
	responseChan := make(chan *GetApprovalResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetApproval(request)
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

// GetApprovalWithCallback invokes the csas.GetApproval API asynchronously
func (client *Client) GetApprovalWithCallback(request *GetApprovalRequest, callback func(response *GetApprovalResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetApprovalResponse
		var err error
		defer close(result)
		response, err = client.GetApproval(request)
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

// GetApprovalRequest is the request struct for api GetApproval
type GetApprovalRequest struct {
	*requests.RpcRequest
	SourceIp   string `position:"Query" name:"SourceIp"`
	ApprovalId string `position:"Query" name:"ApprovalId"`
}

// GetApprovalResponse is the response struct for api GetApproval
type GetApprovalResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Approval  []Data `json:"Approval" xml:"Approval"`
}

// CreateGetApprovalRequest creates a request to invoke GetApproval API
func CreateGetApprovalRequest() (request *GetApprovalRequest) {
	request = &GetApprovalRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("csas", "2023-01-20", "GetApproval", "", "")
	request.Method = requests.GET
	return
}

// CreateGetApprovalResponse creates a response to parse from GetApproval response
func CreateGetApprovalResponse() (response *GetApprovalResponse) {
	response = &GetApprovalResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
