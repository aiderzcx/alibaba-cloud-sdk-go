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

// CancelCollection invokes the quickbi_public.CancelCollection API synchronously
func (client *Client) CancelCollection(request *CancelCollectionRequest) (response *CancelCollectionResponse, err error) {
	response = CreateCancelCollectionResponse()
	err = client.DoAction(request, response)
	return
}

// CancelCollectionWithChan invokes the quickbi_public.CancelCollection API asynchronously
func (client *Client) CancelCollectionWithChan(request *CancelCollectionRequest) (<-chan *CancelCollectionResponse, <-chan error) {
	responseChan := make(chan *CancelCollectionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CancelCollection(request)
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

// CancelCollectionWithCallback invokes the quickbi_public.CancelCollection API asynchronously
func (client *Client) CancelCollectionWithCallback(request *CancelCollectionRequest, callback func(response *CancelCollectionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CancelCollectionResponse
		var err error
		defer close(result)
		response, err = client.CancelCollection(request)
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

// CancelCollectionRequest is the request struct for api CancelCollection
type CancelCollectionRequest struct {
	*requests.RpcRequest
	WorksId     string `position:"Query" name:"WorksId"`
	AccessPoint string `position:"Query" name:"AccessPoint"`
	SignType    string `position:"Query" name:"SignType"`
	UserId      string `position:"Query" name:"UserId"`
}

// CancelCollectionResponse is the response struct for api CancelCollection
type CancelCollectionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    bool   `json:"Result" xml:"Result"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateCancelCollectionRequest creates a request to invoke CancelCollection API
func CreateCancelCollectionRequest() (request *CancelCollectionRequest) {
	request = &CancelCollectionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("quickbi-public", "2022-01-01", "CancelCollection", "quick", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCancelCollectionResponse creates a response to parse from CancelCollection response
func CreateCancelCollectionResponse() (response *CancelCollectionResponse) {
	response = &CancelCollectionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
