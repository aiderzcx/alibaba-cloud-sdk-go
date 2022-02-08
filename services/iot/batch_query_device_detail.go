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

// BatchQueryDeviceDetail invokes the iot.BatchQueryDeviceDetail API synchronously
func (client *Client) BatchQueryDeviceDetail(request *BatchQueryDeviceDetailRequest) (response *BatchQueryDeviceDetailResponse, err error) {
	response = CreateBatchQueryDeviceDetailResponse()
	err = client.DoAction(request, response)
	return
}

// BatchQueryDeviceDetailWithChan invokes the iot.BatchQueryDeviceDetail API asynchronously
func (client *Client) BatchQueryDeviceDetailWithChan(request *BatchQueryDeviceDetailRequest) (<-chan *BatchQueryDeviceDetailResponse, <-chan error) {
	responseChan := make(chan *BatchQueryDeviceDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BatchQueryDeviceDetail(request)
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

// BatchQueryDeviceDetailWithCallback invokes the iot.BatchQueryDeviceDetail API asynchronously
func (client *Client) BatchQueryDeviceDetailWithCallback(request *BatchQueryDeviceDetailRequest, callback func(response *BatchQueryDeviceDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BatchQueryDeviceDetailResponse
		var err error
		defer close(result)
		response, err = client.BatchQueryDeviceDetail(request)
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

// BatchQueryDeviceDetailRequest is the request struct for api BatchQueryDeviceDetail
type BatchQueryDeviceDetailRequest struct {
	*requests.RpcRequest
	RealTenantId      string    `position:"Query" name:"RealTenantId"`
	RealTripartiteKey string    `position:"Query" name:"RealTripartiteKey"`
	IotInstanceId     string    `position:"Query" name:"IotInstanceId"`
	ProductKey        string    `position:"Query" name:"ProductKey"`
	ApiProduct        string    `position:"Body" name:"ApiProduct"`
	ApiRevision       string    `position:"Body" name:"ApiRevision"`
	DeviceName        *[]string `position:"Query" name:"DeviceName"  type:"Repeated"`
}

// BatchQueryDeviceDetailResponse is the response struct for api BatchQueryDeviceDetail
type BatchQueryDeviceDetailResponse struct {
	*responses.BaseResponse
	RequestId    string                       `json:"RequestId" xml:"RequestId"`
	Success      bool                         `json:"Success" xml:"Success"`
	Code         string                       `json:"Code" xml:"Code"`
	ErrorMessage string                       `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         DataInBatchQueryDeviceDetail `json:"Data" xml:"Data"`
}

// CreateBatchQueryDeviceDetailRequest creates a request to invoke BatchQueryDeviceDetail API
func CreateBatchQueryDeviceDetailRequest() (request *BatchQueryDeviceDetailRequest) {
	request = &BatchQueryDeviceDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "BatchQueryDeviceDetail", "", "")
	request.Method = requests.POST
	return
}

// CreateBatchQueryDeviceDetailResponse creates a response to parse from BatchQueryDeviceDetail response
func CreateBatchQueryDeviceDetailResponse() (response *BatchQueryDeviceDetailResponse) {
	response = &BatchQueryDeviceDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
