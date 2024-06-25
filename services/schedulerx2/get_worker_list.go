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

// GetWorkerList invokes the schedulerx2.GetWorkerList API synchronously
func (client *Client) GetWorkerList(request *GetWorkerListRequest) (response *GetWorkerListResponse, err error) {
	response = CreateGetWorkerListResponse()
	err = client.DoAction(request, response)
	return
}

// GetWorkerListWithChan invokes the schedulerx2.GetWorkerList API asynchronously
func (client *Client) GetWorkerListWithChan(request *GetWorkerListRequest) (<-chan *GetWorkerListResponse, <-chan error) {
	responseChan := make(chan *GetWorkerListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetWorkerList(request)
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

// GetWorkerListWithCallback invokes the schedulerx2.GetWorkerList API asynchronously
func (client *Client) GetWorkerListWithCallback(request *GetWorkerListRequest, callback func(response *GetWorkerListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetWorkerListResponse
		var err error
		defer close(result)
		response, err = client.GetWorkerList(request)
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

// GetWorkerListRequest is the request struct for api GetWorkerList
type GetWorkerListRequest struct {
	*requests.RpcRequest
	NamespaceSource string `position:"Query" name:"NamespaceSource"`
	GroupId         string `position:"Query" name:"GroupId"`
	Namespace       string `position:"Query" name:"Namespace"`
}

// GetWorkerListResponse is the response struct for api GetWorkerList
type GetWorkerListResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetWorkerListRequest creates a request to invoke GetWorkerList API
func CreateGetWorkerListRequest() (request *GetWorkerListRequest) {
	request = &GetWorkerListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("schedulerx2", "2019-04-30", "GetWorkerList", "schedulerx2", "openAPI")
	request.Method = requests.GET
	return
}

// CreateGetWorkerListResponse creates a response to parse from GetWorkerList response
func CreateGetWorkerListResponse() (response *GetWorkerListResponse) {
	response = &GetWorkerListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
