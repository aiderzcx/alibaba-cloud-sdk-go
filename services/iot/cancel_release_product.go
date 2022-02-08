package iot

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

// CancelReleaseProduct invokes the iot.CancelReleaseProduct API synchronously
func (client *Client) CancelReleaseProduct(request *CancelReleaseProductRequest) (response *CancelReleaseProductResponse, err error) {
	response = CreateCancelReleaseProductResponse()
	err = client.DoAction(request, response)
	return
}

// CancelReleaseProductWithChan invokes the iot.CancelReleaseProduct API asynchronously
func (client *Client) CancelReleaseProductWithChan(request *CancelReleaseProductRequest) (<-chan *CancelReleaseProductResponse, <-chan error) {
	responseChan := make(chan *CancelReleaseProductResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CancelReleaseProduct(request)
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

// CancelReleaseProductWithCallback invokes the iot.CancelReleaseProduct API asynchronously
func (client *Client) CancelReleaseProductWithCallback(request *CancelReleaseProductRequest, callback func(response *CancelReleaseProductResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CancelReleaseProductResponse
		var err error
		defer close(result)
		response, err = client.CancelReleaseProduct(request)
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

// CancelReleaseProductRequest is the request struct for api CancelReleaseProduct
type CancelReleaseProductRequest struct {
	*requests.RpcRequest
	RealTenantId      string `position:"Query" name:"RealTenantId"`
	RealTripartiteKey string `position:"Query" name:"RealTripartiteKey"`
	IotInstanceId     string `position:"Query" name:"IotInstanceId"`
	ProductKey        string `position:"Query" name:"ProductKey"`
	ApiProduct        string `position:"Body" name:"ApiProduct"`
	ApiRevision       string `position:"Body" name:"ApiRevision"`
	BizTenantId       string `position:"Query" name:"BizTenantId"`
}

// CancelReleaseProductResponse is the response struct for api CancelReleaseProduct
type CancelReleaseProductResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Code         string `json:"Code" xml:"Code"`
}

// CreateCancelReleaseProductRequest creates a request to invoke CancelReleaseProduct API
func CreateCancelReleaseProductRequest() (request *CancelReleaseProductRequest) {
	request = &CancelReleaseProductRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "CancelReleaseProduct", "", "")
	request.Method = requests.POST
	return
}

// CreateCancelReleaseProductResponse creates a response to parse from CancelReleaseProduct response
func CreateCancelReleaseProductResponse() (response *CancelReleaseProductResponse) {
	response = &CancelReleaseProductResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
