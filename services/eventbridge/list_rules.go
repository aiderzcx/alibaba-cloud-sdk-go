package eventbridge

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

// ListRules invokes the eventbridge.ListRules API synchronously
func (client *Client) ListRules(request *ListRulesRequest) (response *ListRulesResponse, err error) {
	response = CreateListRulesResponse()
	err = client.DoAction(request, response)
	return
}

// ListRulesWithChan invokes the eventbridge.ListRules API asynchronously
func (client *Client) ListRulesWithChan(request *ListRulesRequest) (<-chan *ListRulesResponse, <-chan error) {
	responseChan := make(chan *ListRulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListRules(request)
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

// ListRulesWithCallback invokes the eventbridge.ListRules API asynchronously
func (client *Client) ListRulesWithCallback(request *ListRulesRequest, callback func(response *ListRulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListRulesResponse
		var err error
		defer close(result)
		response, err = client.ListRules(request)
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

// ListRulesRequest is the request struct for api ListRules
type ListRulesRequest struct {
	*requests.RpcRequest
	RuleNamePrefix string           `position:"Query" name:"RuleNamePrefix"`
	EventBusName   string           `position:"Query" name:"EventBusName"`
	NextToken      string           `position:"Query" name:"NextToken"`
	Limit          requests.Integer `position:"Query" name:"Limit"`
}

// ListRulesResponse is the response struct for api ListRules
type ListRulesResponse struct {
	*responses.BaseResponse
	Message   string          `json:"Message" xml:"Message"`
	RequestId string          `json:"RequestId" xml:"RequestId"`
	Code      string          `json:"Code" xml:"Code"`
	Success   bool            `json:"Success" xml:"Success"`
	Data      DataInListRules `json:"Data" xml:"Data"`
}

// CreateListRulesRequest creates a request to invoke ListRules API
func CreateListRulesRequest() (request *ListRulesRequest) {
	request = &ListRulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("eventbridge", "2020-04-01", "ListRules", "", "")
	request.Method = requests.POST
	return
}

// CreateListRulesResponse creates a response to parse from ListRules response
func CreateListRulesResponse() (response *ListRulesResponse) {
	response = &ListRulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
