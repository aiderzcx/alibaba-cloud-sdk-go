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

// UpdateImage invokes the mse.UpdateImage API synchronously
func (client *Client) UpdateImage(request *UpdateImageRequest) (response *UpdateImageResponse, err error) {
	response = CreateUpdateImageResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateImageWithChan invokes the mse.UpdateImage API asynchronously
func (client *Client) UpdateImageWithChan(request *UpdateImageRequest) (<-chan *UpdateImageResponse, <-chan error) {
	responseChan := make(chan *UpdateImageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateImage(request)
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

// UpdateImageWithCallback invokes the mse.UpdateImage API asynchronously
func (client *Client) UpdateImageWithCallback(request *UpdateImageRequest, callback func(response *UpdateImageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateImageResponse
		var err error
		defer close(result)
		response, err = client.UpdateImage(request)
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

// UpdateImageRequest is the request struct for api UpdateImage
type UpdateImageRequest struct {
	*requests.RpcRequest
	MseSessionId   string `position:"Query" name:"MseSessionId"`
	ClusterId      string `position:"Query" name:"ClusterId"`
	AcceptLanguage string `position:"Query" name:"AcceptLanguage"`
	VersionCode    string `position:"Query" name:"VersionCode"`
}

// UpdateImageResponse is the response struct for api UpdateImage
type UpdateImageResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateUpdateImageRequest creates a request to invoke UpdateImage API
func CreateUpdateImageRequest() (request *UpdateImageRequest) {
	request = &UpdateImageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "UpdateImage", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateImageResponse creates a response to parse from UpdateImage response
func CreateUpdateImageResponse() (response *UpdateImageResponse) {
	response = &UpdateImageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
