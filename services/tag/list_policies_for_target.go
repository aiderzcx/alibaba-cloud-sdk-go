package tag

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

// ListPoliciesForTarget invokes the tag.ListPoliciesForTarget API synchronously
func (client *Client) ListPoliciesForTarget(request *ListPoliciesForTargetRequest) (response *ListPoliciesForTargetResponse, err error) {
	response = CreateListPoliciesForTargetResponse()
	err = client.DoAction(request, response)
	return
}

// ListPoliciesForTargetWithChan invokes the tag.ListPoliciesForTarget API asynchronously
func (client *Client) ListPoliciesForTargetWithChan(request *ListPoliciesForTargetRequest) (<-chan *ListPoliciesForTargetResponse, <-chan error) {
	responseChan := make(chan *ListPoliciesForTargetResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListPoliciesForTarget(request)
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

// ListPoliciesForTargetWithCallback invokes the tag.ListPoliciesForTarget API asynchronously
func (client *Client) ListPoliciesForTargetWithCallback(request *ListPoliciesForTargetRequest, callback func(response *ListPoliciesForTargetResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListPoliciesForTargetResponse
		var err error
		defer close(result)
		response, err = client.ListPoliciesForTarget(request)
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

// ListPoliciesForTargetRequest is the request struct for api ListPoliciesForTarget
type ListPoliciesForTargetRequest struct {
	*requests.RpcRequest
	TargetId             string           `position:"Query" name:"TargetId"`
	TargetType           string           `position:"Query" name:"TargetType"`
	NextToken            string           `position:"Query" name:"NextToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	MaxResult            requests.Integer `position:"Query" name:"MaxResult"`
}

// ListPoliciesForTargetResponse is the response struct for api ListPoliciesForTarget
type ListPoliciesForTargetResponse struct {
	*responses.BaseResponse
	NextToken string     `json:"NextToken" xml:"NextToken"`
	RequestId string     `json:"RequestId" xml:"RequestId"`
	Data      []DataItem `json:"Data" xml:"Data"`
}

// CreateListPoliciesForTargetRequest creates a request to invoke ListPoliciesForTarget API
func CreateListPoliciesForTargetRequest() (request *ListPoliciesForTargetRequest) {
	request = &ListPoliciesForTargetRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Tag", "2018-08-28", "ListPoliciesForTarget", "tag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListPoliciesForTargetResponse creates a response to parse from ListPoliciesForTarget response
func CreateListPoliciesForTargetResponse() (response *ListPoliciesForTargetResponse) {
	response = &ListPoliciesForTargetResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
