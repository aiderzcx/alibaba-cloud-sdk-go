package schedulerx2

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

// CreateNamespace invokes the schedulerx2.CreateNamespace API synchronously
func (client *Client) CreateNamespace(request *CreateNamespaceRequest) (response *CreateNamespaceResponse, err error) {
	response = CreateCreateNamespaceResponse()
	err = client.DoAction(request, response)
	return
}

// CreateNamespaceWithChan invokes the schedulerx2.CreateNamespace API asynchronously
func (client *Client) CreateNamespaceWithChan(request *CreateNamespaceRequest) (<-chan *CreateNamespaceResponse, <-chan error) {
	responseChan := make(chan *CreateNamespaceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateNamespace(request)
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

// CreateNamespaceWithCallback invokes the schedulerx2.CreateNamespace API asynchronously
func (client *Client) CreateNamespaceWithCallback(request *CreateNamespaceRequest, callback func(response *CreateNamespaceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateNamespaceResponse
		var err error
		defer close(result)
		response, err = client.CreateNamespace(request)
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

// CreateNamespaceRequest is the request struct for api CreateNamespace
type CreateNamespaceRequest struct {
	*requests.RpcRequest
	Description string `position:"Query" name:"Description"`
	Source      string `position:"Query" name:"Source"`
	Uid         string `position:"Query" name:"Uid"`
	Name        string `position:"Query" name:"Name"`
}

// CreateNamespaceResponse is the response struct for api CreateNamespace
type CreateNamespaceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateCreateNamespaceRequest creates a request to invoke CreateNamespace API
func CreateCreateNamespaceRequest() (request *CreateNamespaceRequest) {
	request = &CreateNamespaceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("schedulerx2", "2019-04-30", "CreateNamespace", "schedulerx2", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateNamespaceResponse creates a response to parse from CreateNamespace response
func CreateCreateNamespaceResponse() (response *CreateNamespaceResponse) {
	response = &CreateNamespaceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
