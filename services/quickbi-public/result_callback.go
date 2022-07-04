package quickbi_public

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

// ResultCallback invokes the quickbi_public.ResultCallback API synchronously
func (client *Client) ResultCallback(request *ResultCallbackRequest) (response *ResultCallbackResponse, err error) {
	response = CreateResultCallbackResponse()
	err = client.DoAction(request, response)
	return
}

// ResultCallbackWithChan invokes the quickbi_public.ResultCallback API asynchronously
func (client *Client) ResultCallbackWithChan(request *ResultCallbackRequest) (<-chan *ResultCallbackResponse, <-chan error) {
	responseChan := make(chan *ResultCallbackResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ResultCallback(request)
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

// ResultCallbackWithCallback invokes the quickbi_public.ResultCallback API asynchronously
func (client *Client) ResultCallbackWithCallback(request *ResultCallbackRequest, callback func(response *ResultCallbackResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ResultCallbackResponse
		var err error
		defer close(result)
		response, err = client.ResultCallback(request)
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

// ResultCallbackRequest is the request struct for api ResultCallback
type ResultCallbackRequest struct {
	*requests.RpcRequest
	AccessPoint   string           `position:"Query" name:"AccessPoint"`
	SignType      string           `position:"Query" name:"SignType"`
	HandleReason  string           `position:"Query" name:"HandleReason"`
	ApplicationId string           `position:"Query" name:"ApplicationId"`
	Status        requests.Integer `position:"Query" name:"Status"`
}

// ResultCallbackResponse is the response struct for api ResultCallback
type ResultCallbackResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    bool   `json:"Result" xml:"Result"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateResultCallbackRequest creates a request to invoke ResultCallback API
func CreateResultCallbackRequest() (request *ResultCallbackRequest) {
	request = &ResultCallbackRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("quickbi-public", "2022-01-01", "ResultCallback", "quick", "openAPI")
	request.Method = requests.POST
	return
}

// CreateResultCallbackResponse creates a response to parse from ResultCallback response
func CreateResultCallbackResponse() (response *ResultCallbackResponse) {
	response = &ResultCallbackResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
