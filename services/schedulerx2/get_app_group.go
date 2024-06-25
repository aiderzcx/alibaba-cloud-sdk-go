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

// GetAppGroup invokes the schedulerx2.GetAppGroup API synchronously
func (client *Client) GetAppGroup(request *GetAppGroupRequest) (response *GetAppGroupResponse, err error) {
	response = CreateGetAppGroupResponse()
	err = client.DoAction(request, response)
	return
}

// GetAppGroupWithChan invokes the schedulerx2.GetAppGroup API asynchronously
func (client *Client) GetAppGroupWithChan(request *GetAppGroupRequest) (<-chan *GetAppGroupResponse, <-chan error) {
	responseChan := make(chan *GetAppGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAppGroup(request)
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

// GetAppGroupWithCallback invokes the schedulerx2.GetAppGroup API asynchronously
func (client *Client) GetAppGroupWithCallback(request *GetAppGroupRequest, callback func(response *GetAppGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAppGroupResponse
		var err error
		defer close(result)
		response, err = client.GetAppGroup(request)
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

// GetAppGroupRequest is the request struct for api GetAppGroup
type GetAppGroupRequest struct {
	*requests.RpcRequest
	GroupId   string `position:"Query" name:"GroupId"`
	Namespace string `position:"Query" name:"Namespace"`
}

// GetAppGroupResponse is the response struct for api GetAppGroup
type GetAppGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetAppGroupRequest creates a request to invoke GetAppGroup API
func CreateGetAppGroupRequest() (request *GetAppGroupRequest) {
	request = &GetAppGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("schedulerx2", "2019-04-30", "GetAppGroup", "schedulerx2", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetAppGroupResponse creates a response to parse from GetAppGroup response
func CreateGetAppGroupResponse() (response *GetAppGroupResponse) {
	response = &GetAppGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
