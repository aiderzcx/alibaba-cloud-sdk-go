package beian

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

// ManageAccessorDomainWhiteList invokes the beian.ManageAccessorDomainWhiteList API synchronously
func (client *Client) ManageAccessorDomainWhiteList(request *ManageAccessorDomainWhiteListRequest) (response *ManageAccessorDomainWhiteListResponse, err error) {
	response = CreateManageAccessorDomainWhiteListResponse()
	err = client.DoAction(request, response)
	return
}

// ManageAccessorDomainWhiteListWithChan invokes the beian.ManageAccessorDomainWhiteList API asynchronously
func (client *Client) ManageAccessorDomainWhiteListWithChan(request *ManageAccessorDomainWhiteListRequest) (<-chan *ManageAccessorDomainWhiteListResponse, <-chan error) {
	responseChan := make(chan *ManageAccessorDomainWhiteListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ManageAccessorDomainWhiteList(request)
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

// ManageAccessorDomainWhiteListWithCallback invokes the beian.ManageAccessorDomainWhiteList API asynchronously
func (client *Client) ManageAccessorDomainWhiteListWithCallback(request *ManageAccessorDomainWhiteListRequest, callback func(response *ManageAccessorDomainWhiteListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ManageAccessorDomainWhiteListResponse
		var err error
		defer close(result)
		response, err = client.ManageAccessorDomainWhiteList(request)
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

// ManageAccessorDomainWhiteListRequest is the request struct for api ManageAccessorDomainWhiteList
type ManageAccessorDomainWhiteListRequest struct {
	*requests.RpcRequest
	EndTime   string    `position:"Query" name:"EndTime"`
	Domains   *[]string `position:"Query" name:"Domains"  type:"Repeated"`
	Remark    string    `position:"Query" name:"Remark"`
	StartTime string    `position:"Query" name:"StartTime"`
	Caller    string    `position:"Query" name:"Caller"`
	Operation string    `position:"Query" name:"Operation"`
}

// ManageAccessorDomainWhiteListResponse is the response struct for api ManageAccessorDomainWhiteList
type ManageAccessorDomainWhiteListResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateManageAccessorDomainWhiteListRequest creates a request to invoke ManageAccessorDomainWhiteList API
func CreateManageAccessorDomainWhiteListRequest() (request *ManageAccessorDomainWhiteListRequest) {
	request = &ManageAccessorDomainWhiteListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Beian", "2016-08-10", "ManageAccessorDomainWhiteList", "", "")
	request.Method = requests.POST
	return
}

// CreateManageAccessorDomainWhiteListResponse creates a response to parse from ManageAccessorDomainWhiteList response
func CreateManageAccessorDomainWhiteListResponse() (response *ManageAccessorDomainWhiteListResponse) {
	response = &ManageAccessorDomainWhiteListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
