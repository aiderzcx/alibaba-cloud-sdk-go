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

// ListDataLevelPermissionWhiteList invokes the quickbi_public.ListDataLevelPermissionWhiteList API synchronously
func (client *Client) ListDataLevelPermissionWhiteList(request *ListDataLevelPermissionWhiteListRequest) (response *ListDataLevelPermissionWhiteListResponse, err error) {
	response = CreateListDataLevelPermissionWhiteListResponse()
	err = client.DoAction(request, response)
	return
}

// ListDataLevelPermissionWhiteListWithChan invokes the quickbi_public.ListDataLevelPermissionWhiteList API asynchronously
func (client *Client) ListDataLevelPermissionWhiteListWithChan(request *ListDataLevelPermissionWhiteListRequest) (<-chan *ListDataLevelPermissionWhiteListResponse, <-chan error) {
	responseChan := make(chan *ListDataLevelPermissionWhiteListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListDataLevelPermissionWhiteList(request)
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

// ListDataLevelPermissionWhiteListWithCallback invokes the quickbi_public.ListDataLevelPermissionWhiteList API asynchronously
func (client *Client) ListDataLevelPermissionWhiteListWithCallback(request *ListDataLevelPermissionWhiteListRequest, callback func(response *ListDataLevelPermissionWhiteListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListDataLevelPermissionWhiteListResponse
		var err error
		defer close(result)
		response, err = client.ListDataLevelPermissionWhiteList(request)
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

// ListDataLevelPermissionWhiteListRequest is the request struct for api ListDataLevelPermissionWhiteList
type ListDataLevelPermissionWhiteListRequest struct {
	*requests.RpcRequest
	RuleType    string `position:"Query" name:"RuleType"`
	AccessPoint string `position:"Query" name:"AccessPoint"`
	SignType    string `position:"Query" name:"SignType"`
	CubeId      string `position:"Query" name:"CubeId"`
}

// ListDataLevelPermissionWhiteListResponse is the response struct for api ListDataLevelPermissionWhiteList
type ListDataLevelPermissionWhiteListResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateListDataLevelPermissionWhiteListRequest creates a request to invoke ListDataLevelPermissionWhiteList API
func CreateListDataLevelPermissionWhiteListRequest() (request *ListDataLevelPermissionWhiteListRequest) {
	request = &ListDataLevelPermissionWhiteListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("quickbi-public", "2022-01-01", "ListDataLevelPermissionWhiteList", "quick", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListDataLevelPermissionWhiteListResponse creates a response to parse from ListDataLevelPermissionWhiteList response
func CreateListDataLevelPermissionWhiteListResponse() (response *ListDataLevelPermissionWhiteListResponse) {
	response = &ListDataLevelPermissionWhiteListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
