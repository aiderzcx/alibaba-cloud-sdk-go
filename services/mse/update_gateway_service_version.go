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

// UpdateGatewayServiceVersion invokes the mse.UpdateGatewayServiceVersion API synchronously
func (client *Client) UpdateGatewayServiceVersion(request *UpdateGatewayServiceVersionRequest) (response *UpdateGatewayServiceVersionResponse, err error) {
	response = CreateUpdateGatewayServiceVersionResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateGatewayServiceVersionWithChan invokes the mse.UpdateGatewayServiceVersion API asynchronously
func (client *Client) UpdateGatewayServiceVersionWithChan(request *UpdateGatewayServiceVersionRequest) (<-chan *UpdateGatewayServiceVersionResponse, <-chan error) {
	responseChan := make(chan *UpdateGatewayServiceVersionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateGatewayServiceVersion(request)
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

// UpdateGatewayServiceVersionWithCallback invokes the mse.UpdateGatewayServiceVersion API asynchronously
func (client *Client) UpdateGatewayServiceVersionWithCallback(request *UpdateGatewayServiceVersionRequest, callback func(response *UpdateGatewayServiceVersionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateGatewayServiceVersionResponse
		var err error
		defer close(result)
		response, err = client.UpdateGatewayServiceVersion(request)
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

// UpdateGatewayServiceVersionRequest is the request struct for api UpdateGatewayServiceVersion
type UpdateGatewayServiceVersionRequest struct {
	*requests.RpcRequest
	MseSessionId    string           `position:"Query" name:"MseSessionId"`
	GatewayUniqueId string           `position:"Query" name:"GatewayUniqueId"`
	ServiceVersion  string           `position:"Query" name:"ServiceVersion"`
	AcceptLanguage  string           `position:"Query" name:"AcceptLanguage"`
	ServiceId       requests.Integer `position:"Query" name:"ServiceId"`
}

// UpdateGatewayServiceVersionResponse is the response struct for api UpdateGatewayServiceVersion
type UpdateGatewayServiceVersionResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string `json:"Message" xml:"Message"`
	Code           int    `json:"Code" xml:"Code"`
	Success        bool   `json:"Success" xml:"Success"`
	Data           int64  `json:"Data" xml:"Data"`
}

// CreateUpdateGatewayServiceVersionRequest creates a request to invoke UpdateGatewayServiceVersion API
func CreateUpdateGatewayServiceVersionRequest() (request *UpdateGatewayServiceVersionRequest) {
	request = &UpdateGatewayServiceVersionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "UpdateGatewayServiceVersion", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateGatewayServiceVersionResponse creates a response to parse from UpdateGatewayServiceVersion response
func CreateUpdateGatewayServiceVersionResponse() (response *UpdateGatewayServiceVersionResponse) {
	response = &UpdateGatewayServiceVersionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
