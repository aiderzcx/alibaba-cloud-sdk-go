package idaas_doraemon

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

// RegisterAuthenticator invokes the idaas_doraemon.RegisterAuthenticator API synchronously
func (client *Client) RegisterAuthenticator(request *RegisterAuthenticatorRequest) (response *RegisterAuthenticatorResponse, err error) {
	response = CreateRegisterAuthenticatorResponse()
	err = client.DoAction(request, response)
	return
}

// RegisterAuthenticatorWithChan invokes the idaas_doraemon.RegisterAuthenticator API asynchronously
func (client *Client) RegisterAuthenticatorWithChan(request *RegisterAuthenticatorRequest) (<-chan *RegisterAuthenticatorResponse, <-chan error) {
	responseChan := make(chan *RegisterAuthenticatorResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RegisterAuthenticator(request)
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

// RegisterAuthenticatorWithCallback invokes the idaas_doraemon.RegisterAuthenticator API asynchronously
func (client *Client) RegisterAuthenticatorWithCallback(request *RegisterAuthenticatorRequest, callback func(response *RegisterAuthenticatorResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RegisterAuthenticatorResponse
		var err error
		defer close(result)
		response, err = client.RegisterAuthenticator(request)
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

// RegisterAuthenticatorRequest is the request struct for api RegisterAuthenticator
type RegisterAuthenticatorRequest struct {
	*requests.RpcRequest
	LogParams                  string `position:"Query" name:"LogParams"`
	ClientExtendParamsJson     string `position:"Query" name:"ClientExtendParamsJson"`
	UserId                     string `position:"Query" name:"UserId"`
	ServerExtendParamsJson     string `position:"Query" name:"ServerExtendParamsJson"`
	RegistrationContext        string `position:"Query" name:"RegistrationContext"`
	AuthenticatorName          string `position:"Query" name:"AuthenticatorName"`
	RequireChallengeBase64     string `position:"Query" name:"RequireChallengeBase64"`
	AuthenticatorType          string `position:"Query" name:"AuthenticatorType"`
	ClientExtendParamsJsonSign string `position:"Query" name:"ClientExtendParamsJsonSign"`
	UserSourceIp               string `position:"Query" name:"UserSourceIp"`
	ApplicationExternalId      string `position:"Query" name:"ApplicationExternalId"`
}

// RegisterAuthenticatorResponse is the response struct for api RegisterAuthenticator
type RegisterAuthenticatorResponse struct {
	*responses.BaseResponse
	RequestId         string `json:"RequestId" xml:"RequestId"`
	AuthenticatorUuid string `json:"AuthenticatorUuid" xml:"AuthenticatorUuid"`
	EtasResponseSting string `json:"EtasResponseSting" xml:"EtasResponseSting"`
}

// CreateRegisterAuthenticatorRequest creates a request to invoke RegisterAuthenticator API
func CreateRegisterAuthenticatorRequest() (request *RegisterAuthenticatorRequest) {
	request = &RegisterAuthenticatorRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("idaas-doraemon", "2021-05-20", "RegisterAuthenticator", "idaasauth", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRegisterAuthenticatorResponse creates a response to parse from RegisterAuthenticator response
func CreateRegisterAuthenticatorResponse() (response *RegisterAuthenticatorResponse) {
	response = &RegisterAuthenticatorResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
