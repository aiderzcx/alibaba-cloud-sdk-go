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

// ListListenersByConfig invokes the mse.ListListenersByConfig API synchronously
func (client *Client) ListListenersByConfig(request *ListListenersByConfigRequest) (response *ListListenersByConfigResponse, err error) {
	response = CreateListListenersByConfigResponse()
	err = client.DoAction(request, response)
	return
}

// ListListenersByConfigWithChan invokes the mse.ListListenersByConfig API asynchronously
func (client *Client) ListListenersByConfigWithChan(request *ListListenersByConfigRequest) (<-chan *ListListenersByConfigResponse, <-chan error) {
	responseChan := make(chan *ListListenersByConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListListenersByConfig(request)
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

// ListListenersByConfigWithCallback invokes the mse.ListListenersByConfig API asynchronously
func (client *Client) ListListenersByConfigWithCallback(request *ListListenersByConfigRequest, callback func(response *ListListenersByConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListListenersByConfigResponse
		var err error
		defer close(result)
		response, err = client.ListListenersByConfig(request)
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

// ListListenersByConfigRequest is the request struct for api ListListenersByConfig
type ListListenersByConfigRequest struct {
	*requests.RpcRequest
	MseSessionId   string `position:"Query" name:"MseSessionId"`
	InstanceId     string `position:"Query" name:"InstanceId"`
	DataId         string `position:"Query" name:"DataId"`
	NamespaceId    string `position:"Query" name:"NamespaceId"`
	RequestPars    string `position:"Query" name:"RequestPars"`
	AcceptLanguage string `position:"Query" name:"AcceptLanguage"`
	Group          string `position:"Query" name:"Group"`
}

// ListListenersByConfigResponse is the response struct for api ListListenersByConfig
type ListListenersByConfigResponse struct {
	*responses.BaseResponse
	HttpCode   string     `json:"HttpCode" xml:"HttpCode"`
	TotalCount int        `json:"TotalCount" xml:"TotalCount"`
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	Message    string     `json:"Message" xml:"Message"`
	PageSize   int        `json:"PageSize" xml:"PageSize"`
	PageNumber int        `json:"PageNumber" xml:"PageNumber"`
	ErrorCode  string     `json:"ErrorCode" xml:"ErrorCode"`
	Success    bool       `json:"Success" xml:"Success"`
	Listeners  []Listener `json:"Listeners" xml:"Listeners"`
}

// CreateListListenersByConfigRequest creates a request to invoke ListListenersByConfig API
func CreateListListenersByConfigRequest() (request *ListListenersByConfigRequest) {
	request = &ListListenersByConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "ListListenersByConfig", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListListenersByConfigResponse creates a response to parse from ListListenersByConfig response
func CreateListListenersByConfigResponse() (response *ListListenersByConfigResponse) {
	response = &ListListenersByConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
