package cloudauth

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

// Mobile3MetaSimpleVerify invokes the cloudauth.Mobile3MetaSimpleVerify API synchronously
func (client *Client) Mobile3MetaSimpleVerify(request *Mobile3MetaSimpleVerifyRequest) (response *Mobile3MetaSimpleVerifyResponse, err error) {
	response = CreateMobile3MetaSimpleVerifyResponse()
	err = client.DoAction(request, response)
	return
}

// Mobile3MetaSimpleVerifyWithChan invokes the cloudauth.Mobile3MetaSimpleVerify API asynchronously
func (client *Client) Mobile3MetaSimpleVerifyWithChan(request *Mobile3MetaSimpleVerifyRequest) (<-chan *Mobile3MetaSimpleVerifyResponse, <-chan error) {
	responseChan := make(chan *Mobile3MetaSimpleVerifyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.Mobile3MetaSimpleVerify(request)
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

// Mobile3MetaSimpleVerifyWithCallback invokes the cloudauth.Mobile3MetaSimpleVerify API asynchronously
func (client *Client) Mobile3MetaSimpleVerifyWithCallback(request *Mobile3MetaSimpleVerifyRequest, callback func(response *Mobile3MetaSimpleVerifyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *Mobile3MetaSimpleVerifyResponse
		var err error
		defer close(result)
		response, err = client.Mobile3MetaSimpleVerify(request)
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

// Mobile3MetaSimpleVerifyRequest is the request struct for api Mobile3MetaSimpleVerify
type Mobile3MetaSimpleVerifyRequest struct {
	*requests.RpcRequest
	ParamType   string `position:"Body" name:"ParamType"`
	Mobile      string `position:"Body" name:"Mobile"`
	IdentifyNum string `position:"Body" name:"IdentifyNum"`
	UserName    string `position:"Body" name:"UserName"`
}

// Mobile3MetaSimpleVerifyResponse is the response struct for api Mobile3MetaSimpleVerify
type Mobile3MetaSimpleVerifyResponse struct {
	*responses.BaseResponse
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	Code         string       `json:"Code" xml:"Code"`
	Message      string       `json:"Message" xml:"Message"`
	ResultObject ResultObject `json:"ResultObject" xml:"ResultObject"`
}

// CreateMobile3MetaSimpleVerifyRequest creates a request to invoke Mobile3MetaSimpleVerify API
func CreateMobile3MetaSimpleVerifyRequest() (request *Mobile3MetaSimpleVerifyRequest) {
	request = &Mobile3MetaSimpleVerifyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudauth", "2019-03-07", "Mobile3MetaSimpleVerify", "cloudauth", "openAPI")
	request.Method = requests.POST
	return
}

// CreateMobile3MetaSimpleVerifyResponse creates a response to parse from Mobile3MetaSimpleVerify response
func CreateMobile3MetaSimpleVerifyResponse() (response *Mobile3MetaSimpleVerifyResponse) {
	response = &Mobile3MetaSimpleVerifyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
