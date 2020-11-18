package ga

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

// ListAccelerators invokes the ga.ListAccelerators API synchronously
func (client *Client) ListAccelerators(request *ListAcceleratorsRequest) (response *ListAcceleratorsResponse, err error) {
	response = CreateListAcceleratorsResponse()
	err = client.DoAction(request, response)
	return
}

// ListAcceleratorsWithChan invokes the ga.ListAccelerators API asynchronously
func (client *Client) ListAcceleratorsWithChan(request *ListAcceleratorsRequest) (<-chan *ListAcceleratorsResponse, <-chan error) {
	responseChan := make(chan *ListAcceleratorsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAccelerators(request)
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

// ListAcceleratorsWithCallback invokes the ga.ListAccelerators API asynchronously
func (client *Client) ListAcceleratorsWithCallback(request *ListAcceleratorsRequest, callback func(response *ListAcceleratorsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAcceleratorsResponse
		var err error
		defer close(result)
		response, err = client.ListAccelerators(request)
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

// ListAcceleratorsRequest is the request struct for api ListAccelerators
type ListAcceleratorsRequest struct {
	*requests.RpcRequest
	PageNumber    requests.Integer `position:"Query" name:"PageNumber"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	AcceleratorId string           `position:"Query" name:"AcceleratorId"`
}

// ListAcceleratorsResponse is the response struct for api ListAccelerators
type ListAcceleratorsResponse struct {
	*responses.BaseResponse
	RequestId    string             `json:"RequestId" xml:"RequestId"`
	TotalCount   int                `json:"TotalCount" xml:"TotalCount"`
	PageNumber   int                `json:"PageNumber" xml:"PageNumber"`
	PageSize     int                `json:"PageSize" xml:"PageSize"`
	Accelerators []AcceleratorsItem `json:"Accelerators" xml:"Accelerators"`
}

// CreateListAcceleratorsRequest creates a request to invoke ListAccelerators API
func CreateListAcceleratorsRequest() (request *ListAcceleratorsRequest) {
	request = &ListAcceleratorsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ga", "2019-11-20", "ListAccelerators", "gaplus", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListAcceleratorsResponse creates a response to parse from ListAccelerators response
func CreateListAcceleratorsResponse() (response *ListAcceleratorsResponse) {
	response = &ListAcceleratorsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
