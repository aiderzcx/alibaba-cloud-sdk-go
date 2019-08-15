package cdn

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

// SetHttpHeaderConfig invokes the cdn.SetHttpHeaderConfig API synchronously
// api document: https://help.aliyun.com/api/cdn/sethttpheaderconfig.html
func (client *Client) SetHttpHeaderConfig(request *SetHttpHeaderConfigRequest) (response *SetHttpHeaderConfigResponse, err error) {
	response = CreateSetHttpHeaderConfigResponse()
	err = client.DoAction(request, response)
	return
}

// SetHttpHeaderConfigWithChan invokes the cdn.SetHttpHeaderConfig API asynchronously
// api document: https://help.aliyun.com/api/cdn/sethttpheaderconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetHttpHeaderConfigWithChan(request *SetHttpHeaderConfigRequest) (<-chan *SetHttpHeaderConfigResponse, <-chan error) {
	responseChan := make(chan *SetHttpHeaderConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetHttpHeaderConfig(request)
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

// SetHttpHeaderConfigWithCallback invokes the cdn.SetHttpHeaderConfig API asynchronously
// api document: https://help.aliyun.com/api/cdn/sethttpheaderconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetHttpHeaderConfigWithCallback(request *SetHttpHeaderConfigRequest, callback func(response *SetHttpHeaderConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetHttpHeaderConfigResponse
		var err error
		defer close(result)
		response, err = client.SetHttpHeaderConfig(request)
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

// SetHttpHeaderConfigRequest is the request struct for api SetHttpHeaderConfig
type SetHttpHeaderConfigRequest struct {
	*requests.RpcRequest
	HeaderValue   string           `position:"Query" name:"HeaderValue"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	ConfigId      requests.Integer `position:"Query" name:"ConfigId"`
	DomainName    string           `position:"Query" name:"DomainName"`
	HeaderKey     string           `position:"Query" name:"HeaderKey"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
}

// SetHttpHeaderConfigResponse is the response struct for api SetHttpHeaderConfig
type SetHttpHeaderConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetHttpHeaderConfigRequest creates a request to invoke SetHttpHeaderConfig API
func CreateSetHttpHeaderConfigRequest() (request *SetHttpHeaderConfigRequest) {
	request = &SetHttpHeaderConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "SetHttpHeaderConfig", "cdn", "openAPI")
	return
}

// CreateSetHttpHeaderConfigResponse creates a response to parse from SetHttpHeaderConfig response
func CreateSetHttpHeaderConfigResponse() (response *SetHttpHeaderConfigResponse) {
	response = &SetHttpHeaderConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
