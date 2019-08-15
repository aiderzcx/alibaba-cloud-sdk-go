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

// SetReqAuthConfig invokes the cdn.SetReqAuthConfig API synchronously
// api document: https://help.aliyun.com/api/cdn/setreqauthconfig.html
func (client *Client) SetReqAuthConfig(request *SetReqAuthConfigRequest) (response *SetReqAuthConfigResponse, err error) {
	response = CreateSetReqAuthConfigResponse()
	err = client.DoAction(request, response)
	return
}

// SetReqAuthConfigWithChan invokes the cdn.SetReqAuthConfig API asynchronously
// api document: https://help.aliyun.com/api/cdn/setreqauthconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetReqAuthConfigWithChan(request *SetReqAuthConfigRequest) (<-chan *SetReqAuthConfigResponse, <-chan error) {
	responseChan := make(chan *SetReqAuthConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetReqAuthConfig(request)
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

// SetReqAuthConfigWithCallback invokes the cdn.SetReqAuthConfig API asynchronously
// api document: https://help.aliyun.com/api/cdn/setreqauthconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetReqAuthConfigWithCallback(request *SetReqAuthConfigRequest, callback func(response *SetReqAuthConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetReqAuthConfigResponse
		var err error
		defer close(result)
		response, err = client.SetReqAuthConfig(request)
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

// SetReqAuthConfigRequest is the request struct for api SetReqAuthConfig
type SetReqAuthConfigRequest struct {
	*requests.RpcRequest
	Key1           string           `position:"Query" name:"Key1"`
	Key2           string           `position:"Query" name:"Key2"`
	AuthRemoteDesc string           `position:"Query" name:"AuthRemoteDesc"`
	SecurityToken  string           `position:"Query" name:"SecurityToken"`
	DomainName     string           `position:"Query" name:"DomainName"`
	OwnerId        requests.Integer `position:"Query" name:"OwnerId"`
	TimeOut        string           `position:"Query" name:"TimeOut"`
	AuthType       string           `position:"Query" name:"AuthType"`
}

// SetReqAuthConfigResponse is the response struct for api SetReqAuthConfig
type SetReqAuthConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetReqAuthConfigRequest creates a request to invoke SetReqAuthConfig API
func CreateSetReqAuthConfigRequest() (request *SetReqAuthConfigRequest) {
	request = &SetReqAuthConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "SetReqAuthConfig", "cdn", "openAPI")
	return
}

// CreateSetReqAuthConfigResponse creates a response to parse from SetReqAuthConfig response
func CreateSetReqAuthConfigResponse() (response *SetReqAuthConfigResponse) {
	response = &SetReqAuthConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
