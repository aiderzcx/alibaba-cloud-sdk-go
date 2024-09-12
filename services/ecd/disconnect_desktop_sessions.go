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

// DisconnectDesktopSessions invokes the ecd.DisconnectDesktopSessions API synchronously
func (client *Client) DisconnectDesktopSessions(request *DisconnectDesktopSessionsRequest) (response *DisconnectDesktopSessionsResponse, err error) {
	response = CreateDisconnectDesktopSessionsResponse()
	err = client.DoAction(request, response)
	return
}

// DisconnectDesktopSessionsWithChan invokes the ecd.DisconnectDesktopSessions API asynchronously
func (client *Client) DisconnectDesktopSessionsWithChan(request *DisconnectDesktopSessionsRequest) (<-chan *DisconnectDesktopSessionsResponse, <-chan error) {
	responseChan := make(chan *DisconnectDesktopSessionsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DisconnectDesktopSessions(request)
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

// DisconnectDesktopSessionsWithCallback invokes the ecd.DisconnectDesktopSessions API asynchronously
func (client *Client) DisconnectDesktopSessionsWithCallback(request *DisconnectDesktopSessionsRequest, callback func(response *DisconnectDesktopSessionsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DisconnectDesktopSessionsResponse
		var err error
		defer close(result)
		response, err = client.DisconnectDesktopSessions(request)
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

// DisconnectDesktopSessionsRequest is the request struct for api DisconnectDesktopSessions
type DisconnectDesktopSessionsRequest struct {
	*requests.RpcRequest
	PreCheck requests.Boolean                     `position:"Query" name:"PreCheck"`
	Sessions *[]DisconnectDesktopSessionsSessions `position:"Query" name:"Sessions"  type:"Repeated"`
}

// DisconnectDesktopSessionsSessions is a repeated param struct in DisconnectDesktopSessionsRequest
type DisconnectDesktopSessionsSessions struct {
	EndUserId string `name:"EndUserId"`
	DesktopId string `name:"DesktopId"`
}

// DisconnectDesktopSessionsResponse is the response struct for api DisconnectDesktopSessions
type DisconnectDesktopSessionsResponse struct {
	*responses.BaseResponse
	RequestId       string           `json:"RequestId" xml:"RequestId"`
	InvalidSessions []InvalidSession `json:"InvalidSessions" xml:"InvalidSessions"`
}

// CreateDisconnectDesktopSessionsRequest creates a request to invoke DisconnectDesktopSessions API
func CreateDisconnectDesktopSessionsRequest() (request *DisconnectDesktopSessionsRequest) {
	request = &DisconnectDesktopSessionsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "DisconnectDesktopSessions", "gwsecd", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDisconnectDesktopSessionsResponse creates a response to parse from DisconnectDesktopSessions response
func CreateDisconnectDesktopSessionsResponse() (response *DisconnectDesktopSessionsResponse) {
	response = &DisconnectDesktopSessionsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
