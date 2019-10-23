package qualitycheck

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

// DeleteScoreForApi invokes the qualitycheck.DeleteScoreForApi API synchronously
// api document: https://help.aliyun.com/api/qualitycheck/deletescoreforapi.html
func (client *Client) DeleteScoreForApi(request *DeleteScoreForApiRequest) (response *DeleteScoreForApiResponse, err error) {
	response = CreateDeleteScoreForApiResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteScoreForApiWithChan invokes the qualitycheck.DeleteScoreForApi API asynchronously
// api document: https://help.aliyun.com/api/qualitycheck/deletescoreforapi.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteScoreForApiWithChan(request *DeleteScoreForApiRequest) (<-chan *DeleteScoreForApiResponse, <-chan error) {
	responseChan := make(chan *DeleteScoreForApiResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteScoreForApi(request)
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

// DeleteScoreForApiWithCallback invokes the qualitycheck.DeleteScoreForApi API asynchronously
// api document: https://help.aliyun.com/api/qualitycheck/deletescoreforapi.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteScoreForApiWithCallback(request *DeleteScoreForApiRequest, callback func(response *DeleteScoreForApiResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteScoreForApiResponse
		var err error
		defer close(result)
		response, err = client.DeleteScoreForApi(request)
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

// DeleteScoreForApiRequest is the request struct for api DeleteScoreForApi
type DeleteScoreForApiRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JsonStr         string           `position:"Query" name:"JsonStr"`
}

// DeleteScoreForApiResponse is the response struct for api DeleteScoreForApi
type DeleteScoreForApiResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateDeleteScoreForApiRequest creates a request to invoke DeleteScoreForApi API
func CreateDeleteScoreForApiRequest() (request *DeleteScoreForApiRequest) {
	request = &DeleteScoreForApiRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2019-01-15", "DeleteScoreForApi", "", "")
	return
}

// CreateDeleteScoreForApiResponse creates a response to parse from DeleteScoreForApi response
func CreateDeleteScoreForApiResponse() (response *DeleteScoreForApiResponse) {
	response = &DeleteScoreForApiResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}