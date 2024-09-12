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

// GetCoordinateTicket invokes the ecd.GetCoordinateTicket API synchronously
func (client *Client) GetCoordinateTicket(request *GetCoordinateTicketRequest) (response *GetCoordinateTicketResponse, err error) {
	response = CreateGetCoordinateTicketResponse()
	err = client.DoAction(request, response)
	return
}

// GetCoordinateTicketWithChan invokes the ecd.GetCoordinateTicket API asynchronously
func (client *Client) GetCoordinateTicketWithChan(request *GetCoordinateTicketRequest) (<-chan *GetCoordinateTicketResponse, <-chan error) {
	responseChan := make(chan *GetCoordinateTicketResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetCoordinateTicket(request)
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

// GetCoordinateTicketWithCallback invokes the ecd.GetCoordinateTicket API asynchronously
func (client *Client) GetCoordinateTicketWithCallback(request *GetCoordinateTicketRequest, callback func(response *GetCoordinateTicketResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetCoordinateTicketResponse
		var err error
		defer close(result)
		response, err = client.GetCoordinateTicket(request)
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

// GetCoordinateTicketRequest is the request struct for api GetCoordinateTicket
type GetCoordinateTicketRequest struct {
	*requests.RpcRequest
	CoId      string `position:"Query" name:"CoId"`
	UserType  string `position:"Query" name:"UserType"`
	EndUserId string `position:"Query" name:"EndUserId"`
	TaskId    string `position:"Query" name:"TaskId"`
}

// GetCoordinateTicketResponse is the response struct for api GetCoordinateTicket
type GetCoordinateTicketResponse struct {
	*responses.BaseResponse
	CoId       string `json:"CoId" xml:"CoId"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
	TaskStatus string `json:"TaskStatus" xml:"TaskStatus"`
	TaskId     string `json:"TaskId" xml:"TaskId"`
	Ticket     string `json:"Ticket" xml:"Ticket"`
}

// CreateGetCoordinateTicketRequest creates a request to invoke GetCoordinateTicket API
func CreateGetCoordinateTicketRequest() (request *GetCoordinateTicketRequest) {
	request = &GetCoordinateTicketRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "GetCoordinateTicket", "gwsecd", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetCoordinateTicketResponse creates a response to parse from GetCoordinateTicket response
func CreateGetCoordinateTicketResponse() (response *GetCoordinateTicketResponse) {
	response = &GetCoordinateTicketResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
