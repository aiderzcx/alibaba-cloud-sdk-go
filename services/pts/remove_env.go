package pts

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

// RemoveEnv invokes the pts.RemoveEnv API synchronously
func (client *Client) RemoveEnv(request *RemoveEnvRequest) (response *RemoveEnvResponse, err error) {
	response = CreateRemoveEnvResponse()
	err = client.DoAction(request, response)
	return
}

// RemoveEnvWithChan invokes the pts.RemoveEnv API asynchronously
func (client *Client) RemoveEnvWithChan(request *RemoveEnvRequest) (<-chan *RemoveEnvResponse, <-chan error) {
	responseChan := make(chan *RemoveEnvResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RemoveEnv(request)
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

// RemoveEnvWithCallback invokes the pts.RemoveEnv API asynchronously
func (client *Client) RemoveEnvWithCallback(request *RemoveEnvRequest, callback func(response *RemoveEnvResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RemoveEnvResponse
		var err error
		defer close(result)
		response, err = client.RemoveEnv(request)
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

// RemoveEnvRequest is the request struct for api RemoveEnv
type RemoveEnvRequest struct {
	*requests.RpcRequest
	EnvId string `position:"Query" name:"EnvId"`
}

// RemoveEnvResponse is the response struct for api RemoveEnv
type RemoveEnvResponse struct {
	*responses.BaseResponse
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code           string `json:"Code" xml:"Code"`
	Success        bool   `json:"Success" xml:"Success"`
}

// CreateRemoveEnvRequest creates a request to invoke RemoveEnv API
func CreateRemoveEnvRequest() (request *RemoveEnvRequest) {
	request = &RemoveEnvRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("PTS", "2020-10-20", "RemoveEnv", "", "")
	request.Method = requests.POST
	return
}

// CreateRemoveEnvResponse creates a response to parse from RemoveEnv response
func CreateRemoveEnvResponse() (response *RemoveEnvResponse) {
	response = &RemoveEnvResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
