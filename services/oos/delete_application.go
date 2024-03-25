package oos

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

// DeleteApplication invokes the oos.DeleteApplication API synchronously
func (client *Client) DeleteApplication(request *DeleteApplicationRequest) (response *DeleteApplicationResponse, err error) {
	response = CreateDeleteApplicationResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteApplicationWithChan invokes the oos.DeleteApplication API asynchronously
func (client *Client) DeleteApplicationWithChan(request *DeleteApplicationRequest) (<-chan *DeleteApplicationResponse, <-chan error) {
	responseChan := make(chan *DeleteApplicationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteApplication(request)
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

// DeleteApplicationWithCallback invokes the oos.DeleteApplication API asynchronously
func (client *Client) DeleteApplicationWithCallback(request *DeleteApplicationRequest, callback func(response *DeleteApplicationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteApplicationResponse
		var err error
		defer close(result)
		response, err = client.DeleteApplication(request)
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

// DeleteApplicationRequest is the request struct for api DeleteApplication
type DeleteApplicationRequest struct {
	*requests.RpcRequest
	RetainResource requests.Boolean `position:"Query" name:"RetainResource"`
	Name           string           `position:"Query" name:"Name"`
	Force          requests.Boolean `position:"Query" name:"Force"`
}

// DeleteApplicationResponse is the response struct for api DeleteApplication
type DeleteApplicationResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteApplicationRequest creates a request to invoke DeleteApplication API
func CreateDeleteApplicationRequest() (request *DeleteApplicationRequest) {
	request = &DeleteApplicationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("oos", "2019-06-01", "DeleteApplication", "oos", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteApplicationResponse creates a response to parse from DeleteApplication response
func CreateDeleteApplicationResponse() (response *DeleteApplicationResponse) {
	response = &DeleteApplicationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
