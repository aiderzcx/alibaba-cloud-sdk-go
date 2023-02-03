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

// GetNacosConfig invokes the mse.GetNacosConfig API synchronously
func (client *Client) GetNacosConfig(request *GetNacosConfigRequest) (response *GetNacosConfigResponse, err error) {
	response = CreateGetNacosConfigResponse()
	err = client.DoAction(request, response)
	return
}

// GetNacosConfigWithChan invokes the mse.GetNacosConfig API asynchronously
func (client *Client) GetNacosConfigWithChan(request *GetNacosConfigRequest) (<-chan *GetNacosConfigResponse, <-chan error) {
	responseChan := make(chan *GetNacosConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetNacosConfig(request)
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

// GetNacosConfigWithCallback invokes the mse.GetNacosConfig API asynchronously
func (client *Client) GetNacosConfigWithCallback(request *GetNacosConfigRequest, callback func(response *GetNacosConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetNacosConfigResponse
		var err error
		defer close(result)
		response, err = client.GetNacosConfig(request)
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

// GetNacosConfigRequest is the request struct for api GetNacosConfig
type GetNacosConfigRequest struct {
	*requests.RpcRequest
	MseSessionId   string           `position:"Query" name:"MseSessionId"`
	InstanceId     string           `position:"Query" name:"InstanceId"`
	DataId         string           `position:"Query" name:"DataId"`
	NamespaceId    string           `position:"Query" name:"NamespaceId"`
	AcceptLanguage string           `position:"Query" name:"AcceptLanguage"`
	Beta           requests.Boolean `position:"Query" name:"Beta"`
	Group          string           `position:"Query" name:"Group"`
}

// GetNacosConfigResponse is the response struct for api GetNacosConfig
type GetNacosConfigResponse struct {
	*responses.BaseResponse
	Message       string        `json:"Message" xml:"Message"`
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	ErrorCode     string        `json:"ErrorCode" xml:"ErrorCode"`
	Success       bool          `json:"Success" xml:"Success"`
	Configuration Configuration `json:"Configuration" xml:"Configuration"`
}

// CreateGetNacosConfigRequest creates a request to invoke GetNacosConfig API
func CreateGetNacosConfigRequest() (request *GetNacosConfigRequest) {
	request = &GetNacosConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "GetNacosConfig", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetNacosConfigResponse creates a response to parse from GetNacosConfig response
func CreateGetNacosConfigResponse() (response *GetNacosConfigResponse) {
	response = &GetNacosConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
