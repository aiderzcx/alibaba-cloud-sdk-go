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

// QuerySharesToUserList invokes the quickbi_public.QuerySharesToUserList API synchronously
func (client *Client) QuerySharesToUserList(request *QuerySharesToUserListRequest) (response *QuerySharesToUserListResponse, err error) {
	response = CreateQuerySharesToUserListResponse()
	err = client.DoAction(request, response)
	return
}

// QuerySharesToUserListWithChan invokes the quickbi_public.QuerySharesToUserList API asynchronously
func (client *Client) QuerySharesToUserListWithChan(request *QuerySharesToUserListRequest) (<-chan *QuerySharesToUserListResponse, <-chan error) {
	responseChan := make(chan *QuerySharesToUserListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QuerySharesToUserList(request)
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

// QuerySharesToUserListWithCallback invokes the quickbi_public.QuerySharesToUserList API asynchronously
func (client *Client) QuerySharesToUserListWithCallback(request *QuerySharesToUserListRequest, callback func(response *QuerySharesToUserListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QuerySharesToUserListResponse
		var err error
		defer close(result)
		response, err = client.QuerySharesToUserList(request)
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

// QuerySharesToUserListRequest is the request struct for api QuerySharesToUserList
type QuerySharesToUserListRequest struct {
	*requests.RpcRequest
	AccessPoint string `position:"Query" name:"AccessPoint"`
	SignType    string `position:"Query" name:"SignType"`
	UserId      string `position:"Query" name:"UserId"`
}

// QuerySharesToUserListResponse is the response struct for api QuerySharesToUserList
type QuerySharesToUserListResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    []Data `json:"Result" xml:"Result"`
}

// CreateQuerySharesToUserListRequest creates a request to invoke QuerySharesToUserList API
func CreateQuerySharesToUserListRequest() (request *QuerySharesToUserListRequest) {
	request = &QuerySharesToUserListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("quickbi-public", "2022-01-01", "QuerySharesToUserList", "quick", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQuerySharesToUserListResponse creates a response to parse from QuerySharesToUserList response
func CreateQuerySharesToUserListResponse() (response *QuerySharesToUserListResponse) {
	response = &QuerySharesToUserListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
