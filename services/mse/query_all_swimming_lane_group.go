package mse

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

// QueryAllSwimmingLaneGroup invokes the mse.QueryAllSwimmingLaneGroup API synchronously
func (client *Client) QueryAllSwimmingLaneGroup(request *QueryAllSwimmingLaneGroupRequest) (response *QueryAllSwimmingLaneGroupResponse, err error) {
	response = CreateQueryAllSwimmingLaneGroupResponse()
	err = client.DoAction(request, response)
	return
}

// QueryAllSwimmingLaneGroupWithChan invokes the mse.QueryAllSwimmingLaneGroup API asynchronously
func (client *Client) QueryAllSwimmingLaneGroupWithChan(request *QueryAllSwimmingLaneGroupRequest) (<-chan *QueryAllSwimmingLaneGroupResponse, <-chan error) {
	responseChan := make(chan *QueryAllSwimmingLaneGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryAllSwimmingLaneGroup(request)
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

// QueryAllSwimmingLaneGroupWithCallback invokes the mse.QueryAllSwimmingLaneGroup API asynchronously
func (client *Client) QueryAllSwimmingLaneGroupWithCallback(request *QueryAllSwimmingLaneGroupRequest, callback func(response *QueryAllSwimmingLaneGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryAllSwimmingLaneGroupResponse
		var err error
		defer close(result)
		response, err = client.QueryAllSwimmingLaneGroup(request)
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

// QueryAllSwimmingLaneGroupRequest is the request struct for api QueryAllSwimmingLaneGroup
type QueryAllSwimmingLaneGroupRequest struct {
	*requests.RpcRequest
	MseSessionId   string `position:"Query" name:"MseSessionId"`
	Namespace      string `position:"Query" name:"Namespace"`
	AcceptLanguage string `position:"Query" name:"AcceptLanguage"`
}

// QueryAllSwimmingLaneGroupResponse is the response struct for api QueryAllSwimmingLaneGroup
type QueryAllSwimmingLaneGroupResponse struct {
	*responses.BaseResponse
}

// CreateQueryAllSwimmingLaneGroupRequest creates a request to invoke QueryAllSwimmingLaneGroup API
func CreateQueryAllSwimmingLaneGroupRequest() (request *QueryAllSwimmingLaneGroupRequest) {
	request = &QueryAllSwimmingLaneGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "QueryAllSwimmingLaneGroup", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryAllSwimmingLaneGroupResponse creates a response to parse from QueryAllSwimmingLaneGroup response
func CreateQueryAllSwimmingLaneGroupResponse() (response *QueryAllSwimmingLaneGroupResponse) {
	response = &QueryAllSwimmingLaneGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
