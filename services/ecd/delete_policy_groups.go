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

// DeletePolicyGroups invokes the ecd.DeletePolicyGroups API synchronously
func (client *Client) DeletePolicyGroups(request *DeletePolicyGroupsRequest) (response *DeletePolicyGroupsResponse, err error) {
	response = CreateDeletePolicyGroupsResponse()
	err = client.DoAction(request, response)
	return
}

// DeletePolicyGroupsWithChan invokes the ecd.DeletePolicyGroups API asynchronously
func (client *Client) DeletePolicyGroupsWithChan(request *DeletePolicyGroupsRequest) (<-chan *DeletePolicyGroupsResponse, <-chan error) {
	responseChan := make(chan *DeletePolicyGroupsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeletePolicyGroups(request)
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

// DeletePolicyGroupsWithCallback invokes the ecd.DeletePolicyGroups API asynchronously
func (client *Client) DeletePolicyGroupsWithCallback(request *DeletePolicyGroupsRequest, callback func(response *DeletePolicyGroupsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeletePolicyGroupsResponse
		var err error
		defer close(result)
		response, err = client.DeletePolicyGroups(request)
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

// DeletePolicyGroupsRequest is the request struct for api DeletePolicyGroups
type DeletePolicyGroupsRequest struct {
	*requests.RpcRequest
	PolicyGroupId *[]string `position:"Query" name:"PolicyGroupId"  type:"Repeated"`
}

// DeletePolicyGroupsResponse is the response struct for api DeletePolicyGroups
type DeletePolicyGroupsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeletePolicyGroupsRequest creates a request to invoke DeletePolicyGroups API
func CreateDeletePolicyGroupsRequest() (request *DeletePolicyGroupsRequest) {
	request = &DeletePolicyGroupsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "DeletePolicyGroups", "gwsecd", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeletePolicyGroupsResponse creates a response to parse from DeletePolicyGroups response
func CreateDeletePolicyGroupsResponse() (response *DeletePolicyGroupsResponse) {
	response = &DeletePolicyGroupsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
