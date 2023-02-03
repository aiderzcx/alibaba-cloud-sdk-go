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

// ListSecurityGroup invokes the mse.ListSecurityGroup API synchronously
func (client *Client) ListSecurityGroup(request *ListSecurityGroupRequest) (response *ListSecurityGroupResponse, err error) {
	response = CreateListSecurityGroupResponse()
	err = client.DoAction(request, response)
	return
}

// ListSecurityGroupWithChan invokes the mse.ListSecurityGroup API asynchronously
func (client *Client) ListSecurityGroupWithChan(request *ListSecurityGroupRequest) (<-chan *ListSecurityGroupResponse, <-chan error) {
	responseChan := make(chan *ListSecurityGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListSecurityGroup(request)
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

// ListSecurityGroupWithCallback invokes the mse.ListSecurityGroup API asynchronously
func (client *Client) ListSecurityGroupWithCallback(request *ListSecurityGroupRequest, callback func(response *ListSecurityGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListSecurityGroupResponse
		var err error
		defer close(result)
		response, err = client.ListSecurityGroup(request)
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

// ListSecurityGroupRequest is the request struct for api ListSecurityGroup
type ListSecurityGroupRequest struct {
	*requests.RpcRequest
	MseSessionId    string `position:"Query" name:"MseSessionId"`
	GatewayUniqueId string `position:"Query" name:"GatewayUniqueId"`
	AcceptLanguage  string `position:"Query" name:"AcceptLanguage"`
}

// ListSecurityGroupResponse is the response struct for api ListSecurityGroup
type ListSecurityGroupResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string `json:"Message" xml:"Message"`
	Code           int    `json:"Code" xml:"Code"`
	Success        bool   `json:"Success" xml:"Success"`
	Data           []Sgs  `json:"Data" xml:"Data"`
}

// CreateListSecurityGroupRequest creates a request to invoke ListSecurityGroup API
func CreateListSecurityGroupRequest() (request *ListSecurityGroupRequest) {
	request = &ListSecurityGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "ListSecurityGroup", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListSecurityGroupResponse creates a response to parse from ListSecurityGroup response
func CreateListSecurityGroupResponse() (response *ListSecurityGroupResponse) {
	response = &ListSecurityGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
