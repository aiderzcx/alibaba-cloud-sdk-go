package ecd

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

// UnbindUserDesktop invokes the ecd.UnbindUserDesktop API synchronously
func (client *Client) UnbindUserDesktop(request *UnbindUserDesktopRequest) (response *UnbindUserDesktopResponse, err error) {
	response = CreateUnbindUserDesktopResponse()
	err = client.DoAction(request, response)
	return
}

// UnbindUserDesktopWithChan invokes the ecd.UnbindUserDesktop API asynchronously
func (client *Client) UnbindUserDesktopWithChan(request *UnbindUserDesktopRequest) (<-chan *UnbindUserDesktopResponse, <-chan error) {
	responseChan := make(chan *UnbindUserDesktopResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UnbindUserDesktop(request)
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

// UnbindUserDesktopWithCallback invokes the ecd.UnbindUserDesktop API asynchronously
func (client *Client) UnbindUserDesktopWithCallback(request *UnbindUserDesktopRequest, callback func(response *UnbindUserDesktopResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UnbindUserDesktopResponse
		var err error
		defer close(result)
		response, err = client.UnbindUserDesktop(request)
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

// UnbindUserDesktopRequest is the request struct for api UnbindUserDesktop
type UnbindUserDesktopRequest struct {
	*requests.RpcRequest
	Reason          string           `position:"Query" name:"Reason"`
	DesktopAgentIds *[]string        `position:"Query" name:"DesktopAgentIds"  type:"Repeated"`
	DesktopIds      *[]string        `position:"Query" name:"DesktopIds"  type:"Repeated"`
	DesktopGroupId  string           `position:"Query" name:"DesktopGroupId"`
	Force           requests.Boolean `position:"Query" name:"Force"`
	UserDesktopIds  *[]string        `position:"Query" name:"UserDesktopIds"  type:"Repeated"`
}

// UnbindUserDesktopResponse is the response struct for api UnbindUserDesktop
type UnbindUserDesktopResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUnbindUserDesktopRequest creates a request to invoke UnbindUserDesktop API
func CreateUnbindUserDesktopRequest() (request *UnbindUserDesktopRequest) {
	request = &UnbindUserDesktopRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "UnbindUserDesktop", "gwsecd", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUnbindUserDesktopResponse creates a response to parse from UnbindUserDesktop response
func CreateUnbindUserDesktopResponse() (response *UnbindUserDesktopResponse) {
	response = &UnbindUserDesktopResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
