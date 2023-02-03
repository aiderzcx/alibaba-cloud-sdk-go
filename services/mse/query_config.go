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

// QueryConfig invokes the mse.QueryConfig API synchronously
func (client *Client) QueryConfig(request *QueryConfigRequest) (response *QueryConfigResponse, err error) {
	response = CreateQueryConfigResponse()
	err = client.DoAction(request, response)
	return
}

// QueryConfigWithChan invokes the mse.QueryConfig API asynchronously
func (client *Client) QueryConfigWithChan(request *QueryConfigRequest) (<-chan *QueryConfigResponse, <-chan error) {
	responseChan := make(chan *QueryConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryConfig(request)
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

// QueryConfigWithCallback invokes the mse.QueryConfig API asynchronously
func (client *Client) QueryConfigWithCallback(request *QueryConfigRequest, callback func(response *QueryConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryConfigResponse
		var err error
		defer close(result)
		response, err = client.QueryConfig(request)
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

// QueryConfigRequest is the request struct for api QueryConfig
type QueryConfigRequest struct {
	*requests.RpcRequest
	NeedRunningConf requests.Boolean `position:"Query" name:"NeedRunningConf"`
	MseSessionId    string           `position:"Query" name:"MseSessionId"`
	ConfigType      string           `position:"Query" name:"ConfigType"`
	ClusterId       string           `position:"Query" name:"ClusterId"`
	InstanceId      string           `position:"Query" name:"InstanceId"`
	RequestPars     string           `position:"Query" name:"RequestPars"`
	AcceptLanguage  string           `position:"Query" name:"AcceptLanguage"`
}

// QueryConfigResponse is the response struct for api QueryConfig
type QueryConfigResponse struct {
	*responses.BaseResponse
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           int    `json:"Code" xml:"Code"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Data           Data   `json:"Data" xml:"Data"`
}

// CreateQueryConfigRequest creates a request to invoke QueryConfig API
func CreateQueryConfigRequest() (request *QueryConfigRequest) {
	request = &QueryConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "QueryConfig", "mse", "openAPI")
	request.Method = requests.GET
	return
}

// CreateQueryConfigResponse creates a response to parse from QueryConfig response
func CreateQueryConfigResponse() (response *QueryConfigResponse) {
	response = &QueryConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
