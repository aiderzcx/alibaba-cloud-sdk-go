package quickbi_public

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

// QueryTicketInfo invokes the quickbi_public.QueryTicketInfo API synchronously
func (client *Client) QueryTicketInfo(request *QueryTicketInfoRequest) (response *QueryTicketInfoResponse, err error) {
	response = CreateQueryTicketInfoResponse()
	err = client.DoAction(request, response)
	return
}

// QueryTicketInfoWithChan invokes the quickbi_public.QueryTicketInfo API asynchronously
func (client *Client) QueryTicketInfoWithChan(request *QueryTicketInfoRequest) (<-chan *QueryTicketInfoResponse, <-chan error) {
	responseChan := make(chan *QueryTicketInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryTicketInfo(request)
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

// QueryTicketInfoWithCallback invokes the quickbi_public.QueryTicketInfo API asynchronously
func (client *Client) QueryTicketInfoWithCallback(request *QueryTicketInfoRequest, callback func(response *QueryTicketInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryTicketInfoResponse
		var err error
		defer close(result)
		response, err = client.QueryTicketInfo(request)
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

// QueryTicketInfoRequest is the request struct for api QueryTicketInfo
type QueryTicketInfoRequest struct {
	*requests.RpcRequest
	Ticket      string `position:"Query" name:"Ticket"`
	AccessPoint string `position:"Query" name:"AccessPoint"`
	SignType    string `position:"Query" name:"SignType"`
}

// QueryTicketInfoResponse is the response struct for api QueryTicketInfo
type QueryTicketInfoResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateQueryTicketInfoRequest creates a request to invoke QueryTicketInfo API
func CreateQueryTicketInfoRequest() (request *QueryTicketInfoRequest) {
	request = &QueryTicketInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("quickbi-public", "2022-01-01", "QueryTicketInfo", "quick", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryTicketInfoResponse creates a response to parse from QueryTicketInfo response
func CreateQueryTicketInfoResponse() (response *QueryTicketInfoResponse) {
	response = &QueryTicketInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
