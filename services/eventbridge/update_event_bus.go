package eventbridge

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

// UpdateEventBus invokes the eventbridge.UpdateEventBus API synchronously
func (client *Client) UpdateEventBus(request *UpdateEventBusRequest) (response *UpdateEventBusResponse, err error) {
	response = CreateUpdateEventBusResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateEventBusWithChan invokes the eventbridge.UpdateEventBus API asynchronously
func (client *Client) UpdateEventBusWithChan(request *UpdateEventBusRequest) (<-chan *UpdateEventBusResponse, <-chan error) {
	responseChan := make(chan *UpdateEventBusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateEventBus(request)
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

// UpdateEventBusWithCallback invokes the eventbridge.UpdateEventBus API asynchronously
func (client *Client) UpdateEventBusWithCallback(request *UpdateEventBusRequest, callback func(response *UpdateEventBusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateEventBusResponse
		var err error
		defer close(result)
		response, err = client.UpdateEventBus(request)
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

// UpdateEventBusRequest is the request struct for api UpdateEventBus
type UpdateEventBusRequest struct {
	*requests.RpcRequest
	Description  string `position:"Query" name:"Description"`
	EventBusName string `position:"Query" name:"EventBusName"`
}

// UpdateEventBusResponse is the response struct for api UpdateEventBus
type UpdateEventBusResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      bool   `json:"Data" xml:"Data"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateUpdateEventBusRequest creates a request to invoke UpdateEventBus API
func CreateUpdateEventBusRequest() (request *UpdateEventBusRequest) {
	request = &UpdateEventBusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("eventbridge", "2020-04-01", "UpdateEventBus", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateEventBusResponse creates a response to parse from UpdateEventBus response
func CreateUpdateEventBusResponse() (response *UpdateEventBusResponse) {
	response = &UpdateEventBusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
