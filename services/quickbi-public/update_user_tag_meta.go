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

// UpdateUserTagMeta invokes the quickbi_public.UpdateUserTagMeta API synchronously
func (client *Client) UpdateUserTagMeta(request *UpdateUserTagMetaRequest) (response *UpdateUserTagMetaResponse, err error) {
	response = CreateUpdateUserTagMetaResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateUserTagMetaWithChan invokes the quickbi_public.UpdateUserTagMeta API asynchronously
func (client *Client) UpdateUserTagMetaWithChan(request *UpdateUserTagMetaRequest) (<-chan *UpdateUserTagMetaResponse, <-chan error) {
	responseChan := make(chan *UpdateUserTagMetaResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateUserTagMeta(request)
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

// UpdateUserTagMetaWithCallback invokes the quickbi_public.UpdateUserTagMeta API asynchronously
func (client *Client) UpdateUserTagMetaWithCallback(request *UpdateUserTagMetaRequest, callback func(response *UpdateUserTagMetaResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateUserTagMetaResponse
		var err error
		defer close(result)
		response, err = client.UpdateUserTagMeta(request)
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

// UpdateUserTagMetaRequest is the request struct for api UpdateUserTagMeta
type UpdateUserTagMetaRequest struct {
	*requests.RpcRequest
	TagDescription string `position:"Query" name:"TagDescription"`
	TagName        string `position:"Query" name:"TagName"`
	TagId          string `position:"Query" name:"TagId"`
	AccessPoint    string `position:"Query" name:"AccessPoint"`
	SignType       string `position:"Query" name:"SignType"`
}

// UpdateUserTagMetaResponse is the response struct for api UpdateUserTagMeta
type UpdateUserTagMetaResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    bool   `json:"Result" xml:"Result"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateUpdateUserTagMetaRequest creates a request to invoke UpdateUserTagMeta API
func CreateUpdateUserTagMetaRequest() (request *UpdateUserTagMetaRequest) {
	request = &UpdateUserTagMetaRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("quickbi-public", "2022-01-01", "UpdateUserTagMeta", "quick", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateUserTagMetaResponse creates a response to parse from UpdateUserTagMeta response
func CreateUpdateUserTagMetaResponse() (response *UpdateUserTagMetaResponse) {
	response = &UpdateUserTagMetaResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
