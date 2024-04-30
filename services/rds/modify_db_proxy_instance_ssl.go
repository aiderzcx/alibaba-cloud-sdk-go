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

// ModifyDbProxyInstanceSsl invokes the rds.ModifyDbProxyInstanceSsl API synchronously
func (client *Client) ModifyDbProxyInstanceSsl(request *ModifyDbProxyInstanceSslRequest) (response *ModifyDbProxyInstanceSslResponse, err error) {
	response = CreateModifyDbProxyInstanceSslResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDbProxyInstanceSslWithChan invokes the rds.ModifyDbProxyInstanceSsl API asynchronously
func (client *Client) ModifyDbProxyInstanceSslWithChan(request *ModifyDbProxyInstanceSslRequest) (<-chan *ModifyDbProxyInstanceSslResponse, <-chan error) {
	responseChan := make(chan *ModifyDbProxyInstanceSslResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDbProxyInstanceSsl(request)
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

// ModifyDbProxyInstanceSslWithCallback invokes the rds.ModifyDbProxyInstanceSsl API asynchronously
func (client *Client) ModifyDbProxyInstanceSslWithCallback(request *ModifyDbProxyInstanceSslRequest, callback func(response *ModifyDbProxyInstanceSslResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDbProxyInstanceSslResponse
		var err error
		defer close(result)
		response, err = client.ModifyDbProxyInstanceSsl(request)
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

// ModifyDbProxyInstanceSslRequest is the request struct for api ModifyDbProxyInstanceSsl
type ModifyDbProxyInstanceSslRequest struct {
	*requests.RpcRequest
	DbProxySslEnabled    string `position:"Query" name:"DbProxySslEnabled"`
	DbProxyConnectString string `position:"Query" name:"DbProxyConnectString"`
	DbInstanceId         string `position:"Query" name:"DbInstanceId"`
	DBProxyEngineType    string `position:"Query" name:"DBProxyEngineType"`
	DbProxyEndpointId    string `position:"Query" name:"DbProxyEndpointId"`
}

// ModifyDbProxyInstanceSslResponse is the response struct for api ModifyDbProxyInstanceSsl
type ModifyDbProxyInstanceSslResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyDbProxyInstanceSslRequest creates a request to invoke ModifyDbProxyInstanceSsl API
func CreateModifyDbProxyInstanceSslRequest() (request *ModifyDbProxyInstanceSslRequest) {
	request = &ModifyDbProxyInstanceSslRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "ModifyDbProxyInstanceSsl", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyDbProxyInstanceSslResponse creates a response to parse from ModifyDbProxyInstanceSsl response
func CreateModifyDbProxyInstanceSslResponse() (response *ModifyDbProxyInstanceSslResponse) {
	response = &ModifyDbProxyInstanceSslResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
