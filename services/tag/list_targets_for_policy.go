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

// ListTargetsForPolicy invokes the tag.ListTargetsForPolicy API synchronously
func (client *Client) ListTargetsForPolicy(request *ListTargetsForPolicyRequest) (response *ListTargetsForPolicyResponse, err error) {
	response = CreateListTargetsForPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// ListTargetsForPolicyWithChan invokes the tag.ListTargetsForPolicy API asynchronously
func (client *Client) ListTargetsForPolicyWithChan(request *ListTargetsForPolicyRequest) (<-chan *ListTargetsForPolicyResponse, <-chan error) {
	responseChan := make(chan *ListTargetsForPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListTargetsForPolicy(request)
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

// ListTargetsForPolicyWithCallback invokes the tag.ListTargetsForPolicy API asynchronously
func (client *Client) ListTargetsForPolicyWithCallback(request *ListTargetsForPolicyRequest, callback func(response *ListTargetsForPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListTargetsForPolicyResponse
		var err error
		defer close(result)
		response, err = client.ListTargetsForPolicy(request)
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

// ListTargetsForPolicyRequest is the request struct for api ListTargetsForPolicy
type ListTargetsForPolicyRequest struct {
	*requests.RpcRequest
	PolicyId             string           `position:"Query" name:"PolicyId"`
	NextToken            string           `position:"Query" name:"NextToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	MaxResult            requests.Integer `position:"Query" name:"MaxResult"`
}

// ListTargetsForPolicyResponse is the response struct for api ListTargetsForPolicy
type ListTargetsForPolicyResponse struct {
	*responses.BaseResponse
	IsRd      bool     `json:"IsRd" xml:"IsRd"`
	NextToken string   `json:"NextToken" xml:"NextToken"`
	RdId      string   `json:"RdId" xml:"RdId"`
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Targets   []Target `json:"Targets" xml:"Targets"`
}

// CreateListTargetsForPolicyRequest creates a request to invoke ListTargetsForPolicy API
func CreateListTargetsForPolicyRequest() (request *ListTargetsForPolicyRequest) {
	request = &ListTargetsForPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Tag", "2018-08-28", "ListTargetsForPolicy", "tag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListTargetsForPolicyResponse creates a response to parse from ListTargetsForPolicy response
func CreateListTargetsForPolicyResponse() (response *ListTargetsForPolicyResponse) {
	response = &ListTargetsForPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
