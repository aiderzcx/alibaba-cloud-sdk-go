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

// DeleteUserGroupMembers invokes the quickbi_public.DeleteUserGroupMembers API synchronously
func (client *Client) DeleteUserGroupMembers(request *DeleteUserGroupMembersRequest) (response *DeleteUserGroupMembersResponse, err error) {
	response = CreateDeleteUserGroupMembersResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteUserGroupMembersWithChan invokes the quickbi_public.DeleteUserGroupMembers API asynchronously
func (client *Client) DeleteUserGroupMembersWithChan(request *DeleteUserGroupMembersRequest) (<-chan *DeleteUserGroupMembersResponse, <-chan error) {
	responseChan := make(chan *DeleteUserGroupMembersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteUserGroupMembers(request)
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

// DeleteUserGroupMembersWithCallback invokes the quickbi_public.DeleteUserGroupMembers API asynchronously
func (client *Client) DeleteUserGroupMembersWithCallback(request *DeleteUserGroupMembersRequest, callback func(response *DeleteUserGroupMembersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteUserGroupMembersResponse
		var err error
		defer close(result)
		response, err = client.DeleteUserGroupMembers(request)
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

// DeleteUserGroupMembersRequest is the request struct for api DeleteUserGroupMembers
type DeleteUserGroupMembersRequest struct {
	*requests.RpcRequest
	AccessPoint  string `position:"Query" name:"AccessPoint"`
	SignType     string `position:"Query" name:"SignType"`
	UserGroupIds string `position:"Query" name:"UserGroupIds"`
	UserId       string `position:"Query" name:"UserId"`
}

// DeleteUserGroupMembersResponse is the response struct for api DeleteUserGroupMembers
type DeleteUserGroupMembersResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    bool   `json:"Result" xml:"Result"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateDeleteUserGroupMembersRequest creates a request to invoke DeleteUserGroupMembers API
func CreateDeleteUserGroupMembersRequest() (request *DeleteUserGroupMembersRequest) {
	request = &DeleteUserGroupMembersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("quickbi-public", "2022-01-01", "DeleteUserGroupMembers", "quick", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteUserGroupMembersResponse creates a response to parse from DeleteUserGroupMembers response
func CreateDeleteUserGroupMembersResponse() (response *DeleteUserGroupMembersResponse) {
	response = &DeleteUserGroupMembersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
