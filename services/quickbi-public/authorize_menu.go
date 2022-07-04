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

// AuthorizeMenu invokes the quickbi_public.AuthorizeMenu API synchronously
func (client *Client) AuthorizeMenu(request *AuthorizeMenuRequest) (response *AuthorizeMenuResponse, err error) {
	response = CreateAuthorizeMenuResponse()
	err = client.DoAction(request, response)
	return
}

// AuthorizeMenuWithChan invokes the quickbi_public.AuthorizeMenu API asynchronously
func (client *Client) AuthorizeMenuWithChan(request *AuthorizeMenuRequest) (<-chan *AuthorizeMenuResponse, <-chan error) {
	responseChan := make(chan *AuthorizeMenuResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AuthorizeMenu(request)
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

// AuthorizeMenuWithCallback invokes the quickbi_public.AuthorizeMenu API asynchronously
func (client *Client) AuthorizeMenuWithCallback(request *AuthorizeMenuRequest, callback func(response *AuthorizeMenuResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AuthorizeMenuResponse
		var err error
		defer close(result)
		response, err = client.AuthorizeMenu(request)
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

// AuthorizeMenuRequest is the request struct for api AuthorizeMenu
type AuthorizeMenuRequest struct {
	*requests.RpcRequest
	DataPortalId    string           `position:"Query" name:"DataPortalId"`
	UserIds         string           `position:"Query" name:"UserIds"`
	AuthPointsValue requests.Integer `position:"Query" name:"AuthPointsValue"`
	AccessPoint     string           `position:"Query" name:"AccessPoint"`
	SignType        string           `position:"Query" name:"SignType"`
	UserGroupIds    string           `position:"Query" name:"UserGroupIds"`
	MenuIds         string           `position:"Query" name:"MenuIds"`
}

// AuthorizeMenuResponse is the response struct for api AuthorizeMenu
type AuthorizeMenuResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    int    `json:"Result" xml:"Result"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateAuthorizeMenuRequest creates a request to invoke AuthorizeMenu API
func CreateAuthorizeMenuRequest() (request *AuthorizeMenuRequest) {
	request = &AuthorizeMenuRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("quickbi-public", "2022-01-01", "AuthorizeMenu", "quick", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAuthorizeMenuResponse creates a response to parse from AuthorizeMenu response
func CreateAuthorizeMenuResponse() (response *AuthorizeMenuResponse) {
	response = &AuthorizeMenuResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
