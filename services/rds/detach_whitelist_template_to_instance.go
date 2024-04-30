package rds

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

// DetachWhitelistTemplateToInstance invokes the rds.DetachWhitelistTemplateToInstance API synchronously
func (client *Client) DetachWhitelistTemplateToInstance(request *DetachWhitelistTemplateToInstanceRequest) (response *DetachWhitelistTemplateToInstanceResponse, err error) {
	response = CreateDetachWhitelistTemplateToInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// DetachWhitelistTemplateToInstanceWithChan invokes the rds.DetachWhitelistTemplateToInstance API asynchronously
func (client *Client) DetachWhitelistTemplateToInstanceWithChan(request *DetachWhitelistTemplateToInstanceRequest) (<-chan *DetachWhitelistTemplateToInstanceResponse, <-chan error) {
	responseChan := make(chan *DetachWhitelistTemplateToInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DetachWhitelistTemplateToInstance(request)
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

// DetachWhitelistTemplateToInstanceWithCallback invokes the rds.DetachWhitelistTemplateToInstance API asynchronously
func (client *Client) DetachWhitelistTemplateToInstanceWithCallback(request *DetachWhitelistTemplateToInstanceRequest, callback func(response *DetachWhitelistTemplateToInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DetachWhitelistTemplateToInstanceResponse
		var err error
		defer close(result)
		response, err = client.DetachWhitelistTemplateToInstance(request)
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

// DetachWhitelistTemplateToInstanceRequest is the request struct for api DetachWhitelistTemplateToInstance
type DetachWhitelistTemplateToInstanceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	TemplateId           requests.Integer `position:"Query" name:"TemplateId"`
	ResourceGroupId      string           `position:"Query" name:"ResourceGroupId"`
	InsName              string           `position:"Query" name:"InsName"`
}

// DetachWhitelistTemplateToInstanceResponse is the response struct for api DetachWhitelistTemplateToInstance
type DetachWhitelistTemplateToInstanceResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Data           Data   `json:"Data" xml:"Data"`
}

// CreateDetachWhitelistTemplateToInstanceRequest creates a request to invoke DetachWhitelistTemplateToInstance API
func CreateDetachWhitelistTemplateToInstanceRequest() (request *DetachWhitelistTemplateToInstanceRequest) {
	request = &DetachWhitelistTemplateToInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DetachWhitelistTemplateToInstance", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDetachWhitelistTemplateToInstanceResponse creates a response to parse from DetachWhitelistTemplateToInstance response
func CreateDetachWhitelistTemplateToInstanceResponse() (response *DetachWhitelistTemplateToInstanceResponse) {
	response = &DetachWhitelistTemplateToInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
