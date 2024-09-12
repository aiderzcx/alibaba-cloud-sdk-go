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

// SetDesktopGroupTimerStatus invokes the ecd.SetDesktopGroupTimerStatus API synchronously
func (client *Client) SetDesktopGroupTimerStatus(request *SetDesktopGroupTimerStatusRequest) (response *SetDesktopGroupTimerStatusResponse, err error) {
	response = CreateSetDesktopGroupTimerStatusResponse()
	err = client.DoAction(request, response)
	return
}

// SetDesktopGroupTimerStatusWithChan invokes the ecd.SetDesktopGroupTimerStatus API asynchronously
func (client *Client) SetDesktopGroupTimerStatusWithChan(request *SetDesktopGroupTimerStatusRequest) (<-chan *SetDesktopGroupTimerStatusResponse, <-chan error) {
	responseChan := make(chan *SetDesktopGroupTimerStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetDesktopGroupTimerStatus(request)
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

// SetDesktopGroupTimerStatusWithCallback invokes the ecd.SetDesktopGroupTimerStatus API asynchronously
func (client *Client) SetDesktopGroupTimerStatusWithCallback(request *SetDesktopGroupTimerStatusRequest, callback func(response *SetDesktopGroupTimerStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetDesktopGroupTimerStatusResponse
		var err error
		defer close(result)
		response, err = client.SetDesktopGroupTimerStatus(request)
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

// SetDesktopGroupTimerStatusRequest is the request struct for api SetDesktopGroupTimerStatus
type SetDesktopGroupTimerStatusRequest struct {
	*requests.RpcRequest
	TimerType      requests.Integer `position:"Query" name:"TimerType"`
	DesktopGroupId string           `position:"Query" name:"DesktopGroupId"`
	Status         requests.Integer `position:"Query" name:"Status"`
}

// SetDesktopGroupTimerStatusResponse is the response struct for api SetDesktopGroupTimerStatus
type SetDesktopGroupTimerStatusResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetDesktopGroupTimerStatusRequest creates a request to invoke SetDesktopGroupTimerStatus API
func CreateSetDesktopGroupTimerStatusRequest() (request *SetDesktopGroupTimerStatusRequest) {
	request = &SetDesktopGroupTimerStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "SetDesktopGroupTimerStatus", "gwsecd", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSetDesktopGroupTimerStatusResponse creates a response to parse from SetDesktopGroupTimerStatus response
func CreateSetDesktopGroupTimerStatusResponse() (response *SetDesktopGroupTimerStatusResponse) {
	response = &SetDesktopGroupTimerStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
